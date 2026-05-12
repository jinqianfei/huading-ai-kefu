# Skill: 订单模板转换 (skill-order-template-converter)

> **版本：** 1.0.0
> **日期：** 2026年4月28日
> **触发方式：** 客服上传文件 + 确认识别结果
> **核心功能：** 将客户订单转换为华鼎标准出库单模板
> **复杂度：** 高
> **优先级：** ⭐⭐⭐ 核心首发功能

---

## 功能说明

本 Skill 是 AI客服赋能系统 的核心功能，负责将客户发来的非标准订单 Excel/截图，
快速准确地转换为华鼎标准的出库单格式。

### 核心能力

1. **文件格式识别**：支持 Excel、PDF、图片（OCR）等多种格式
2. **5类数据映射**：
   - 客户商品编码 → 系统SKU编码
   - 门店简称 → 门店全称
   - 规格描述 → 系统计量单位
   - 字段位置映射
   - 客户 → 默认仓库/配送规则
3. **智能置信度评估**：每个字段都有置信度，高置信度可跳过，低置信度高亮提醒
4. **特殊场景处理**：合单、自提、二次追加等

---

## INPUT_SCHEMA

```json
{
  "skill_id": "skill-order-template-converter",
  "version": "1.0.0",
  "input_fields": [
    {
      "field_id": "order_file",
      "type": "file",
      "label": "上传订单文件",
      "accepted_formats": [".xlsx", ".xls", ".pdf", ".png", ".jpg", ".jpeg"],
      "max_size_mb": 50,
      "required": true,
      "description": "支持Excel、PDF、图片格式，单个文件不超过50MB"
    },
    {
      "field_id": "customer_id",
      "type": "select",
      "label": "选择客户",
      "options_source": "api:/customers",
      "searchable": true,
      "required": true,
      "placeholder": "请输入客户名称搜索"
    },
    {
      "field_id": "order_type",
      "type": "radio",
      "label": "订单类型",
      "options": [
        { "value": "standard", "label": "标准配送单" },
        { "value": "self_pickup", "label": "顾客自提" },
        { "value": "combined", "label": "合单" }
      ],
      "required": true,
      "default": "standard"
    }
  ]
}
```

---

## OUTPUT_SCHEMA

