# -*- coding: utf-8 -*-
"""
单位转换模块

【重要】单位转换逻辑：
1. 先根据SKU的spec判断规格（大单位、小单位、换算比）
2. 再判断出库单填写的单位是什么
3. 只有当出库单单位 = 规格中的大单位时，才需要使用换算比转换为小单位
4. 如果出库单单位已经是小单位，则不需要转换

示例：
- spec="规格：12罐/件" → 1件 = 12罐，大单位=件，小单位=罐
- 客户填写"件" → 需要转换：数量 × 12 = 罐的数量
- 客户填写"罐" → 不需要转换
"""

import re
from typing import Dict, List, Optional, Tuple, Any


# 华鼎标准单位类型
UNIT_PIECE = "PIECE"    # 件/个
UNIT_BOX = "BOX"        # 箱/盒
UNIT_BAG = "BAG"        # 袋/包
UNIT_BOTTLE = "BOTTLE"  # 瓶/罐
UNIT_WEIGHT = "WEIGHT"  # 重量(kg/g)
UNIT_VOLUME = "VOLUME"  # 体积(L/ml)
UNIT_PALLET = "PALLET"  # 托/托盘
UNIT_CASE = "CASE"      # 大件
UNIT_SET = "SET"        # 套/组
UNIT_UNKNOWN = "UNKNOWN" # 未知


# 单位映射表
UNIT_MAPPINGS = [
    # 件/个类
    {
        "customer_units": ["件", "个", "只", "支", "条", "根", "台", "辆", "张", "把", "块", "片", "粒", "颗", "本", "份"],
        "standard_type": UNIT_PIECE,
        "category": "单品",
        "priority": 1,
    },
    # 箱/盒类
    {
        "customer_units": ["箱", "盒", "carton", "box", "case", "ctn"],
        "standard_type": UNIT_BOX,
        "category": "包装",
        "priority": 1,
    },
    # 袋/包类
    {
        "customer_units": ["袋", "包", "sack", "bag", "pkg", "package"],
        "standard_type": UNIT_BAG,
        "category": "包装",
        "priority": 1,
    },
    # 瓶/罐类
    {
        "customer_units": ["瓶", "罐", "听", "杯", "桶", "壶", "bottle", "can", "jar", "drum"],
        "standard_type": UNIT_BOTTLE,
        "category": "容器",
        "priority": 1,
    },
    # 重量类 - 千克
    {
        "customer_units": ["千克", "公斤", "kg", "kgs", "kilogram", "kilograms", "kilo"],
        "standard_type": UNIT_WEIGHT,
        "category": "重量",
        "priority": 1,
    },
    # 重量类 - 克
    {
        "customer_units": ["克", "g", "gram", "grams", "gm", "gms"],
        "standard_type": UNIT_WEIGHT,
        "category": "重量",
        "priority": 2,
    },
    # 重量类 - 吨
    {
        "customer_units": ["吨", "t", "ton", "tons", "tonne", "tonnes"],
        "standard_type": UNIT_WEIGHT,
        "category": "重量",
        "priority": 3,
    },
    # 重量类 - 斤/两
    {
        "customer_units": ["斤", "市斤", "两", "市两"],
        "standard_type": UNIT_WEIGHT,
        "category": "重量(市制)",
        "priority": 4,
    },
    # 体积类 - 升
    {
        "customer_units": ["升", "l", "L", "liter", "liters", "litre", "litres"],
        "standard_type": UNIT_VOLUME,
        "category": "体积",
        "priority": 1,
    },
    # 体积类 - 毫升
    {
        "customer_units": ["毫升", "ml", "mL", "milliliter", "milliliters"],
        "standard_type": UNIT_VOLUME,
        "category": "体积",
        "priority": 2,
    },
    # 托盘类
    {
        "customer_units": ["托", "托盘", "板", "pallet", "plt", "plts"],
        "standard_type": UNIT_PALLET,
        "category": "运输",
        "priority": 1,
    },
    # 大件类
    {
        "customer_units": ["大件", "大包装", "bulk"],
        "standard_type": UNIT_CASE,
        "category": "包装",
        "priority": 2,
    },
    # 套/组类
    {
        "customer_units": ["套", "组", "副", "对", "双", "set", "sets", "kit", "kits", "pair", "pairs"],
        "standard_type": UNIT_SET,
        "category": "组合",
        "priority": 1,
    },
]

