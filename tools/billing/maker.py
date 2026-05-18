# billing_tools/maker.py
# 账单制作工具 - 完整版 v2.3

"""
账单制作工具 - 完整版 v2.3

支持所有费用类型（PRD v2.3）：
- 存储费
- 操作费（整件/拆零）
- 增值服务费
- 运输费（5仓库×多目的地，含困难配送加收）
- 调拨费（无最低收费）
- 送货进店费
- 超标件系数

标准件计算规则（v2.3）：
1. 从商品信息表（calc_sku_coefficients计算后）获取SKU计费系数
2. 固定SKU（合同fixed_skus表中）系数固定为1.0
3. 其他SKU按合同超规格规则: max(weight_kg/max_weight, volume_m3/max_volume) × 件数
4. 向上取整

重要更新（v2.3）：
- 费用计算使用的计费系数来自商品信息表（calc_sku_coefficients计算后保存到规则库）
- 不依赖出库明细表中的任何预计算系数
"""

import sys
import os
import json

# 添加src到路径
_tools_dir = os.path.dirname(os.path.abspath(__file__))
_src_dir = os.path.join(_tools_dir, 'src')
_knowledge_dir = os.path.expanduser("~/openclaw-workspaces/ai-kefu/knowledge")
if _src_dir not in sys.path:
    sys.path.insert(0, _src_dir)

from calculator.transport import calc_transport_fee
from calculator.transfer import calc_transfer_fee
from calculator.storage import calc_storage_fee, get_storage_price
from .calculator.sku_coefficient import compute_oversized_coefficient

def make_bill(customer_id: str, period: str, excel_path: str = None) -> dict:
    """
    根据出库数据和规则生成完整账单
    
    支持费用类型：
    - 存储费
    - 操作费（整件1.5元/拆零0.3元）
    - 增值服务费
    - 运输费（5仓库×多目的地）
    - 困难配送加收20%
    - 调拨费（无最低收费）
    - 送货进店费1元/杯
    
    Args:
        customer_id: 客户ID
        period: 账期（如 "2026-02"）
        excel_path: Excel文件路径
        
    Returns:
        {
            "success": bool,
            "bill": dict,
            "summary": dict,
            "message": str
        }
    """
    try:
        global _customer_id_for_products
        _customer_id_for_products = customer_id
        _product_table_cache = None  # 重置缓存
        
        # 1. 加载规则
        rules = _load_rules(customer_id)
        if not rules:
            return {
                "success": False,
                "bill": None,
                "summary": None,
                "message": f"未找到客户规则（ID: {customer_id}）"
            }
        
        # 2. 如果没有Excel，返回规则确认
        if not excel_path or not os.path.exists(excel_path):
            return {
                "success": True,
                "bill": None,
                "summary": _get_rule_summary(rules),
                "message": "规则已加载，请上传Excel文件进行账单制作"
            }
        
        # 3. 读取Excel并计算
        try:
            from excel.reader import ExcelReader
            
            reader = ExcelReader(excel_path)
            
            # 读取运输费用明细
            _, transport_data = reader.read_sheet('运输费用计算明细')
            
            # 读取仓储费
            _, storage_data = reader.read_sheet('仓储透视表')
            
            reader.close()
            
            # 4. 计算各项费用
            fees = _calculate_all_fees(transport_data, storage_data, rules)
            
            return {
                "success": True,
                "bill": {
                    "customer_id": customer_id,
                    "period": period,
                    "customer_name": rules.get("customer_name"),
                    "transport_data": transport_data[:5],  # 前5条预览
                    "storage_data": storage_data[:5],
                    "fees": fees
                },
                "summary": fees,
                "message": "账单生成成功"
            }
            
        except Exception as e:
            return {
                "success": False,
                "bill": None,
                "summary": None,
                "message": f"读取Excel失败: {str(e)}"
            }
        
    except Exception as e:
        return {
            "success": False,
            "bill": None,
            "summary": None,
            "message": f"生成失败: {str(e)}"
        }


def _load_rules(customer_id: str) -> dict:
    """加载客户规则"""
    rules_path = os.path.join(_knowledge_dir, "rules")
    
    for f in os.listdir(rules_path):
        if customer_id in f and f.endswith('.json'):
            with open(os.path.join(rules_path, f), 'r', encoding='utf-8') as fp:
                return json.load(fp)
    return None


