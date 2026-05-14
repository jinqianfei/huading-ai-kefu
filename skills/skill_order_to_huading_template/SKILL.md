---
name: skill-order-to-huading-template
description: 将客户订单转换为华鼎出库单模板（31字段），支持Excel/图片/PDF/文字多种输入格式。当用户说"处理订单"、"生成模板"、"订单转华鼎"时触发。
metadata:
  openclaw:
    triggers:
      - 处理订单
      - 生成模板
      - 订单转华鼎
      - 导出模板
      - 生成华鼎模板
      - 转换订单
      - 上传订单
      - 扫描订单
    requires:
      config:
        - db_config
---

# skill-order-to-huading-template

## Overview

**Skill Name**: 订单转华鼎出库单模板

**Description**: 将客户订单（Excel/图片/PDF/文字）转换为华鼎出库单模板（31字段）的完整流程。包括订单解析、门店匹配、货主确认、SKU匹配、字段映射等。

**Version**: 3.0 (2026-05-14)

**Author**: AI客服

**Created**: 2026-05-12

**Updated**: 2026-05-14 - 图片OCR集成到skill内部，通过agent的image工具识别

---

## 1. 功能说明

支持多种格式的订单输入，自动提取订单数据后走标准映射流程：

| 输入格式 | 处理方式 |
|----------|---------|
| Excel (.xlsx/.xls) | 直接读取，解析订单数据 |
| 图片 (.jpg/.png/.pdf截图) | OCR识别，自动提取文字 |
| PDF (.pdf) | PDF解析 + OCR识别 |
| 文字（直接粘贴） | 自然语言解析，提取关键字段 |

### 输入

- 订单文件路径或文件对象
- 或 订单图片/PDF
- 或 直接粘贴的订单文字
- 数据库连接 (db_config)

### 输出

- 华鼎出库单模板Excel文件路径

---

## 2. ⚙️ 配置参数

### 2.1 必填配置

| 配置项 | 说明 | 示例 |
|--------|------|------|
| **数据库连接 (db_config)** | PostgreSQL连接信息 | host: localhost, port: 5432, database: ai_cs_support, user: xxx |

### 2.2 可选配置

| 配置项 | 默认值 | 说明 |
|--------|--------|------|
| **输出目录** | `./output` | 生成的模板文件存放目录 |

---

## 3. 处理流程（完整版）

```
┌─────────────────────────────────────────────────────────────────┐
│ Step 1: 订单输入解析                                             │
│   ├── Excel文件 → 直接读取解析                                    │
│   ├── 图片 → OCR识别 → 提取文字                                   │
│   ├── PDF → PDF解析/OCR → 提取文字                               │
│   └── 文字粘贴 → 自然语言解析 → 提取结构化数据                    │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│ Step 2: 订单数据标准化                                            │
│   └── 提取：订单号、门店名、联系人、电话、商品列表                │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│ Step 3: 门店匹配 → 获取货主ID                                     │
│   ├── 查询store_list表                                           │
│   ├── 获取：门店编号、仓库、收货人、电话、地址                    │
│   └── ⭐ 获取owner_code（货主ID）                                │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│ Step 4: 货主确认（必须！）                                        │
│   ├── 展示匹配到的货主信息                                        │
│   ├── 询问用户确认货主是否正确                                    │
│   └── 等待用户确认后继续                                         │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│ Step 5: SKU匹配（用owner_code过滤）                              │
│   ├── 查询shipper_sku_mapping表（按货主ID过滤）                   │
│   ├── 获取系统SKU编号                                             │
│   └── 判断单位类型（大/中/小单位）                               │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│ Step 6: 生成华鼎模板                                             │
│   ├── 按31字段模板填充数据                                        │
│   └── 输出Excel文件                                               │
└─────────────────────────────────────────────────────────────────┘
```

---

## 4. 输入格式处理详解

### 4.1 Excel文件处理

