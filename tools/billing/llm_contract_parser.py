# billing_tools/llm_contract_parser.py
# LLM合同解析器 - 使用视觉模型从合同图片中提取规则

"""
LLM合同解析

功能：
- 读取合同图片（jpg/png/pdf）
- 将图片编码为base64
- 调用视觉LLM提取结构化规则
- 输出与 contract.py 兼容的规则结构

支持模型：
- OpenClaw 配置的默认模型（支持vision）
- 或者显式指定模型如 minimax/gpt-4o
"""

import sys
import os
import json
import base64

# 添加src到路径
_tools_dir = os.path.dirname(os.path.abspath(__file__))
_src_dir = os.path.join(_tools_dir, 'src')
if _src_dir not in sys.path:
    sys.path.insert(0, _src_dir)


# 系统提示词模板
SYSTEM_PROMPT = """你是华鼎冷链仓储物流服务的合同规则提取专家。

请从合同图片中提取计费规则，输出JSON格式：

{
  "storage": {"冷冻": 3.2, "冷藏": 3.2, "常温": 1.0},  // 存储费：元/托/天
  "handling": {"整件": 1.5, "拆零": 0.3},  // 操作费：元/标准件
  "transfer": {"长沙->广州": 5.0, "长沙->成都": 8.0},  // 调拨费：元/标准件
  "delivery": {
    "长沙": {"长沙市": {"price": 3.5, "min": 40}},
    "广州": {"珠三角": {"price": 3.5, "min": 45}}
  },  // 共配费：price=元/标准件, min=最低收费
  "value_added": {"灌装分箱": 0.5, "贴标": 0.3},
  "oversized": {  // 超规格规则（从合同提取）
    "基准重量_kg": 15,
    "基准体积_m3": 0.05,
    "threshold_tier1": {"weight_kg": {"min": 15, "max": 25}, "volume_m3": {"min": 0.05, "max": 0.08}},
    "threshold_tier2": {"weight_kg": {"min": 25, "max": 50}, "volume_m3": {"min": 0.08, "max": 0.1}},
    "threshold_tier3": {"weight_kg": {"min": 50}, "volume_m3": {"min": 0.1}},
    "coef_tier1": 1.2,
    "coef_tier2": 2.0
  },
  "special": {"困难配送加收": 0.2, "送货进店费": 1.0},
  "tax_rate": {"仓储": 0, "运输": 0.09},
  "fixed_skus": {  // 合同中固定的SKU列表
    "GGK01": {"name": "广承康宁", "spec": "10张/包", "volume_m3": 0.001, "weight_kg": 0.5}
  }
}

规则说明：
- oversized.threshold_tier1: 15kg < 重量 ≤ 25kg 或 0.05m³ < 体积 ≤ 0.08m³ → coef_tier1倍
- oversized.threshold_tier2: 25kg < 重量 ≤ 50kg 或 0.08m³ < 体积 ≤ 0.1m³ → coef_tier2倍
- oversized.threshold_tier3: 重量 > 50kg 或 体积 > 0.1m³ → max(重量/基准重量, 体积/基准体积)
- 标准件：重量 ≤ 基准重量 且 体积 ≤ 基准体积 → 系数1.0
- 如果合同没有某项规则，设为null或空对象
"""


def read_image_base64(image_path: str) -> str:
    """读取图片并转为base64"""
    if not os.path.exists(image_path):
        raise FileNotFoundError(f"图片不存在: {image_path}")
    
    with open(image_path, "rb") as f:
        image_data = f.read()
    
    # 判断文件类型
    ext = os.path.splitext(image_path)[1].lower()
    mime_types = {
        ".jpg": "image/jpeg",
        ".jpeg": "image/jpeg",
        ".png": "image/png",
        ".gif": "image/gif",
        ".webp": "image/webp",
        ".pdf": "application/pdf"
    }
    mime_type = mime_types.get(ext, "image/jpeg")
    
    return base64.b64encode(image_data).decode("utf-8")


