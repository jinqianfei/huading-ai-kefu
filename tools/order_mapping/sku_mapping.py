# -*- coding: utf-8 -*-
"""
SKU映射模块 - 从Go工具翻译为Python（已连接数据库）
实现4层匹配策略：精确匹配→历史匹配→模糊匹配→语义匹配
"""

import re
import json
from typing import List, Dict, Optional, Any, Tuple
from dataclasses import dataclass
from difflib import SequenceMatcher

# 导入数据库模块
from . import db


@dataclass
class SystemSkuCandidate:
    """系统SKU候选结果"""
    sku_code: str
    sku_name: str
    confidence: float
    match_type: str  # EXACT/HISTORY/FUZZY/SEMANTIC/NONE


# 匹配级别常量
MATCH_EXACT = "EXACT"
MATCH_HISTORY = "HISTORY"
MATCH_FUZZY = "FUZZY"
MATCH_SEMANTIC = "SEMANTIC"
MATCH_NONE = "NONE"


# ==================== 核心匹配函数 ====================

def get_system_sku_by_product_name(product_name: str, customer_id: str = "") -> List[Dict]:
    """
    根据商品名称获取系统SKU候选列表
    实现4层匹配策略：精确匹配→历史匹配→模糊匹配→语义匹配
    """
    if not product_name:
        return []
    
    candidates = []
    seen = set()
    
    # 第1层：精确匹配
    exact_matches = match_exact(product_name, customer_id, seen)
    candidates.extend(exact_matches)
    
    # 第2层：历史匹配
    if not candidates or candidates[0]['confidence'] < 0.95:
        history_matches = match_history(product_name, customer_id, seen)
        candidates.extend(history_matches)
    
    # 第3层：模糊匹配
    if not candidates or candidates[0]['confidence'] < 0.85:
        fuzzy_matches = match_fuzzy(product_name, customer_id, seen)
        candidates.extend(fuzzy_matches)
    
    # 第4层：语义匹配（当前使用降级方案）
    if not candidates or candidates[0]['confidence'] < 0.70:
        semantic_matches = match_semantic(product_name, customer_id, seen)
        candidates.extend(semantic_matches)
    
    # 排序并去重
    sorted_candidates = sort_and_deduplicate_candidates(candidates)
    
    return sorted_candidates[:10]


def match_customer_sku(customer_sku_code: str, customer_id: str = None) -> Dict[str, Any]:
    """
    客户SKU编码匹配华鼎系统SKU编码（Tool接口）
    """
    try:
        # 尝试商品名称匹配
        candidates = get_system_sku_by_product_name(customer_sku_code, customer_id)
        
        if candidates:
            top = candidates[0]
            return {
                "success": True,
                "huading_sku_code": top['sku_code'],
                "confidence": top['confidence'],
                "match_method": top['match_type'],
                "message": f"匹配成功，置信度 {top['confidence']:.2f}"
            }
        
        return {
            "success": True,
            "huading_sku_code": None,
            "confidence": 0.0,
            "match_method": "none",
            "message": f"SKU '{customer_sku_code}' 未找到匹配"
        }
        
    except Exception as e:
        return {
            "success": False,
            "error": str(e),
            "huading_sku_code": None,
            "confidence": 0.0,
            "match_method": "none"
        }


# ==================== 第1层：精确匹配 ====================

def match_exact(product_name: str, customer_id: str, seen: set) -> List[Dict]:
    """第1层：精确匹配"""
    candidates = []
    
    # 1.1 查询customer_sku_mapping中的精确匹配
    mapping = db.get_customer_sku_mapping_by_name(customer_id, product_name)
    if mapping and mapping['system_sku_code'] not in seen:
        seen.add(mapping['system_sku_code'])
        candidates.append({
            "sku_code": mapping['system_sku_code'],
            "sku_name": mapping['system_sku_name'],
            "confidence": mapping['confidence'],
            "match_type": MATCH_EXACT
        })
    
    # 1.2 查询system_sku表中的product_name_lower精确匹配
    product_name_lower = product_name.lower().strip()
    skus = db.get_system_sku_by_product_name_lower(product_name_lower)
    for sku in skus:
        if sku['sku_code'] not in seen:
            seen.add(sku['sku_code'])
            candidates.append({
                "sku_code": sku['sku_code'],
                "sku_name": sku['sku_name'],
                "confidence": 1.0,
                "match_type": MATCH_EXACT
            })
    
    return candidates