# 单位转换规则
CONVERSION_RULES = [
    # 重量转换
    {"from": "g", "to": "kg", "ratio": 0.001, "description": "克转千克"},
    {"from": "kg", "to": "g", "ratio": 1000, "description": "千克转克"},
    {"from": "t", "to": "kg", "ratio": 1000, "description": "吨转千克"},
    {"from": "kg", "to": "t", "ratio": 0.001, "description": "千克转吨"},
    {"from": "斤", "to": "kg", "ratio": 0.5, "description": "市斤转千克"},
    {"from": "两", "to": "kg", "ratio": 0.05, "description": "市两转千克"},
    # 体积转换
    {"from": "ml", "to": "l", "ratio": 0.001, "description": "毫升转升"},
    {"from": "l", "to": "ml", "ratio": 1000, "description": "升转毫升"},
]


def get_unit_type(customer_unit: str) -> Dict[str, Any]:
    """
    解析客户单位，返回华鼎标准单位类型（Tool接口）
    
    参数:
        customer_unit: 客户单位（如"箱"、"件"、"kg"）
        
    返回:
        {
            "success": bool,
            "unit_type": str,           # PIECE/BOX/BAG/BOTTLE/WEIGHT/VOLUME/PALLET/CASE/SET/UNKNOWN
            "category": str,            # 单位分类
            "confidence": float,        # 置信度
            "message": str
        }
    """
    if not customer_unit:
        return {"success": False, "error": "customer_unit cannot be empty"}
    
    # 标准化输入
    unit = normalize_unit(customer_unit)
    
    # 精确匹配
    unit_type, confidence = match_unit_exact(unit)
    if confidence == 1.0:
        return {
            "success": True,
            "unit_type": unit_type,
            "category": get_unit_category(unit_type),
            "confidence": 1.0,
            "message": f"精确匹配: {customer_unit} → {unit_type}"
        }
    
    # 模糊匹配
    unit_type, confidence = match_unit_fuzzy_with_score(unit)
    if confidence >= 0.7:
        return {
            "success": True,
            "unit_type": unit_type,
            "category": get_unit_category(unit_type),
            "confidence": confidence,
            "message": f"模糊匹配: {customer_unit} → {unit_type} (置信度{confidence:.2f})"
        }
    
    # 尝试提取单位（去除数字）
    extracted_unit = extract_unit_from_string(unit)
    if extracted_unit and extracted_unit != unit:
        unit_type, confidence = match_unit_exact(extracted_unit)
        if confidence == 1.0:
            return {
                "success": True,
                "unit_type": unit_type,
                "category": get_unit_category(unit_type),
                "confidence": 0.9,
                "message": f"提取匹配: {customer_unit} → {unit_type}"
            }
    
    return {
        "success": True,
        "unit_type": UNIT_UNKNOWN,
        "category": "未知",
        "confidence": 0.0,
        "message": f"无法识别单位: {customer_unit}"
    }


def normalize_unit(unit: str) -> str:
    """标准化单位字符串"""
    unit = unit.strip()
    # 转小写
    return unit.lower()


def match_unit_exact(unit: str) -> Tuple[str, float]:
    """精确匹配单位"""
    for mapping in UNIT_MAPPINGS:
        for candidate in mapping["customer_units"]:
            if unit.lower() == candidate.lower():
                return mapping["standard_type"], 1.0
    return UNIT_UNKNOWN, 0.0


def match_unit_fuzzy_with_score(unit: str) -> Tuple[str, float]:
    """模糊匹配单位并返回置信度"""
    best_type = UNIT_UNKNOWN
    best_score = 0.0
    
    for mapping in UNIT_MAPPINGS:
        for candidate in mapping["customer_units"]:
            score = calculate_unit_similarity(unit, candidate)
            if score > best_score:
                best_score = score
                best_type = mapping["standard_type"]
    
    return best_type, best_score