def _calculate_all_fees(transport_data: list, storage_data: list, rules: dict) -> dict:
    """计算所有费用"""
    all_rules = rules.get('rules', {})
    transport_rules = all_rules.get('transport', {})
    handling_rules = all_rules.get('handling', {})
    value_added_rules = all_rules.get('value_added', {})
    special_rules = all_rules.get('special', {})
    oversized_rules = all_rules.get('oversized', {})
    
    # 获取标准件计算规则
    standard_unit_rules = all_rules.get('standard_unit', {})
    max_weight_kg = standard_unit_rules.get('max_weight_kg', 15)
    max_volume_m3 = standard_unit_rules.get('max_volume_m3', 0.05)
    
    # 获取产品固定系数
    product_coefficients = all_rules.get('product_coefficients', {})
    products_map = product_coefficients.get('products', {})
    
    # 费用汇总
    total_transport = 0      # 运输费
    total_handling = 0       # 操作费
    total_value_added = 0    # 增值服务费
    total_difficult = 0       # 困难配送加收
    total_transfer = 0        # 调拨费
    total_delivery = 0       # 送货进店费
    
    # 遍历订单计算
    for order in transport_data:
        # 获取商品标识和数量
        sku_code = order.get('SKU编码') or order.get('商品三方sku编码') or ''
        product_name = order.get('商品名称', '')
        weight_kg = float(order.get('重量(kg)') or 0)
        volume_m3 = float(order.get('体积(m³)') or 0)
        qty = float(order.get('实际出库数量') or order.get('计划出库数量') or 1)
        
        # 计算标准件数（使用商品信息表的计费系数）
        std_units = _calc_standard_units(
            sku_code, product_name, weight_kg, volume_m3, qty,
            products_map, max_weight_kg, max_volume_m3,
            oversized_rules=oversized_rules,
            product_coefficient=_get_product_coefficient_from_table(sku_code)
        )
        
        # 获取费用字段
        shared_fee = float(order.get('共配段运费') or 0)
        transfer_fee = float(order.get('调拨段运费') or 0)
        store_name = order.get('门店名称') or order.get('单位名称', '')
        
        # 1. 运输费 = 共配段 + 调拨段
        total_transport += shared_fee + transfer_fee
        
        # 2. 操作费（整件）
        if handling_rules.get('整件'):
            total_handling += std_units * handling_rules['整件']
        
        # 3. 困难配送加收
        if special_rules.get('困难配送加收'):
            # 简单判断：门店名称含特定关键词
            difficult_keywords = ['超市', '商场', '购物中心', '步行街', '市场']
            if any(kw in store_name for kw in difficult_keywords):
                total_difficult += shared_fee * special_rules['困难配送加收']
        
        # 4. 送货进店费
        if special_rules.get('送货进店费'):
            total_delivery += std_units * special_rules['送货进店费']
        
        # 5. 调拨费已包含在transport_fee中
    
    # 6. 存储费（按仓库+温区计算）
    storage_rules = all_rules.get('storage', {})
    total_storage = 0
    total_outbound = 0
    
    for s in storage_data:
        warehouse = s.get('配送中心', s.get('仓库', ''))
        temp_zone = s.get('存储方式', '')
        fee_type = s.get('费用项', '')
        tray_count = float(s.get('结算托盘数') or s.get('托盘数') or 0)
        
        if '存储费' in fee_type:
            # 使用calc_storage_fee按仓库+温区计算
            amount = calc_storage_fee(warehouse, temp_zone, int(tray_count), storage_rules)
            total_storage += amount
        elif '出库' in fee_type or '统配' in fee_type:
            # 出库费按托盘数计算
            outbound_price = s.get('单价', 0)
            total_outbound += tray_count * float(outbound_price)
    
    # 汇总
    return {
        "transport_fee": round(total_transport, 2),
        "handling_fee": round(total_handling, 2),
        "value_added_fee": round(total_value_added, 2),
        "difficult_surcharge": round(total_difficult, 2),
        "transfer_fee": round(total_transfer, 2),
        "delivery_fee": round(total_delivery, 2),
        "storage_fee": round(total_storage, 2),
        "outbound_fee": round(total_outbound, 2),
        "total": round(
            total_transport + total_handling + total_value_added + 
            total_difficult + total_transfer + total_delivery + 
            total_storage + total_outbound, 
            2
        )
    }


