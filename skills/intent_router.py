"""
意图路由模块

自动识别用户意图并路由到对应 Skill
"""

import re
from typing import Dict, Optional, List, Tuple


# 意图关键词配置
INTENT_KEYWORDS = {
    "order_processing": {
        "keywords": [
            "处理订单", "生成模板", "订单转华鼎", "导出模板", 
            "生成华鼎模板", "转换订单", "订单转换"
        ],
        "skill": "skill-order-to-huading-template",
        "required_params": ["order_file", "shipper_name"],
        "optional_params": ["output_file"]
    },
    "bill_calculation": {
        "keywords": [
            "计算账单", "生成账单", "制作账单", "月度账单",
            "某账期账单", "这个月账单", "账单计算"
        ],
        "skill": "billing-workflow.step4",
        "required_params": ["shipper_name", "period"],
        "optional_params": ["outbound_data"]
    },
    "contract_parsing": {
        "keywords": [
            "更新合同", "上传合同", "解析合同", "提取合同规则",
            "合同规则", "合同解析"
        ],
        "skill": "billing-workflow.step2",
        "required_params": ["contract_file", "shipper_name"],
        "optional_params": []
    },
    "sku_coefficient": {
        "keywords": [
            "更新SKU", "更新商品系数", "重新计算系数",
            "SKU系数", "商品系数"
        ],
        "skill": "billing-workflow.step3",
        "required_params": ["product_excel", "shipper_name"],
        "optional_params": []
    },
    "bill_audit": {
        "keywords": [
            "审核账单", "比对账单", "账单差异", "账单核对"
        ],
        "skill": "billing-workflow.step5",
        "required_params": ["shipper_name", "period"],
        "optional_params": ["customer_bill_excel"]
    },
    "full_update": {
        "keywords": [
            "全面更新", "完整流程", "重新制作", "全部更新"
        ],
        "skill": "billing-workflow",
        "required_params": ["contract_file", "product_excel", "shipper_name", "period"],
        "optional_params": ["outbound_data"]
    }
}


def match_intent(user_input: str) -> Tuple[Optional[str], float]:
    """
    匹配用户意图
    
    Args:
        user_input: 用户输入文本
    
    Returns:
        (intent_name, confidence_score)
    """
    user_input_lower = user_input.lower()
    best_match = None
    best_score = 0.0
    
    for intent_name, intent_config in INTENT_KEYWORDS.items():
        for keyword in intent_config["keywords"]:
            # 计算相似度
            score = calculate_keyword_score(user_input, keyword)
            if score > best_score:
                best_match = intent_name
                best_score = score
    
    return best_match, best_score


def calculate_keyword_score(text: str, keyword: str) -> float:
    """
    计算关键词匹配得分
    
    Args:
        text: 用户输入
        keyword: 关键词
    
    Returns:
        0.0 ~ 1.0 的得分
    """
    text_lower = text.lower()
    keyword_lower = keyword.lower()
    
    # 完全包含
    if keyword_lower in text_lower:
        return 1.0
    
    # 关键词包含在文本中
    if text_lower in keyword_lower:
        return 0.9
    
    # 简单相似度计算
    common_chars = set(text_lower) & set(keyword_lower)
    if len(common_chars) > 0:
        return len(common_chars) / max(len(set(keyword_lower)), 1) * 0.8
    
    return 0.0


def extract_params(user_input: str) -> Dict[str, str]:
    """
    从用户输入中提取参数
    
    Args:
        user_input: 用户输入文本
    
    Returns:
        参数字典
    """
    params = {}
    
    # 提取货主名称（需要结合上下文和数据库）
    shipper_name = extract_shipper_name(user_input)
    if shipper_name:
        params["shipper_name"] = shipper_name
    
    # 提取账期
    period = extract_period(user_input)
    if period:
        params["period"] = period
    
    return params


