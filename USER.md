# USER.md - 用户信息

- **Name:** 华鼎云仓
- **What to call them:** 客服用户
- **Timezone:** Asia/Shanghai (UTC+8)
- **Channel:** 飞书机器人

---

## 业务背景

AI客服赋能系统 - 基于OpenClaw智能体平台

### 核心业务

| 业务 | 说明 |
|------|------|
| **订单处理** | 客户上传订单Excel → 生成华鼎31字段出库单模板 |
| **账单管理** | 合同规则提取 → 月度账单计算 → 账单审核 |

### 核心技术栈

- OpenClaw Agent Platform (port 18789)
- PostgreSQL 数据库 (`ai_cs_support`)
- 2个Skills：`skill_order_to_huading_template`、`billing-workflow`

---

## 货主-品牌对照表（重要）

| 货主公司全称 | 品牌名称 | 货主ID | 说明 |
|--------------|---------|--------|------|
| 河南上黎供应链管理有限公司 | 制茶青年 | HZ2023061500002 | 制茶青年品牌门店 |
| 天津王口镇店 | - | CUSTOMER-WANGKOU | 独立货主 |
| 长沙广承供应链有限公司 | - | CUSTOMER-GUANGCHENG | 独立货主 |
| 沧州任丘长丰镇店 | - | CUSTOMER-CANGZHOU | 独立货主 |

---

## 数据库字段对照

| 表 | 货主ID字段 | 说明 |
|---|-----------|------|
| customer | customer_id | 货主ID |
| store_list | owner_code | 门店所属货主ID |
| shipper_sku_mapping | shipper_id | SKU所属货主ID |

---

## 仓库编码

| 仓库名称 | 编码 |
|----------|------|
| 廊坊仓 | LFWW / 5 |
| 天津仓 | 8 |
| 济南仓 | 2 |
| 南京仓 | 3 |
| 郑州仓 | 4 |
| 杭州仓 | 5 |
| 西安仓 | 6 |
| 武汉仓 | 7 |
| 太原仓 | 9 |
| 长沙仓 | 10 |
| 合肥仓 | 13 |