def calculate_unit_similarity(s1: str, s2: str) -> float:
    """计算单位相似度"""
    if not s1 or not s2:
        return 0.0
    
    s1 = s1.lower()
    s2 = s2.lower()
    
    # 完全匹配
    if s1 == s2:
        return 1.0
    
    # 包含关系
    if s1 in s2 or s2 in s1:
        longer = max(len(s1), len(s2))
        shorter = min(len(s1), len(s2))
        return 0.8 + 0.2 * (shorter / longer)
    
    # 编辑距离相似度
    distance = levenshtein_distance(s1, s2)
    max_len = max(len(s1), len(s2))
    if max_len == 0:
        return 0.0
    
    return 1.0 - distance / max_len


def levenshtein_distance(s1: str, s2: str) -> int:
    """计算编辑距离"""
    m, n = len(s1), len(s2)
    if m == 0:
        return n
    if n == 0:
        return m
    
    # 滚动数组优化
    prev = list(range(n + 1))
    curr = [0] * (n + 1)
    
    for i in range(1, m + 1):
        curr[0] = i
        for j in range(1, n + 1):
            cost = 0 if s1[i-1] == s2[j-1] else 1
            curr[j] = min(prev[j] + 1, curr[j-1] + 1, prev[j-1] + cost)
        prev, curr = curr, prev
    
    return prev[n]


def extract_unit_from_string(s: str) -> str:
    """从字符串中提取单位部分"""
    # 去除数字和小数点
    result = ""
    for c in s:
        if c.isdigit() or c in '.':  # 保留小数点
            continue
        result += c
    return result.strip()


def get_unit_category(unit_type: str) -> str:
    """获取单位分类"""
    categories = {
        UNIT_PIECE: "单品",
        UNIT_BOX: "包装",
        UNIT_BAG: "包装",
        UNIT_BOTTLE: "容器",
        UNIT_WEIGHT: "重量",
        UNIT_VOLUME: "体积",
        UNIT_PALLET: "运输",
        UNIT_CASE: "包装",
        UNIT_SET: "组合",
        UNIT_UNKNOWN: "未知",
    }
    return categories.get(unit_type, "未知")


def convert_unit(value: float, from_unit: str, to_unit: str) -> Dict[str, Any]:
    """
    单位转换
    
    参数:
        value: 数值
        from_unit: 源单位
        to_unit: 目标单位
        
    返回:
        {
            "success": bool,
            "result": float,
            "message": str
        }
    """
    if value < 0:
        return {"success": False, "error": "value cannot be negative"}
    
    from_type_result = get_unit_type(from_unit)
    to_type_result = get_unit_type(to_unit)
    
    if from_type_result["confidence"] == 0:
        return {"success": False, "error": f"invalid from unit: {from_unit}"}
    if to_type_result["confidence"] == 0:
        return {"success": False, "error": f"invalid to unit: {to_unit}"}
    
    # 同类型直接查找转换规则
    if from_type_result["unit_type"] == to_type_result["unit_type"]:
        for rule in CONVERSION_RULES:
            if rule["from"].lower() == from_unit.lower() and rule["to"].lower() == to_unit.lower():
                return {
                    "success": True,
                    "result": value * rule["ratio"],
                    "message": f"{value} {from_unit} = {value * rule['ratio']} {to_unit}"
                }
    
    return {"success": False, "error": f"conversion from {from_unit} to {to_unit} not supported"}


def get_standard_unit_name(unit_type: str) -> str:
    """获取标准单位名称"""
    names = {
        UNIT_PIECE: "件",
        UNIT_BOX: "箱",
        UNIT_BAG: "袋",
        UNIT_BOTTLE: "瓶",
        UNIT_WEIGHT: "千克",
        UNIT_VOLUME: "升",
        UNIT_PALLET: "托",
        UNIT_CASE: "大件",
        UNIT_SET: "套",
        UNIT_UNKNOWN: "未知",
    }
    return names.get(unit_type, "未知")


def parse_quantity_with_unit(input_str: str) -> Dict[str, Any]:
    """
    解析带单位的数量字符串
    
    例如: "10箱" → {quantity: 10.0, unit: "箱"}
    """
    if not input_str:
        return {"success": False, "error": "input cannot be empty"}
    
    # 提取数字部分
    match = re.match(r'^([\d.,]+)\s*(.*)$', input_str.strip())
    if not match:
        return {"success": False, "error": f"no number found in input: {input_str}"}
    
    num_str = match.group(1).replace(',', '')
    try:
        quantity = float(num_str)
    except ValueError:
        return {"success": False, "error": f"failed to parse quantity: {num_str}"}
    
    unit = match.group(2).strip()
    
    return {
        "success": True,
        "quantity": quantity,
        "unit": unit,
        "message": f"解析成功: {quantity} {unit}" if unit else f"解析成功: {quantity}"
    }


