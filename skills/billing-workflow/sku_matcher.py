"""
SKU匹配模块 - 6层匹配策略

为billing-workflow提供SKU匹配功能
"""

import re
from typing import Dict, List, Tuple, Optional


class SKUMatcher:
    """
    SKU匹配器 - 支持6层匹配策略
    """
    
    VERSION = "1.0"
    
    def __init__(self, db_config: Dict):
        """
        初始化SKU匹配器
        
        Args:
            db_config: 数据库连接配置
                {
                    "host": "localhost",
                    "port": 5432,
                    "database": "ai_cs_support",
                    "user": "xxx",
                    "password": "xxx"  # 可选
                }
        """
        self.db_config = db_config
    
    def match(self, product_name: str, shipper_id: str) -> Optional[Dict]:
        """
        匹配SKU
        
        Args:
            product_name: 商品名称（来自客户订单）
            shipper_id: 货主ID
        
        Returns:
            {
                "sku_code": "SK230904000009",
                "match_method": "前缀匹配",
                "confidence": 0.9
            }
            或 None（未匹配）
        """
        conn = None
        try:
            import psycopg2
            conn = psycopg2.connect(**self.db_config)
            cur = conn.cursor()
            
            original_name = product_name
            result = None
            
            # ========== 1. 精确匹配 ==========
            result = self._exact_match(cur, original_name, shipper_id)
            if result:
                result["match_method"] = "精确匹配"
                result["confidence"] = 0.99
            
            # ========== 2. 清理后匹配 ==========
            if not result:
                clean_name = self._clean_product_name(original_name)
                if clean_name != original_name:
                    result = self._exact_match(cur, clean_name, shipper_id)
                    if result:
                        result["match_method"] = f"清理匹配"
                        result["confidence"] = 0.95
            
            # ========== 3. 前缀匹配 ==========
            if not result:
                clean_name = self._clean_product_name(original_name)
                result = self._prefix_match(cur, clean_name, shipper_id)
                if result:
                    result["match_method"] = "前缀匹配"
                    result["confidence"] = 0.90
            
            # ========== 4. 后缀匹配 ==========
            if not result:
                clean_name = self._clean_product_name(original_name)
                result = self._suffix_match(cur, clean_name, shipper_id)
                if result:
                    result["match_method"] = "后缀匹配"
                    result["confidence"] = 0.85
            
            # ========== 5. 包含匹配 ==========
            if not result:
                clean_name = self._clean_product_name(original_name)
                result = self._include_match(cur, clean_name, shipper_id)
                if result:
                    result["match_method"] = "包含匹配"
                    result["confidence"] = 0.80
            
            # ========== 6. 反向包含 ==========
            if not result:
                clean_name = self._clean_product_name(original_name)
                result = self._reverse_include_match(cur, clean_name, shipper_id)
                if result:
                    result["match_method"] = "反向包含"
                    result["confidence"] = 0.75
            
            return result
            
        except Exception as e:
            print(f"SKU匹配错误: {e}")
            return None
        finally:
            if conn:
                conn.close()
    
    def _clean_product_name(self, name: str) -> str:
        """
        清理商品名称，去除特殊字符
        """
        # 去除首尾的特殊字符
        name = name.strip().rstrip('-').rstrip('/').rstrip('.')
        # 去除括号及其内容
        name = re.sub(r'[（(].*[)）]', '', name)
        # 去除空格
        name = name.strip()
        return name
    
    def _exact_match(self, cur, product_name: str, shipper_id: str) -> Optional[Dict]:
        """精确匹配"""
        cur.execute("""
            SELECT system_sku_code,
                   (unit_conversion_rule->>'ratio')::numeric as ratio
            FROM shipper_sku_mapping 
            WHERE customer_sku_name = %s 
            AND shipper_id = %s
            ORDER BY ratio DESC
            LIMIT 1
        """, (product_name, shipper_id))
        
        row = cur.fetchone()
        if row:
            return {
                "sku_code": row[0],
                "original_name": product_name
            }
        return None
    
    def _prefix_match(self, cur, product_name: str, shipper_id: str) -> Optional[Dict]:
        """前缀匹配"""
        cur.execute("""
            SELECT system_sku_code,
                   (unit_conversion_rule->>'ratio')::numeric as ratio,
                   customer_sku_name
            FROM shipper_sku_mapping 
            WHERE customer_sku_name LIKE %s
            AND shipper_id = %s
            AND customer_sku_name != %s
            ORDER BY ratio DESC
            LIMIT 1
        """, (f"{product_name}%", shipper_id, product_name))
        
        row = cur.fetchone()
        if row:
            return {
                "sku_code": row[0],
                "original_name": row[2],  # 返回实际匹配到的名称
                "matched_name": product_name
            }
        return None
    
    def _suffix_match(self, cur, product_name: str, shipper_id: str) -> Optional[Dict]:
        """后缀匹配"""
        cur.execute("""
            SELECT system_sku_code,
                   (unit_conversion_rule->>'ratio')::numeric as ratio,
                   customer_sku_name
            FROM shipper_sku_mapping 
            WHERE customer_sku_name LIKE %s
            AND shipper_id = %s
            AND customer_sku_name != %s
            ORDER BY ratio DESC
            LIMIT 1
        """, (f"%{product_name}", shipper_id, product_name))
        
        row = cur.fetchone()
        if row:
            return {
                "sku_code": row[0],
                "original_name": row[2],
                "matched_name": product_name
            }
        return None
    
    def _include_match(self, cur, product_name: str, shipper_id: str) -> Optional[Dict]:
        """包含匹配（订单名称包含系统名称）"""
        # 先获取所有候选
        cur.execute("""
            SELECT system_sku_code,
                   (unit_conversion_rule->>'ratio')::numeric as ratio,
                   customer_sku_name
            FROM shipper_sku_mapping 
            WHERE shipper_id = %s
            ORDER BY ratio DESC
        """, (shipper_id,))
        
        rows = cur.fetchall()
        for row in rows:
            if row[2] in product_name:
                return {
                    "sku_code": row[0],
                    "original_name": row[2],
                    "matched_name": product_name
                }
        return None
    
    def _reverse_include_match(self, cur, product_name: str, shipper_id: str) -> Optional[Dict]:
        """反向包含（系统名称包含订单名称）"""
        # 先获取所有候选
        cur.execute("""
            SELECT system_sku_code,
                   (unit_conversion_rule->>'ratio')::numeric as ratio,
                   customer_sku_name
            FROM shipper_sku_mapping 
            WHERE shipper_id = %s
            ORDER BY ratio DESC
        """, (shipper_id,))
        
        rows = cur.fetchall()
        for row in rows:
            if product_name in row[2]:
                return {
                    "sku_code": row[0],
                    "original_name": row[2],
                    "matched_name": product_name
                }
        return None
    
    def batch_match(self, items: List[Dict], shipper_id: str) -> Tuple[List[Dict], List[Dict]]:
        """
        批量匹配SKU
        
        Args:
            items: 商品列表，每项包含 product_name
            shipper_id: 货主ID
        
        Returns:
            (matched_items, unmatched_items)
        """
        matched = []
        unmatched = []
        
        for item in items:
            result = self.match(item["product_name"], shipper_id)
            
            if result:
                matched.append({
                    **item,
                    "sku_code": result["sku_code"],
                    "match_method": result["match_method"],
                    "confidence": result["confidence"]
                })
            else:
                unmatched.append(item)
        
        return matched, unmatched


# 测试
if __name__ == "__main__":
    matcher = SKUMatcher({
        "host": "localhost",
        "port": 5432,
        "database": "ai_cs_support",
        "user": "jinqianfei"
    })
    
    # 测试
    test_names = ["果糖-", "果糖/", "柠檬（常温）", "不存在的商品"]
    
    print("=== SKU匹配测试 ===\n")
    for name in test_names:
        result = matcher.match(name, "HZ2023061500002")
        if result:
            print(f"'{name}' → {result['sku_code']} ({result['match_method']})")
        else:
            print(f"'{name}' → 未匹配")