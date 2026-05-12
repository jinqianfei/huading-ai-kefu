# Skill: 账单计算 (skill-bill-calculator)

> **版本：** 1.0.0
> **日期：** 2026年4月28日
> **触发方式：** 客服手动触发 / 定时任务
> **核心功能：** 根据合同计费规则，自动计算客户账单金额
> **复杂度：** 高
> **优先级：** ⭐⭐⭐ Phase 2 核心功能（M8）
> **前置条件：** 必须先完成M7合同规则前置录入

---

## 功能说明

本 Skill 负责根据合同规则库中的计费规则，计算客户的月度账单金额。

### 支持的计费模型

**模型1：长沙广承类（共配运费）**
```
运输费 = 最低收费 + max(0, 标准件数量 - 阈值) × 超量单价
```

**模型2：广州依然类（运输费）**
```
运输费 = 标准件数量 × 运输单价
```

**模型3：仓储费**
```
仓储费 = 托盘数 × 存储单价 × 存储天数
出库费 = 整件数 × 整件单价 + 拆零数 × 拆零单价
```

**模型4：区域定价**
```
不同区域不同价格：
- 成都：40元/3.5元
- 重庆：120元/9元
- 昆明：260元/15元
```

---

## INPUT_SCHEMA

```json
{
  "skill_id": "skill-bill-calculator",
  "version": "1.0.0",
  "input_fields": [
    {
      "field_id": "customer_id",
      "type": "select",
      "label": "选择客户",
      "options_source": "api:/customers",
      "searchable": true,
      "required": true
    },
    {
      "field_id": "billing_period",
      "type": "month",
      "label": "账期",
      "required": true,
      "default": "当前月份"
    },
    {
      "field_id": "data_source",
      "type": "radio",
      "label": "数据来源",
      "options": [
        { "value": "auto", "label": "自动获取（从OMS拉取出库记录）" },
        { "value": "manual", "label": "手动上传（Excel导入）" }
      ],
      "default": "auto",
      "required": true
    },
    {
      "field_id": "outbound_file",
      "type": "file",
      "label": "出库记录文件",
      "accepted_formats": [".xlsx", ".xls"],
      "max_size_mb": 20,
      "required": false,
      "description": "当数据来源选择手动上传时需要"
    }
  ]
}
```

---

## OUTPUT_SCHEMA

```json
{
  "skill_id": "skill-bill-calculator",
  "version": "1.0.0",
  "output_type": "bill_calculation",
  "fields": [
    {
      "section_id": "bill_header",
      "title": "账单头信息",
      "confidence_threshold": 0.9,
      "fields": [
        { "field_id": "customer_name", "label": "客户名称", "confidence": 1.0, "status": "CONFIRMED" },
        { "field_id": "billing_period", "label": "账期", "confidence": 1.0, "status": "CONFIRMED" },
        { "field_id": "outbound_records_count", "label": "出库记录数", "confidence": 1.0, "status": "CONFIRMED" }
      ]
    },
    {
      "section_id": "fee_breakdown",
      "title": "费用明细",
      "confidence_threshold": 0.95,
      "row_based": true,
      "columns": [
        { "col_id": "fee_type", "label": "费用类型", "width": "120px" },
        { "col_id": "calculation_rule", "label": "计算规则", "width": "200px" },
        { "col_id": "quantity", "label": "数量", "width": "100px", "align": "right" },
        { "col_id": "unit_price", "label": "单价", "width": "100px", "align": "right" },
        { "col_id": "amount", "label": "金额", "width": "120px", "align": "right" },
        { "col_id": "source", "label": "来源", "width": "120px" }
      ]
    },
    {
      "section_id": "bill_summary",
      "title": "账单汇总",
      "confidence_threshold": 0.99,
      "fields": [
        { "field_id": "storage_fee", "label": "仓储费", "confidence": 0.99, "status": "CONFIRMED" },
        { "field_id": "transport_fee", "label": "运输费", "confidence": 0.99, "status": "CONFIRMED" },
        { "field_id": "outbound_fee", "label": "出库费", "confidence": 0.99, "status": "CONFIRMED" },
        { "field_id": "other_fee", "label": "其他费用", "confidence": 0.95, "status": "CONFIRMED" },
        { "field_id": "total_amount", "label": "账单总额", "confidence": 0.99, "status": "CONFIRMED" }
      ]
    },
    {
      "section_id": "audit_info",
      "title": "审核信息",
      "fields": [
        { "field_id": "ai_confidence", "label": "AI置信度", "confidence": 0.95 },
        { "field_id": "rule_match_status", "label": "规则匹配状态", "confidence": 0.98 }
      ]
    }
  ],
  "system_actions": [
    { "action_id": "confirm_bill", "label": "确认账单", "type": "primary", "icon": "check" },
    { "action_id": "override_bill", "label": "人工修正", "type": "warning", "icon": "edit" },
    { "action_id": "submit_bill", "label": "提交BMS", "type": "primary", "icon": "upload", "enabled_when": "status = confirmed" },
    { "action_id": "reject_bill", "label": "驳回重算", "type": "danger", "icon": "close" }
  ],
  "metadata": {
    "task_id": "string",
    "calculation_time_ms": "number",
    "rules_applied": "number"
  }
}
```

