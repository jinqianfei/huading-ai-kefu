# Skill: 账单审核 (skill-bill-audit)

> **版本：** 1.0.0
> **日期：** 2026年4月28日
> **触发方式：** 账单提交后自动触发 / 客服手动触发
> **核心功能：** AI审核账单，标红异常项，与合同规则逐一比对
> **复杂度：** 高
> **优先级：** ⭐⭐⭐ Phase 2 核心功能（M8审核）
> **目标：** 100%覆盖，零漏报

---

## 功能说明

本 Skill 负责AI自动审核账单，将AI重新计算的金额与原始账单金额进行比对，
发现异常并标注，确保账单准确性。

### 审核规则

```
一致性审核：
- AI计算金额 vs 账单原始金额
- 差异率 = |AI计算 - 账单金额| / max(AI计算, 账单金额)
- 差异率 < 0.1% → PASS（绿色）
- 差异率 0.1%-5% → WARNING（橙色）
- 差异率 > 5% → ERROR（红色）

环比审核：
- 同店同费用项环比
- 差异 > 30% → WARNING
- 差异 > 100% → ERROR
```

---

## INPUT_SCHEMA

```json
{
  "skill_id": "skill-bill-audit",
  "version": "1.0.0",
  "input_fields": [
    {
      "field_id": "bill_no",
      "type": "text",
      "label": "账单编号",
      "placeholder": "输入账单编号，如：BILL-202604-001",
      "required": true
    },
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
      "required": true
    },
    {
      "field_id": "audit_scope",
      "type": "radio",
      "label": "审核范围",
      "options": [
        { "value": "full", "label": "完整审核（全部费用项）" },
        { "value": "amount_only", "label": "仅金额审核" },
        { "value": "anomaly_only", "label": "仅异常项" }
      ],
      "default": "full",
      "required": true
    }
  ]
}
```

---

## OUTPUT_SCHEMA

```json
{
  "skill_id": "skill-bill-audit",
  "version": "1.0.0",
  "output_type": "bill_audit_result",
  "fields": [
    {
      "section_id": "audit_summary",
      "title": "审核结果汇总",
      "confidence_threshold": 0.9,
      "fields": [
        { "field_id": "audit_passed", "label": "审核状态", "confidence": 1.0 },
        { "field_id": "consistency_score", "label": "一致性得分", "confidence": 0.95 },
        { "field_id": "total_items", "label": "总审核项", "confidence": 1.0 },
        { "field_id": "passed_items", "label": "通过项", "confidence": 1.0 },
        { "field_id": "warning_items", "label": "警告项", "confidence": 1.0 },
        { "field_id": "error_items", "label": "错误项", "confidence": 1.0 }
      ]
    },
    {
      "section_id": "audit_details",
      "title": "审核明细",
      "confidence_threshold": 0.85,
      "row_based": true,
      "columns": [
        { "col_id": "fee_type", "label": "费用类型", "width": "100px" },
        { "col_id": "ai_amount", "label": "AI计算金额", "width": "120px", "align": "right" },
        { "col_id": "bill_amount", "label": "账单金额", "width": "120px", "align": "right" },
        { "col_id": "diff", "label": "差异", "width": "100px", "align": "right" },
        { "col_id": "diff_rate", "label": "差异率", "width": "100px", "align": "right" },
        { "col_id": "status", "label": "状态", "width": "80px" },
        { "col_id": "severity", "label": "严重程度", "width": "80px" }
      ]
    },
    {
      "section_id": "anomaly_details",
      "title": "异常详情",
      "confidence_threshold": 0.8,
      "row_based": true,
      "columns": [
        { "col_id": "anomaly_type", "label": "异常类型", "width": "120px" },
        { "col_id": "description", "label": "描述", "width": "250px" },
        { "col_id": "suggestion", "label": "处理建议", "width": "200px" },
        { "col_id": "action", "label": "操作", "width": "80px" }
      ],
      "row_actions": [
        { "action_id": "view_detail", "label": "查看", "icon": "eye", "type": "default" },
        { "action_id": "ignore", "label": "忽略", "icon": "close", "type": "default" }
      ]
    },
    {
      "section_id": "audit_actions",
      "title": "审核操作记录",
      "fields": [
        { "field_id": "auditor", "label": "审核人" },
        { "field_id": "audit_time", "label": "审核时间" },
        { "field_id": "final_status", "label": "最终状态" }
      ]
    }
  ],
  "system_actions": [
    { "action_id": "approve_bill", "label": "通过审核", "type": "primary", "icon": "check", "enabled_when": "error_items = 0" },
    { "action_id": "reject_bill", "label": "驳回账单", "type": "danger", "icon": "close" },
    { "action_id": "export_report", "label": "导出审核报告", "type": "secondary", "icon": "download" }
  ],
  "metadata": {
    "task_id": "string",
    "audit_time_ms": "number"
  }
}
```

