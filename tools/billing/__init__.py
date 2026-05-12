# billing_tools/__init__.py
# AI账单助手工具包 v2.3

"""
AI账单助手工具包

提供给AI客服Agent调用

支持功能（PRD v2.3）：
- 合同解析（支持调拨报价/共配报价分离）
- 账单制作（含操作费/增值服务/困难配送/超标件）
- 账单审核
- 差异标记
- 订单查询

规则类型（v2.3新增产品固定系数）：
- storage: 仓储费（按仓库+温区）
- handling: 操作费（整件/拆零）
- transfer: 调拨费（仓库间调拨）⭐ v2.1新增
- delivery: 共配费（仓库到门店）⭐ v2.1新增
- value_added: 增值服务费
- oversized: 超标件计费系数
- special: 特殊规则
- product_coefficients: 标准件系数 ⭐ v2.3新增
  - 标准件定义：合同约定的标准件体积/重量阈值（如 ≤15kg 且 ≤0.05m³）
  - 标准件系数：每个SKU的计费系数（固定SKU=1.0，非固定SKU按超规格规则计算）

使用方式:
    from billing_tools import parse_contract, make_bill, review_bill, query_order
    
    # 解析合同
    result = parse_contract("/path/to/contract.jpg")
    # result['rules'] 包含: storage, handling, transfer, delivery, product_coefficients, ...
    
    # 生成账单
    result = make_bill("CUST_001", "2026-02", "/path/to/excel.xlsx")
    
    # 审核账单
    result = review_bill("CUST_001", "2026-02")
    
    # 查询订单
    result = query_order(store_name="宫小燕", period="2026-02")
    
    # 格式化账单汇总
    from billing_tools import format_bill_summary
    print(format_bill_summary(result['summary']))
"""

from .contract import parse_contract, save_rule
from .maker import make_bill, format_bill_summary
from .reviewer import review_bill, mark_adjustment
from .query import query_order, query_store_summary
from .calculator.sku_coefficient import calc_sku_coefficients, batch_calc_sku_coefficients

__all__ = [
    'parse_contract',
    'save_rule',
    'make_bill',
    'format_bill_summary',
    'review_bill',
    'mark_adjustment',
    'query_order',
    'query_store_summary',
    'calc_sku_coefficients',
    'batch_calc_sku_coefficients',
]

__version__ = "2.3.1"
