# billing_tools/src/calculator/value_added.py
# 增值服务费计算模块

"""
增值服务费计算模块

服务类型及单价（来源：规则库）：
- 灌装分箱：0.5 元/件
- 再包装：3.0 元/件
- 贴标：0.3 元/件
- 标签：0.2 元/个

参考PRD v1.2 T-6.2
"""

from dataclasses import dataclass
from typing import Dict, List


# 增值服务费率表（元/件 或 元/个）
value_added_rules: Dict[str, float] = {
    "灌装分箱": 0.5,
    "再包装":  3.0,
    "贴标":    0.3,
    "标签":    0.2,
}


@dataclass
class ValueAddedFeeResult:
    """增值服务费计算结果"""
    service_type: str        # 服务类型
    quantity: int             # 数量
    unit_price: float        # 单价
    total_fee: float         # 费用


def calc_value_added_fee(service_type: str, quantity: int) -> float:
    """
    计算单类增值服务费

    Args:
        service_type: 服务类型，"灌装分箱" | "再包装" | "贴标" | "标签"
        quantity: 数量（件或个）

    Returns:
        float: 费用，元，精确到分
    """
    if quantity <= 0:
        return 0.0

    unit_price = value_added_rules.get(service_type, 0.0)
    if unit_price <= 0:
        return 0.0

    return round(unit_price * quantity, 2)


def calc_all_value_added(order_details: Dict[str, int]) -> Dict[str, float]:
    """
    计算订单所有增值服务费

    Args:
        order_details: {"灌装分箱": 数量, "再包装": 数量, "贴标": 数量, "标签": 数量}
                       只填写需要计费的服务，数量为0或未填写的服务不计费

    Returns:
        dict: {服务类型: 费用}，只包含有费用的项目
    """
    result = {}
    for service_type, quantity in order_details.items():
        fee = calc_value_added_fee(service_type, quantity)
        if fee > 0:
            result[service_type] = fee

    return result


def get_value_added_summary(order_details: Dict[str, int]) -> Dict[str, any]:
    """
    获取增值服务费汇总信息（含明细和合计）

    Args:
        order_details: 同 calc_all_value_added

    Returns:
        dict: 包含明细、合计、公式的完整报告
    """
    details = calc_all_value_added(order_details)
    total = round(sum(details.values()), 2)

    breakdown = []
    for service_type, fee in details.items():
        qty = order_details.get(service_type, 0)
        unit_price = value_added_rules.get(service_type, 0)
        breakdown.append({
            "service_type": service_type,
            "quantity": qty,
            "unit_price": unit_price,
            "fee": fee,
            "formula": f"{qty} × {unit_price} = {fee}"
        })

    return {
        "breakdown": breakdown,
        "total": total,
        "has_charges": total > 0
    }


if __name__ == "__main__":
    print("=== 增值服务费测试 ===")

    # 测试1：单类计费
    fee = calc_value_added_fee("灌装分箱", 100)
    print(f"灌装分箱 100件: {fee}元")

    fee = calc_value_added_fee("再包装", 50)
    print(f"再包装 50件: {fee}元")

    # 测试2：批量计算
    order = {
        "灌装分箱": 200,
        "再包装": 30,
        "贴标": 150,
        "标签": 100,
    }
    summary = get_value_added_summary(order)
    print(f"\n订单增值服务费明细:")
    for item in summary["breakdown"]:
        print(f"  {item['formula']} ({item['service_type']})")
    print(f"合计: {summary['total']}元")