---

## UI_RENDER_RULES

```json
{
  "output_type": "bill_calculation",
  "layout": "two_column",
  "sections": [
    {
      "section_id": "bill_header",
      "layout": "card",
      "collapsible": false,
      "fields": [
        { "field_id": "customer_name", "render_as": "readonly_text" },
        { "field_id": "billing_period", "render_as": "readonly_text" },
        { "field_id": "outbound_records_count", "render_as": "readonly_text" }
      ]
    },
    {
      "section_id": "fee_breakdown",
      "layout": "table",
      "collapsible": false,
      "columns": [
        { "col_id": "fee_type", "label": "费用类型", "width": "120px" },
        { "col_id": "calculation_rule", "label": "计算规则", "width": "200px" },
        { "col_id": "quantity", "label": "数量", "width": "100px", "align": "right" },
        { "col_id": "unit_price", "label": "单价", "width": "100px", "align": "right" },
        { "col_id": "amount", "label": "金额", "width": "120px", "align": "right" },
        { "col_id": "source", "label": "来源", "width": "120px" }
      ],
      "striped": true,
      "hoverable": true,
      "summary": true
    },
    {
      "section_id": "bill_summary",
      "layout": "card",
      "collapsible": false,
      "fields": [
        { "field_id": "storage_fee", "render_as": "readonly_text" },
        { "field_id": "transport_fee", "render_as": "readonly_text" },
        { "field_id": "outbound_fee", "render_as": "readonly_text" },
        { "field_id": "other_fee", "render_as": "readonly_text" },
        { "field_id": "total_amount", "render_as": "readonly_text", "class": "total-amount" }
      ],
      "summary_style": true
    },
    {
      "section_id": "audit_info",
      "layout": "card",
      "collapsible": true,
      "defaultCollapsed": true,
      "fields": [
        { "field_id": "ai_confidence", "render_as": "confidence_badge" },
        { "field_id": "rule_match_status", "render_as": "readonly_text" }
      ]
    }
  ],
  "system_actions_bar": {
    "position": "bottom_fixed",
    "background": "#ffffff",
    "shadow": "0 -2px 8px rgba(0,0,0,0.1)",
    "padding": "16px 24px",
    "actions": [
      { "action_id": "confirm_bill", "render_as": "primary_button", "label": "确认账单", "icon": "check" },
      { "action_id": "override_bill", "render_as": "secondary_button", "label": "人工修正", "icon": "edit" },
      { "action_id": "submit_bill", "render_as": "primary_button", "label": "提交BMS", "icon": "upload" },
      { "action_id": "reject_bill", "render_as": "text_button", "label": "驳回重算", "danger": true, "icon": "close" }
    ]
  }
}
```

---

## ACTION_HANDLERS

