# -*- coding: utf-8 -*-
"""
数据库连接模块 - 订单映射工具包
"""

import psycopg2
from typing import Optional, List, Dict, Any
from contextlib import contextmanager


# 数据库配置
DB_CONFIG = {
    "host": "localhost",
    "port": 5432,
    "database": "ai_cs_support",
    "user": "postgres",
    "password": "openclaw123"
}


@contextmanager
def get_db_connection():
    """获取数据库连接的上下文管理器"""
    conn = None
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        yield conn
    finally:
        if conn:
            conn.close()


@contextmanager  
def get_db_cursor():
    """获取数据库游标的上下文管理器"""
    with get_db_connection() as conn:
        cursor = conn.cursor()
        try:
            yield cursor
            conn.commit()
        except Exception as e:
            conn.rollback()
            raise e
        finally:
            cursor.close()


# ==================== SKU映射查询 ====================

def get_customer_sku_mapping_by_name(customer_id: str, product_name: str) -> Optional[Dict]:
    """查询客户SKU映射（按商品名称）"""
    sql = """
        SELECT customer_sku_code, customer_sku_name, system_sku_code, system_sku_name, confidence
        FROM customer_sku_mapping
        WHERE customer_id = %s AND customer_sku_name = %s AND status = 'ACTIVE'
        LIMIT 1
    """
    with get_db_cursor() as cursor:
        cursor.execute(sql, (customer_id, product_name))
        row = cursor.fetchone()
        if row:
            return {
                "customer_sku_code": row[0],
                "customer_sku_name": row[1],
                "system_sku_code": row[2],
                "system_sku_name": row[3],
                "confidence": float(row[4]) if row[4] else 0.99
            }
    return None


def get_system_sku_by_product_name_lower(product_name_lower: str) -> List[Dict]:
    """查询系统SKU（按商品名称小写）"""
    sql = """
        SELECT sku_code, sku_name FROM system_sku
        WHERE product_name_lower = %s AND status = 'ACTIVE'
        LIMIT 10
    """
    with get_db_cursor() as cursor:
        cursor.execute(sql, (product_name_lower,))
        return [{"sku_code": r[0], "sku_name": r[1]} for r in cursor.fetchall()]


def get_system_sku_by_keywords(keyword: str) -> List[Dict]:
    """查询系统SKU（按关键词）"""
    sql = """
        SELECT sku_code, sku_name FROM system_sku
        WHERE status = 'ACTIVE' 
        AND (sku_name LIKE %s OR product_keywords LIKE %s)
        LIMIT 20
    """
    pattern = f"%{keyword}%"
    with get_db_cursor() as cursor:
        cursor.execute(sql, (pattern, pattern))
        return [{"sku_code": r[0], "sku_name": r[1]} for r in cursor.fetchall()]


def get_all_system_skus() -> List[Dict]:
    """获取所有系统SKU"""
    sql = "SELECT sku_code, sku_name FROM system_sku WHERE status = 'ACTIVE'"
    with get_db_cursor() as cursor:
        cursor.execute(sql)
        return [{"sku_code": r[0], "sku_name": r[1]} for r in cursor.fetchall()]


# ==================== 门店映射查询 ====================

def get_huading_store_by_code_or_name(customer_id: str, store_code_or_name: str) -> List[Dict]:
    """查询华鼎门店映射（按代码或名称）"""
    sql = """
        SELECT huading_store_code, huading_store_name, customer_store_code, customer_store_name, match_confidence
        FROM huading_store_mapping
        WHERE customer_id = %s 
        AND status = 'ACTIVE'
        AND (customer_store_code = %s OR customer_store_name = %s)
        LIMIT 5
    """
    with get_db_cursor() as cursor:
        cursor.execute(sql, (customer_id, store_code_or_name, store_code_or_name))
        return [{
            "huading_store_code": r[0],
            "huading_store_name": r[1],
            "customer_store_code": r[2],
            "customer_store_name": r[3],
            "match_confidence": float(r[4]) if r[4] else 1.0
        } for r in cursor.fetchall()]