# ==================== 基于SKU规格的单位转换 ====================

def parse_spec_to_unit_info(spec: str) -> Dict[str, Any]:
    """
    解析SKU的spec字段，提取大单位、小单位和换算比
    
    Args:
        spec: SKU规格字符串，如：
            - "规格：12罐/件"
            - "件：1kg*10袋/件"
            - "规格：6.25kg*4/件"
            - "件：1L*12盒/件"
            
    Returns:
        {
            "success": bool,
            "big_unit": str,      # 大单位，如"件"
            "small_unit": str,    # 小单位，如"罐"
            "ratio": float,       # 换算比（大单位→小单位），如12
            "big_unit_type": str, # 大单位类型，如"PIECE"
            "small_unit_type": str, # 小单位类型，如"BOTTLE"
            "message": str
        }
    """
    if not spec:
        return {
            "success": False,
            "big_unit": "",
            "small_unit": "",
            "ratio": 1.0,
            "big_unit_type": UNIT_UNKNOWN,
            "small_unit_type": UNIT_UNKNOWN,
            "message": "spec为空"
        }
    
    # 匹配模式1: "规格：12罐/件" 或 "件：12罐/件" 或 "规格：X个/件"
    # 匹配模式2: "件：1kg*10袋/件" 或 "件：6.25kg*4/件" 或 "规格：6.25kg*4/件"
    # 匹配模式3: "件：10包＊100支/件" 等
    patterns = [
        # 模式1: "规格：12罐/件" 或 "件：12罐/件" 或 "规格：X个/件" - 数字+小单位/大单位
        r'[规格件：:][\s]*(\d+(?:\.\d+)?)\s*([罐瓶袋个只支条根台辆张把块片粒颗本份箱盒杯桶壶]+)\s*/\s*([件箱盒杯桶袋])',
        # 模式2: "件：1kg*10袋/件" 或 "件：1kg*10袋" - 有重量/体积*数量+小单位/件，可选/件后缀
        r'[规格件：:]\s*(\d+(?:\.\d+)?)\s*(kg|g|L|ml)[\s]*[\*＊][\s]*(\d+(?:\.\d+)?)\s*([袋瓶罐个只支条根箱盒杯桶壶包]+)(?:/件)?',
        # 模式3: "件：10包＊100支" 或 "件：10包*100支/件" - 小单位*比值+单位，可选/件后缀
        r'[规格件：:]\s*(\d+(?:\.\d+)?)\s*([袋瓶罐个只支条根箱盒杯桶壶包]+)[\s]*[\*＊][\s]*(\d+(?:\.\d+)?)\s*([袋瓶罐个只支条根箱盒杯桶壶包]+)(?:/件)?',
        # 模式4: "规格：6.25kg*4/件" - 重量*倍数/大单位，只匹配倍数和大单位 (无小单位)
        r'[规格件：:]\s*(\d+(?:\.\d+)?)\s*(kg|g|L|ml)[\s]*[\*＊][\s]*(\d+(?:\.\d+)?)/件',
    ]
    
    for idx, pattern in enumerate(patterns):
        match = re.search(pattern, spec)
        if match:
            # 根据不同的pattern解析不同的group组合
            if idx == 0:
                # 模式1: "规格：12罐/件" - 数字+小单位/大单位
                # group(1)=比值, group(2)=小单位, group(3)=大单位
                ratio = float(match.group(1))
                small_unit = normalize_unit_name(match.group(2))
                big_unit = normalize_unit_name(match.group(3))
            elif idx == 1:
                # 模式2: "件：1kg*10袋/件" - 数字+单位*比值+小单位/件
                # group(1)=数字, group(2)=单位, group(3)=比值, group(4)=小单位
                # 解析结果: 1件 = 数字*比值 小单位
                base_num = float(match.group(1))
                base_unit = normalize_unit_name(match.group(2))
                ratio = float(match.group(3))
                small_unit = normalize_unit_name(match.group(4))
                big_unit = "件"  # 固定大单位为件
            elif idx == 2:
                # 模式3: "件：10包*100支/件" 或 "件：10包＊100支" - 数字+小单位*数字+单位，可选/件后缀
                # group(1)=第一个数字, group(2)=第一个单位, group(3)=第二个数字, group(4)=第二个单位
                # 解析结果: 1件 = 第一个数字 * 第二个数字 第二个单位
                # 例如: 件：10包*100支/件 → 1件 = 10 * 100 支 = 1000支
                num1 = float(match.group(1))
                unit1 = normalize_unit_name(match.group(2))
                num2 = float(match.group(3))
                unit2 = normalize_unit_name(match.group(4)) if match.group(4) else "支"
                ratio = num1 * num2
                small_unit = unit2  # 第二单位作为小单位（更小的那个）
                big_unit = "件"  # 固定大单位为件
            elif idx == 3:
                # 模式4: "规格：6.25kg*4/件" - 重量*倍数/大单位
                # group(1)=比值, group(2)=单位(kg/g/L), group(3)=倍数
                # 这类spec没有小单位，返回失败让用户确认
                return {
                    "success": False,
                    "big_unit": "",
                    "small_unit": "",
                    "ratio": 1.0,
                    "big_unit_type": UNIT_UNKNOWN,
                    "small_unit_type": UNIT_UNKNOWN,
                    "message": f"无法解析spec(重量格式无小单位): {spec}"
                }
            else:
                continue
            
            # 标准化单位名称
            small_unit = normalize_unit_name(small_unit)
            big_unit = normalize_unit_name(big_unit)
            
            # 获取单位类型
            small_unit_type = get_unit_type(small_unit)["unit_type"]
            big_unit_type = get_unit_type(big_unit)["unit_type"]
            
            # 对于模式2，需要重新计算ratio (数字 * 比值)
            if idx == 1:
                ratio = ratio  # ratio已经是正确的比值
            
            return {
                "success": True,
                "big_unit": big_unit,
                "small_unit": small_unit,
                "ratio": ratio,
                "big_unit_type": big_unit_type,
                "small_unit_type": small_unit_type,
                "message": f"解析成功: 1{big_unit} = {ratio}{small_unit}"
            }
    
    return {
        "success": False,
        "big_unit": "",
        "small_unit": "",
        "ratio": 1.0,
        "big_unit_type": UNIT_UNKNOWN,
        "small_unit_type": UNIT_UNKNOWN,
        "message": f"无法解析spec: {spec}"
    }