```python
# 直接读取Excel
import pandas as pd
df = pd.read_excel(file_path, header=None)

# 解析结构（标准订货单格式）
# 行0: 标题 "订货单"
# 行1: 订单编号 | xxx | 下单日期 | xxx | 预交货日期 | xxx
# 行2: 公司名称 | xxx | ... | 联系人 | xxx | xxx
# 行3: 收货地址 | xxx | ... | 联系电话 | xxx
# 行4: 商品明细（标题行）
# 行5+: 商品数据
```

### 4.2 图片OCR处理

图片识别集成在skill内部，通过agent的image工具进行识别。当skill遇到图片输入时，会返回 `need_ocr=True`，由调用方（agent）使用image工具识别后，将结果回传给skill继续处理。

```python
# skill执行流程（图片输入）
result = skill.execute(order_input="/path/to/order.jpg", order_type="image")

if result.get("need_ocr"):
    # 调用方使用image工具识别图片
    ocr_prompt = """请仔细识别这张订单图片，提取以下信息并返回JSON格式：
{
    "order_no": "订单编号（如果没有则为空字符串）",
    "store_name": "门店名称或收货人姓名（用于货主匹配）",
    "contact_person": "联系人姓名",
    "phone": "联系电话",
    "address": "收货地址",
    "items": [
        {"seq": 1, "product_name": "商品名称", "spec": "规格（如有）", "quantity": 数量, "unit": "单位"}
    ]
}
重要：
- product_name请填写能用于SKU匹配的完整商品名称
- quantity必须是数字"""
    
    # 调用openclaw image工具识别（通过agent工具调用）
    # ocr_result = image_tool.analyze(image_path=result["image_path"], prompt=ocr_prompt)
    
    # 将OCR结果回传给skill继续处理
    result = skill.execute(
        order_input=result["image_path"],
        order_type="image",
        ocr_result=ocr_result  # 传入识别结果
    )
```

### 4.3 PDF处理

类似图片处理，PDF解析也返回 `need_ocr=True` 标记，由调用方使用pdf工具识别后回传。

### 4.4 文字直接解析

```python
# 解析用户粘贴的订单文字
def parse_order_text(text: str) -> dict:
    """
    输入示例：
    订单号：DH-O-20260423-278070
    门店：任丘三中
    联系人：任建华
    电话：18154355555
    商品：
    - 椰果果粒 3件
    - 蓝莓果肉果酱 1件
    ...
    """
    # 自然语言解析，提取结构化数据
    ...
```

---

## 5. SKU多层匹配规则

当 `shipper_sku_mapping` 表中无法通过客户SKU编码精确匹配时，按以下层级依次尝试：

### 第1层：精确匹配（优先）
```sql
SELECT system_sku_code, system_sku_name 
FROM shipper_sku_mapping 
WHERE shipper_id = %s AND customer_sku_code = %s
```
- 使用订单中的商品编码与 `customer_sku_code` 精确匹配
- 匹配成功后直接使用 `system_sku_code`

### 第2层：名称匹配（SKU编码未匹配时）
```sql
SELECT sku_code, sku_name, package_spec 
FROM system_sku 
WHERE sku_name LIKE %s AND (条件)
```
- 用商品名称（含关键词）在 `system_sku` 表中模糊匹配
- 可能返回多条，需结合规格进一步筛选

### 第3层：名称+规格匹配（多条匹配时）
```sql
SELECT sku_code, sku_name, package_spec 
FROM system_sku 
WHERE sku_name LIKE %s AND package_spec LIKE %s
```
- 在名称匹配基础上，用规格参数（如 `2kg*6瓶/件`、`1kg*20袋`）进一步精确筛选
- 优先选择规格完全吻合的记录

### 匹配优先级
1. 第1层命中 → 直接使用
2. 第2层命中1条 → 使用该条
3. 第2层命中多条 → 进入第3层筛选
4. 第3层命中 → 使用该条
5. 均未命中 → SKU编码留空，标记⚠️待确认

