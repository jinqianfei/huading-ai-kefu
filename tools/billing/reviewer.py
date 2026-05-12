# billing_tools/reviewer.py
# 账单审核工具

"""
账单审核工具

用法:
    review_bill(customer_id, period) -> {report}
    mark_adjustment(customer_id, period, adjustment) -> {success, message}
"""

import sys
import os
import json

# 添加src到路径
_tools_dir = os.path.dirname(os.path.abspath(__file__))
_src_dir = os.path.join(_tools_dir, 'src')
_knowledge_dir = os.path.expanduser("~/openclaw-workspaces/ai-kefu/knowledge")
sys.path.insert(0, _src_dir)

from calculator.transport import calc_transport_fee

# 告警阈值
THRESHOLD = 0.01


def review_bill(customer_id: str, period: str = None, bill_data: dict = None) -> dict:
    """
    审核账单，标记差异
    
    Args:
        customer_id: 客户ID
        period: 账期（如 "2026-02"）
        bill_data: 账单数据（可选，不传则使用规则验证）
        
    Returns:
        {
            "success": bool,
            "report": dict,
            "message": str
        }
    """
    try:
        # 1. 加载规则
        rules_path = os.path.join(_knowledge_dir, "rules")
        rule_file = None
        
        for f in os.listdir(rules_path):
            if customer_id in f and f.endswith('.json'):
                rule_file = os.path.join(rules_path, f)
                break
        
        if not rule_file:
            return {
                "success": False,
                "report": None,
                "message": f"未找到客户规则（ID: {customer_id}）"
            }
        
        with open(rule_file, 'r', encoding='utf-8') as f:
            rules = json.load(f)
        
        customer_name = rules.get("customer_name", "")
        transport_rules = rules.get("rules", {}).get("transport", {})
        unit_price = transport_rules.get("price_per_unit", 0.15)
        min_fee = transport_rules.get("min_fee", 40)
        
        # 2. 如果没有传入账单数据，返回审核模板
        if not bill_data:
            return {
                "success": True,
                "report": {
                    "customer_id": customer_id,
                    "customer_name": customer_name,
                    "period": period or "2026-02",
                    "rules_loaded": True,
                    "unit_price": unit_price,
                    "min_fee": min_fee,
                    "threshold": THRESHOLD,
                    "total_count": 0,
                    "passed_count": 0,
                    "alert_count": 0,
                    "alerts": [],
                    "message": "规则已加载，请提供Excel文件进行审核"
                },
                "message": f"已加载{customer_name}的规则，阈值{THRESHOLD}元"
            }
        
        # 3. 审核账单数据
        alerts = []
        passed_count = 0
        total_count = 0
        
        for item in bill_data.get("transport_fees", []):
            total_count += 1
            std_units = float(item.get("std_units") or 0)
            bill_value = float(item.get("fee") or 0)
            
            # AI重新计算
            result = calc_transport_fee(std_units, unit_price, min_fee)
            ai_value = result.calculated_fee
            
            diff = abs(ai_value - bill_value)
            
            if diff > THRESHOLD:
                alerts.append({
                    "order_no": item.get("order_no"),
                    "store": item.get("store"),
                    "std_units": std_units,
                    "ai_value": round(ai_value, 2),
                    "bill_value": round(bill_value, 2),
                    "diff": round(diff, 2),
                    "severity": "high" if diff > 1 else "medium"
                })
            else:
                passed_count += 1
        
        report = {
            "customer_id": customer_id,
            "customer_name": customer_name,
            "period": period or "2026-02",
            "total_count": total_count,
            "passed_count": passed_count,
            "alert_count": len(alerts),
            "alerts": alerts[:10],  # 最多显示10条
            "pass_rate": round(passed_count / total_count * 100, 2) if total_count > 0 else 100
        }
        
        return {
            "success": True,
            "report": report,
            "message": f"审核完成：通过{passed_count}/{total_count}，告警{len(alerts)}项"
        }
        
    except Exception as e:
        return {
            "success": False,
            "report": None,
            "message": f"审核失败: {str(e)}"
        }


def mark_adjustment(customer_id: str, period: str, adjustment: dict) -> dict:
    """
    标记人工调整
    
    Args:
        customer_id: 客户ID
        period: 账期
        adjustment: {"type": str, "amount": float, "reason": str}
        
    Returns:
        {"success": bool, "message": str}
    """
    try:
        adj_dir = os.path.join(_knowledge_dir, "adjustments")
        os.makedirs(adj_dir, exist_ok=True)
        
        adj_file = os.path.join(adj_dir, f"{customer_id}_adjustments.json")
        
        # 读取现有调整
        adjustments = []
        if os.path.exists(adj_file):
            with open(adj_file, 'r', encoding='utf-8') as f:
                adjustments = json.load(f)
        
        # 添加新调整
        adjustments.append({
            "period": period,
            "type": adjustment.get("type"),
            "amount": adjustment.get("amount"),
            "reason": adjustment.get("reason"),
            "marked_at": "2026-05-08"
        })
        
        # 保存
        with open(adj_file, 'w', encoding='utf-8') as f:
            json.dump(adjustments, f, ensure_ascii=False, indent=2)
        
        return {
            "success": True,
            "message": f"已标记人工调整：{adjustment.get('reason')}"
        }
        
    except Exception as e:
        return {
            "success": False,
            "message": f"标记失败: {str(e)}"
        }


if __name__ == "__main__":
    print("账单审核工具")
    print(review_bill("CUST_001", "2026-02"))
