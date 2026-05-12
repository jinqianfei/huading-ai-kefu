# Skill: SKU智能匹配 (skill-sku-matcher)

> **版本：** 1.0.0
> **日期：** 2026年4月28日
> **触发方式：** 自动触发（订单处理时）/ 手动触发（客服查询时）
> **核心功能：** 客户商品编码/名称 → 系统SKU编码的智能匹配
> **复杂度：** 中

---

## 功能说明

本 Skill 负责将客户订单中的商品信息（客户商品编码或商品名称）
智能匹配到华鼎系统的 SKU 编码。

### 匹配策略（优先级排序）

1. **精确匹配**：客户商品编码 → 系统SKU编码（99%置信度）
2. **历史匹配**：查历史匹配记录（80-95%置信度）
3. **语义匹配**：商品名称相似度计算（60-85%置信度）
4. **人工兜底**：无法匹配时展示备选，客服手动选择

### 语义相似度计算公式

```
相似度 = 0.4 × 字符相似度 + 0.6 × 向量相似度

字符相似度 = 1 - Levenshtein距离 / max(len(A), len(B))
向量相似度 = CosineSimilarity(Embedding(商品名A), Embedding(商品名B))
```

---

## INPUT_SCHEMA

```json
{
  "skill_id": "skill-sku-matcher",
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
      "field_id": "customer_sku_code",
      "type": "text",
      "label": "客户商品编码",
      "placeholder": "输入客户商品编码，如 P1556689351",
      "required": false
    },
    {
      "field_id": "customer_sku_name",
      "type": "text",
      "label": "客户商品名称",
      "placeholder": "输入客户商品名称",
      "required": false
    },
    {
      "field_id": "quantity",
      "type": "number",
      "label": "数量",
      "min": 0,
      "required": false
    }
  ]
}
```

---

## OUTPUT_SCHEMA

```json
{
  "skill_id": "skill-sku-matcher",
  "version": "1.0.0",
  "output_type": "sku_match_result",
  "fields": [
    {
      "section_id": "match_result",
      "title": "匹配结果",
      "confidence_threshold": 0.85,
      "fields": [
        {
          "field_id": "match_type",
          "label": "匹配类型",
          "value": "精确匹配",
          "confidence": 0.99,
          "status": "CONFIRMED",
          "source": "查表得出"
        },
        {
          "field_id": "system_sku_code",
          "label": "系统SKU编码",
          "value": "SKU-20240615-0042",
          "confidence": 0.99,
          "status": "CONFIRMED",
          "source": "查表得出"
        },
        {
          "field_id": "system_sku_name",
          "label": "系统商品名",
          "value": "1000塑杯（600个装）",
          "confidence": 0.99,
          "status": "CONFIRMED",
          "source": "查表得出"
        },
        {
          "field_id": "unit_conversion",
          "label": "单位换算",
          "value": "1件 = 600个",
          "confidence": 0.95,
          "status": "CONFIRMED",
          "source": "查表得出"
        },
        {
          "field_id": "fee_coefficient",
          "label": "计费系数",
          "value": "1.0",
          "confidence": 0.99,
          "status": "CONFIRMED",
          "source": "查表得出"
        }
      ]
    },
    {
      "section_id": "alternatives",
      "title": "备选方案",
      "confidence_threshold": 0.6,
      "fields": []
    }
  ],
  "metadata": {
    "task_id": "",
    "process_time_ms": 0,
    "model_used": "qwen-turbo"
  }
}
```

---

## UI_RENDER_RULES