```json
{
  "skill_id": "skill-order-template-converter",
  "version": "1.0.0",
  "output_type": "order_confirmation",
  "fields": [
    {
      "section_id": "order_header",
      "title": "订单头信息",
      "confidence_threshold": 0.8,
      "fields": [
        {
          "field_id": "order_no",
          "label": "订单编号",
          "value": "DH-O-20260423-283742",
          "confidence": 0.99,
          "status": "CONFIRMED",
          "source": "直接引用"
        },
        {
          "field_id": "customer_name",
          "label": "客户名称",
          "value": "任丘北辛庄店",
          "original_value": "任丘北辛庄店",
          "confidence": 0.95,
          "status": "CONFIRMED",
          "matched_alias": "沧州任丘北辛庄店",
          "source": "查表得出"
        },
        {
          "field_id": "contact_person",
          "label": "联系人",
          "value": "肖文月",
          "confidence": 0.98,
          "status": "CONFIRMED",
          "source": "直接引用"
        },
        {
          "field_id": "delivery_address",
          "label": "收货地址",
          "value": "河北省沧州市任丘市北辛庄镇北辛庄村制茶青年",
          "confidence": 0.92,
          "status": "CONFIRMED",
          "source": "直接引用"
        },
        {
          "field_id": "contact_phone",
          "label": "联系电话",
          "value": "13171700566",
          "confidence": 0.99,
          "status": "CONFIRMED",
          "source": "直接引用"
        },
        {
          "field_id": "delivery_date",
          "label": "发货日期",
          "value": "2026-04-24",
          "confidence": 0.99,
          "status": "CONFIRMED",
          "source": "直接引用"
        }
      ]
    },
    {
      "section_id": "order_items",
      "title": "商品明细",
      "confidence_threshold": 0.8,
      "fields": [
        {
          "field_id": "item_1",
          "row_index": 1,
          "confidence": 0.97,
          "status": "CONFIRMED",
          "columns": [
            { "col_id": "sku_code", "label": "商品编码", "value": "P1556689351", "confidence": 0.99, "status": "CONFIRMED" },
            { "col_id": "sku_name", "label": "商品名称", "value": "1000塑杯", "confidence": 0.95, "status": "CONFIRMED" },
            { "col_id": "spec", "label": "规格", "value": "规格：600个", "confidence": 0.98, "status": "CONFIRMED" },
            { "col_id": "quantity", "label": "数量", "value": "1", "confidence": 0.99, "status": "CONFIRMED" },
            { "col_id": "unit", "label": "单位", "value": "件", "confidence": 0.99, "status": "CONFIRMED" },
            { "col_id": "matched_sku", "label": "系统SKU", "value": "SKU-20240615-0042", "confidence": 0.99, "status": "AUTO_MATCHED" },
            { "col_id": "matched_sku_name", "label": "系统商品名", "value": "1000塑杯（600个装）", "confidence": 0.95, "status": "AUTO_MATCHED" }
          ]
        },
        {
          "field_id": "item_2",
          "row_index": 2,
          "confidence": 0.85,
          "status": "LOW_CONFIDENCE",
          "columns": [
            { "col_id": "sku_code", "label": "商品编码", "value": "P969550376", "confidence": 0.99, "status": "CONFIRMED" },
            { "col_id": "sku_name", "label": "商品名称", "value": "速冻桑葚/新", "confidence": 0.85, "status": "LOW_CONFIDENCE", "notes": "商品名称可能有误，请确认" },
            { "col_id": "spec", "label": "规格", "value": "件：1kg*10袋/件", "confidence": 0.90, "status": "CONFIRMED" },
            { "col_id": "quantity", "label": "数量", "value": "1", "confidence": 0.99, "status": "CONFIRMED" },
            { "col_id": "unit", "label": "单位", "value": "件", "confidence": 0.99, "status": "CONFIRMED" },
            { "col_id": "matched_sku", "label": "系统SKU", "value": "SKU-20240615-0089", "confidence": 0.88, "status": "LOW_CONFIDENCE" },
            { "col_id": "matched_sku_name", "label": "系统商品名", "value": "速冻桑葚（1kg*10袋）", "confidence": 0.88, "status": "LOW_CONFIDENCE" }
          ]
        }
      ]
    },
    {
      "section_id": "summary",
      "title": "汇总信息",
      "confidence_threshold": 0.8,
      "fields": [
        { "field_id": "total_items", "label": "商品种数", "value": "2", "confidence": 1.0, "status": "CONFIRMED" },
        { "field_id": "total_quantity", "label": "商品总件数", "value": "2", "confidence": 1.0, "status": "CONFIRMED" },
        { "field_id": "overall_confidence", "label": "整体置信度", "value": "0.95", "confidence": 0.95, "status": "CONFIRMED" }
      ]
    }
  ],
  "system_actions": [
    {
      "action_id": "generate_huading_template",
      "label": "生成华鼎标准模板",
      "type": "primary",
      "icon": "download",
      "enabled_when": "所有CONFIRMED字段 ≥ 95%"
    },
    {
      "action_id": "export_excel",
      "label": "导出Excel",
      "type": "secondary",
      "icon": "file-excel"
    },
    {
      "action_id": "reject_all",
      "label": "全部重填",
      "type": "danger",
      "icon": "refresh"
    }
  ],
  "metadata": {
    "task_id": "task_20260428_001",
    "process_time_ms": 2340,
    "model_used": "qwen-turbo"
  }
}
```

---

## UI_RENDER_RULES

