# Skill: 订单识别 (skill-order-recognition)

> **版本：** 1.2.0
> **日期：** 2026年4月28日
> **触发方式：** 客服上传文件
> **核心功能：** 识别文件格式，提取字段，计算置信度，返回结构化数据
> **复杂度：** 高

---

## INPUT_SCHEMA

参见 [INPUT_SCHEMA.json](./INPUT_SCHEMA.json)

---

## OUTPUT_SCHEMA

参见 [OUTPUT_SCHEMA.json](./OUTPUT_SCHEMA.json)

**核心特性：**
- 每个字段必须包含 `confidence` 和 `status`
- 订单明细每行有独立的行级置信度和状态
- 行置信度 = 加权平均（sku_code×0.25 + sku_name×0.25 + spec×0.15 + quantity×0.20 + unit×0.15）

---

## UI_RENDER_RULES

参见 [UI_RENDER_RULES.json](./UI_RENDER_RULES.json)

**渲染特性：**
- 订单明细表格显示置信度徽章和状态标签
- 行根据状态自动高亮（绿色/橙色/红色）
- 低置信度字段自动高亮提示

---

## ACTION_HANDLERS

参见 [ACTION_HANDLERS.json](./ACTION_HANDLERS.json)

---

## PROCESS_PROMPT

参见 [PROCESS_PROMPT.md](./PROCESS_PROMPT.md)

**置信度计算规则：**

| 字段 | 置信度规则 |
|------|-----------|
| sku_code | 直接引用=0.99，部分识别=0.85，无法识别=0.0 |
| sku_name | 直接引用=0.95，语义推断=0.80-0.90，无法识别=0.0 |
| spec | 规则识别成功=0.90，部分识别=0.75，无法识别=0.0 |
| quantity | 直接引用=0.99，计算得出=0.90，无法识别=0.0 |
| unit | 规则识别成功=0.95，推断=0.80，无法识别=0.0 |

**状态判定：**
- ≥95% → CONFIRMED（绿色）
- 80%-95% → LOW_CONFIDENCE（橙色）
- <80% → NEED_INPUT（红色）

---

## 依赖工具

| 工具名称 | 说明 |
|---------|------|
| read_excel | 读取Excel文件内容 |
| match_customer_sku | 客户商品编码匹配（可选） |
| match_store_alias | 门店简称匹配全称 |

## 状态

- [x] SKILL.md 完成
- [x] INPUT_SCHEMA.json 完成
- [x] OUTPUT_SCHEMA.json 完成（v1.2.0）
- [x] UI_RENDER_RULES.json 完成
- [x] ACTION_HANDLERS.json 完成
- [x] PROCESS_PROMPT.md 完成（v1.1）
- [ ] 单元测试待编写