```json
{
  "output_type": "sku_match_result",
  "layout": "two_column",
  "sections": [
    {
      "section_id": "match_result",
      "layout": "card",
      "collapsible": false,
      "fields": [
        { "field_id": "match_type", "render_as": "readonly_text", "show_confidence": true },
        { "field_id": "system_sku_code", "render_as": "editable_text", "highlight_on_status": ["LOW_CONFIDENCE"], "show_confidence": true },
        { "field_id": "system_sku_name", "render_as": "editable_text", "highlight_on_status": ["LOW_CONFIDENCE"], "show_confidence": true },
        { "field_id": "unit_conversion", "render_as": "readonly_text", "show_confidence": true },
        { "field_id": "fee_coefficient", "render_as": "readonly_text", "show_confidence": true }
      ],
      "confidence_visual": {
        "high": { "color": "#52c41a", "icon": "check-circle" },
        "medium": { "color": "#faad14", "icon": "exclamation-circle" },
        "low": { "color": "#f5222d", "icon": "close-circle" }
      }
    }
  ],
  "system_actions_bar": {
    "position": "bottom_fixed",
    "actions": [
      { "action_id": "confirm_match", "render_as": "primary_button", "label": "确认匹配", "icon": "check" },
      { "action_id": "manual_select", "render_as": "secondary_button", "label": "手动选择", "icon": "edit" }
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
      "action_id": "confirm_match",
      "trigger": "button_click",
      "request": {
        "action": "CONFIRM_SKU_MATCH",
        "payload": {
          "task_id": "{{task_id}}",
          "customer_id": "{{customer_id}}",
          "customer_sku_code": "{{customer_sku_code}}",
          "system_sku_code": "{{system_sku_code}}",
          "match_type": "{{match_type}}",
          "confidence": "{{confidence}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "show_message",
          "message": "匹配已确认，已记录到历史匹配表",
          "message_type": "success"
        }
      },
      "loading": {
        "show": true,
        "text": "正在保存匹配记录..."
      }
    },
    {
      "action_id": "manual_select",
      "trigger": "button_click",
      "request": {
        "action": "SHOW_MANUAL_SELECT",
        "payload": {
          "task_id": "{{task_id}}",
          "customer_id": "{{customer_id}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "refresh_table"
        }
      }
    }
  ]
}
```

---

## PROCESS_PROMPT

```markdown
# Skill: SKU智能匹配 v1.0

## 1. 角色定义

你是一个专业的商品匹配助手。你的任务是将客户订单中的商品信息，
智能匹配到华鼎系统的SKU编码。

## 2. 核心原则

1. **精确优先**：有精确匹配时绝不推荐模糊匹配
2. **有据可查**：每个匹配结果都要说明匹配路径
3. **透明置信度**：置信度要量化，让人知道可靠程度

## 3. 输入定义

- 客户ID
- 客户商品编码（可选，有则优先使用）
- 客户商品名称
- 数量（可选）

## 4. 处理流程

### 步骤1：精确匹配（客户商品编码）
```
查 customer_sku_mapping 表
条件：customer_id + customer_sku_code
命中 → 返回系统SKU编码，置信度99%
```

### 步骤2：历史匹配
```
查 match_history 表
条件：customer_id + customer_sku_code
命中 → 检查历史纠正情况
  - 无纠正 → 置信度95%
  - 有1次纠正 → 置信度90%
  - 有3次以上纠正 → 置信度85%
```

### 步骤3：语义匹配
```
如果以上都未命中，使用语义相似度匹配

计算公式：
相似度 = 0.4 × 字符相似度 + 0.6 × 向量相似度

匹配结果：
- 相似度 ≥ 85% → 自动推荐，置信度 = 相似度
- 相似度 60-85% → 列出Top3备选，置信度 = 相似度
- 相似度 < 60% → 标注"无法匹配"
```

### 步骤4：热度衰减
```
根据最近匹配时间调整置信度：
- 3个月内成功匹配过 → 热度系数1.0
- 3-6个月前成功匹配 → 热度系数0.9
- 6个月以上未匹配 → 热度系数0.8
- 新商品 → 热度系数0.8

最终置信度 = 基础置信度 × 热度系数
```

## 5. 置信度分级处理规则

| 置信度范围 | 字段状态 | 说明 |
|-----------|---------|------|
| ≥95% | CONFIRMED | 高可信，自动采纳 |
| 85%-95% | LOW_CONFIDENCE | 可信，建议确认 |
| 60%-85% | NOT_FOUND | 需客服选择 |
| <60% | NOT_FOUND | 无法匹配，强制人工 |

## 6. 输出规范

返回JSON格式：
```json
{
  "match_type": "EXACT/HISTORY/SEMANTIC/MANUAL",
  "system_sku_code": "SKU-xxx",
  "system_sku_name": "商品名称",
  "confidence": 0.99,
  "alternatives": []  // 当置信度60-85%时的备选
}
```
```

---

## 依赖工具

| 工具名称 | 说明 |
|---------|------|
| match_customer_sku | 客户商品编码精确匹配系统SKU |
| match_store_alias | 门店简称匹配全称（可选） |

## 状态

- [x] SKILL.md 完成
- [ ] INPUT_SCHEMA.json 待实现
- [ ] OUTPUT_SCHEMA.json 待实现
- [ ] UI_RENDER_RULES.json 待实现
- [ ] ACTION_HANDLERS.json 待实现
- [ ] PROCESS_PROMPT.md 待实现
- [ ] 单元测试待编写
