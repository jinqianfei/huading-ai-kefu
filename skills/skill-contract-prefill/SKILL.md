# Skill: 合同规则预填 (skill-contract-prefill)

> **版本：** 2.4.0
> **日期：** 2026年5月11日
> **更新内容：** 新增标准件系数提取（产品固定系数升级版），支持步骤2 SKU系数计算
> **触发方式：** 客服上传合同PDF + AI预填 + 逐条确认
> **核心功能：** 读取合同PDF，自动提取计费规则，录入系统生成规则库
> **复杂度：** 高
> **优先级：** ⭐⭐⭐ Phase 2 核心前置（M7）
> **前置条件：** M7未完成则M8无法启动

---

## 功能说明

本 Skill 是账单智能计算（M8）的前置条件，负责将合同PDF中的计费规则
结构化提取并录入系统，为后续账单计算提供规则依据。

### 核心能力

1. **PDF文本提取**：Word可编辑PDF（文字提取）+ 扫描件PDF（表格OCR）
2. **LLM规则提取**：Tool Use模式，每个字段标注合同来源（第X页第X段）
3. **交叉验证**：阈值=最低收费÷超量单价，自动推算并验证公式成立
4. **双人复核**：关键字段（单价/最低收费）需双人确认
5. **版本管理**：新规则覆盖旧版本，保留历史快照

---

## 计费模式区分

必须从合同中识别计费模式，不同模式决定了规则结构的差异：

### 模式识别

| 计费模式 | 识别特征 | 适用客户 |
|---------|---------|---------|
| `pallet` | 提及"托盘"、"元/托/天"、按托盘计费 | 长沙广承 |
| `standard_unit` | 提及"标准件"、"30天/周期"、按件计费 | 广州依然 |
| `hybrid` | 同时支持托盘和标准件两种模式 | - |

---

## 计费规则类型

### 1. 仓储费规则（storage_rules）

| 字段 | 说明 | 必填 |
|------|------|------|
| period_days | 计费周期天数（标准件模式） | 看模式 |
| storage_pricing | 按仓库+温区的单价 | ✅ |
| operation_mode | 操作费模式 | ✅ |
| overdue_enabled | 是否收取超期费 | standard_unit |
| overdue_grace_days | 超期宽限期 | 看模式 |
| overdue_fee_per_unit | 超期费单价 | 看模式 |

### 2. 调拨报价（transfer_rules）

仓库间调拨，必须与共配报价分开。

| 字段 | 说明 |
|------|------|
| transfer_min_order_qty | 起运量 |
| transfer_routes | 路线表 |
| transfer_price_type | 计费类型 |
| transfer_unit | 单位 |

### 3. 共配报价（delivery_rules）

仓库到门店配送，是重点提取部分。

#### 必填字段

| 字段 | 说明 | 必填 |
|------|------|------|
| delivery_pricing_type | 报价类型 | ✅ |
| min_order_qty | 起运量 | ✅ |
| unit_price | 单价 | ✅ |

#### 起运量相关（必须提取）

| 字段 | 说明 |
|------|------|
| min_order_qty | 起运量（件） |
| min_order_qty_by_warehouse | 按仓库的起运量（不同仓库可能不同） |
| min_order_qty_billing_mode | 不足起运量计费方式 |

**起运量计费方式：**
- `按实际数量计费`：不足也按实际计费
- `按起运量计费`：不足按起运量计费（常见）
- `按起运量整数倍计费`：如不足按起运量的2倍计费

#### 单价相关（必须提取）

| 字段 | 说明 |
|------|------|
| unit_price | 单价（元/标准件） |
| unit_price_by_warehouse | 按仓库的单价 |
| unit_price_by_zone | 按区域的单价 |

#### 最低收费相关（必须提取）

| 字段 | 说明 |
|------|------|
| min_fee | 最低收费金额 |
| min_fee_rule | 最低收费规则描述 |

**最低收费规则示例：**
- "不足起运量时，按起运量×单价计费"
- "最低收费50元/票"

### 4. 超规格计费规则（oversized_rules）

超出标准件定义（重量/体积）的计费规则。

| 字段 | 说明 | 必填 |
|------|------|------|
| oversized_tiers | 超标分级定义 | ✅ |
| oversized_tier1_threshold | 一级超标阈值 | ✅ |
| oversized_tier1_coef | 一级超标系数（默认1.2） | ✅ |
| oversized_tier2_threshold | 二级超标阈值 | 可选 |
| oversized_tier2_coef | 二级超标系数（默认2.0） | 可选 |
| oversized_tier3_* | 三级超标（如果有） | 可选 |
| extra_weight_fee | 额外超重费用 | 可选 |
| oversized_billing_note | 超规格计费特殊说明 | 可选 |

**超标分级示例：**
```
一级超标：15kg<重量≤25kg 或 0.05m³<体积≤0.08m³ → 系数1.2
二级超标：25kg<重量≤30kg 或 0.08m³<体积≤0.1m³ → 系数2.0
```

### 5. 标准件计算规则（standard_unit_rules）

| 字段 | 说明 | 默认值 |
|------|------|-------|
| max_weight_kg | 最大重量 | 15 kg |
| max_volume_m3 | 最大体积 | 0.05 m³ |
| calculation_rule | 计算规则 | max(重量/最大重量, 体积/最大体积) |

