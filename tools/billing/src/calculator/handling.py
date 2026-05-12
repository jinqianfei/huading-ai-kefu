"""
操作费计算模块

计算整件操作费和拆零操作费
"""

# 操作费费率
HANDLING_FEE_STANDARD = 1.5  # 整件操作费（元/标准件）
HANDLING_FEE_SPLIT = 0.3      # 拆零操作费（元/最小单位）


def calc_handling_fee(std_units: float, is_split: bool = False) -> float:
    """
    计算操作费
    
    Args:
        std_units: 标准件数量（浮点数，支持小数）
        is_split: 是否拆零，默认False（整件）
    
    Returns:
        float: 操作费（元）
    """
    if std_units <= 0:
        return 0.0
    
    if is_split:
        # 拆零操作费
        return round(std_units * HANDLING_FEE_SPLIT, 2)
    else:
        # 整件操作费
        return round(std_units * HANDLING_FEE_STANDARD, 2)


def calc_handling_fee_detail(std_units: float, is_split: bool = False) -> dict:
    """
    计算操作费（详细版）
    
    Args:
        std_units: 标准件数量
        is_split: 是否拆零
    
    Returns:
        dict: 包含详细信息的字典
    """
    if std_units <= 0:
        return {
            "fee": 0.0,
            "std_units": std_units,
            "fee_type": "整件" if not is_split else "拆零",
            "unit_price": HANDLING_FEE_STANDARD if not is_split else HANDLING_FEE_SPLIT,
            "calculation": "0 件 × 0 元 = 0 元"
        }
    
    unit_price = HANDLING_FEE_STANDARD if not is_split else HANDLING_FEE_SPLIT
    fee = round(std_units * unit_price, 2)
    fee_type = "整件" if not is_split else "拆零"
    
    return {
        "fee": fee,
        "std_units": std_units,
        "fee_type": fee_type,
        "unit_price": unit_price,
        "calculation": f"{std_units} 件 × {unit_price} 元 = {fee} 元"
    }


def get_handling_fee_rates() -> dict:
    """
    获取操作费费率
    
    Returns:
        dict: 费率信息
    """
    return {
        "standard": {
            "name": "整件操作费",
            "rate": HANDLING_FEE_STANDARD,
            "unit": "元/标准件"
        },
        "split": {
            "name": "拆零操作费",
            "rate": HANDLING_FEE_SPLIT,
            "unit": "元/最小单位"
        }
    }