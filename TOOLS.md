# TOOLS.md - 工具配置

## 账单助手工具 (billing_tools)

### 工具路径
```
tools/billing/
├── __init__.py      # 入口
├── contract.py       # 合同解析
├── maker.py          # 账单制作
├── reviewer.py       # 账单审核
├── query.py          # 订单查询
└── src/             # 核心模块
    ├── calculator/   # 计费计算
    ├── knowledge/    # 规则库
    └── excel/        # Excel读取
```

### 可用函数

```python
# 合同解析
from billing_tools import parse_contract
parse_contract(file_path) -> {success, rules, confidence, message}

# 规则保存（必须确认）
from billing_tools import save_rule
save_rule(customer_id, rules, confirmed=True) -> {success, message}

# 账单制作
from billing_tools import make_bill
make_bill(customer_id, period, excel_path) -> {success, bill, summary, message}

# 账单审核
from billing_tools import review_bill
review_bill(customer_id, period) -> {success, report, message}

# 差异标记
from billing_tools import mark_adjustment
mark_adjustment(customer_id, period, {type, amount, reason}) -> {success, message}

# 订单查询
from billing_tools import query_order
query_order(customer_id, store_name, period) -> {success, orders, summary, message}
```

---

## 订单映射工具 (Order Mapping Tools)

### 工具路径
```
tools/
├── sku_mapping.go        # SKU匹配
├── store_mapping.go      # 门店映射
├── warehouse_mapping.go   # 仓库映射
├── unit_converter.go     # 单位转换
├── tools.go             # 工具注册
└── generate_huading_template.py  # 华鼎模板生成
```

### 可用Tool接口

| 工具名 | 功能 | 优先级 |
|--------|------|--------|
| `match_customer_sku` | 客户SKU→系统SKU | P0 |
| `get_huading_store_code` | 门店名称→华鼎门店编号 | P0 |
| `get_warehouse_code` | 门店→仓库编码 | P0 |
| `get_unit_type` | 客户单位→华鼎单位 | P0 |
| `generate_huading_template` | 生成31字段华鼎模板 | P0 |
| `validate_order_data` | 验证订单数据 | P2 |

---

## 知识库路径

```
knowledge/
├── rules/           # 客户规则（账单用）
├── mapping/         # 订单映射规则
│   ├── sku/        # SKU映射表
│   ├── store/       # 门店映射表
│   └── warehouse/   # 仓库映射表
├── bills/           # 历史账单
└── adjustments/    # 人工调整记录
```

---

## 告警阈值

- **账单**：差异 > 0.01元 → 必须告警
- **订单映射**：置信度 < 80% → 需要人工确认
