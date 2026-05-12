# -*- coding: utf-8 -*-
"""
订单映射工具包 - 整合版
将Go工具实现完整翻译为Python，保持逻辑不变

包含：
- SKU映射（4层匹配策略）
- 门店映射（3级匹配策略）
- 仓库映射
- 单位转换
"""

__version__ = "1.0.0"
__all__ = [
    # SKU映射
    "get_system_sku_by_product_name",
    "match_customer_sku",
    "save_sku_match_result",
    "update_sku_mapping_from_feedback",
    # 门店映射
    "get_huading_store_code",
    "get_huading_store_code_with_detail",
    "batch_get_huading_store_codes",
    # 仓库映射
    "get_warehouse_code",
    "get_warehouse_code_with_detail",
    # 单位转换
    "get_unit_type",
    "convert_unit",
    "parse_quantity_with_unit",
    # 辅助函数
    "extract_keywords",
    "extract_core_words",
    "calculate_similarity",
    "calculate_levenshtein_similarity",
]

# 导入子模块
from . import sku_mapping
from . import store_mapping
from . import warehouse_mapping
from . import unit_converter

# SKU映射
get_system_sku_by_product_name = sku_mapping.get_system_sku_by_product_name
match_customer_sku = sku_mapping.match_customer_sku
save_sku_match_result = sku_mapping.save_sku_match_result
update_sku_mapping_from_feedback = sku_mapping.update_sku_mapping_from_feedback

# 门店映射
get_huading_store_code = store_mapping.get_huading_store_code
get_huading_store_code_with_detail = store_mapping.get_huading_store_code_with_detail
batch_get_huading_store_codes = store_mapping.batch_get_huading_store_codes

# 仓库映射
get_warehouse_code = warehouse_mapping.get_warehouse_code
get_warehouse_code_with_detail = warehouse_mapping.get_warehouse_code_with_detail

# 单位转换
get_unit_type = unit_converter.get_unit_type
convert_unit = unit_converter.convert_unit
parse_quantity_with_unit = unit_converter.parse_quantity_with_unit

# 辅助函数
extract_keywords = sku_mapping.extract_keywords
extract_core_words = sku_mapping.extract_core_words
calculate_similarity = sku_mapping.calculate_similarity
calculate_levenshtein_similarity = sku_mapping.calculate_levenshtein_similarity


def order_mapping_tool(tool_name: str, **kwargs) -> dict:
    """
    统一的订单映射工具调用入口
    
    Args:
        tool_name: 工具名
        **kwargs: 工具参数
        
    Returns:
        工具执行结果
        
    Supported tools:
        - match_customer_sku: SKU匹配
        - get_huading_store_code: 门店编码查询
        - get_warehouse_code: 仓库编码查询
        - get_unit_type: 单位类型转换
    """
    tools = {
        "match_customer_sku": match_customer_sku,
        "get_huading_store_code": get_huading_store_code,
        "get_warehouse_code": get_warehouse_code,
        "get_unit_type": get_unit_type,
    }
    
    if tool_name not in tools:
        return {
            "success": False,
            "error": f"Unknown tool: {tool_name}",
            "available_tools": list(tools.keys())
        }
    
    try:
        return tools[tool_name](**kwargs)
    except Exception as e:
        return {
            "success": False,
            "error": str(e),
            "tool": tool_name
        }


# ==================== 测试入口 ====================

if __name__ == "__main__":
    print("=" * 60)
    print("订单映射工具包 - 整合版")
    print("=" * 60)
    
    print("\n【SKU匹配】")
    print(match_customer_sku("1000塑杯"))
    
    print("\n【门店编码查询】")
    print(get_huading_store_code("CUST_001", "沧州任丘北辛庄店"))
    
    print("\n【仓库编码查询】")
    print(get_warehouse_code("CUST_001", "CHANGSHA001"))
    
    print("\n【单位类型转换】")
    print(get_unit_type("箱"))
    print(get_unit_type("kg"))
    
    print("\n" + "=" * 60)
    print("✅ 订单映射工具包测试完成")
    print("=" * 60)

# 导入skill_order_mapping
from .skill_order_mapping import OrderMappingSkill, process_order, OrderMappingResult
from .order_reader import read_order
from .template_generator import generate_huating_template, create_empty_template

# 添加到__all__
__all__.extend([
    "OrderMappingSkill",
    "process_order", 
    "OrderMappingResult",
    "read_order",
    "generate_huating_template",
    "create_empty_template",
])