### 代码示例
```python
def match_sku(db_conn, shipper_id, customer_sku_code, product_name, spec):
    # 第1层：精确匹配
    cur.execute('''
        SELECT system_sku_code, system_sku_name 
        FROM shipper_sku_mapping 
        WHERE shipper_id = %s AND customer_sku_code = %s
    ''', (shipper_id, customer_sku_code))
    row = cur.fetchone()
    if row:
        return row['system_sku_code'], row['system_sku_name'], 'exact'
    
    # 第2层：名称匹配
    cur.execute('''
        SELECT sku_code, sku_name, package_spec 
        FROM system_sku 
        WHERE sku_name LIKE %s
    ''', (f'%{product_name}%',))
    candidates = cur.fetchall()
    
    if not candidates:
        return None, None, 'not_found'
    
    if len(candidates) == 1:
        return candidates[0]['sku_code'], candidates[0]['sku_name'], 'name_match'
    
    # 第3层：名称+规格匹配
    for c in candidates:
        if spec and c['package_spec'] and spec.split('*')[0] in c['package_spec']:
            return c['sku_code'], c['sku_name'], 'name_spec_match'
    
    # 兜底：返回第一个
    return candidates[0]['sku_code'], candidates[0]['sku_name'], 'name_fallback'
```

---

## 6. 货主-品牌对照

| 货主公司全称 | 品牌名称 | 货主ID | 说明 |
|--------------|---------|--------|------|
| 河南上黎供应链管理有限公司 | 制茶青年 | HZ2023061500002 | 制茶青年品牌门店 |
| 天津王口镇店 | - | CUSTOMER-WANGKOU | 独立货主 |
| 长沙广承供应链有限公司 | - | CUSTOMER-GUANGCHENG | 独立货主 |
| 沧州任丘长丰镇店 | - | CUSTOMER-CANGZHOU | 独立货主 |

---

## 7. 字段映射规则

### ❌ 无默认值字段（必须通过匹配/提取获取）

| 字段 | 数据来源 | 说明 |
|------|---------|------|
| 门店编号 | match_store | 从门店数据获取 |
| 仓库编码 | match_store | 从门店的warehouse字段映射 |
| 商品SKU编号 | match_sku | 必须通过SKU匹配获取（按货主ID过滤） |
| 单位类型 | match_sku | 根据ratio判断（最大=大单位） |
| 备注 | 订单提取 | 从订单备注字段提取 |
| 三方单号 | 订单提取 | 从订单编号字段获取 |
| 收货人 | match_store | 从门店数据获取 |
| 联系电话 | match_store | 从门店数据获取 |
| 收货地址 | match_store | 从门店数据获取 |
| 货主ID (owner_code) | match_store | 从门店数据获取 |

### ✅ 有默认值字段

| 字段 | 默认值 | 说明 |
|------|--------|------|
| 加急程度 | 0 | 0=普通，1=加急 |
| 指定库存状态 | "正常" | 必填 |
| 出库类型 | 201 | 201=销售订单 |
| 配送方式 | "共配" | 必填 |
| 是否垫付 | "否" | 必填 |
| 是否制定批次 | "否" | 0=否，1=是 |

### ❌ 固定为空字段

- 门店三方编码、商品三方SPEC编号、指定车型、快递公司、单价、总金额、批次号、生产日期、生产厂家编号、门店收货地址编码、业务模式、业务类型、C端快递公司

---

## 7. 单位类型判断规则

```python
# SKU配置中有多个单位时，按ratio判断：
# ratio最大的 → 大单位
# ratio中等的 → 中单位
# ratio最小的 → 小单位
```

---

## 8. 调用示例

### 8.1 初始化（不需要shipper_id）

```python
from skills.skill_order_to_huading_template import OrderToHuadingTemplate

skill = OrderToHuadingTemplate(
    db_config={
        "host": "localhost",
        "port": 5432,
        "database": "ai_cs_support",
        "user": "your_username"
    }
)
```

### 8.2 处理Excel订单

```python
result = skill.execute(order_input="/path/to/order.xlsx")
```

### 8.3 处理图片订单（两步OCR流程）