def fuzzy_match_store(customer_id: str, store_name: str) -> List[Dict]:
    """模糊匹配门店"""
    sql = """
        SELECT huading_store_code, huading_store_name, customer_store_code, customer_store_name, match_confidence
        FROM huading_store_mapping
        WHERE customer_id = %s 
        AND status = 'ACTIVE'
        AND customer_store_name LIKE %s
        LIMIT 10
    """
    with get_db_cursor() as cursor:
        cursor.execute(sql, (customer_id, f"%{store_name}%"))
        return [{
            "huading_store_code": r[0],
            "huading_store_name": r[1],
            "customer_store_code": r[2],
            "customer_store_name": r[3],
            "match_confidence": float(r[4]) if r[4] else 0.8
        } for r in cursor.fetchall()]


# ==================== 仓库映射查询 ====================

def get_warehouse_by_store(customer_id: str, huading_store_code: str) -> Optional[Dict]:
    """查询仓库映射（按门店）"""
    sql = """
        SELECT cwm.warehouse_dict_code, w.warehouse_name, cwm.priority
        FROM customer_warehouse_mapping cwm
        JOIN warehouse_dict w ON cwm.warehouse_dict_code = w.warehouse_code
        WHERE cwm.customer_id = %s 
        AND cwm.enabled = true
        ORDER BY cwm.priority ASC
        LIMIT 1
    """
    with get_db_cursor() as cursor:
        # 先尝试按门店查
        cursor.execute(sql, (customer_id,))
        row = cursor.fetchone()
        if row:
            return {
                "warehouse_dict_code": row[0],
                "warehouse_name": row[1],
                "priority": row[2]
            }
    return None


def get_default_warehouse(customer_id: str) -> Optional[str]:
    """获取客户默认仓库"""
    sql = """
        SELECT warehouse_dict_code FROM customer_warehouse_mapping
        WHERE customer_id = %s AND enabled = true
        ORDER BY priority ASC
        LIMIT 1
    """
    with get_db_cursor() as cursor:
        cursor.execute(sql, (customer_id,))
        row = cursor.fetchone()
        if row:
            return row[0]
    return "WH_CHANGSHA"  # 默认值


# ==================== 匹配历史 ====================

def save_match_history(customer_id: str, customer_sku_name: str, 
                      system_sku_code: str, system_sku_name: str,
                      match_type: str, confidence: float) -> bool:
    """保存匹配历史"""
    sql = """
        INSERT INTO match_history (customer_id, customer_sku_name, system_sku_code, system_sku_name, match_type, confidence)
        VALUES (%s, %s, %s, %s, %s, %s)
    """
    try:
        with get_db_cursor() as cursor:
            cursor.execute(sql, (customer_id, customer_sku_name, system_sku_code, system_sku_name, match_type, confidence))
        return True
    except Exception as e:
        print(f"保存匹配历史失败: {e}")
        return False


if __name__ == "__main__":
    print("=== 数据库连接测试 ===")
    
    # 测试SKU映射
    print("\n【1. SKU映射查询】")
    result = get_customer_sku_mapping_by_name("CUST_001", "1000塑杯")
    print(f"  1000塑杯: {result}")
    
    # 测试门店映射
    print("\n【2. 门店映射查询】")
    result = get_huading_store_by_code_or_name("CUST_002", "沧州任丘北辛庄店")
    print(f"  沧州任丘北辛庄店: {result}")
    
    # 测试仓库映射
    print("\n【3. 仓库映射查询】")
    result = get_warehouse_by_store("CUST_001", "CHANGSHA001")
    print(f"  CUST_001-CHANGSHA001: {result}")
    
    default_wh = get_default_warehouse("CUST_001")
    print(f"  CUST_001默认仓库: {default_wh}")
    
    print("\n✅ 数据库连接测试完成！")