---

## UI_RENDER_RULES

```json
{
  "output_type": "bill_audit_result",
  "layout": "two_column",
  "sections": [
    {
      "section_id": "audit_summary",
      "layout": "card",
      "collapsible": false,
      "fields": [
        { "field_id": "audit_passed", "render_as": "alert_message", "show_status_icon": true },
        { "field_id": "consistency_score", "render_as": "confidence_badge" },
        { "field_id": "total_items", "render_as": "readonly_text" },
        { "field_id": "passed_items", "render_as": "readonly_text", "class": "text-success" },
        { "field_id": "warning_items", "render_as": "readonly_text", "class": "text-warning" },
        { "field_id": "error_items", "render_as": "readonly_text", "class": "text-danger" }
      ]
    },
    {
      "section_id": "audit_details",
      "layout": "table",
      "collapsible": false,
      "columns": [
        { "col_id": "fee_type", "label": "费用类型", "width": "100px" },
        { "col_id": "ai_amount", "label": "AI计算金额", "width": "120px", "align": "right" },
        { "col_id": "bill_amount", "label": "账单金额", "width": "120px", "align": "right" },
        { "col_id": "diff", "label": "差异", "width": "100px", "align": "right" },
        { "col_id": "diff_rate", "label": "差异率", "width": "100px", "align": "right" },
        { "col_id": "status", "label": "状态", "width": "80px" },
        { "col_id": "severity", "label": "严重程度", "width": "80px" }
      ],
      "row_highlight": {
        "ERROR": "#fee2e2",
        "WARNING": "#fef9c3"
      },
      "striped": true
    },
    {
      "section_id": "anomaly_details",
      "layout": "table",
      "collapsible": true,
      "defaultCollapsed": false,
      "title_suffix": "（需关注）",
      "columns": [
        { "col_id": "anomaly_type", "label": "异常类型", "width": "120px" },
        { "col_id": "description", "label": "描述", "width": "250px" },
        { "col_id": "suggestion", "label": "处理建议", "width": "200px" },
        { "col_id": "action", "label": "操作", "width": "80px" }
      ],
      "row_actions": [
        { "action_id": "view_detail", "label": "查看", "icon": "eye", "type": "default" },
        { "action_id": "ignore", "label": "忽略", "icon": "close", "type": "default" }
      ],
      "empty_state": "无异常项"
    },
    {
      "section_id": "audit_actions",
      "layout": "card",
      "collapsible": true,
      "defaultCollapsed": true,
      "fields": [
        { "field_id": "auditor", "render_as": "readonly_text" },
        { "field_id": "audit_time", "render_as": "readonly_text" },
        { "field_id": "final_status", "render_as": "readonly_text" }
      ]
    }
  ],
  "system_actions_bar": {
    "position": "bottom_fixed",
    "background": "#ffffff",
    "shadow": "0 -2px 8px rgba(0,0,0,0.1)",
    "padding": "16px 24px",
    "actions": [
      { "action_id": "approve_bill", "render_as": "primary_button", "label": "通过审核", "icon": "check" },
      { "action_id": "reject_bill", "render_as": "danger_button", "label": "驳回账单", "icon": "close" },
      { "action_id": "export_report", "render_as": "secondary_button", "label": "导出审核报告", "icon": "download" }
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
      "action_id": "approve_bill",
      "trigger": "button_click",
      "request": {
        "action": "APPROVE_BILL",
        "payload": {
          "task_id": "{{task_id}}",
          "bill_no": "{{bill_no}}",
          "auditor": "{{current_user}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "show_message",
          "message": "账单审核已通过",
          "message_type": "success"
        },
        "render_type_2": "update_field",
        "update_fields": { "final_status": "APPROVED" }
      },
      "confirm": {
        "required": true,
        "title": "确认通过",
        "message": "确定要通过此账单审核吗？"
      }
    },
    {
      "action_id": "reject_bill",
      "trigger": "button_click",
      "request": {
        "action": "REJECT_BILL",
        "payload": {
          "task_id": "{{task_id}}",
          "bill_no": "{{bill_no}}",
          "reason": "{{user_reason}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "show_message",
          "message": "账单已驳回，请检查异常项后重新提交",
          "message_type": "warning"
        }
      },
      "confirm": {
        "required": true,
        "title": "确认驳回",
        "message": "确定要驳回此账单吗？请填写驳回原因。"
      }
    },
    {
      "action_id": "export_report",
      "trigger": "button_click",
      "request": {
        "action": "EXPORT_AUDIT_REPORT",
        "payload": {
          "task_id": "{{task_id}}",
          "bill_no": "{{bill_no}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "download_file",
          "file_field": "report_file"
        }
      },
      "loading": {
        "show": true,
        "text": "正在生成审核报告..."
      }
    },
    {
      "action_id": "view_detail",
      "trigger": "row_button_click",
      "request": {
        "action": "SHOW_ANOMALY_DETAIL",
        "payload": {
          "task_id": "{{task_id}}",
          "anomaly_id": "{{row.anomaly_id}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "show_dialog",
          "dialog_type": "anomaly_detail"
        }
      }
    },
    {
      "action_id": "ignore",
      "trigger": "row_button_click",
      "request": {
        "action": "IGNORE_ANOMALY",
        "payload": {
          "task_id": "{{task_id}}",
          "anomaly_id": "{{row.anomaly_id}}",
          "reason": "{{user_reason}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "update_row",
          "update_fields": { "status": "IGNORED" }
        }
      },
      "confirm": {
        "required": true,
        "title": "确认忽略",
        "message": "确定要忽略此异常项吗？"
      }
    }
  ]
}
```

