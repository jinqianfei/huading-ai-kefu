# 意图路由配置

**版本**: v1.0  
**日期**: 2026-05-12  
**用途**: 自动识别用户意图并路由到对应 Skill

---

## 一、核心规则

### 1. 参数自动获取

| 参数 | 来源 | 说明 |
|------|------|------|
| **货主名称** | 用户输入 | 必须包含货主名称 |
| **shipper_id** | 数据库查询 | 根据货主名称从 `shipper` 表查询 |
| period | 用户输入 | 账期（如"2026年2月"、"2月"） |
| order_file | 用户上传 | 订单Excel文件 |
| contract_file | 用户上传 | 合同文件 |
| product_excel | 用户上传 | 商品信息Excel |

### 2. 货主名称查询

```python
# 用户输入"长沙广承" → 查询数据库 → 获取 shipper_id
SELECT shipper_id FROM shipper WHERE shipper_name LIKE '%长沙广承%'
```

---

## 二、意图路由表

### 意图类型 1: 订单处理

| 项目 | 内容 |
|------|------|
| **触发关键词** | 处理订单、生成模板、订单转华鼎、导出模板、生成华鼎模板、转换订单 |
| **目标 Skill** | `skill-order-to-huading-template` |
| **必填参数** | order_file, 货主名称→shipper_id |
| **可选参数** | output_file |

### 意图类型 2: 账单计算

| 项目 | 内容 |
|------|------|
| **触发关键词** | 计算账单、生成账单、制作账单、某账期账单、这个月账单 |
| **目标 Skill** | `billing-workflow` (步骤4) |
| **必填参数** | 货主名称→shipper_id, period |
| **可选参数** | 出库数据Excel |

### 意图类型 3: 合同解析

| 项目 | 内容 |
|------|------|
| **触发关键词** | 更新合同、上传合同、解析合同、提取合同规则、合同规则 |
| **目标 Skill** | `billing-workflow` (步骤2) |
| **必填参数** | contract_file, 货主名称→shipper_id |
| **可选参数** | - |

### 意图类型 4: SKU系数计算

| 项目 | 内容 |
|------|------|
| **触发关键词** | 更新SKU、更新商品系数、重新计算系数、SKU系数、商品系数 |
| **目标 Skill** | `billing-workflow` (步骤3) |
| **必填参数** | product_excel, 货主名称→shipper_id |
| **可选参数** | - |

### 意图类型 5: 账单审核

| 项目 | 内容 |
|------|------|
| **触发关键词** | 审核账单、比对账单、账单差异、账单核对 |
| **目标 Skill** | `billing-workflow` (步骤5) |
| **必填参数** | 货主名称→shipper_id, period |
| **可选参数** | 客户账单Excel |

### 意图类型 6: 全面更新

| 项目 | 内容 |
|------|------|
| **触发关键词** | 全面更新、完整流程、重新制作 |
| **目标 Skill** | `billing-workflow` (步骤2→3→4→5) |
| **必填参数** | contract_file, product_excel, 货主名称→shipper_id, period |
| **可选参数** | 出库数据Excel |

---

## 三、意图冲突处理

### 优先级规则

| 优先级 | 意图类型 | 说明 |
|--------|----------|------|
| 1 | 账单审核 | 更具体，需要比对 |
| 2 | 账单计算 | 包含period信息 |
| 3 | 合同解析 | 基础设施更新 |
| 4 | SKU系数 | 系数更新 |
| 5 | 订单处理 | 订单转换 |

### 冲突情况

| 用户输入 | 匹配 | 处理 |
|----------|------|------|
| "上传合同并计算账单" | 合同解析+账单计算 | 先执行合同解析，再执行账单计算 |
| "生成长沙广承账单" | 账单计算 | 识别货主名称="长沙广承"，period=当前月 |

### 参数缺失处理

| 缺失参数 | 处理方式 |
|----------|----------|
| 货主名称 | 询问用户："请问是哪个货主/客户？" |
| period | 默认当前月份 或 询问用户 |
| order_file | 询问用户上传订单文件 |
| contract_file | 询问用户上传合同文件 |

---

## 四、数据库依赖

### 货主表 (shipper)

```sql
SELECT shipper_id, shipper_name FROM shipper;
```

### 门店表关联 (store_list)

```sql
SELECT store_code, store_name, owner_name 
FROM store_list 
WHERE owner_name LIKE '%货主名称%';
```

### SKU映射关联 (shipper_sku_mapping)

```sql
SELECT customer_sku_name, system_sku_code 
FROM shipper_sku_mapping 
WHERE shipper_id = '查询到的shipper_id';
```

---

## 五、执行示例

### 示例 1: 订单处理

**用户输入**: "帮我处理这个订单，生成华鼎模板"

**处理流程**:
```
1. 识别意图: 订单处理
2. 提取货主名称: 从上下文/上传文件推断
3. 查询 shipper_id: SELECT shipper_id FROM shipper WHERE ...
4. 调用 Skill: skill-order-to-huading-template(order_file, shipper_id)
5. 返回结果: 华鼎模板文件
```

### 示例 2: 账单计算

**用户输入**: "生成长沙广承2026年2月账单"

**处理流程**:
```
1. 识别意图: 账单计算
2. 提取货主名称: "长沙广承"
3. 提取账期: "2026年2月"
4. 查询 shipper_id: SELECT shipper_id FROM shipper WHERE shipper_name LIKE '%长沙广承%'
5. 调用 Skill: billing-workflow(shipper_id, period="2026年2月")
6. 返回结果: 账单计算结果
```

### 示例 3: 合同解析

**用户输入**: "上传新合同，解析规则"

**处理流程**:
```
1. 识别意图: 合同解析
2. 确认上传: contract_file 已上传
3. 询问货主名称: "请问是哪个货主的合同？"
4. 用户回复: "河南上黎"
5. 查询 shipper_id: SELECT shipper_id FROM shipper WHERE shipper_name LIKE '%河南上黎%'
6. 调用 Skill: billing-workflow.step2(contract_file, shipper_id)
7. 返回结果: 解析后的规则，供确认
```

---

## 六、意图识别算法

### 关键词匹配

```python
def match_intent(user_input):
    # 精确匹配
    for intent, keywords in INTENT_KEYWORDS.items():
        for keyword in keywords:
            if keyword in user_input:
                return intent
    
    # 模糊匹配（相似度）
    best_match = None
    best_score = 0
    for intent, keywords in INTENT_KEYWORDS.items():
        for keyword in keywords:
            score = similarity(user_input, keyword)
            if score > best_score and score > 0.6:
                best_match = intent
                best_score = score
    
    return best_match
```

### 参数提取

```python
def extract_params(user_input):
    params = {}
    
    # 提取货主名称（贪婪匹配）
    shipper_names = extract_shipper_names(user_input)  # 需要实体识别
    if shipper_names:
        params['shipper_name'] = shipper_names[0]
    
    # 提取账期
    period = extract_period(user_input)
    if period:
        params['period'] = period
    
    return params
```

---

*最后更新：2026-05-12 v1.0*