```python
# 第一次调用：触发图片OCR流程
result = skill.execute(
    order_input="/path/to/order.jpg",  # 或 .png
    order_type="image"  # 也可设为 "auto" 自动检测
)

if result.get("need_ocr"):
    # 调用方使用image工具识别图片
    ocr_prompt = """请仔细识别这张订单图片，提取以下信息并返回JSON格式：
{
    "order_no": "订单编号（如果没有则为空字符串）",
    "store_name": "门店名称或收货人姓名（用于货主匹配）",
    "contact_person": "联系人姓名",
    "phone": "联系电话",
    "address": "收货地址",
    "items": [
        {"seq": 1, "product_name": "商品名称", "spec": "规格（如有）", "quantity": 数量, "unit": "单位"}
    ]
}
重要：
- product_name请填写能用于SKU匹配的完整商品名称
- quantity必须是数字"""
    
    # 调用openclaw image工具识别（通过agent工具调用）
    # ocr_result = image_tool.analyze(image_path=result["image_path"], prompt=ocr_prompt)
    
    # 第二次调用：将OCR结果回传给skill
    result = skill.execute(
        order_input=result["image_path"],
        order_type="image",
        ocr_result=ocr_result  # 传入识别结果
    )
```

### 8.4 处理PDF订单

```python
result = skill.execute(
    order_input="/path/to/order.pdf",
    order_type="pdf"
)
```

### 8.5 处理文字输入

```python
order_text = """
订单号：DH-O-20260423-278070
门店：任丘三中
联系人：任建华
电话：18154355555
商品：
- 椰果果粒 3件
- 蓝莓果肉果酱 1件
"""

result = skill.execute(
    order_input=order_text,  # 直接传文字
    order_type="text"
)
```

### 8.6 自动识别输入类型（推荐）

```python
# 传入文件路径或文字，order_type="auto" 自动判断类型
result = skill.execute(
    order_input="/path/to/order.jpg",  # 或 .xlsx，或直接粘贴的文字
    order_type="auto"
)
```

### 8.6 返回结果

```python
{
    "success": True,
    "output_file": "/path/to/output/huading_template.xlsx",
    "order_no": "DH-O-20260423-294092",
    "store_code": "KH2024070300038",
    "store_name": "制茶青年-任丘三中",
    "owner_code": "HZ2023061500002",
    "owner_name": "河南上黎供应链管理有限公司（制茶青年）",
    "warehouse_code": "LFWW",
    "warehouse_name": "廊坊仓",
    "contact_person": "任建华",
    "contact_phone": "18154355555",
    "contact_address": "会站北道任丘市第三中学对面制茶青年",
    "item_count": 7,
    "unmatched_count": 0,
    "unmatched_items": [],
    "extracted_from": "image",  // 订单来源：excel/image/pdf/text
    "message": "模板生成成功"
}
```

---

## 9. 数据库依赖

| 表名 | 关联字段 | 用途 |
|------|---------|------|
| store_list | owner_code | 门店列表，按门店名匹配 |
| shipper_sku_mapping | shipper_id | SKU精确匹配，按货主ID过滤 |
| system_sku | sku_code | SKU名称/规格匹配（备用层） |
| warehouse_code_mapping | - | 仓库编码映射 |
| customer | customer_id | 货主信息 |

---

## 10. 错误处理

| 错误情况 | 处理方式 |
|----------|----------|
| 订单文件不存在 | 返回错误：文件不存在 |
| **图片需要OCR** | 返回 `need_ocr=True` 和 `image_path`，由调用方识别后回传 |
| **PDF需要OCR** | 返回 `need_ocr=True` 和 `pdf_path`，由调用方识别后回传 |
| 门店匹配失败 | 提示用户门店不在数据库中 |
| 货主未确认 | 暂停流程，等待用户确认 |
| SKU匹配失败 | SKU编号为空，记录unmatched_items |

---

## 11. 配置检查清单

使用前请确认以下配置已就绪：

- [ ] **数据库连接** - 必填
- [ ] **store_list表** - 必填，门店数据已导入
- [ ] **shipper_sku_mapping表** - 必填，SKU映射已导入
- [ ] **warehouse_code_mapping表** - 必填，仓库编码已配置
- [ ] **customer表** - 必填，货主信息已导入
- [ ] **OCR能力** - 图片/PDF由调用方使用image/pdf工具识别后回传