```json
{
  "output_type": "order_confirmation",
  "layout": "two_column",
  "sections": [
    {
      "section_id": "order_header",
      "layout": "vertical_form",
      "collapsible": false,
      "fields": [
        { "field_id": "order_no", "render_as": "readonly_text", "show_confidence": false },
        { "field_id": "customer_name", "render_as": "editable_text", "highlight_on_status": ["LOW_CONFIDENCE", "MANUAL_MODIFIED"], "show_confidence": true },
        { "field_id": "contact_person", "render_as": "editable_text", "highlight_on_status": ["LOW_CONFIDENCE"], "show_confidence": true },
        { "field_id": "delivery_address", "render_as": "editable_text", "highlight_on_status": ["LOW_CONFIDENCE"], "show_confidence": true },
        { "field_id": "contact_phone", "render_as": "readonly_text", "show_confidence": true },
        { "field_id": "delivery_date", "render_as": "editable_text", "show_confidence": true }
      ],
      "confidence_visual": {
        "high": { "color": "#52c41a", "icon": "check-circle" },
        "medium": { "color": "#faad14", "icon": "exclamation-circle" },
        "low": { "color": "#f5222d", "icon": "close-circle" }
      }
    },
    {
      "section_id": "order_items",
      "layout": "table",
      "collapsible": false,
      "columns": [
        { "col_id": "sku_code", "label": "商品编码", "width": "120px", "sortable": true },
        { "col_id": "sku_name", "label": "商品名称", "width": "180px" },
        { "col_id": "spec", "label": "规格", "width": "140px" },
        { "col_id": "quantity", "label": "数量", "width": "80px", "align": "right" },
        { "col_id": "unit", "label": "单位", "width": "80px" },
        { "col_id": "matched_sku", "label": "系统SKU", "width": "140px" },
        { "col_id": "matched_sku_name", "label": "系统商品名", "width": "180px" }
      ],
      "row_actions": [
        { "action_id": "confirm_item", "label": "确认", "icon": "check", "type": "primary" },
        { "action_id": "edit_item", "label": "修改", "icon": "edit", "type": "default" },
        { "action_id": "delete_item", "label": "删除", "icon": "delete", "type": "danger" }
      ],
      "striped": true,
      "hoverable": true,
      "empty_state": "暂无商品明细"
    },
    {
      "section_id": "summary",
      "layout": "card",
      "collapsible": true,
      "defaultCollapsed": false,
      "fields": [
        { "field_id": "total_items", "render_as": "readonly_text" },
        { "field_id": "total_quantity", "render_as": "readonly_text" },
        { "field_id": "overall_confidence", "render_as": "confidence_badge" }
      ]
    }
  ],
  "system_actions_bar": {
    "position": "bottom_fixed",
    "background": "#ffffff",
    "shadow": "0 -2px 8px rgba(0,0,0,0.1)",
    "padding": "16px 24px",
    "actions": [
      { "action_id": "generate_huading_template", "render_as": "primary_button", "label": "生成华鼎标准模板", "icon": "download" },
      { "action_id": "export_excel", "render_as": "secondary_button", "label": "导出Excel", "icon": "file-excel" },
      { "action_id": "reject_all", "render_as": "text_button", "label": "全部重填", "danger": true, "icon": "refresh" }
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
      "action_id": "generate_huading_template",
      "trigger": "button_click",
      "request": {
        "action": "GENERATE_TEMPLATE",
        "payload": {
          "task_id": "{{task_id}}",
          "confirmed_fields": "{{all_confirmed_fields}}",
          "output_format": "xlsx"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "download_file",
          "file_field": "template_file"
        },
        "error": {
          "render_type": "error_message",
          "message": "生成模板失败，请重试",
          "retry": true
        }
      },
      "loading": {
        "show": true,
        "text": "正在生成华鼎标准模板..."
      }
    },
    {
      "action_id": "confirm_item",
      "trigger": "row_button_click",
      "request": {
        "action": "CONFIRM_ITEM",
        "payload": {
          "task_id": "{{task_id}}",
          "item_index": "{{row_index}}",
          "confirmed": true
        }
      },
      "response_handling": {
        "success": {
          "render_type": "update_row",
          "update_fields": { "status": "CONFIRMED", "confidence": 1.0 }
        }
      }
    },
    {
      "action_id": "edit_item",
      "trigger": "row_button_click",
      "request": {
        "action": "EDIT_ITEM",
        "payload": {
          "task_id": "{{task_id}}",
          "item_index": "{{row_index}}",
          "edit_data": "{{user_input}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "update_row",
          "update_fields": { "status": "MANUAL_MODIFIED" }
        }
      },
      "confirm": {
        "required": true,
        "title": "确认修改",
        "message": "确定要修改这条记录吗？"
      }
    },
    {
      "action_id": "delete_item",
      "trigger": "row_button_click",
      "request": {
        "action": "DELETE_ITEM",
        "payload": {
          "task_id": "{{task_id}}",
          "item_index": "{{row_index}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "delete_row"
        }
      },
      "confirm": {
        "required": true,
        "title": "确认删除",
        "message": "确定要删除这条记录吗？此操作不可撤销。"
      }
    },
    {
      "action_id": "reject_all",
      "trigger": "button_click",
      "request": {
        "action": "REJECT_ALL",
        "payload": {
          "task_id": "{{task_id}}",
          "reason": "{{user_reason}}"
        }
      },
      "response_handling": {
        "success": {
          "render_type": "reset_form",
          "message": "已重置，请重新上传订单"
        }
      },
      "confirm": {
        "required": true,
        "title": "确认重填",
        "message": "确定要清空所有数据重新填写吗？之前的操作将被丢弃。",
        "okText": "确认重填",
        "cancelText": "取消"
      }
    }
  ]
}
```