def parse_contract_with_vision(
    image_path: str,
    model: str = None,
    api_key: str = None,
    base_url: str = None
) -> dict:
    """
    使用视觉LLM解析合同图片，提取计费规则
    
    Args:
        image_path: 合同图片路径（jpg/png/pdf）
        model: 使用的模型（默认使用OpenClaw配置）
        api_key: API密钥（默认使用环境变量）
        base_url: API地址（默认使用环境变量）
        
    Returns:
        {
            "success": bool,
            "rules": dict,
            "confidence": float,
            "missing_fields": list,
            "message": str
        }
    """
    try:
        # 1. 读取图片
        image_base64 = read_image_base64(image_path)
        image_url = f"data:image/jpeg;base64,{image_base64}"
        
        # 2. 调用LLM
        rules = _call_vision_llm(image_url, model, api_key, base_url)
        
        # 3. 验证和补充
        rules = _normalize_rules(rules)
        
        # 4. 检查必填字段
        missing = _check_required_fields(rules)
        confidence = _calculate_confidence(rules, missing)
        
        return {
            "success": len(missing) == 0,
            "rules": rules,
            "confidence": confidence,
            "missing_fields": missing,
            "message": "解析完成" if not missing else f"缺少: {', '.join(missing)}"
        }
        
    except Exception as e:
        return {
            "success": False,
            "rules": None,
            "confidence": 0,
            "missing_fields": [],
            "message": f"解析失败: {str(e)}"
        }


def _call_vision_llm(
    image_url: str,
    model: str = None,
    api_key: str = None,
    base_url: str = None
) -> dict:
    """
    调用视觉LLM提取规则
    
    Args:
        image_url: base64编码的图片URL
        model: 模型名称
        api_key: API密钥
        base_url: API地址
        
    Returns:
        提取的规则字典
    """
    try:
        # 尝试使用 OpenAI SDK 呼叫支持 vision 的模型
        try:
            from openai import OpenAI
        except ImportError:
            raise ImportError("需要安装 openai: pip install openai")
        
        # 读取环境变量或使用传入的值
        if not api_key:
            api_key = os.environ.get("OPENAI_API_KEY") or os.environ.get("MINIMAX_API_KEY")
        if not base_url:
            base_url = os.environ.get("OPENAI_API_BASE") or os.environ.get("MINIMAX_BASE_URL", "https://api.minimax.chat")
        
        if not api_key:
            raise ValueError("未设置API密钥，请设置 OPENAI_API_KEY 或 MINIMAX_API_KEY 环境变量")
        
        client = OpenAI(api_key=api_key, base_url=base_url)
        
        # 选择模型（支持 vision 的模型）
        if not model:
            # 优先使用支持 vision 的模型
            model = os.environ.get("VISION_MODEL") or "minimax/MiniMax-V01"
        
        # 构建消息
        messages = [
            {"role": "system", "content": SYSTEM_PROMPT},
            {
                "role": "user",
                "content": [
                    {
                        "type": "image_url",
                        "image_url": {"url": image_url, "detail": "high"}
                    },
                    {
                        "type": "text",
                        "text": "请从这张合同图片中提取所有计费规则，输出JSON格式。"
                    }
                ]
            }
        ]
        
        # 调用API
        response = client.chat.completions.create(
            model=model,
            messages=messages,
            temperature=0.1,  # 低温度确保稳定性
            max_tokens=4096
        )
        
        # 解析响应
        content = response.choices[0].message.content
        
        # 提取JSON（可能在markdown代码块中）
        content = content.strip()
        if content.startswith("```"):
            # 去掉markdown代码块
            lines = content.split("\n")
            content = "\n".join(lines[1:-1])  # 去掉 ```json 和 ```
        
        rules = json.loads(content)
        return rules
        
    except json.JSONDecodeError as e:
        raise ValueError(f"LLM返回的不是有效JSON: {str(e)}\n原始内容: {content[:500]}")
    except Exception as e:
        raise RuntimeError(f"LLM调用失败: {str(e)}")