### 6. 标准件系数（product_coefficients）

**含义**：每个SKU的计费系数

**区分两类SKU**：
- **固定SKU**：合同附件商品参数表中列出的SKU，系数固定为1.0，不适用超规格计费规则
- **非固定SKU**：按标准件定义 + 超规格规则计算系数

**计算逻辑**：
```
固定SKU（合同附件商品参数表中）→ 系数 = 1.0
非固定SKU:
  - 重量>30kg 或 体积>0.1m³ → 系数 = max(重量/15, 体积/0.05)
  - 25kg<重量≤30kg 或 0.08m³<体积≤0.1m³ → 系数 = 2.0
  - 15kg<重量≤25kg 或 0.05m³<体积≤0.08m³ → 系数 = 1.2
  - 重量≤15kg 且 体积≤0.05m³ → 系数 = 1.0
```

**提取要求**：
- 识别合同附件中的商品参数表（通常包含：SKU编码、名称、规格、体积m³、重量kg）
- 表格中列出的SKU，系数固定为1.0
- 标注来源位置（第X页第X个表格）

**输出格式**：
```json
{
  "product_coefficients": {
    "description": "产品标准件系数",
    "coefficient_unit": "倍",
    "standard_unit_definition": {
      "max_weight_kg": 15,
      "max_volume_m3": 0.05
    },
    "fixed_skus": {
      "GGK01": { "name": "Review Card", "spec": "...", "volume_m3": 0.001, "weight_kg": 0.5 },
      "STD01": { "name": "Hand Bag", "spec": "...", "volume_m3": 0.002, "weight_kg": 0.3 }
    },
    "fixed_sku_count": 15,
    "oversized_rules": {
      "tier1": { "threshold": "15kg<重量≤25kg 或 0.05m³<体积≤0.08m³", "coef": 1.2 },
      "tier2": { "threshold": "25kg<重量≤30kg 或 0.08m³<体积≤0.1m³", "coef": 2.0 },
      "tier3": { "threshold": "重量>30kg 或 体积>0.1m³", "coef": "max(重量/15, 体积/0.05)" }
    }
  }
}
```

**示例（长沙广承合同）**：
| SKU编码 | 名称 | 规格 | 体积m³ | 重量kg | 标准件系数 |
|--------|------|------|--------|--------|----------|
| GGK01 | Review Card | 10张/包 | 0.001 | 0.5 | 1.0（固定）|
| STD01 | Hand Bag | 1个 | 0.002 | 0.3 | 1.0（固定）|
| （其他SKU） | - | - | - | - | 按超规格规则 |

**如果合同没有商品参数表**：
- 设置 `fixed_skus: {}`
- 所有SKU按超规格规则计算标准件系数

**如果合同没有商品参数表**：
- 设置 `products: {}`
- 标记 `NOT_FOUND`
- 所有SKU按超规格规则计算

### 7. 增值服务费（value_added_rules）

| 字段 | 说明 |
|------|------|
| sorting_fee | 混装分拣费 |
| repackaging_fee | 再包装费 |
| labeling_fee | 贴标费 |
| delivery_to_store_fee | 送货进店费 |
| market_distribution_fee | 代发市场费用（如：20件起，2元/件） |

### 8. 特殊条款（special_rules）

| 字段 | 说明 |
|------|------|
| difficult_delivery_surcharge | 困难配送加收 |
| delivery_period | 配送周期（一周六配、一周三配等） |
| delivery_time_limit | 配送时效（1-2天） |
| other_special | 其他特殊条款 |

### 9. 返利/折扣条款（rebate_discount_rules）

⚠️ 返利折扣需人工确认后生效

---

## 输出结构（按 OUTPUT_SCHEMA v2.2）

```
contract_info          → 合同基本信息
billing_mode           → 计费模式
storage_rules          → 仓储费规则
transfer_rules         → 调拨报价
delivery_rules         → 共配报价（含超规格/起运量/最低收费）
oversized_rules        → 超规格计费规则 ⭐ 新增
standard_unit_rules    → 标准件计算规则
product_coefficients   → 产品固定系数
value_added_rules      → 增值服务费
special_rules          → 特殊条款
rebate_discount_rules  → 返利/折扣条款
validation             → 验证结果
summary                → 摘要
```

---

## 疑点自动标记规则

| 触发条件 | 标记 | 严重度 |
|---------|------|--------|
| 调拨/共配未区分 | NOT_DISTINGUISHED | 🔴 高 |
| 计费模式未识别 | NOT_IDENTIFIED | 🔴 高 |
| 起运量未明确 | min_order_qty_NOT_IDENTIFIED | 🔴 高 |
| 超规格规则未明确 | oversized_NOT_IDENTIFIED | 🔴 高 |
| 必填字段缺失 | NOT_FOUND | 🔴 高 |
| 单价与最低收费矛盾 | CONFLICT | 🔴 高 |
| 含"待定"/"面议" | PENDING | ⚠️ 中 |

---

## 版本历史

| 版本 | 日期 | 更新内容 |
|------|------|---------|
| 2.0.0 | 2026-05-09 | 新增调拨/共配分离、多仓库多温区等 |
| 2.1.0 | 2026-05-09 | 新增计费模式区分、存储费周期、操作费归属 |
| 2.2.0 | 2026-05-09 | 新增超规格计费规则、起运量规则、最低收费规则 |