def _calc_standard_units(sku_code: str, product_name: str, weight_kg: float, 
                          volume_m3: float, qty: float,
                          products_map: dict, max_weight_kg: float, 
                          max_volume_m3: float,
                          oversized_rules: dict = None,
                          product_coefficient: float = None) -> float:
    """
    计算标准件数
    
    优先级（v2.3）：
    1. 若有商品信息表的计费系数（calc_sku_coefficients计算后），直接使用
    2. 若SKU/品名在合同fixed_skus表中，系数固定为1.0
    3. 否则按合同超规格规则计算（使用 compute_oversized_coefficient）
    
    Args:
        sku_code: SKU编码
        product_name: 商品名称
        weight_kg: 重量(kg)
        volume_m3: 体积(m³)
        qty: 件数
        products_map: 合同fixed_skus表
        max_weight_kg: 标准件最大重量
        max_volume_m3: 标准件最大体积
        oversized_rules: 超规格规则（从合同提取，v2.1+格式）
        product_coefficient: 商品信息表中的计费系数（calc_sku_coefficients计算后）
    
    Returns:
        标准件数（向上取整）
    """
    import math
    
    # 1. 优先使用商品信息表的计费系数（calc_sku_coefficients计算后）
    if product_coefficient is not None:
        return math.ceil(qty * product_coefficient)
    
    # 2. 检查合同fixed_skus表（固定SKU系数为1.0）
    if products_map:
        # 按SKU匹配
        for prod_sku, prod_info in products_map.items():
            if sku_code and prod_sku in sku_code:
                # 固定SKU系数为1.0
                return math.ceil(qty * 1.0)
        
        # 按品名匹配
        for prod_sku, prod_info in products_map.items():
            prod_name = prod_info.get('name', '')
            if prod_name and prod_name in product_name:
                # 固定SKU系数为1.0
                return math.ceil(qty * 1.0)
    
    # 3. 按合同超规格规则计算（使用 compute_oversized_coefficient）
    if weight_kg > 0 or volume_m3 > 0:
        # 使用 compute_oversized_coefficient 从 oversized_rules 读取阈值
        coef = compute_oversized_coefficient(
            weight_kg, volume_m3, max_weight_kg, max_volume_m3, oversized_rules
        )
        std_units = coef * qty
    else:
        # 无重量体积数据，默认1:1
        std_units = qty
    
    return math.ceil(std_units)


def _get_product_coefficient_from_table(sku_code: str) -> float:
    """
    从商品信息表获取SKU的计费系数
    
    商品信息表存储在规则库中，通过load_rule加载
    格式：{customer_id}_products.json
    
    Args:
        sku_code: SKU编码
        
    Returns:
        计费系数（float），若未找到返回None
    """
    import os
    global _product_table_cache
    
    if not sku_code:
        return None
    
    # 如果缓存为空，加载商品信息表
    if _product_table_cache is None:
        _product_table_cache = _load_product_table()
    
    if _product_table_cache is None:
        return None
    
    # 从缓存中查找SKU系数
    # 商品信息表格式：{商品编码: {name, volume_m3, weight_kg, coefficient}}
    for product_sku, product_info in _product_table_cache.items():
        if sku_code and product_sku in sku_code:
            return product_info.get('coefficient')
    
    return None


def _load_product_table() -> dict:
    """
    加载商品信息表
    
    商品信息表路径：{knowledge_dir}/products/{customer_id}_products.json
    由calc_sku_coefficients生成
    
    Returns:
        商品信息字典 {商品编码: {coefficient, ...}}
    """
    global _customer_id_for_products
    
    if _customer_id_for_products is None:
        return None
    
    products_dir = os.path.join(_knowledge_dir, "products")
    product_file = os.path.join(products_dir, f"{_customer_id_for_products}_products.json")
    
    if not os.path.exists(product_file):
        return None
    
    try:
        with open(product_file, 'r', encoding='utf-8') as f:
            return json.load(f)
    except:
        return None


# 全局缓存
_product_table_cache = None
_customer_id_for_products = None


def _get_rule_summary(rules: dict) -> dict:
    """获取规则摘要"""
    if not rules:
        return {}
    
    r = rules.get('rules', {})
    
    return {
        "customer_name": rules.get("customer_name"),
        "contract_no": rules.get("contract_no"),
        "storage": r.get("storage"),
        "handling": r.get("handling"),
        "value_added": r.get("value_added"),
        "transport_warehouses": list(r.get("transport", {}).keys()),
        "oversized": r.get("oversized"),
        "special": r.get("special")
    }


def format_bill_summary(summary: dict) -> str:
    """格式化账单汇总为可读字符串"""
    if not summary:
        return "无账单数据"
    
    lines = [
        "**账单汇总**",
        f"├─ 运输费：{summary.get('transport_fee', 0):,.2f}元",
    ]
    
    if summary.get('handling_fee', 0) > 0:
        lines.append(f"├─ 操作费：{summary.get('handling_fee', 0):,.2f}元")
    
    if summary.get('value_added_fee', 0) > 0:
        lines.append(f"├─ 增值服务费：{summary.get('value_added_fee', 0):,.2f}元")
    
    if summary.get('difficult_surcharge', 0) > 0:
        lines.append(f"├─ 困难配送加收：{summary.get('difficult_surcharge', 0):,.2f}元")
    
    if summary.get('transfer_fee', 0) > 0:
        lines.append(f"├─ 调拨费：{summary.get('transfer_fee', 0):,.2f}元")
    
    if summary.get('delivery_fee', 0) > 0:
        lines.append(f"├─ 送货进店费：{summary.get('delivery_fee', 0):,.2f}元")
    
    lines.append(f"├─ 存储费：{summary.get('storage_fee', 0):,.2f}元")
    lines.append(f"├─ 出库费：{summary.get('outbound_fee', 0):,.2f}元")
    lines.append(f"└─ **合计：{summary.get('total', 0):,.2f}元**")
    
    return "\n".join(lines)


if __name__ == "__main__":
    print("账单制作工具 v2.0")
    print(_get_rule_summary({"customer_name": "测试"}))