# ==================== 第2层：历史匹配 ====================

def match_history(product_name: str, customer_id: str, seen: set) -> List[Dict]:
    """第2层：历史匹配"""
    # 从customer_sku_mapping查询该客户的历史匹配
    mapping = db.get_customer_sku_mapping_by_name(customer_id, product_name)
    if mapping and mapping['system_sku_code'] not in seen:
        seen.add(mapping['system_sku_code'])
        # 历史匹配降权
        confidence = mapping['confidence'] * 0.95
        return [{
            "sku_code": mapping['system_sku_code'],
            "sku_name": mapping['system_sku_name'],
            "confidence": confidence,
            "match_type": MATCH_HISTORY
        }]
    return []


# ==================== 第3层：模糊匹配 ====================

def match_fuzzy(product_name: str, customer_id: str, seen: set) -> List[Dict]:
    """第3层：模糊匹配"""
    candidates = []
    
    # 提取关键词
    keywords = extract_keywords(product_name)
    if not keywords:
        keywords = [product_name]
    
    # 基于关键词匹配
    for keyword in keywords:
        if len(keyword) < 2:
            continue
        skus = db.get_system_sku_by_keywords(keyword)
        for sku in skus:
            if sku['sku_code'] in seen:
                continue
            similarity = calculate_similarity(product_name, sku['sku_name'])
            if similarity >= 0.6:
                seen.add(sku['sku_code'])
                candidates.append({
                    "sku_code": sku['sku_code'],
                    "sku_name": sku['sku_name'],
                    "confidence": similarity,
                    "match_type": MATCH_FUZZY
                })
    
    # 基于编辑距离的模糊匹配
    all_skus = db.get_all_system_skus()
    for sku in all_skus:
        if sku['sku_code'] in seen:
            continue
        similarity = calculate_levenshtein_similarity(product_name, sku['sku_name'])
        if similarity >= 0.7:
            seen.add(sku['sku_code'])
            candidates.append({
                "sku_code": sku['sku_code'],
                "sku_name": sku['sku_name'],
                "confidence": similarity,
                "match_type": MATCH_FUZZY
            })
    
    return candidates


# ==================== 第4层：语义匹配 ====================

def match_semantic(product_name: str, customer_id: str, seen: set) -> List[Dict]:
    """第4层：语义匹配（降级方案：基于关键词）"""
    candidates = []
    
    core_words = extract_core_words(product_name)
    if not core_words:
        return candidates
    
    all_skus = db.get_all_system_skus()
    for sku in all_skus:
        if sku['sku_code'] in seen:
            continue
        sku_core_words = extract_core_words(sku['sku_name'])
        match_score = calculate_word_match_score(core_words, sku_core_words)
        if match_score >= 0.5:
            seen.add(sku['sku_code'])
            candidates.append({
                "sku_code": sku['sku_code'],
                "sku_name": sku['sku_name'],
                "confidence": match_score,
                "match_type": MATCH_SEMANTIC
            })
    
    return candidates


# ==================== 辅助函数 ====================

def extract_keywords(product_name: str) -> List[str]:
    """提取关键词"""
    cleaned = clean_product_name(product_name)
    
    separators = [" ", "-", "_", "|", "，", ",", "、", "（", "）", "(", ")"]
    words = [cleaned]
    
    for sep in separators:
        new_words = []
        for w in words:
            parts = w.split(sep)
            for p in parts:
                p = p.strip()
                if p and len(p) >= 2:
                    new_words.append(p)
        words = new_words
    
    return words


def extract_core_words(product_name: str) -> List[str]:
    """提取核心词"""
    stop_words = {
        "个", "件", "箱", "盒", "瓶", "袋",
        "kg", "g", "ml", "l", "升", "克",
        "大", "中", "小", "新", "老",
    }
    
    words = extract_keywords(product_name)
    core_words = []
    
    for w in words:
        w_lower = w.lower()
        if w_lower not in stop_words and len(w) >= 2:
            core_words.append(w_lower)
    
    return core_words


