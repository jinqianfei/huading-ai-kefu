# TOOLS.md - 工具配置

## 工具目录

```
tools/
├── billing/                    # billing-workflow 相关
│   ├── __init__.py
│   ├── contract.py             # 合同解析
│   ├── maker.py                # 账单制作
│   ├── reviewer.py             # 账单审核
│   ├── query.py                # 订单查询
│   ├── calculator/
│   │   └── sku_coefficient.py  # SKU系数计算
│   └── llm_contract_parser.py # LLM视觉合同解析
│
└── (skill_order_to_huading_template 在 skills/ 目录下)

skills/
├── skill_order_to_huading_template/
│   ├── __init__.py             # OrderToHuadingTemplate类
│   └── SKILL.md
│
└── billing-workflow/
    └── SKILL.md
```

---

## Skill 1: skill_order_to_huading_template

### 工具类
```python
from skills.skill_order_to_huading_template import OrderToHuadingTemplate

# 初始化时只需要db_config，不传shipper_id
skill = OrderToHuadingTemplate(
    db_config={
        "host": "localhost",
        "port": 5432,
        "database": "ai_cs_support",
        "user": "your_username"
    }
)

# execute时会自动通过门店名匹配货主
result = skill.execute(order_file="/path/to/order.xlsx")
```

### 返回结果
```python
{
    "success": True,
    "output_file": "/path/to/huading_template.xlsx",
    "order_no": "DH-O-20260423-278070",
    "store_code": "KH2024070300038",
    "owner_code": "HZ2023061500002",  # 自动匹配到的货主ID
    "warehouse_code": "8",
    "item_count": 10,
    "unmatched_count": 0,
    "unmatched_items": [],
    "message": "模板生成成功"
}
```

### 货主匹配流程

```
1. 解析订单Excel → 获取门店名
2. 查询store_list表 → 匹配门店
3. 获取owner_code（货主ID）→ 返回给用户确认
4. 用户确认后 → 用owner_code过滤SKU匹配
5. 生成模板
```

---

## Skill 2: billing-workflow

### 可用函数

```python
# 1. 合同解析（步骤1）
from billing_tools import parse_contract, save_rule
result = parse_contract("/path/to/contract.jpg")
# result['rules'] 包含: storage, handling, transfer, delivery, product_coefficients...

# 保存规则（必须确认）
save_rule(customer_id, rules, confirmed=True)

# 2. SKU系数计算（步骤2）
from billing_tools import calc_sku_coefficients
result = calc_sku_coefficients(customer_id, product_excel_path)

# 3. 账单制作（步骤3）
from billing_tools import make_bill, format_bill_summary
result = make_bill(customer_id, period, excel_path)
print(format_bill_summary(result['summary']))

# 4. 账单审核（步骤4）
from billing_tools import review_bill, mark_adjustment
result = review_bill(customer_id, period)

# 5. 订单查询
from billing_tools import query_order
result = query_order(store_name="天津王口镇店", period="2026-02")
```

### 规则类型 (v2.4)
| 类型 | 说明 |
|------|------|
| storage | 仓储费（按仓库+温区） |
| handling | 操作费（整件/拆零） |
| transfer | 调拨费 |
| delivery | 共配费 |
| value_added | 增值服务费 |
| oversized | 超标件系数 |
| product_coefficients | 标准件系数 |

---

## 数据库表

| 表名 | 货主ID字段 | 用途 |
|------|-----------|------|
| store_list | owner_code | 门店列表，按owner_code过滤 |
| shipper_sku_mapping | shipper_id | SKU映射，按shipper_id过滤 |
| warehouse_code_mapping | - | 仓库编码（warehouse_name → warehouse_code） |
| customer | customer_id | 货主信息 |

---

## 货主-品牌对照

| 货主公司全称 | 品牌名称 | 货主ID |
|--------------|---------|--------|
| 河南上黎供应链管理有限公司 | 制茶青年 | HZ2023061500002 |
| 天津王口镇店 | - | CUSTOMER-WANGKOU |
| 长沙广承供应链有限公司 | - | CUSTOMER-GUANGCHENG |
| 沧州任丘长丰镇店 | - | CUSTOMER-CANGZHOU |

---

## 告警阈值

- **账单**：差异 > 0.01元 → 必须告警
- **订单映射**：置信度 < 80% → 需要人工确认
- **货主匹配**：必须用户确认后才能继续