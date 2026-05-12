# billing_tools/src/calculator/storage.py
# 仓储费计算模块

"""
仓储费计算模块

支持按仓库+温区维度计算存储费
"""

from typing import Dict, Optional


def calc_storage_fee(
    warehouse: str,
    temperature_zone: str,
    tray_count: int,
    storage_rules: Dict
) -> float:
    """
    计算仓储费
    
    支持两种规则结构：
    1. 按温区统一价格：
       {"常温": 1.7, "冷冻": 3.2, "冷藏": 3.2}
    
    2. 按仓库+温区区分：
       {
         "长沙仓": {"常温": 1.7, "冷冻": 3.2, "冷藏": 3.2},
         "成都仓": {"常温": 1.8, "冷冻": 3.5, "冷藏": 3.5},
         ...
       }
    
    Args:
        warehouse: 仓库名称（如"成都仓"）
        temperature_zone: 温区（如"冷冻"、"冷藏"、"常温"）
        tray_count: 托盘数
        storage_rules: 存储费规则
        
    Returns:
        float: 仓储费
    """
    price = get_storage_price(warehouse, temperature_zone, storage_rules)
    return round(tray_count * price, 2)


def get_storage_price(
    warehouse: str,
    temperature_zone: str,
    storage_rules: Dict
) -> float:
    """
    获取存储单价
    
    优先按仓库+温区查询，如果没有则按温区统一查询
    
    Args:
        warehouse: 仓库名称
        temperature_zone: 温区
        storage_rules: 存储费规则
        
    Returns:
        float: 单价，如果未找到返回0
    """
    # 1. 尝试按仓库+温区查找
    if isinstance(storage_rules, dict):
        # 检查是否是仓库维度结构
        if warehouse in storage_rules:
            warehouse_rules = storage_rules[warehouse]
            if isinstance(warehouse_rules, dict):
                if temperature_zone in warehouse_rules:
                    return float(warehouse_rules[temperature_zone] or 0)
        
        # 2. 尝试按温区查找（统一价格）
        if temperature_zone in storage_rules:
            value = storage_rules[temperature_zone]
            if isinstance(value, (int, float)):
                return float(value)
            elif isinstance(value, dict):
                # 统一价格结构
                return float(value.get('price', 0))
    
    return 0


def calc_storage_fee_detail(
    warehouse: str,
    temperature_zone: str,
    tray_count: int,
    storage_rules: Dict
) -> Dict:
    """
    计算仓储费（详细版）
    
    Returns:
        {
            "warehouse": str,
            "temperature_zone": str,
            "tray_count": int,
            "unit_price": float,
            "total_fee": float,
            "rule_source": str  # "warehouse" or "default"
        }
    """
    # 确定使用的规则来源
    rule_source = "default"
    unit_price = 0
    
    if isinstance(storage_rules, dict) and warehouse in storage_rules:
        warehouse_rules = storage_rules[warehouse]
        if isinstance(warehouse_rules, dict) and temperature_zone in warehouse_rules:
            unit_price = float(warehouse_rules[temperature_zone] or 0)
            rule_source = "warehouse"
    
    if unit_price == 0:
        if isinstance(storage_rules, dict) and temperature_zone in storage_rules:
            value = storage_rules[temperature_zone]
            if isinstance(value, (int, float)):
                unit_price = float(value)
            elif isinstance(value, dict):
                unit_price = float(value.get('price', 0))
    
    total_fee = round(tray_count * unit_price, 2)
    
    return {
        "warehouse": warehouse,
        "temperature_zone": temperature_zone,
        "tray_count": tray_count,
        "unit_price": unit_price,
        "total_fee": total_fee,
        "rule_source": rule_source
    }


# 测试
if __name__ == "__main__":
    print("=" * 50)
    print("仓储费计算测试")
    print("=" * 50)
    
    # 测试1：统一价格结构
    rules1 = {"冷冻": 3.2, "冷藏": 3.2, "常温": 1.7}
    print("\n【统一价格】")
    print(f"  成都仓-冷冻-10托: {calc_storage_fee('成都仓', '冷冻', 10, rules1)}元")
    
    # 测试2：按仓库区分结构
    rules2 = {
        "成都仓": {"冷冻": 3.2, "冷藏": 3.2, "常温": 1.7},
        "长沙仓": {"冷冻": 3.0, "冷藏": 3.0, "常温": 1.5},
        "廊坊仓": {"冷冻": 3.5, "冷藏": 3.5, "常温": 2.0}
    }
    print("\n【按仓库区分】")
    print(f"  成都仓-冷冻-10托: {calc_storage_fee('成都仓', '冷冻', 10, rules2)}元")
    print(f"  长沙仓-冷冻-10托: {calc_storage_fee('长沙仓', '冷冻', 10, rules2)}元")
    print(f"  廊坊仓-冷冻-10托: {calc_storage_fee('廊坊仓', '冷冻', 10, rules2)}元")
    
    # 测试3：详细计算
    print("\n【详细计算】")
    detail = calc_storage_fee_detail("长沙仓", "冷冻", 10, rules2)
    print(f"  {detail}")
    
    print("\n" + "=" * 50)