def normalize_unit_name(unit: str) -> str:
    """
    标准化单位名称（去除空格等）
    
    Args:
        unit: 单位字符串，如"罐 "、" 瓶"
        
    Returns:
        标准化后的单位名称，如"罐"、"瓶"
    """
    return unit.strip()


def convert_quantity_by_spec(
    quantity: float,
    order_unit: str,
    spec: str
) -> Dict[str, Any]:
    """
    根据SKU的spec规格，判断是否需要单位转换
    
    【核心逻辑】
    1. 先根据spec解析出大单位、小单位、换算比
    2. 再判断出库单填写的单位(order_unit)是什么
    3. 只有当 order_unit = 大单位 时，才需要转换为小单位
    4. 如果 order_unit = 小单位，则不需要转换
    
    Args:
        quantity: 出库单上的数量
        order_unit: 出库单上填写的单位（如"件"、"箱"、"罐"）
        spec: SKU的规格定义（如"规格：12罐/件"）
        
    Returns:
        {
            "success": bool,
            "original_quantity": float,  # 原始数量
            "original_unit": str,        # 原始单位
            "converted_quantity": float,  # 转换后数量
            "converted_unit": str,        # 转换后单位（华鼎标准小单位）
            "need_conversion": bool,      # 是否进行了转换
            "conversion_ratio": float,    # 换算比（仅need_conversion=True时有效）
            "message": str
        }
    """
    # 解析spec
    spec_info = parse_spec_to_unit_info(spec)
    
    # 标准化出库单单位
    order_unit_normalized = normalize_unit_name(order_unit)
    
    # 无法解析spec时，不做转换
    if not spec_info["success"]:
        return {
            "success": True,
            "original_quantity": quantity,
            "original_unit": order_unit_normalized,
            "converted_quantity": quantity,
            "converted_unit": order_unit_normalized,
            "need_conversion": False,
            "conversion_ratio": 1.0,
            "message": f"无法解析spec: {spec}，不进行单位转换"
        }
    
    big_unit = spec_info["big_unit"]
    small_unit = spec_info["small_unit"]
    ratio = spec_info["ratio"]
    
    # 判断出库单单位是否是大单位
    if order_unit_normalized == big_unit:
        # 需要转换：大单位 → 小单位
        return {
            "success": True,
            "original_quantity": quantity,
            "original_unit": order_unit_normalized,
            "converted_quantity": quantity * ratio,
            "converted_unit": small_unit,
            "need_conversion": True,
            "conversion_ratio": ratio,
            "message": f"{quantity}{big_unit} = {quantity * ratio}{small_unit}，已按规格转换为小单位"
        }
    elif order_unit_normalized == small_unit:
        # 不需要转换：已经是小单位
        return {
            "success": True,
            "original_quantity": quantity,
            "original_unit": order_unit_normalized,
            "converted_quantity": quantity,
            "converted_unit": small_unit,
            "need_conversion": False,
            "conversion_ratio": 1.0,
            "message": f"出库单单位已是小单位({small_unit})，无需转换"
        }
    else:
        # 出库单单位既不是大单位也不是小单位，无法判断
        return {
            "success": True,
            "original_quantity": quantity,
            "original_unit": order_unit_normalized,
            "converted_quantity": quantity,
            "converted_unit": order_unit_normalized,
            "need_conversion": False,
            "conversion_ratio": 1.0,
            "message": f"出库单单位({order_unit_normalized})既不是大单位({big_unit})也不是小单位({small_unit})，无法判断是否需要转换"
        }


