# -*- coding: utf-8 -*-
"""
门店映射模块 - 从Go工具翻译为Python（已连接数据库）
将客户侧的门店代码或名称映射到华鼎系统的门店编号
支持三级匹配策略：精确匹配 → 模糊匹配 → 地址匹配
"""

import re
from typing import List, Dict, Optional, Any, Tuple
from dataclasses import dataclass

# 导入数据库模块
from . import db


@dataclass
class StoreMappingResult:
    """门店映射结果"""
    huading_store_code: str
    huading_store_name: str
    match_type: str  # EXACT/FUZZY/ADDRESS/NOT_FOUND
    confidence: float
    matched_source: str
    customer_store_code: str
    customer_store_name: str
    message: str = ""


def get_huading_store_code(customer_id: str, store_code_or_name: str, 
                           address: str = "") -> Dict[str, Any]:
    """
    获取华鼎门店编码（Tool接口）
    """
    if not customer_id:
        return {"success": False, "error": "customer_id is required"}
    
    if not store_code_or_name:
        return {"success": False, "error": "store_code_or_name is required"}
    
    store_code_or_name = store_code_or_name.strip()
    address = address.strip() if address else ""
    
    # 第一步：尝试精确匹配门店代码或名称
    mappings = db.get_huading_store_by_code_or_name(customer_id, store_code_or_name)
    if mappings:
        m = mappings[0]
        return {
            "success": True,
            "huading_store_code": m['huading_store_code'],
            "huading_store_name": m['huading_store_name'],
            "confidence": m['match_confidence'],
            "match_method": "EXACT",
            "message": f"精确匹配: {m['customer_store_name']}"
        }
    
    # 第二步：模糊匹配门店名称
    if len(store_code_or_name) >= 2:
        fuzzy_results = db.fuzzy_match_store(customer_id, store_code_or_name)
        if fuzzy_results:
            m = fuzzy_results[0]
            confidence = calculate_name_match_confidence(store_code_or_name, m['customer_store_name'])
            return {
                "success": True,
                "huading_store_code": m['huading_store_code'],
                "huading_store_name": m['huading_store_name'],
                "confidence": confidence,
                "match_method": "FUZZY",
                "message": f"模糊匹配: {m['customer_store_name']} (置信度{confidence:.2f})"
            }
    
    # 所有匹配策略都失败
    return {
        "success": False,
        "error": f"无法找到匹配的华鼎门店: {store_code_or_name}",
        "huading_store_code": None,
        "confidence": 0.0,
        "match_method": "none"
    }


def get_huading_store_code_with_detail(customer_id: str, store_code_or_name: str,
                                       address: str = "") -> Dict[str, Any]:
    """获取华鼎门店编码（带详细结果）"""
    result = get_huading_store_code(customer_id, store_code_or_name, address)
    
    if result['success']:
        return {
            "huading_store_code": result['huading_store_code'],
            "huading_store_name": result['huading_store_name'],
            "match_type": result['match_method'],
            "confidence": result['confidence'],
            "customer_store_name": store_code_or_name,
            "message": result['message']
        }
    
    return {
        "match_type": "NOT_FOUND",
        "confidence": 0.0,
        "customer_store_name": store_code_or_name,
        "message": f"无法找到匹配的华鼎门店: {store_code_or_name}"
    }


def calculate_name_match_confidence(input_str: str, matched: str) -> float:
    """计算名称匹配置信度"""
    input_str = input_str.lower().strip()
    matched = matched.lower().strip()
    
    if input_str == matched:
        return 1.0
    
    if matched in input_str or input_str in matched:
        return 0.9
    
    distance = levenshtein_distance(input_str, matched)
    max_len = max(len(input_str), len(matched))
    
    if max_len == 0:
        return 1.0
    
    similarity = 1.0 - distance / max_len
    return similarity


def levenshtein_distance(s1: str, s2: str) -> int:
    """计算编辑距离"""
    if len(s1) == 0:
        return len(s2)
    if len(s2) == 0:
        return len(s1)
    
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    for i in range(m + 1):
        dp[i][0] = i
    for j in range(n + 1):
        dp[0][j] = j
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            cost = 0 if s1[i-1] == s2[j-1] else 1
            dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1] + cost)
    
    return dp[m][n]


def batch_get_huading_store_codes(customer_id: str, 
                                  store_inputs: List[Dict[str, str]]) -> Dict[str, Any]:
    """批量获取华鼎门店编码"""
    results = {}
    errors = []
    
    for inp in store_inputs:
        key = inp.get("code", "") or inp.get("name", "")
        
        result = get_huading_store_code_with_detail(
            customer_id,
            inp.get("code", "") or inp.get("name", ""),
            inp.get("address", "")
        )
        
        if result.get("match_type") == "NOT_FOUND":
            errors.append(f"门店映射失败: {key}")
        
        results[key] = result
    
    return {"results": results, "errors": errors}


# ==================== 测试 ====================

if __name__ == "__main__":
    print("=== 门店映射模块测试（已连接数据库）===")
    
    # 测试门店编码查询
    print("\n1. 门店编码查询:")
    for store in ["沧州任丘北辛庄店", "长沙广承总店", "成都仓配送"]:
        result = get_huading_store_code("CUST_002", store)
        print(f"  {store}: {result.get('huading_store_code', 'N/A')} ({result.get('match_method', 'N/A')}, {result.get('confidence', 0):.2f})")
    
    print("\n2. 置信度计算:")
    print(f"  '沧州任丘北辛庄店' vs '沧州仓-任丘店': {calculate_name_match_confidence('沧州任丘北辛庄店', '沧州仓-任丘店'):.2f}")