```json
{
  "handlers": [
    {
      "action_id": "confirm_bill",
      "trigger": "button_click",
      "request": {
        "action": "CONFIRM_BILL",
        "payload": {
          "task_id": "{{task_id}}",
          "customer_id": "{{customer_id}}",
          "billing_period": "{{billing_period}}",
          "total_amount": "{{total_amount}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "update_field",
          "update_fields": { "status": "CONFIRMED" },
          "render_type_2": "show_message",
          "message": "账单已确认",
          "message_type": "success"
        }
      },
      "loading": {
        "show": true,
        "text": "正在确认账单..."
      }
    },
    {
      "action_id": "override_bill",
      "trigger": "button_click",
      "request": {
        "action": "SHOW_OVERRIDE_FORM",
        "payload": {
          "task_id": "{{task_id}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "show_dialog",
          "dialog_type": "override_form"
        }
      }
    },
    {
      "action_id": "submit_bill",
      "trigger": "button_click",
      "request": {
        "action": "SUBMIT_TO_BMS",
        "payload": {
          "task_id": "{{task_id}}",
          "customer_id": "{{customer_id}}",
          "billing_period": "{{billing_period}}",
          "bill_summary": "{{all_confirmed_fields}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "show_message",
          "message": "账单已提交至BMS系统",
          "message_type": "success"
        }
      },
      "loading": {
        "show": true,
        "text": "正在提交账单..."
      },
      "confirm": {
        "required": true,
        "title": "确认提交",
        "message": "确定要将账单提交至BMS系统吗？提交后将无法撤回。",
        "okText": "确认提交",
        "cancelText": "取消"
      }
    },
    {
      "action_id": "reject_bill",
      "trigger": "button_click",
      "request": {
        "action": "REJECT_BILL",
        "payload": {
          "task_id": "{{task_id}}",
          "reason": "{{user_reason}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "reset_form",
          "message": "账单已驳回，请检查数据后重新计算"
        }
      },
      "confirm": {
        "required": true,
        "title": "确认驳回",
        "message": "确定要驳回此账单吗？"
      }
    }
  ]
}
```

---

## PROCESS_PROMPT

```markdown
# Skill: 账单计算 v1.0

## 1. 角色定义

你是一个专业的账单计算助手。你的任务是：
1. 从数据库获取客户合同规则
2. 从OMS系统获取出库记录
3. 按计费规则计算账单金额
4. 输出详细的费用明细和汇总

## 2. 核心原则

1. **精确计算**：每笔费用都要有计算公式和来源
2. **规则匹配**：优先使用客户合同规则
3. **透明可查**：费用明细清晰可追踪
4. **AI可校验**：计算结果可供M9账单审核使用

## 3. 输入定义

- 客户ID
- 账期（月份）
- 数据来源（自动/手动）

## 4. 处理流程

### 步骤1：获取合同规则
```
查 customer_contract_rules 表
条件：customer_id + 有效期内
输出：计费规则JSON
```

### 步骤2：获取出库记录
```
查 outbound_records 表
条件：customer_id + 账期内
输出：出库记录列表
```

### 步骤3：费用计算

**仓储费计算**
```
存储费 = Σ(托盘数 × 存储单价 × 存储天数)
入库费 = Σ(入库托盘数 × 入库单价)
出库费_整件 = Σ(整件数 × 整件单价)
出库费_拆零 = Σ(拆零数 × 拆零单价)
```

**运输费计算**
```
共配模式：运输费 = 最低收费 + max(0, 标准件 - 阈值) × 超量单价
直发模式：运输费 = 标准件数 × 运输单价
```

**区域定价**
```
按配送区域匹配对应单价
```

### 步骤4：汇总输出
```json
{
  "storage_fee": 1234.00,
  "transport_fee": 5678.00,
  "outbound_fee": 890.00,
  "other_fee": 0.00,
  "total_amount": 7802.00,
  "fee_details": [...]
}
```

## 5. 置信度计算

```
置信度 = 规则匹配度 × 数据完整度

规则匹配度：
- 100%规则可用 → 1.0
- 90%规则可用 → 0.9
- <50%规则可用 → 0.5

数据完整度：
- 出库记录完整 → 1.0
- 出库记录缺失 <10% → 0.9
- 出库记录缺失 >20% → 0.7
```

## 6. 特殊情况

### 无合同规则
输出提示"该客户未配置计费规则，请先完成M7合同规则录入"

### 数据不完整
部分数据缺失时，标注数据缺失比例，计算最大可能金额

### 计费规则变更
跨账期时，按规则变更日期分段计算
```

---

## 依赖工具

| 工具名称 | 说明 |
|---------|------|
| calculate_bill | 按计费规则计算账单 |
| get_outbound_records | 获取出库记录 |
| get_contract_rules | 获取合同规则 |

---

## 状态

- [x] SKILL.md 完成
- [ ] 其他规范文件待实现
- [ ] 单元测试待编写