---

## PROCESS_PROMPT

```markdown
# Skill: 订单模板自动转换 v1.0

## 1. 角色定义

你是一个专业的物流订单处理助手。你的任务是将客户发来的非标准订单Excel/截图，
快速准确地转换为华鼎标准的出库单格式。

## 2. 核心原则

1. **高准确率**：你处理的每一条数据都会影响财务结算，必须严谨
2. **有据可查**：每个字段标注来源，置信度透明
3. **人机协作**：置信度低的字段不要自己猜，标注出来让客服确认
4. **不静默丢失**：无法处理的字段必须明确标注，不可静默跳过

## 3. 输入定义

- 客户上传的订单文件（Excel/图片/PDF）
- 客户ID（已知客户可查映射表）
- 订单类型（标准配送单/顾客自提/合单）

## 4. 处理流程

### 步骤1：识别文件格式
- Excel → 使用POI直接读取表格数据
- PDF/图片 → 调用OCR识别

### 步骤2：识别客户模板类型
- 优先通过客户ID匹配已知模板配置
- 未知客户 → 通过表头字段签名识别模板类型

### 步骤3：字段提取与映射
识别并提取以下字段：

| 字段类型 | 提取要求 | 置信度规则 |
|---------|---------|-----------|
| 订单编号 | 直接引用 | 99% |
| 下单日期 | 直接引用 | 99% |
| 客户名称 | 匹配门店简称→全称映射 | 查表为准 |
| 收货地址 | 直接引用 | 95%+地址完整性校验 |
| 联系电话 | 直接引用 | 99%+手机号正则校验 |
| 商品编码 | 查客户商品编码→系统SKU映射表 | 精确匹配99% |
| 商品名称 | 直接引用，必要时用商品名称匹配SKU | 85-95% |
| 规格/单位 | 识别大单位/小单位，提取换算率 | 80-95% |
| 数量 | 直接引用 | 99% |

### 步骤4：5类数据映射

**映射1：客户商品编码 → 系统SKU编码（精确匹配）**
```
优先精确查 customer_sku_mapping 表
命中 → 直接返回系统SKU编码，置信度99%
未命中 → 进入映射2
```

**映射2：客户商品名 → 系统商品名（语义匹配）**
```
计算商品名称相似度（字符相似度×40% + 向量相似度×60%）
相似度≥85% → 自动推荐
相似度60-85% → 列出Top3备选，供客服选择
相似度<60% → 标注"无法匹配"，客服手动选择
```

**映射3：门店简称 → 门店全称**
```
查 customer_alias_mapping 表
命中 → 返回全称+匹配类型
未命中 → 标注"疑似新门店"，客服确认
```

**映射4：规格描述 → 系统计量单位**
```
规则1：含"件："前缀 → 大单位出库，提取换算率（例："件：1kg*10袋/件"→比率10）
规则2：含"规格："前缀 → 按实际数量计
规则3：无法识别 → 标注"单位待确认"
```

**映射5：客户 → 默认仓库/配送规则**
```
查 customer_warehouse_mapping 表
命中 → 返回默认仓库+配送类型
未命中 → 标注"请选择仓库"
```

### 步骤5：生成OUTPUT_SCHEMA结构
将处理结果按以下格式输出（JSON）：

```json
{
  "output_type": "order_confirmation",
  "task_id": "自动生成",
  "fields": [...],
  "system_actions": [...],
  "metadata": { "process_time_ms": xxx, "model_used": "xxx" }
}
```

## 5. 置信度分级处理规则

| 置信度范围 | 字段状态 | 前端渲染 | 客服操作 |
|-----------|---------|---------|---------|
| ≥95% | CONFIRMED | 绿色勾 | 可跳过 |
| 80%-95% | LOW_CONFIDENCE | 橙色高亮 | 建议确认 |
| <80% | NEED_INPUT | 红色+空 | 必须填写 |
| 无法匹配 | NOT_FOUND | 红色标注 | 手动选择 |

## 6. 特殊场景处理

### 合单（多客户合并发货）
识别逻辑：
- 同一文件中出现多个"客户名称"字段
- 识别"合单"、"合并"关键字
处理：拆分为多个订单，分别处理

### 自提订单
识别关键字："自提"、"顾客来拿"、"到店自取"
处理：在备注字段自动标注"【自提】"，不分配配送

### 二次报货
识别关键字："引用"、"追加"、"续报"
处理：找到原订单号，追加商品明细

### 大小单位混合下单
识别规格中的单位，自动拆分为两条系统记录

## 7. 输出规范

1. 所有字段必须标注来源（原文引用/计算得出/查表得出）
2. 置信度必须量化，不可留空
3. 无法处理的字段必须明确标注，不可静默跳过
4. 返回的结构必须符合OUTPUT_SCHEMA.json定义
```

---

## 依赖工具

| 工具名称 | 说明 |
|---------|------|
| read_excel | 读取Excel文件内容 |
| match_customer_sku | 客户商品编码精确匹配系统SKU |
| match_store_alias | 门店简称匹配全称 |
| generate_huading_template | 生成华鼎标准出库单Excel |

## 置信度计算公式

```
字段置信度 = 匹配类型得分 × 60% + 历史一致性得分 × 40%

其中：
- 精确匹配 = 1.0
- 高相似度 = 0.85-0.99
- 中相似度 = 0.70-0.84
- 低相似度 = <0.70
```

## 状态

- [x] SKILL.md 完成
- [ ] INPUT_SCHEMA.json 待实现
- [ ] OUTPUT_SCHEMA.json 待实现
- [ ] UI_RENDER_RULES.json 待实现
- [ ] ACTION_HANDLERS.json 待实现
- [ ] PROCESS_PROMPT.md 待实现
- [ ] 单元测试待编写
