# billing_tools/query.py
# 订单查询工具

"""
订单查询工具

用法:
    query_order(customer_id=None, store_name=None, period=None, excel_path=None) -> {success, orders, summary, message}
"""

import sys
import os

# 添加src到路径
_tools_dir = os.path.dirname(os.path.abspath(__file__))
_src_dir = os.path.join(_tools_dir, 'src')
sys.path.insert(0, _src_dir)


def query_order(
    customer_id: str = None,
    store_name: str = None,
    period: str = None,
    excel_path: str = None
) -> dict:
    """
    查询订单明细
    
    Args:
        customer_id: 客户ID（可选）
        store_name: 门店名称（可选，模糊匹配）
        period: 账期（可选）
        excel_path: Excel文件路径（可选）
        
    Returns:
        {
            "success": bool,
            "orders": list,
            "summary": dict,
            "message": str
        }
    """
    try:
        if excel_path and os.path.exists(excel_path):
            # 从Excel读取
            from excel.reader import ExcelReader
            reader = ExcelReader(excel_path)
            orders = reader.read_order_details()
            
            # 筛选
            if store_name:
                orders = [o for o in orders if store_name in o.get("门店名称", "")]
            if period:
                orders = [o for o in orders if period in str(o.get("业务日期", ""))]
        else:
            orders = []
        
        # 汇总
        total_fee = sum(o.get("运费合计", 0) for o in orders)
        total_units = sum(o.get("标准件数量", 0) for o in orders)
        
        summary = {
            "count": len(orders),
            "total_units": round(total_units, 2),
            "total_fee": round(total_fee, 2)
        }
        
        return {
            "success": True,
            "orders": orders,
            "summary": summary,
            "message": f"找到{len(orders)}笔订单"
        }
        
    except Exception as e:
        return {
            "success": False,
            "orders": [],
            "summary": None,
            "message": f"查询失败: {str(e)}"
        }


def query_store_summary(
    customer_id: str,
    store_name: str,
    period: str,
    excel_path: str = None
) -> dict:
    """
    查询某门店的月度汇总
    """
    result = query_order(
        customer_id=customer_id,
        store_name=store_name,
        period=period,
        excel_path=excel_path
    )
    
    if not result["success"]:
        return result
    
    orders = result["orders"]
    if not orders:
        return {
            "success": False,
            "message": f"未找到{store_name}的订单"
        }
    
    total_fee = sum(o.get("运费合计", 0) for o in orders)
    total_units = sum(o.get("标准件数量", 0) for o in orders)
    
    return {
        "success": True,
        "store": store_name,
        "period": period,
        "order_count": len(orders),
        "total_units": round(total_units, 2),
        "total_fee": round(total_fee, 2),
        "details": orders,
        "message": f"{store_name} {period}共{len(orders)}笔，合计{total_fee}元"
    }


if __name__ == "__main__":
    print("订单查询工具")
    print(query_order(store_name="宫小燕", period="2026-02"))