---

## PROCESS_PROMPT

```markdown
# Skill: 账单审核 v1.0

## 1. 角色定义

你是一个专业的账单审核助手。你的任务是：
1. AI重新计算账单各项费用
2. 与原始账单金额进行比对
3. 标红异常项，生成审核报告
4. 提供处理建议

## 2. 核心原则

1. **零漏报**：所有账单都必须经过审核
2. **精准定位**：异常项必须精确到费用类型和金额
3. **有据可查**：每个异常都要有计算依据
4. **建议明确**：提供可执行的处理建议

## 3. 输入定义

- 账单编号
- 客户ID
- 账期
- 审核范围

## 4. 处理流程

### 步骤1：获取账单原始数据
```
查 bill_summary 表
获取：各项费用原始金额
```

### 步骤2：获取合同规则
```
查 customer_contract_rules 表
获取：计费规则
```

### 步骤3：AI重新计算
```
按合同规则重新计算每项费用：
仓储费 = 计算结果
运输费 = 计算结果
...
```

### 步骤4：一致性比对
```
差异 = AI计算金额 - 账单金额
差异率 = |差异| / max(AI计算, 账单金额)

状态判断：
- 差异率 < 0.1% → PASS（绿色）
- 差异率 0.1%-5% → WARNING（橙色）
- 差异率 > 5% → ERROR（红色）
```

### 步骤5：环比分析
```
获取历史同期账单数据
比对同费用项变化
异常变化 → 标注WARNING
```

### 步骤6：生成审核报告
```json
{
  "audit_passed": false,
  "consistency_score": 0.87,
  "total_items": 5,
  "passed_items": 3,
  "warning_items": 1,
  "error_items": 1,
  "anomaly_details": [...]
}
```

## 5. 置信度分级

| 审核结果 | 置信度 | 处理方式 |
|---------|-------|---------|
| PASS | ≥95% | 自动通过 |
| WARNING | 85%-95% | 建议确认 |
| ERROR | <85% | 必须处理 |

## 6. Override机制

当客服确认忽略某异常项时：
1. 记录Override原因
2. 标注Override标记
3. 不阻断提交流程
4. 保留审计记录
```

---

## 依赖工具

| 工具名称 | 说明 |
|---------|------|
| calculate_bill | AI重新计算账单 |
| get_contract_rules | 获取合同规则 |
| audit_bill | 执行一致性审核 |

---

## 状态

- [x] SKILL.md 完成
- [ ] 其他规范文件待实现
- [ ] 单元测试待编写
