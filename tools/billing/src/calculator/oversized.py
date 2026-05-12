"""
超标件系数计算模块

根据重量和体积计算超标件计费系数
"""

def calc_oversized_coefficient(weight_kg: float, volume_m3: float) -> float:
    """
    计算超标件系数
    
    超标件计费规则:
    - 普通件: 重量 ≤ 15kg 且 体积 ≤ 0.05m³ → 系数1.0
    - 15kg < 重量 ≤ 25kg 或 0.05m³ < 体积 ≤ 0.08m³ → 系数1.2
    - 25kg < 重量 ≤ 50kg 或 0.08m³ < 体积 ≤ 0.1m³ → 系数2.0
    - 重量 > 50kg 或 体积 > 0.1m³ → 按[重量(kg)/15]或[体积(m³)/0.05]取较大值
    
    Args:
        weight_kg: 重量（千克）
        volume_m3: 体积（立方米）
    
    Returns:
        float: 超标件系数
    """
    # 普通件判断
    if weight_kg <= 15 and volume_m3 <= 0.05:
        return 1.0
    
    # 轻微超标: 15kg < 重量 ≤ 25kg 或 0.05m³ < 体积 ≤ 0.08m³
    weight_tier1 = 15 < weight_kg <= 25
    volume_tier1 = 0.05 < volume_m3 <= 0.08
    if weight_tier1 or volume_tier1:
        return 1.2
    
    # 中度超标: 25kg < 重量 ≤ 50kg 或 0.08m³ < 体积 ≤ 0.1m³
    weight_tier2 = 25 < weight_kg <= 50
    volume_tier2 = 0.08 < volume_m3 <= 0.1
    if weight_tier2 or volume_tier2:
        return 2.0
    
    # 重度超标: 重量 > 50kg 或 体积 > 0.1m³
    weight_tier3 = weight_kg > 50
    volume_tier3 = volume_m3 > 0.1
    if weight_tier3 or volume_tier3:
        weight_coefficient = weight_kg / 15
        volume_coefficient = volume_m3 / 0.05
        return max(weight_coefficient, volume_coefficient)
    
    # 默认返回1.0（普通件）
    return 1.0


def get_oversized_tier(weight_kg: float, volume_m3: float) -> str:
    """
    获取超标件等级描述
    
    Args:
        weight_kg: 重量（千克）
        volume_m3: 体积（立方米）
    
    Returns:
        str: 超标件等级
    """
    if weight_kg <= 15 and volume_m3 <= 0.05:
        return "普通件"
    elif (15 < weight_kg <= 25) or (0.05 < volume_m3 <= 0.08):
        return "轻微超标件"
    elif (25 < weight_kg <= 50) or (0.08 < volume_m3 <= 0.1):
        return "中度超标件"
    else:
        return "重度超标件"