def clean_product_name(name: str) -> str:
    """清理商品名称"""
    import re
    name = name.strip()
    name = re.sub(r'\s+', ' ', name)
    return name


def calculate_similarity(s1: str, s2: str) -> float:
    """计算字符串相似度（基于最长公共子串）"""
    if not s1 or not s2:
        return 0.0
    
    s1 = s1.lower()
    s2 = s2.lower()
    
    if s1 == s2:
        return 1.0
    
    max_len = longest_common_substring(s1, s2)
    ref_len = max(len(s1), len(s2))
    
    if ref_len == 0:
        return 0.0
    
    return max_len / ref_len


def longest_common_substring(s1: str, s2: str) -> int:
    """计算最长公共子串长度"""
    m, n = len(s1), len(s2)
    if m == 0 or n == 0:
        return 0
    
    prev = [0] * (n + 1)
    max_len = 0
    
    for i in range(1, m + 1):
        curr = [0] * (n + 1)
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                curr[j] = prev[j-1] + 1
                if curr[j] > max_len:
                    max_len = curr[j]
        prev = curr
    
    return max_len


def calculate_levenshtein_similarity(s1: str, s2: str) -> float:
    """计算编辑距离相似度"""
    if not s1 and not s2:
        return 1.0
    if not s1 or not s2:
        return 0.0
    
    s1 = s1.lower()
    s2 = s2.lower()
    
    distance = levenshtein_distance(s1, s2)
    max_len = max(len(s1), len(s2))
    
    return 1.0 - distance / max_len


def levenshtein_distance(s1: str, s2: str) -> int:
    """计算编辑距离"""
    m, n = len(s1), len(s2)
    if m == 0:
        return n
    if n == 0:
        return m
    
    prev = list(range(n + 1))
    curr = [0] * (n + 1)
    
    for i in range(1, m + 1):
        curr[0] = i
        for j in range(1, n + 1):
            cost = 0 if s1[i-1] == s2[j-1] else 1
            curr[j] = min(prev[j] + 1, curr[j-1] + 1, prev[j-1] + cost)
        prev, curr = curr, prev
    
    return prev[n]


def calculate_word_match_score(words1: List[str], words2: List[str]) -> float:
    """计算词匹配分数"""
    if not words1 or not words2:
        return 0.0
    
    match_count = 0
    for w1 in words1:
        for w2 in words2:
            if w1 == w2 or w1 in w2 or w2 in w1:
                match_count += 1
                break
    
    min_len = min(len(words1), len(words2))
    return match_count / min_len if min_len > 0 else 0.0


def sort_and_deduplicate_candidates(candidates: List[Dict]) -> List[Dict]:
    """排序并去重候选结果"""
    if not candidates:
        return []
    
    candidates.sort(key=lambda x: x['confidence'], reverse=True)
    
    seen = set()
    result = []
    for c in candidates:
        if c['sku_code'] not in seen:
            seen.add(c['sku_code'])
            result.append(c)
    
    return result


# ==================== 测试 ====================

if __name__ == "__main__":
    print("=== SKU映射模块测试（已连接数据库）===")
    
    # 测试SKU匹配
    print("\n1. SKU匹配:")
    for product in ["1000塑杯", "600个塑杯", "伊利安慕希希腊酸奶", "可口可乐500ml"]:
        result = match_customer_sku(product, "CUST_001")
        print(f"  {product}: {result['huading_sku_code']} ({result['match_method']}, {result['confidence']:.2f})")


# ==================== 历史记录操作 ====================

def save_sku_match_result(customer_id: str, customer_sku_name: str, 
                          system_sku_code: str, system_sku_name: str,
                          match_type: str, confidence: float) -> bool:
    """保存SKU匹配结果到历史记录"""
    return db.save_match_history(customer_id, customer_sku_name, system_sku_code, system_sku_name, match_type, confidence)


def update_sku_mapping_from_feedback(customer_id: str, customer_sku_name: str, 
                                     correct_sku_code: str) -> bool:
    """根据用户反馈更新SKU映射"""
    # TODO: 实现反馈更新逻辑
    return True