def convert_quantity_by_spec_with_sku(
    quantity: float,
    order_unit: str,
    spec: str
) -> Dict[str, Any]:
    """
    根据SKU的spec规格进行单位转换的便捷函数
    
    这是convert_quantity_by_spec的别名，保持与原有函数名的兼容性
    """
    return convert_quantity_by_spec(quantity, order_unit, spec)


# ==================== 测试 ====================

if __name__ == "__main__":
    print("=== 单位转换模块测试 ===")
    
    print("\n1. 单位类型识别:")
    test_units = ["箱", "件", "kg", "千克", "升", "ml", "托", "袋", "UNKNOWN_UNIT"]
    for unit in test_units:
        result = get_unit_type(unit)
        print(f"  {unit}: {result['unit_type']} ({result['confidence']:.2f})")
    
    print("\n2. 单位转换:")
    print(f"  1000g → kg: {convert_unit(1000, 'g', 'kg')}")
    print(f"  1kg → g: {convert_unit(1, 'kg', 'g')}")
    print(f"  1吨 → kg: {convert_unit(1, 't', 'kg')}")
    
    print("\n3. 带单位的数量解析:")
    test_inputs = ["10箱", "100kg", "5.5L", "20件"]
    for inp in test_inputs:
        print(f"  {inp}: {parse_quantity_with_unit(inp)}")
    
    print("\n4. 基于SKU规格的单位转换:")
    test_specs = [
        "规格：12罐/件",
        "件：1kg*10袋/件",
        "规格：6.25kg*4/件",
        "件：1L*12盒/件",
        "",
        None
    ]
    for spec in test_specs:
        print(f"\n  spec: {spec}")
        spec_info = parse_spec_to_unit_info(spec)
        print(f"    解析结果: {spec_info}")
        
        # 测试转换场景
        if spec_info["success"]:
            # 场景1：出库单填大单位
            result1 = convert_quantity_by_spec(10, spec_info["big_unit"], spec)
            print(f"    出库单填{result1['original_unit']}→{result1['converted_quantity']}{result1['converted_unit']}: {result1['need_conversion']}")
            
            # 场景2：出库单填小单位
            result2 = convert_quantity_by_spec(120, spec_info["small_unit"], spec)
            print(f"    出库单填{result2['original_unit']}→{result2['converted_quantity']}{result2['converted_unit']}: {result2['need_conversion']}")
