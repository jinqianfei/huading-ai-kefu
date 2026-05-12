-- ============================================================
-- AI客服赋能系统 · 种子数据
-- 版本：V1.0
-- 日期：2026年4月29日
-- 说明：基于真实业务数据的测试种子数据
-- ============================================================

-- ============================================================
-- 1. 客户主数据
-- ============================================================
INSERT INTO customer (customer_id, customer_name, customer_type, contact_person, contact_phone, address, warehouse_code, warehouse_name, delivery_type, default_region, status)
VALUES
  ('CUSTOMER-WANGKOU', '天津王口镇店', '微信订单', '王静', '15133626603', '天津市静海区王口镇制茶青年', 'WH-TJ', '天津仓', 'SELF_PICKUP', '天津', 'ACTIVE'),
  ('CUSTOMER-GUANGCHENG', '长沙广承供应链有限公司', 'ERP导出', '张经理', '13800138000', '长沙市雨花区环保路88号', 'WH-CD', '成都仓', 'DELIVERY', '成都', 'ACTIVE'),
  ('CUSTOMER-CANGZHOU', '沧州任丘长丰镇店', '微信订单', '史帅男', '13900001111', '河北省沧州市任丘市长丰镇', 'WH-CZ', '沧州仓', 'SELF_PICKUP', '沧州', 'ACTIVE');

-- ============================================================
-- 2. 客户商品编码 → 系统SKU映射
-- ============================================================
INSERT INTO customer_sku_mapping (customer_id, customer_sku_code, customer_sku_name, system_sku_code, system_sku_name, unit_conversion_rule, confidence, source, status)
VALUES
  -- 天津王口镇店 SKU
  ('CUSTOMER-WANGKOU', 'P3056699182', '凤梨味果酱', 'SKU-101', '凤梨味果酱2kg', '{"大单位":"件","小单位":"瓶","比率":6}'::jsonb, 0.99, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-WANGKOU', 'P8365640612', '原味晶球', 'SKU-102', '原味晶球1kg', '{"大单位":"件","小单位":"袋","比率":20}'::jsonb, 0.99, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-WANGKOU', 'P2369520534', '速煮珍珠粉圆（黑珍珠）', 'SKU-103', '速煮珍珠粉圆1kg', '{"大单位":"件","小单位":"袋","比率":16}'::jsonb, 0.99, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-WANGKOU', 'P9264773821', '牛奶冰淇淋粉', 'SKU-104', '牛奶冰淇淋粉3kg', '{"大单位":"件","小单位":"袋","比率":7}'::jsonb, 0.98, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-WANGKOU', 'P968563703', '速冻芒果浆', 'SKU-105', '速冻芒果浆1kg', '{"大单位":"件","小单位":"瓶","比率":12}'::jsonb, 0.99, 'AUTO', 'ACTIVE'),

  -- 长沙广承供应链 SKU
  ('CUSTOMER-GUANGCHENG', 'P1556689351', '1000塑杯', 'SKU-001', '1000塑杯(一次性)', '{"大单位":"件","小单位":"个","比率":1000}'::jsonb, 0.99, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-GUANGCHENG', 'P1556689352', '500ml饮料杯', 'SKU-002', '500ml饮料杯', '{"大单位":"件","小单位":"个","比率":50}'::jsonb, 0.98, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-GUANGCHENG', 'P1556689353', '塑料吸管', 'SKU-003', '塑料吸管', '{"大单位":"包","小单位":"根","比率":100}'::jsonb, 0.87, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-GUANGCHENG', 'P1556689354', '1000塑杯(加厚)', 'SKU-004', '1000塑杯加厚版', '{"大单位":"件","小单位":"个","比率":1000}'::jsonb, 0.95, 'AUTO', 'ACTIVE'),
  ('CUSTOMER-GUANGCHENG', 'P1556689355', '瓦楞纸杯托', 'SKU-005', '瓦楞纸杯托', '{"大单位":"个","小单位":"个","比率":1}'::jsonb, 0.90, 'AUTO', 'ACTIVE');

-- ============================================================
-- 3. 门店简称 → 全称映射
-- ============================================================
INSERT INTO customer_alias_mapping (customer_id, short_name, full_name, match_type, confidence)
VALUES
  (NULL, '王口镇店', '天津王口镇店', 'FUZZY', 0.95),
  (NULL, '天津王口', '天津王口镇店', 'FUZZY', 0.92),
  ('CUSTOMER-WANGKOU', '王口镇', '天津王口镇店', 'FUZZY', 0.98),
  ('CUSTOMER-GUANGCHENG', '广承', '长沙广承供应链有限公司', 'FUZZY', 0.95),
  ('CUSTOMER-GUANGCHENG', '长沙广承', '长沙广承供应链有限公司', 'FUZZY', 0.98),
  (NULL, '任丘长丰', '沧州任丘长丰镇店', 'FUZZY', 0.93),
  (NULL, '长丰镇店', '沧州任丘长丰镇店', 'FUZZY', 0.90),
  (NULL, '成都双流店', '成都双流分店', 'EXACT', 0.99),
  (NULL, '成都武侯店', '成都武侯分店', 'EXACT', 0.99);

-- ============================================================
-- 4. 系统SKU基础数据
-- ============================================================
INSERT INTO system_sku (sku_code, sku_name, category, unit, weight_kg, fee_coefficient, warehouse_code, status)
VALUES
  ('SKU-001', '1000塑杯(一次性)', '杯子类', '件', 5.0, 1.0, 'WH-CD', 'ACTIVE'),
  ('SKU-002', '500ml饮料杯', '杯子类', '件', 3.0, 1.0, 'WH-CD', 'ACTIVE'),
  ('SKU-003', '塑料吸管', '配件类', '包', 0.5, 0.8, 'WH-CD', 'ACTIVE'),
  ('SKU-004', '1000塑杯加厚版', '杯子类', '件', 6.0, 1.2, 'WH-CD', 'ACTIVE'),
  ('SKU-005', '瓦楞纸杯托', '配件类', '个', 0.3, 0.6, 'WH-CD', 'ACTIVE'),
  ('SKU-101', '凤梨味果酱2kg', '酱料类', '件', 12.0, 1.0, 'WH-TJ', 'ACTIVE'),
  ('SKU-102', '原味晶球1kg', '小料类', '件', 20.0, 1.0, 'WH-TJ', 'ACTIVE'),
  ('SKU-103', '速煮珍珠粉圆1kg', '小料类', '件', 16.0, 1.0, 'WH-TJ', 'ACTIVE'),
  ('SKU-104', '牛奶冰淇淋粉3kg', '粉类', '件', 21.0, 1.0, 'WH-TJ', 'ACTIVE'),
  ('SKU-105', '速冻芒果浆1kg', '酱料类', '件', 12.0, 1.0, 'WH-TJ', 'ACTIVE');

-- ============================================================
-- 5. 客户合同规则（长沙广承）
-- ============================================================
INSERT INTO customer_contract_rules (customer_id, contract_no, rule_category, rule_subcategory, field_name, field_value, field_unit, confidence, status, effective_date)
VALUES
  -- 仓储费规则
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '仓储费', '成都仓', '存储费单价', '3.2', '元/托/天', 0.95, 'CONFIRMED', '2025-01-01'),
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '仓储费', '成都仓', '最低仓储费', '50', '元/月', 0.90, 'CONFIRMED', '2025-01-01'),
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '仓储费', '成都仓', '计费天数', '30', '天/月', 0.99, 'CONFIRMED', '2025-01-01'),

  -- 运输费规则
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '运输费', '成都仓', '超量单价', '0.5', '元/件', 0.95, 'CONFIRMED', '2025-01-01'),
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '运输费', '成都仓', '最低运费', '100', '元/次', 0.90, 'CONFIRMED', '2025-01-01'),

  -- 出库费规则
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '出库费', '成都仓', '出库单价', '1.0', '元/件', 0.95, 'CONFIRMED', '2025-01-01'),
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '出库费', '成都仓', '最低出库收费', '30', '元/单', 0.90, 'CONFIRMED', '2025-01-01'),

  -- 其他费用
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '其他费用', '成都仓', '管理费', '200', '元/月', 0.85, 'CONFIRMED', '2025-01-01'),
  ('CUSTOMER-GUANGCHENG', 'HDLL-CP-2025-286', '其他费用', '成都仓', '服务费', '100', '元/月', 0.85, 'CONFIRMED', '2025-01-01');

