# -*- coding: utf-8 -*-
"""
仓库映射模块 - 从Go工具翻译为Python（已连接数据库）
根据客户ID和华鼎门店编码查询对应的仓库字典编码
"""

from typing import Dict, Optional, Any
from dataclasses import dataclass

# 导入数据库模块
from . import db


@dataclass
class WarehouseMappingResult:
    """仓库映射结果"""
    warehouse_dict_code: str
    warehouse_dict_name: str = ""
    warehouse_code: str = ""
    warehouse_name: str = ""
    customer_id: str = ""
    huading_store_code: str = ""
    match_type: str = "NOT_FOUND"
    confidence: float = 0.0
    priority: int = 0
    message: str = ""


def get_warehouse_code(customer_id: str, huading_store_code: str) -> Dict[str, Any]:
    """
    获取仓库字典编码（Tool接口）
    """
    if not customer_id:
        return {"success": False, "error": "customer_id is required"}
    
    if not huading_store_code:
        return {"success": False, "error": "huading_store_code is required"}
    
    customer_id = customer_id.strip()
    huading_store_code = huading_store_code.strip()
    
    # 调用数据库查询
    warehouse = db.get_warehouse_by_store(customer_id, huading_store_code)
    
    if warehouse:
        return {
            "success": True,
            "warehouse_code": warehouse['warehouse_dict_code'],
            "warehouse_name": warehouse['warehouse_name'],
            "match_type": "DIRECT",
            "confidence": 1.0,
            "priority": warehouse.get('priority', 1),
            "message": "仓库映射成功"
        }
    
    # 尝试获取客户默认仓库
    default_code = db.get_default_warehouse(customer_id)
    if default_code:
        return {
            "success": True,
            "warehouse_code": default_code,
            "warehouse_name": "",
            "match_type": "DEFAULT",
            "confidence": 0.8,
            "message": "使用客户默认仓库"
        }
    
    return {
        "success": False,
        "error": f"无法找到仓库映射: customer_id={customer_id}, huading_store_code={huading_store_code}",
        "warehouse_code": None
    }


def get_warehouse_code_with_detail(customer_id: str, huading_store_code: str) -> Dict[str, Any]:
    """获取仓库映射（带详细结果）"""
    result = get_warehouse_code(customer_id, huading_store_code)
    
    if result['success']:
        return {
            "warehouse_dict_code": result['warehouse_code'],
            "warehouse_dict_name": result.get('warehouse_name', ''),
            "match_type": result['match_type'],
            "confidence": result['confidence'],
            "message": result['message']
        }
    
    return {
        "warehouse_dict_code": None,
        "match_type": "NOT_FOUND",
        "confidence": 0.0,
        "message": f"无法找到仓库映射: customer_id={customer_id}, huading_store_code={huading_store_code}"
    }


# ==================== 测试 ====================

if __name__ == "__main__":
    print("=== 仓库映射模块测试（已连接数据库）===")
    
    print("\n1. 仓库编码查询:")
    test_cases = [
        ("CUST_001", "CHANGSHA001"),
        ("CUST_001", "CHENGDU001"),
        ("CUST_002", "CANGZHOU001"),
    ]
    for customer_id, store_code in test_cases:
        result = get_warehouse_code(customer_id, store_code)
        status = "✅" if result['success'] else "❌"
        print(f"  {customer_id}-{store_code}: {result.get('warehouse_code', 'N/A')} ({result.get('match_type', 'N/A')})")