def _normalize_rules(rules: dict) -> dict:
    """
    规范化规则结构，补充默认值
    
    Args:
        rules: LLM提取的规则
        
    Returns:
        规范化后的规则
    """
    normalized = {}
    
    # storage
    normalized["storage"] = rules.get("storage", {})
    
    # handling
    normalized["handling"] = rules.get("handling", {})
    
    # transfer
    normalized["transfer"] = rules.get("transfer", {"min_order_qty": 50, "unit": "元/标准件"})
    
    # delivery
    normalized["delivery"] = rules.get("delivery", {})
    
    # value_added
    normalized["value_added"] = rules.get("value_added", {})
    
    # oversized（核心：超规格规则）
    oversized = rules.get("oversized", {})
    if oversized:
        normalized["oversized"] = {
            "基准重量_kg": oversized.get("基准重量_kg", 15),
            "基准体积_m3": oversized.get("基准体积_m3", 0.05),
            "threshold_tier1": oversized.get("threshold_tier1", {
                "weight_kg": {"min": 15, "max": 25},
                "volume_m3": {"min": 0.05, "max": 0.08}
            }),
            "threshold_tier2": oversized.get("threshold_tier2", {
                "weight_kg": {"min": 25, "max": 50},
                "volume_m3": {"min": 0.08, "max": 0.1}
            }),
            "threshold_tier3": oversized.get("threshold_tier3", {
                "weight_kg": {"min": 50},
                "volume_m3": {"min": 0.1}
            }),
            "coef_tier1": oversized.get("coef_tier1", 1.2),
            "coef_tier2": oversized.get("coef_tier2", 2.0)
        }
    else:
        # 如果LLM没提取到，使用默认值
        normalized["oversized"] = {
            "基准重量_kg": 15,
            "基准体积_m3": 0.05,
            "threshold_tier1": {"weight_kg": {"min": 15, "max": 25}, "volume_m3": {"min": 0.05, "max": 0.08}},
            "threshold_tier2": {"weight_kg": {"min": 25, "max": 50}, "volume_m3": {"min": 0.08, "max": 0.1}},
            "threshold_tier3": {"weight_kg": {"min": 50}, "volume_m3": {"min": 0.1}},
            "coef_tier1": 1.2,
            "coef_tier2": 2.0
        }
    
    # special
    normalized["special"] = rules.get("special", {})
    
    # tax_rate
    normalized["tax_rate"] = rules.get("tax_rate", {"仓储": 0, "运输": 0.09})
    
    # product_coefficients（包含fixed_skus）
    product_coefficients = rules.get("product_coefficients", {})
    if not product_coefficients and rules.get("fixed_skus"):
        product_coefficients = {
            "fixed_skus": rules.get("fixed_skus"),
            "standard_unit_definition": {
                "max_weight_kg": normalized["oversized"]["基准重量_kg"],
                "max_volume_m3": normalized["oversized"]["基准体积_m3"]
            }
        }
    normalized["product_coefficients"] = product_coefficients
    
    return normalized


def _check_required_fields(rules: dict) -> list:
    """检查必填字段"""
    missing = []
    
    if not rules.get("storage"):
        missing.append("存储费")
    
    if not rules.get("delivery") and not rules.get("transport"):
        missing.append("共配报价")
    
    return missing


def _calculate_confidence(rules: dict, missing: list) -> float:
    """计算置信度"""
    if not rules:
        return 0.0
    
    score = 1.0 - len(missing) * 0.2
    
    # 检查oversized是否完整
    oversized = rules.get("oversized", {})
    required_fields = ["基准重量_kg", "基准体积_m3", "coef_tier1", "coef_tier2"]
    for field in required_fields:
        if field not in oversized:
            score -= 0.05
    
    return round(max(min(score, 1.0), 0.0), 2)


def batch_parse_contracts(
    image_paths: list,
    model: str = None,
    api_key: str = None,
    base_url: str = None
) -> list:
    """
    批量解析多个合同图片
    
    Args:
        image_paths: 图片路径列表
        model: 模型名称
        api_key: API密钥
        base_url: API地址
        
    Returns:
        结果列表，每个元素是 parse_contract_with_vision 的返回值
    """
    results = []
    for path in image_paths:
        result = parse_contract_with_vision(path, model, api_key, base_url)
        results.append({
            "image_path": path,
            "result": result
        })
    return results


if __name__ == "__main__":
    import argparse
    
    parser = argparse.ArgumentParser(description="LLM合同解析")
    parser.add_argument("image_path", help="合同图片路径")
    parser.add_argument("--model", "-m", help="模型名称", default=None)
    parser.add_argument("--api-key", "-k", help="API密钥", default=None)
    parser.add_argument("--base-url", "-u", help="API地址", default=None)
    
    args = parser.parse_args()
    
    result = parse_contract_with_vision(
        args.image_path,
        model=args.model,
        api_key=args.api_key,
        base_url=args.base_url
    )
    
    print(f"成功: {result['success']}")
    print(f"置信度: {result['confidence']}")
    print(f"消息: {result['message']}")
    print(f"\n规则:\n{json.dumps(result['rules'], indent=2, ensure_ascii=False)}")