-- ============================================================
-- 6. 天津王口镇店合同规则（简化版）
-- ============================================================
INSERT INTO customer_contract_rules (customer_id, contract_no, rule_category, rule_subcategory, field_name, field_value, field_unit, confidence, status, effective_date)
VALUES
  ('CUSTOMER-WANGKOU', 'HDLL-CP-2025-100', '仓储费', '天津仓', '存储费单价', '2.5', '元/托/天', 0.95, 'CONFIRMED', '2025-01-01'),
  ('CUSTOMER-WANGKOU', 'HDLL-CP-2025-100', '出库费', '天津仓', '出库单价', '0.8', '元/件', 0.95, 'CONFIRMED', '2025-01-01');

-- ============================================================
-- 7. 客户仓库映射
-- ============================================================
INSERT INTO customer_warehouse_mapping (customer_id, warehouse_code, warehouse_name, delivery_type, default_region, priority, enabled)
VALUES
  ('CUSTOMER-WANGKOU', 'WH-TJ', '天津仓', 'SELF_PICKUP', '天津', 1, true),
  ('CUSTOMER-GUANGCHENG', 'WH-CD', '成都仓', 'DELIVERY', '成都', 1, true),
  ('CUSTOMER-CANGZHOU', 'WH-CZ', '沧州仓', 'SELF_PICKUP', '沧州', 1, true);

-- ============================================================
-- 8. 订单模板配置
-- ============================================================
INSERT INTO order_template_config (customer_id, template_name, template_type, file_pattern, enabled)
VALUES
  ('CUSTOMER-WANGKOU', '微信订单模板', '微信文字', '%王口镇店%', true),
  ('CUSTOMER-GUANGCHENG', 'ERP导出模板', 'ERP导出', '%广承供应链%', true),
  ('CUSTOMER-CANGZHOU', '微信订单模板', '微信文字', '%任丘%', true);

-- ============================================================
-- 验证插入数据
-- ============================================================
SELECT '客户数据' as table_name, COUNT(*) as count FROM customer
UNION ALL
SELECT 'SKU映射', COUNT(*) FROM customer_sku_mapping
UNION ALL
SELECT '门店别名', COUNT(*) FROM customer_alias_mapping
UNION ALL
SELECT '系统SKU', COUNT(*) FROM system_sku
UNION ALL
SELECT '合同规则', COUNT(*) FROM customer_contract_rules
UNION ALL
SELECT '仓库映射', COUNT(*) FROM customer_warehouse_mapping
UNION ALL
SELECT '订单模板', COUNT(*) FROM order_template_config;
