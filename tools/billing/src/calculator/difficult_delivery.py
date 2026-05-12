# billing_tools/src/calculator/difficult_delivery.py
# 困难配送加收计算模块

"""
困难配送加收计算模块

困难配送定义：
- 超市、商场、步行街、临牌市场、mall 等车辆无法进入的区域
- 加收规则：运费 × 20%

参考PRD v1.2 T-6.3
"""

from dataclasses import dataclass
from typing import Tuple


# 困难配送关键词列表（不区分大小写匹配）
difficult_keywords = [
    "超市", "商场", "广场", "步行街", "市场", "mall",
    "超市内", "商场内", "步行区内",
]


@dataclass
class DifficultDeliveryResult:
    """困难配送判定结果"""
    is_difficult: bool        # 是否困难配送
    store_name: str          # 店铺名称
    matched_keyword: str     # 匹配的关键词（若有）
    surcharge_rate: float    # 加收比例
    surcharge_amount: float  # 加收金额


def is_difficult_delivery(store_name: str) -> Tuple[bool, str]:
    """
    判断是否为困难配送

    Args:
        store_name: 门店名称

    Returns:
        Tuple[bool, str]: (是否困难配送, 匹配的关键词)
    """
    if not store_name:
        return False, ""

    name_lower = store_name.lower()
    for keyword in difficult_keywords:
        if keyword.lower() in name_lower:
            return True, keyword

    return False, ""


def calc_difficult_surcharge(base_fee: float, store_name: str) -> float:
    """
    计算困难配送加收

    Args:
        base_fee: 基础运费（元）
        store_name: 门店名称

    Returns:
        float: 加收金额。
               困难配送返回 base_fee × 20%；
               否则返回 0
    """
    difficult, _ = is_difficult_delivery(store_name)
    if not difficult:
        return 0.0

    return round(base_fee * 0.20, 2)


def get_difficult_delivery_info(
    base_fee: float,
    store_name: str
) -> DifficultDeliveryResult:
    """
    获取完整的困难配送信息

    Args:
        base_fee: 基础运费
        store_name: 门店名称

    Returns:
        DifficultDeliveryResult: 包含完整判定和计算结果
    """
    is_difficult, matched_keyword = is_difficult_delivery(store_name)
    surcharge_rate = 0.20 if is_difficult else 0.0
    surcharge_amount = round(base_fee * surcharge_rate, 2) if is_difficult else 0.0

    return DifficultDeliveryResult(
        is_difficult=is_difficult,
        store_name=store_name,
        matched_keyword=matched_keyword,
        surcharge_rate=surcharge_rate,
        surcharge_amount=surcharge_amount,
    )


if __name__ == "__main__":
    print("=== 困难配送判定测试 ===")

    test_cases = [
        "永辉超市",
        "万达广场",
        "成都春熙路步行街",
        "普通便利店",
        "家乐福超市成都高新店",
        "",
    ]

    for name in test_cases:
        difficult, keyword = is_difficult_delivery(name)
        info = get_difficult_delivery_info(100.0, name)
        status = "困难配送" if difficult else "普通配送"
        surcharge = f"加收{info.surcharge_amount}元" if difficult else "无加收"
        print(f"[{name or '(空)'}] {status} {surcharge}")
