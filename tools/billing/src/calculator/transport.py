"""
运输费计算模块

公式：运输费 = MAX(标准件 × 单价, 最低收费)
"""

from typing import Union, Optional
from dataclasses import dataclass


@dataclass
class TransportFeeResult:
    """运输费计算结果"""
    std_units: float          # 标准件数
    unit_price: float         # 单价
    min_fee: float            # 最低收费
    calculated_fee: float     # 计算费用
    used_min_fee: bool         # 是否使用了最低收费
    formula: str              # 使用的公式


def calc_transport_fee(
    std_units: Union[int, float], 
    unit_price: Union[int, float], 
    min_fee: Union[int, float],
    validate: bool = True
) -> TransportFeeResult:
    """
    计算运输费
    
    公式：运输费 = MAX(标准件 × 单价, 最低收费)
    
    Args:
        std_units: 标准件数（实际数量 × 转换比 × 计费系数）
        unit_price: 单价（元/标准件）
        min_fee: 最低收费（元）
        validate: 是否进行参数校验
        
    Returns:
        TransportFeeResult: 计算结果
        
    Raises:
        ValueError: 参数校验失败
        TypeError: 参数类型错误
    """
    # 类型检查
    if not isinstance(std_units, (int, float)):
        raise TypeError(f"std_units必须是数字类型，当前类型: {type(std_units).__name__}")
    if not isinstance(unit_price, (int, float)):
        raise TypeError(f"unit_price必须是数字类型，当前类型: {type(unit_price).__name__}")
    if not isinstance(min_fee, (int, float)):
        raise TypeError(f"min_fee必须是数字类型，当前类型: {type(min_fee).__name__}")
    
    # 参数校验
    if validate:
        if std_units < 0:
            raise ValueError(f"标准件数不能为负数: {std_units}")
        if unit_price < 0:
            raise ValueError(f"单价不能为负数: {unit_price}")
        if min_fee < 0:
            raise ValueError(f"最低收费不能为负数: {min_fee}")
    
    # 计算费用
    calculated_fee = std_units * unit_price
    used_min_fee = calculated_fee < min_fee
    
    # 如果计算费用低于最低收费，使用最低收费
    if used_min_fee:
        final_fee = min_fee
        formula = f"MAX({std_units} × {unit_price}, {min_fee}) = {min_fee} (使用最低收费)"
    else:
        final_fee = calculated_fee
        formula = f"{std_units} × {unit_price} = {final_fee}"
    
    return TransportFeeResult(
        std_units=float(std_units),
        unit_price=float(unit_price),
        min_fee=float(min_fee),
        calculated_fee=round(final_fee, 2),
        used_min_fee=used_min_fee,
        formula=formula
    )


def batch_calc_transport_fee(
    records: list[dict]
) -> list[TransportFeeResult]:
    """
    批量计算运输费
    
    Args:
        records: 记录列表，每条记录包含std_units, unit_price, min_fee
        
    Returns:
        list[TransportFeeResult]: 计算结果列表
        
    Raises:
        ValueError: 参数校验失败
    """
    if not isinstance(records, list):
        raise TypeError(f"records必须是列表类型，当前类型: {type(records).__name__}")
    
    results = []
    for i, record in enumerate(records):
        try:
            result = calc_transport_fee(
                std_units=record["std_units"],
                unit_price=record["unit_price"],
                min_fee=record["min_fee"]
            )
            results.append(result)
        except (KeyError, TypeError, ValueError) as e:
            raise ValueError(f"第{i+1}条记录计算失败: {str(e)}")
    
    return results


# ============================================================
# 以下为兼容旧接口的函数（保留但标记为deprecated）
# ============================================================

import warnings

def calculate_transport_fee(
    std_units: Union[int, float], 
    unit_price: Union[int, float], 
    min_fee: Union[int, float]
) -> float:
    """
    [已废弃] 请使用 calc_transport_fee 函数
    
    计算运输费
    
    Args:
        std_units: 标准件数
        unit_price: 单价
        min_fee: 最低收费
        
    Returns:
        float: 运输费
    """
    warnings.warn(
        "calculate_transport_fee 已废弃，请使用 calc_transport_fee 函数",
        DeprecationWarning,
        stacklevel=2
    )
    result = calc_transport_fee(std_units, unit_price, min_fee)
    return result.calculated_fee


if __name__ == "__main__":
    # 示例用法
    print("=" * 50)
    print("运输费计算示例")
    print("=" * 50)
    
    # 示例1：正常计算
    result1 = calc_transport_fee(std_units=10, unit_price=2.5, min_fee=10)
    print(f"\n示例1：标准件=10, 单价=2.5, 最低收费=10")
    print(f"计算结果：{result1.calculated_fee}元")
    print(f"公式：{result1.formula}")
    print(f"是否使用最低收费：{'是' if result1.used_min_fee else '否'}")
    
    # 示例2：使用最低收费
    result2 = calc_transport_fee(std_units=2, unit_price=2, min_fee=10)
    print(f"\n示例2：标准件=2, 单价=2, 最低收费=10")
    print(f"计算结果：{result2.calculated_fee}元")
    print(f"公式：{result2.formula}")
    print(f"是否使用最低收费：{'是' if result2.used_min_fee else '否'}")
    
    # 示例3：批量计算
    print("\n" + "=" * 50)
    print("批量计算示例")
    print("=" * 50)
    
    records = [
        {"std_units": 10, "unit_price": 2.5, "min_fee": 10},
        {"std_units": 20, "unit_price": 1.5, "min_fee": 15},
        {"std_units": 5, "unit_price": 3.0, "min_fee": 20},
    ]
    
    try:
        results = batch_calc_transport_fee(records)
        for i, r in enumerate(results):
            print(f"记录{i+1}: {r.std_units}件 × {r.unit_price}元 = {r.calculated_fee}元")
    except ValueError as e:
        print(f"批量计算失败: {e}")
