# billing_tools/src/calculator/transfer.py
# 调拨段运费计算

"""
调拨段运费计算模块

调拨出库计费规则：
- 无最低收费
- 调拨段单独计算
- 标准件数量 × 调拨单价

参考PRD 4.2.3：
- 调拨单价：从合同规则中获取
- 调拨段运费 = 标准件数量 × 调拨单价
"""

from dataclasses import dataclass
from typing import Optional


@dataclass
class TransferFeeResult:
    """调拨费计算结果"""
    std_units: float          # 标准件数量
    transfer_price: float     # 调拨单价
    calculated_fee: float     # 计算出的运费
    formula: str             # 计算公式


def calc_transfer_fee(
    std_units: float,
    transfer_price: float
) -> TransferFeeResult:
    """
    计算调拨段运费
    
    调拨出库无最低收费，直接按 标准件数量 × 调拨单价 计算
    
    Args:
        std_units: 标准件数量
        transfer_price: 调拨单价（元/标准件）
        
    Returns:
        TransferFeeResult: 包含计算结果的 dataclass
    """
    if std_units <= 0 or transfer_price <= 0:
        return TransferFeeResult(
            std_units=std_units,
            transfer_price=transfer_price,
            calculated_fee=0.0,
            formula="无调拨"
        )
    
    calculated_fee = std_units * transfer_price
    
    return TransferFeeResult(
        std_units=std_units,
        transfer_price=transfer_price,
        calculated_fee=round(calculated_fee, 2),
        formula=f"{std_units} × {transfer_price} = {calculated_fee}"
    )


def calc_combined_transport_fee(
    std_units: float,
    shared_price: float,
    shared_min_fee: float,
    shared_fee: float,
    transfer_price: float,
    transfer_fee: float
) -> dict:
    """
    计算综合运输费（共配段 + 调拨段）
    
    用于订单同时包含共配和调拨两段的情况
    
    Args:
        std_units: 标准件数量
        shared_price: 共配单价
        shared_min_fee: 共配最低收费
        shared_fee: 共配段计算费用
        transfer_price: 调拨单价
        transfer_fee: 调拨段费用
        
    Returns:
        dict: 综合运输费结果
    """
    # 共配段费用（有最低收费限制）
    if shared_fee < shared_min_fee:
        actual_shared_fee = shared_min_fee
        used_min_fee = True
    else:
        actual_shared_fee = shared_fee
        used_min_fee = False
    
    # 调拨段费用（无最低收费）
    actual_transfer_fee = transfer_fee if transfer_fee > 0 else 0
    
    total_fee = actual_shared_fee + actual_transfer_fee
    
    return {
        "std_units": std_units,
        "shared_fee": actual_shared_fee,
        "used_min_fee": used_min_fee,
        "transfer_fee": actual_transfer_fee,
        "total_fee": round(total_fee, 2),
        "detail": {
            "共配段": actual_shared_fee,
            "调拨段": actual_transfer_fee
        }
    }


if __name__ == "__main__":
    # 测试
    print("=== 调拨费计算测试 ===")
    
    # 测试1：普通调拨
    result = calc_transfer_fee(50, 0.15)
    print(f"50件 × 0.15元 = {result.calculated_fee}元")
    print(f"公式: {result.formula}")
    
    # 测试2：零值
    result = calc_transfer_fee(0, 0.15)
    print(f"0件 × 0.15元 = {result.calculated_fee}元")
    
    print("\n=== 综合运输费测试 ===")
    result = calc_combined_transport_fee(
        std_units=50,
        shared_price=15,
        shared_min_fee=260,
        shared_fee=750,
        transfer_price=0.15,
        transfer_fee=7.5
    )
    print(f"共配段: {result['shared_fee']}元 (最低{result['used_min_fee']})")
    print(f"调拨段: {result['transfer_fee']}元")
    print(f"合计: {result['total_fee']}元")