def extract_shipper_name(text: str) -> Optional[str]:
    """
    提取货主名称
    
    Args:
        text: 用户输入文本
    
    Returns:
        货主名称或None
    """
    # 常见货主名称模式
    patterns = [
        r"([\u4e00-\u9fa5]+(?:上黎|广承|供应链|物流|商贸))",
        r"([\u4e00-\u9fa5]+公司)",
        r"货主[是为]?([\u4e00-\u9fa5]+)",
        r"客户[是为]?([\u4e00-\u9fa5]+)",
    ]
    
    for pattern in patterns:
        match = re.search(pattern, text)
        if match:
            return match.group(1)
    
    return None


def extract_period(text: str) -> Optional[str]:
    """
    提取账期
    
    Args:
        text: 用户输入文本
    
    Returns:
        账期字符串或None
    """
    # 账期模式
    patterns = [
        r"(\d{4})年(\d{1,2})月",      # 2026年2月
        r"(\d{4})-(\d{1,2})",          # 2026-02
        r"(\d{1,2})月",               # 2月
        r"本月|当月|本月",            # 本月
        r"上月|上个月",               # 上月
    ]
    
    for pattern in patterns:
        match = re.search(pattern, text)
        if match:
            if match.lastindex == 2:
                return f"{match.group(1)}年{match.group(2)}月"
            elif match.lastindex == 1:
                return f"{match.group(1)}"
            else:
                return match.group(0)
    
    return None


def get_intent_config(intent_name: str) -> Optional[Dict]:
    """
    获取意图配置
    
    Args:
        intent_name: 意图名称
    
    Returns:
        意图配置字典
    """
    return INTENT_KEYWORDS.get(intent_name)


def validate_params(intent_name: str, params: Dict) -> Tuple[bool, List[str]]:
    """
    验证参数是否完整
    
    Args:
        intent_name: 意图名称
        params: 提取的参数
    
    Returns:
        (is_valid, missing_params)
    """
    config = get_intent_config(intent_name)
    if not config:
        return False, ["未知意图"]
    
    missing = []
    for required_param in config["required_params"]:
        if required_param not in params or not params[required_param]:
            missing.append(required_param)
    
    return len(missing) == 0, missing


def resolve_intent(user_input: str, available_data: Dict = None) -> Dict:
    """
    解析用户意图并返回路由信息
    
    Args:
        user_input: 用户输入文本
        available_data: 可用的上下文数据（如已上传的文件等）
    
    Returns:
        {
            "intent": intent_name,
            "confidence": score,
            "skill": target_skill,
            "params": extracted_params,
            "missing_params": [],
            "suggestion": "下一步建议"
        }
    """
    available_data = available_data or {}
    
    # 1. 匹配意图
    intent_name, score = match_intent(user_input)
    
    if not intent_name:
        return {
            "intent": None,
            "confidence": 0.0,
            "skill": None,
            "params": {},
            "missing_params": [],
            "suggestion": "无法理解您的意图，请尝试：处理订单、计算账单、解析合同等"
        }
    
    # 2. 提取参数
    params = extract_params(user_input)
    params.update(available_data)  # 合并可用数据
    
    # 3. 验证参数
    is_valid, missing = validate_params(intent_name, params)
    
    # 4. 获取Skill信息
    config = get_intent_config(intent_name)
    
    # 5. 生成建议
    suggestion = ""
    if missing:
        suggestion = f"请提供：{'、'.join(missing)}"
    else:
        suggestion = f"将执行{config['skill']}，请问确认吗？"
    
    return {
        "intent": intent_name,
        "confidence": score,
        "skill": config["skill"],
        "params": params,
        "missing_params": missing,
        "suggestion": suggestion
    }


# 测试
if __name__ == "__main__":
    test_inputs = [
        "帮我处理这个订单，生成华鼎模板",
        "生成长沙广承2026年2月账单",
        "上传合同，解析规则",
        "计算账单",
    ]
    
    for inp in test_inputs:
        result = resolve_intent(inp)
        print(f"\n输入: {inp}")
        print(f"意图: {result['intent']}")
        print(f"置信度: {result['confidence']}")
        print(f"Skill: {result['skill']}")
        print(f"参数: {result['params']}")
        print(f"缺失: {result['missing_params']}")
        print(f"建议: {result['suggestion']}")