-- 商品信息导入
INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241120000107',
    '柠檬（常温）',
    '启用',
    'GG241120000106',
    NULL,
    '柠檬（常温）',
    '15kg/箱',
    '启用',
    'SK241120000197',
    NULL,
    '箱',
    1.0,
    '启用',
    '12',
    '月',
    '保鲜',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2026-05-06 10:00:36',
    '2024-11-20 17:16:58'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260506000013',
    '速冻芭乐浆',
    '启用',
    'GG260506000013',
    NULL,
    '速冻芭乐浆',
    '1kg*12瓶/件',
    '启用',
    'SK260506000023',
    NULL,
    '件',
    12.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-05-06 09:53:36',
    '2026-05-06 09:53:36'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260506000013',
    '速冻芭乐浆',
    '启用',
    'GG260506000013',
    NULL,
    '速冻芭乐浆',
    '瓶',
    '启用',
    'SK260506000022',
    NULL,
    '瓶',
    1.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-05-06 09:53:36',
    '2026-05-06 09:53:36'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260506000009',
    '120口径封口膜（4卷装）',
    '启用',
    'GG260506000009',
    NULL,
    '120口径封口膜（4卷装）',
    '4卷/件',
    '启用',
    'SK260506000016',
    NULL,
    '件',
    4.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-05-06 09:45:26',
    '2026-05-06 09:45:26'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260506000009',
    '120口径封口膜（4卷装）',
    '启用',
    'GG260506000009',
    NULL,
    '120口径封口膜（4卷装）',
    '卷',
    '启用',
    'SK260506000015',
    NULL,
    '卷',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-05-06 09:45:26',
    '2026-05-06 09:45:26'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260506000008',
    '90口径封口膜（4卷装）',
    '启用',
    'GG260506000008',
    NULL,
    '90口径封口膜（4卷装）',
    '4卷/件',
    '启用',
    'SK260506000014',
    NULL,
    '件',
    4.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-05-06 09:44:55',
    '2026-05-06 09:44:55'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260506000008',
    '90口径封口膜（4卷装）',
    '启用',
    'GG260506000008',
    NULL,
    '90口径封口膜（4卷装）',
    '卷',
    '启用',
    'SK260506000013',
    NULL,
    '卷',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-05-06 09:44:55',
    '2026-05-06 09:44:55'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260410000037',
    '常温五合一芝士奶盖',
    '启用',
    'GG260410000037',
    NULL,
    '常温五合一芝士奶盖',
    '1kg*12盒/件',
    '启用',
    'SK260410000066',
    NULL,
    '件',
    12.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-04-10 10:05:20',
    '2026-04-10 10:01:12'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260410000037',
    '常温五合一芝士奶盖',
    '启用',
    'GG260410000037',
    NULL,
    '常温五合一芝士奶盖',
    '盒',
    '启用',
    'SK260410000065',
    NULL,
    '盒',
    1.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-04-10 10:05:20',
    '2026-04-10 10:01:12'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260410000038',
    '奶冻粉蛋白固体饮料',
    '启用',
    'GG260410000038',
    NULL,
    '奶冻粉蛋白固体饮料',
    '500克*16袋/箱',
    '启用',
    'SK260410000068',
    NULL,
    '箱',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-04-10 10:01:44',
    '2026-04-10 10:01:44'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260410000038',
    '奶冻粉蛋白固体饮料',
    '启用',
    'GG260410000038',
    NULL,
    '奶冻粉蛋白固体饮料',
    '袋',
    '启用',
    'SK260410000067',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-04-10 10:01:44',
    '2026-04-10 10:01:44'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240826000012',
    '三孔吸管',
    '启用',
    'GG240826000012',
    NULL,
    '三孔吸管',
    '100支*10包/件',
    '启用',
    'SK240826000018',
    NULL,
    '件',
    10.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-30 11:47:49',
    '2024-08-26 09:11:55'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240826000012',
    '三孔吸管',
    '启用',
    'GG240826000012',
    NULL,
    '三孔吸管',
    '包',
    '启用',
    'SK240826000017',
    NULL,
    '包',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-30 11:47:49',
    '2024-08-26 09:11:55'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260329000053',
    '可可饼干碎（500g*16袋）',
    '启用',
    'GG260329000053',
    NULL,
    '可可饼干碎（500g*16袋）',
    '500g*16袋/件',
    '启用',
    'SK260329000070',
    NULL,
    '件',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-29 16:34:53',
    '2026-03-29 16:34:53'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260329000053',
    '可可饼干碎（500g*16袋）',
    '启用',
    'GG260329000053',
    NULL,
    '可可饼干碎（500g*16袋）',
    '袋',
    '启用',
    'SK260329000069',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-29 16:34:53',
    '2026-03-29 16:34:53'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260326000026',
    '冷冻蜜瓜浆',
    '启用',
    'GG260326000026',
    NULL,
    '冷冻蜜瓜浆',
    '1kg*12瓶/件',
    '启用',
    'SK260326000033',
    NULL,
    '件',
    12.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-26 09:39:47',
    '2026-03-26 09:39:47'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260326000026',
    '冷冻蜜瓜浆',
    '启用',
    'GG260326000026',
    NULL,
    '冷冻蜜瓜浆',
    '瓶',
    '启用',
    'SK260326000032',
    NULL,
    '瓶',
    1.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-26 09:39:47',
    '2026-03-26 09:39:47'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260326000025',
    '冷冻草莓浆',
    '启用',
    'GG260326000025',
    NULL,
    '冷冻草莓浆',
    '1kg*12瓶/件',
    '启用',
    'SK260326000031',
    NULL,
    '件',
    12.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-26 09:39:10',
    '2026-03-26 09:39:10'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260326000025',
    '冷冻草莓浆',
    '启用',
    'GG260326000025',
    NULL,
    '冷冻草莓浆',
    '瓶',
    '启用',
    'SK260326000030',
    NULL,
    '瓶',
    1.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-26 09:39:10',
    '2026-03-26 09:39:10'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260319000046',
    '玫珑瓜纸杯',
    '启用',
    'GG260319000045',
    NULL,
    '玫珑瓜纸杯',
    '250个/件',
    '启用',
    'SK260319000080',
    NULL,
    '件',
    250.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-19 11:27:39',
    '2026-03-19 11:27:39'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260319000046',
    '玫珑瓜纸杯',
    '启用',
    'GG260319000045',
    NULL,
    '玫珑瓜纸杯',
    '个',
    '启用',
    'SK260319000079',
    NULL,
    '个',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-19 11:27:39',
    '2026-03-19 11:27:39'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260318000049',
    '玫珑瓜冰淇淋粉',
    '启用',
    'GG260318000049',
    NULL,
    '玫珑瓜冰淇淋粉',
    '3kg*6袋/箱',
    '启用',
    'SK260318000059',
    NULL,
    '箱',
    6.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-18 10:47:40',
    '2026-03-18 10:47:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260318000049',
    '玫珑瓜冰淇淋粉',
    '启用',
    'GG260318000049',
    NULL,
    '玫珑瓜冰淇淋粉',
    '3kg/袋',
    '启用',
    'SK260318000058',
    NULL,
    '袋',
    1.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-03-18 10:47:40',
    '2026-03-18 10:47:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260225000034',
    '马蹄爆爆珠(850g*12罐)',
    '启用',
    'GG260225000034',
    NULL,
    '马蹄爆爆珠(850g*12罐)',
    '850g*12罐/件',
    '启用',
    'SK260225000045',
    NULL,
    '件',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-02-25 12:07:42',
    '2026-02-25 12:07:42'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260225000034',
    '马蹄爆爆珠(850g*12罐)',
    '启用',
    'GG260225000034',
    NULL,
    '马蹄爆爆珠(850g*12罐)',
    '850g/罐',
    '启用',
    'SK260225000044',
    NULL,
    '罐',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-02-25 12:07:42',
    '2026-02-25 12:07:42'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260224000040',
    '圣代勺（1000支）',
    '启用',
    'GG260224000039',
    NULL,
    '圣代勺（1000支）',
    '100个*10包/件',
    '启用',
    'SK260224000073',
    NULL,
    '件',
    10.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-02-24 15:06:21',
    '2026-02-24 15:06:21'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260224000040',
    '圣代勺（1000支）',
    '启用',
    'GG260224000039',
    NULL,
    '圣代勺（1000支）',
    '100个/包',
    '启用',
    'SK260224000072',
    NULL,
    '包',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-02-24 15:06:21',
    '2026-02-24 15:06:21'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260120000037',
    '新年旋转主题-500纸杯',
    '启用',
    'GG260120000037',
    NULL,
    '新年旋转主题-500纸杯',
    '500个/件',
    '启用',
    'SK260120000062',
    NULL,
    '件',
    500.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-01-20 11:26:36',
    '2026-01-20 11:26:11'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP260120000037',
    '新年旋转主题-500纸杯',
    '启用',
    'GG260120000037',
    NULL,
    '新年旋转主题-500纸杯',
    '500个',
    '启用',
    'SK260120000061',
    NULL,
    '个',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2026-01-20 11:26:36',
    '2026-01-20 11:26:11'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251208000055',
    '红西柚颗粒罐头（520g*24罐）',
    '启用',
    'GG251208000054',
    NULL,
    '红西柚颗粒罐头（520g*24罐）',
    '520g*24罐/箱',
    '启用',
    'SK251208000081',
    NULL,
    '箱',
    24.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-12-08 11:16:40',
    '2025-12-08 11:16:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251208000055',
    '红西柚颗粒罐头（520g*24罐）',
    '启用',
    'GG251208000054',
    NULL,
    '红西柚颗粒罐头（520g*24罐）',
    '520g/罐',
    '启用',
    'SK251208000080',
    NULL,
    '罐',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-12-08 11:16:40',
    '2025-12-08 11:16:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251127000168',
    '轻乳茶大杯（纸杯）',
    '启用',
    'GG251127000168',
    NULL,
    '轻乳茶大杯（纸杯）',
    '20只*25条/件',
    '启用',
    'SK251127000312',
    NULL,
    '件',
    25.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-11-27 15:41:20',
    '2025-11-27 15:41:20'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251127000168',
    '轻乳茶大杯（纸杯）',
    '启用',
    'GG251127000168',
    NULL,
    '轻乳茶大杯（纸杯）',
    '20只/条',
    '启用',
    'SK251127000311',
    NULL,
    '条',
    1.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-11-27 15:41:20',
    '2025-11-27 15:41:20'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251110000021',
    '牛乳茶纸杯',
    '启用',
    'GG251110000021',
    NULL,
    '牛乳茶纸杯',
    '25只*20条/箱',
    '启用',
    'SK251110000035',
    NULL,
    '箱',
    20.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-11-10 09:44:25',
    '2025-11-10 09:44:25'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251110000021',
    '牛乳茶纸杯',
    '启用',
    'GG251110000021',
    NULL,
    '牛乳茶纸杯',
    '25只/条',
    '启用',
    'SK251110000034',
    NULL,
    '条',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-11-10 09:44:25',
    '2025-11-10 09:44:25'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251105000336',
    '速煮珍珠粉圆（黑珍珠）',
    '启用',
    'GG251105000335',
    NULL,
    '速煮珍珠粉圆（黑珍珠）',
    '1kg*16袋/件',
    '启用',
    'SK251105000474',
    NULL,
    '件',
    16.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-11-05 17:09:39',
    '2025-11-05 16:55:30'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251105000336',
    '速煮珍珠粉圆（黑珍珠）',
    '启用',
    'GG251105000335',
    NULL,
    '速煮珍珠粉圆（黑珍珠）',
    '1kg/袋',
    '启用',
    'SK251105000473',
    NULL,
    '袋',
    1.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-11-05 17:09:39',
    '2025-11-05 16:55:30'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251025001295',
    '制茶青年甄选老红糖',
    '启用',
    'GG251025001294',
    NULL,
    '制茶青年甄选老红糖',
    '1kg*16袋/件',
    '启用',
    'SK251025001546',
    NULL,
    '件',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-10-25 16:07:05',
    '2025-10-25 16:07:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251025001295',
    '制茶青年甄选老红糖',
    '启用',
    'GG251025001294',
    NULL,
    '制茶青年甄选老红糖',
    '1kg/袋',
    '启用',
    'SK251025001545',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-10-25 16:07:05',
    '2025-10-25 16:07:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251021000093',
    '呱呱蛙糯米蛋黄锅巴(藤椒麻辣味)',
    '启用',
    'GG251021000092',
    NULL,
    '呱呱蛙糯米蛋黄锅巴(藤椒麻辣味)',
    '40g*100袋/件',
    '启用',
    'SK251021000148',
    NULL,
    '件',
    100.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-10-21 14:50:48',
    '2025-10-21 14:50:48'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251021000093',
    '呱呱蛙糯米蛋黄锅巴(藤椒麻辣味)',
    '启用',
    'GG251021000092',
    NULL,
    '呱呱蛙糯米蛋黄锅巴(藤椒麻辣味)',
    '40g/袋',
    '启用',
    'SK251021000147',
    NULL,
    '袋',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-10-21 14:50:48',
    '2025-10-21 14:50:48'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251021000091',
    '呱呱蛙糯米蛋黄锅巴(海苔肉松味)',
    '启用',
    'GG251021000090',
    NULL,
    '呱呱蛙糯米蛋黄锅巴(海苔肉松味)',
    '40g*100袋/件',
    '启用',
    'SK251021000144',
    NULL,
    '件',
    100.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-10-21 14:49:50',
    '2025-10-21 14:49:50'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251021000091',
    '呱呱蛙糯米蛋黄锅巴(海苔肉松味)',
    '启用',
    'GG251021000090',
    NULL,
    '呱呱蛙糯米蛋黄锅巴(海苔肉松味)',
    '40g/袋',
    '启用',
    'SK251021000143',
    NULL,
    '袋',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-10-21 14:49:50',
    '2025-10-21 14:49:50'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251015000101',
    '外卖杯托',
    '启用',
    'GG251015000098',
    NULL,
    '外卖杯托',
    '300个/套',
    '启用',
    'SK251015000174',
    NULL,
    '套',
    300.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-15 14:19:23',
    '2025-10-15 14:19:23'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251015000101',
    '外卖杯托',
    '启用',
    'GG251015000098',
    NULL,
    '外卖杯托',
    '个',
    '启用',
    'SK251015000173',
    NULL,
    '个',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-15 14:19:23',
    '2025-10-15 14:19:23'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251013000123',
    '大杯纸杯(茉莉小白专用)',
    '启用',
    'GG251013000120',
    NULL,
    '大杯纸杯(茉莉小白专用)',
    '500个/件',
    '启用',
    'SK251013000197',
    NULL,
    '件',
    500.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-13 17:50:52',
    '2025-10-13 17:50:52'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251013000123',
    '大杯纸杯(茉莉小白专用)',
    '启用',
    'GG251013000120',
    NULL,
    '大杯纸杯(茉莉小白专用)',
    '个',
    '启用',
    'SK251013000196',
    NULL,
    '个',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-13 17:50:52',
    '2025-10-13 17:50:52'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251013000122',
    '制茶基底乳',
    '启用',
    'GG251013000119',
    NULL,
    '制茶基底乳',
    '1L*12盒/箱',
    '启用',
    'SK251013000195',
    NULL,
    '箱',
    12.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-13 17:48:04',
    '2025-10-13 17:48:04'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251013000122',
    '制茶基底乳',
    '启用',
    'GG251013000119',
    NULL,
    '制茶基底乳',
    '1L/盒',
    '启用',
    'SK251013000194',
    NULL,
    '盒',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-13 17:48:04',
    '2025-10-13 17:48:04'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251013000117',
    '速冻桑葚/新',
    '启用',
    'GG251013000114',
    NULL,
    '速冻桑葚/新',
    '1kg*10袋/件',
    '启用',
    'SK251013000189',
    NULL,
    '件',
    10.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-13 17:34:46',
    '2025-10-13 17:34:46'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP251013000117',
    '速冻桑葚/新',
    '启用',
    'GG251013000114',
    NULL,
    '速冻桑葚/新',
    '1kg/袋',
    '启用',
    'SK251013000188',
    NULL,
    '袋',
    1.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-10-13 17:34:46',
    '2025-10-13 17:34:46'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250920001595',
    '草莓果肉饮料',
    '启用',
    'GG250920001595',
    NULL,
    '草莓果肉饮料',
    '1kg*8袋/件',
    '启用',
    'SK250920002581',
    NULL,
    '件',
    8.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-22 14:44:15',
    '2025-09-20 17:05:43'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250920001595',
    '草莓果肉饮料',
    '启用',
    'GG250920001595',
    NULL,
    '草莓果肉饮料',
    '1kg/袋',
    '启用',
    'SK250920002580',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-22 14:44:15',
    '2025-09-20 17:05:43'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250921000019',
    '石榴鲜果',
    '启用',
    'GG250921000019',
    NULL,
    '石榴鲜果',
    '20公斤/框',
    '启用',
    'SK250921000029',
    NULL,
    '框',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-21 11:54:27',
    '2025-09-21 11:53:03'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250918000016',
    '石榴果酱',
    '启用',
    'GG250918000016',
    NULL,
    '石榴果酱',
    '1.2kg*12罐/箱',
    '启用',
    'SK250918000028',
    NULL,
    '箱',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-18 10:20:30',
    '2025-09-18 10:20:30'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250918000016',
    '石榴果酱',
    '启用',
    'GG250918000016',
    NULL,
    '石榴果酱',
    '1.2kg/罐',
    '启用',
    'SK250918000027',
    NULL,
    '罐',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-18 10:20:30',
    '2025-09-18 10:20:30'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250911000243',
    '外卖打包四杯袋',
    '启用',
    'GG250911000243',
    NULL,
    '外卖打包四杯袋',
    '50个*6扎/件',
    '启用',
    'SK250911000445',
    NULL,
    '件',
    6.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-11 12:11:43',
    '2025-09-11 12:11:43'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250911000243',
    '外卖打包四杯袋',
    '启用',
    'GG250911000243',
    NULL,
    '外卖打包四杯袋',
    '50个/扎',
    '启用',
    'SK250911000444',
    NULL,
    '扎',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-11 12:11:43',
    '2025-09-11 12:11:43'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250911000242',
    '外卖打包双杯袋',
    '启用',
    'GG250911000242',
    NULL,
    '外卖打包双杯袋',
    '50个*6扎/件',
    '启用',
    'SK250911000443',
    NULL,
    '件',
    6.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-11 12:09:57',
    '2025-09-11 12:09:57'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250911000242',
    '外卖打包双杯袋',
    '启用',
    'GG250911000242',
    NULL,
    '外卖打包双杯袋',
    '50个/扎',
    '启用',
    'SK250911000442',
    NULL,
    '扎',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-11 12:09:57',
    '2025-09-11 12:09:57'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250911000241',
    '外卖打包单杯袋',
    '启用',
    'GG250911000241',
    NULL,
    '外卖打包单杯袋',
    '50个*6扎/件',
    '启用',
    'SK250911000441',
    NULL,
    '件',
    6.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-11 12:08:14',
    '2025-09-11 12:08:14'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250911000241',
    '外卖打包单杯袋',
    '启用',
    'GG250911000241',
    NULL,
    '外卖打包单杯袋',
    '50个/扎',
    '启用',
    'SK250911000440',
    NULL,
    '扎',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-09-11 12:08:14',
    '2025-09-11 12:08:14'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077269',
    '蛋奶布丁粉',
    '启用',
    'SP2023070400031',
    NULL,
    '蛋奶布丁粉',
    '1kg/袋*12袋/箱',
    '启用',
    'OWL2023070400029',
    NULL,
    '箱',
    12.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2025-08-29 09:13:20',
    '2023-07-04 11:30:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077269',
    '蛋奶布丁粉',
    '启用',
    'SP2023070400031',
    NULL,
    '蛋奶布丁粉',
    '1kg',
    '启用',
    'OSWL2023070400020',
    NULL,
    '袋',
    1.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2025-08-29 09:13:20',
    '2023-07-04 11:30:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250819000003',
    '制茶青年高山花韵红茶',
    '启用',
    'GG250819000003',
    NULL,
    '制茶青年高山花韵红茶',
    '100g*50袋/件',
    '启用',
    'SK250819000017',
    NULL,
    '件',
    50.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-08-19 09:07:19',
    '2025-08-19 09:07:19'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250819000003',
    '制茶青年高山花韵红茶',
    '启用',
    'GG250819000003',
    NULL,
    '制茶青年高山花韵红茶',
    '100g/袋',
    '启用',
    'SK250819000016',
    NULL,
    '袋',
    1.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2025-08-19 09:07:19',
    '2025-08-19 09:07:19'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250723000184',
    '大桶杯（盖子+提手圈+吸管）',
    '启用',
    'GG250723000179',
    NULL,
    '大桶杯（盖子+提手圈+吸管）',
    '300套/箱',
    '启用',
    'SK250723000324',
    NULL,
    '箱',
    300.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-23 20:07:05',
    '2025-07-23 20:07:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250723000184',
    '大桶杯（盖子+提手圈+吸管）',
    '启用',
    'GG250723000179',
    NULL,
    '大桶杯（盖子+提手圈+吸管）',
    '套',
    '启用',
    'SK250723000323',
    NULL,
    '套',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-23 20:07:05',
    '2025-07-23 20:07:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250723000182',
    '大桶杯（98制茶青年）',
    '启用',
    'GG250723000177',
    NULL,
    '大桶杯（98制茶青年）',
    '25*12条/箱',
    '启用',
    'SK250723000317',
    NULL,
    '箱',
    12.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-23 19:18:28',
    '2025-07-23 19:18:28'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250723000182',
    '大桶杯（98制茶青年）',
    '启用',
    'GG250723000177',
    NULL,
    '大桶杯（98制茶青年）',
    '25/条',
    '启用',
    'SK250723000316',
    NULL,
    '条',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-23 19:18:28',
    '2025-07-23 19:18:28'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250717000106',
    '500纸杯（联名款）',
    '启用',
    'GG250717000104',
    NULL,
    '500纸杯（联名款）',
    '25只*20条/箱',
    '启用',
    'SK250717000180',
    NULL,
    '箱',
    20.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-17 18:24:37',
    '2025-07-17 18:24:37'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250717000106',
    '500纸杯（联名款）',
    '启用',
    'GG250717000104',
    NULL,
    '500纸杯（联名款）',
    '条',
    '启用',
    'SK250717000179',
    NULL,
    '条',
    1.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-17 18:24:37',
    '2025-07-17 18:24:37'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250716000003',
    '面筋片（零食）',
    '启用',
    'GG250716000003',
    NULL,
    '面筋片（零食）',
    '26g*5包*4袋/件',
    '启用',
    'SK250716000004',
    NULL,
    '件',
    4.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-16 09:06:27',
    '2025-07-16 09:06:27'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250716000003',
    '面筋片（零食）',
    '启用',
    'GG250716000003',
    NULL,
    '面筋片（零食）',
    '袋',
    '启用',
    'SK250716000003',
    NULL,
    '袋',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-16 09:06:27',
    '2025-07-16 09:06:27'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250715000139',
    '500磨砂塑杯（联名呱呱）',
    '启用',
    'GG250715000139',
    NULL,
    '500磨砂塑杯（联名呱呱）',
    '25条/箱',
    '启用',
    'SK250715000228',
    NULL,
    '箱',
    25.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-15 15:50:06',
    '2025-07-15 15:50:06'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250715000139',
    '500磨砂塑杯（联名呱呱）',
    '启用',
    'GG250715000139',
    NULL,
    '500磨砂塑杯（联名呱呱）',
    '条',
    '启用',
    'SK250715000227',
    NULL,
    '条',
    1.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-15 15:50:06',
    '2025-07-15 15:50:06'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250715000138',
    '制茶青年700磨砂塑杯（联名呱呱）',
    '启用',
    'GG250715000138',
    NULL,
    '制茶青年700磨砂塑杯（联名呱呱）',
    '20袋*25个/箱',
    '启用',
    'SK250715000226',
    NULL,
    '箱',
    25.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-15 15:48:51',
    '2025-07-15 15:48:51'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250715000138',
    '制茶青年700磨砂塑杯（联名呱呱）',
    '启用',
    'GG250715000138',
    NULL,
    '制茶青年700磨砂塑杯（联名呱呱）',
    '20袋/个',
    '启用',
    'SK250715000225',
    NULL,
    '个',
    1.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-15 15:48:51',
    '2025-07-15 15:48:51'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250702000104',
    '薄荷糖浆',
    '启用',
    'GG250702000104',
    NULL,
    '薄荷糖浆',
    '1.2kg*4桶/件',
    '启用',
    'SK250702000122',
    NULL,
    '件',
    4.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-02 09:12:16',
    '2025-07-02 09:12:16'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250702000104',
    '薄荷糖浆',
    '启用',
    'GG250702000104',
    NULL,
    '薄荷糖浆',
    '桶',
    '启用',
    'SK250702000121',
    NULL,
    '桶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-07-02 09:12:16',
    '2025-07-02 09:12:16'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250623000102',
    '春香茉莉花茶/新',
    '启用',
    'GG250623000101',
    NULL,
    '春香茉莉花茶/新',
    '100g*50袋/件',
    '启用',
    'SK250623000189',
    NULL,
    '件',
    50.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-06-23 18:19:18',
    '2025-06-23 18:19:18'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250623000102',
    '春香茉莉花茶/新',
    '启用',
    'GG250623000101',
    NULL,
    '春香茉莉花茶/新',
    '100g/袋',
    '启用',
    'SK250623000188',
    NULL,
    '袋',
    1.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-06-23 18:19:18',
    '2025-06-23 18:19:18'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250518000047',
    '红糖珍珠粉圆',
    '启用',
    'GG250518000047',
    NULL,
    '制茶青年-红糖珍珠粉圆',
    '1kg*20袋/件',
    '启用',
    'SK250518000069',
    NULL,
    '件',
    20.0,
    '启用',
    '12',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-05-19 10:42:04',
    '2025-05-18 19:19:17'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250518000047',
    '红糖珍珠粉圆',
    '启用',
    'GG250518000047',
    NULL,
    '制茶青年-红糖珍珠粉圆',
    '1kg/袋',
    '启用',
    'SK250518000068',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-05-19 10:42:04',
    '2025-05-18 19:19:17'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250511000044',
    '桑葚果酱/新',
    '启用',
    'GG250511000043',
    NULL,
    '桑葚果酱/新',
    '1kg*16袋/件',
    '启用',
    'SK250511000077',
    NULL,
    '件',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-05-11 17:50:06',
    '2025-05-11 17:50:06'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250511000044',
    '桑葚果酱/新',
    '启用',
    'GG250511000043',
    NULL,
    '桑葚果酱/新',
    '1kg/袋',
    '启用',
    'SK250511000076',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-05-11 17:50:06',
    '2025-05-11 17:50:06'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250510000015',
    '速冻桑葚',
    '启用',
    'GG250510000015',
    NULL,
    '速冻桑葚',
    '1KG*20袋/件',
    '启用',
    'SK250510000026',
    NULL,
    '件',
    20.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-05-10 10:06:15',
    '2025-05-10 09:52:12'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250510000015',
    '速冻桑葚',
    '启用',
    'GG250510000015',
    NULL,
    '速冻桑葚',
    '1KG/袋',
    '启用',
    'SK250510000025',
    NULL,
    '袋',
    1.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-05-10 10:06:15',
    '2025-05-10 09:52:12'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250408000062',
    '500ml小胖杯（含500ml小胖杯盖）',
    '启用',
    'GG250408000062',
    NULL,
    '500ml小胖杯（含500ml小胖杯盖）',
    '500套（杯子，盖子）/件',
    '启用',
    'SK250408000106',
    NULL,
    '件',
    500.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-08 12:44:54',
    '2025-04-08 12:44:54'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250408000062',
    '500ml小胖杯（含500ml小胖杯盖）',
    '启用',
    'GG250408000062',
    NULL,
    '500ml小胖杯（含500ml小胖杯盖）',
    '套',
    '启用',
    'SK250408000105',
    NULL,
    '套',
    1.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-08 12:44:54',
    '2025-04-08 12:44:54'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250404000024',
    '草莓果酱',
    '启用',
    'GG250404000023',
    NULL,
    '草莓果酱',
    '1kg*16袋/件',
    '启用',
    'SK250404000032',
    NULL,
    '件',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-05 13:38:52',
    '2025-04-04 12:15:42'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250404000024',
    '草莓果酱',
    '启用',
    'GG250404000023',
    NULL,
    '草莓果酱',
    '袋',
    '启用',
    'SK250404000031',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-05 13:38:52',
    '2025-04-04 12:15:42'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250405000007',
    '500膜内贴塑杯（粉色）',
    '启用',
    'GG250405000007',
    NULL,
    '500膜内贴塑杯（粉色）',
    '25个*20条/件',
    '启用',
    'SK250405000010',
    NULL,
    '件',
    20.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-05 10:59:49',
    '2025-04-05 10:59:49'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250405000007',
    '500膜内贴塑杯（粉色）',
    '启用',
    'GG250405000007',
    NULL,
    '500膜内贴塑杯（粉色）',
    '25个/条',
    '启用',
    'SK250405000009',
    NULL,
    '条',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-05 10:59:49',
    '2025-04-05 10:59:49'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250405000006',
    '90防漏盖（粉色）',
    '启用',
    'GG250405000006',
    NULL,
    '90防漏盖（粉色）',
    '20个*25条/箱',
    '启用',
    'SK250405000008',
    NULL,
    '箱',
    25.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-05 10:57:40',
    '2025-04-05 10:57:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250405000006',
    '90防漏盖（粉色）',
    '启用',
    'GG250405000006',
    NULL,
    '90防漏盖（粉色）',
    '20个/条',
    '启用',
    'SK250405000007',
    NULL,
    '条',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-05 10:57:40',
    '2025-04-05 10:57:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250404000023',
    '白桃果酱',
    '启用',
    'GG250404000022',
    NULL,
    '白桃果酱',
    '1kg*16袋/件',
    '启用',
    'SK250404000030',
    NULL,
    '件',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-04 12:15:56',
    '2025-04-04 12:14:43'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250404000023',
    '白桃果酱',
    '启用',
    'GG250404000022',
    NULL,
    '白桃果酱',
    '1kg/袋',
    '启用',
    'SK250404000029',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-04 12:15:56',
    '2025-04-04 12:14:43'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250402000155',
    '制茶青年白桃乌龙味龙茶',
    '启用',
    'GG250402000154',
    NULL,
    '制茶青年白桃乌龙味龙茶',
    '100g*50包/件',
    '启用',
    'SK250402000287',
    NULL,
    '件',
    50.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-02 14:10:21',
    '2025-04-02 14:10:21'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250402000155',
    '制茶青年白桃乌龙味龙茶',
    '启用',
    'GG250402000154',
    NULL,
    '制茶青年白桃乌龙味龙茶',
    '100g/包',
    '启用',
    'SK250402000286',
    NULL,
    '包',
    1.0,
    '启用',
    '2',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-04-02 14:10:21',
    '2025-04-02 14:10:21'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250324000153',
    '意式咖啡豆',
    '启用',
    'GG250324000153',
    NULL,
    '意式咖啡豆',
    '500g*20袋/件',
    '启用',
    'SK250324000296',
    NULL,
    '件',
    20.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-03-24 11:41:38',
    '2025-03-24 11:41:38'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250324000153',
    '意式咖啡豆',
    '启用',
    'GG250324000153',
    NULL,
    '意式咖啡豆',
    '500g/袋',
    '启用',
    'SK250324000295',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-03-24 11:41:38',
    '2025-03-24 11:41:38'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250226000078',
    '制茶青年甄选纯牛奶',
    '启用',
    'GG250226000078',
    NULL,
    '制茶青年甄选纯牛奶',
    '1L*12盒/件',
    '启用',
    'SK250226000140',
    NULL,
    '件',
    12.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-02-26 14:57:31',
    '2025-02-26 14:57:31'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP250226000078',
    '制茶青年甄选纯牛奶',
    '启用',
    'GG250226000078',
    NULL,
    '制茶青年甄选纯牛奶',
    '盒',
    '启用',
    'SK250226000139',
    NULL,
    '盒',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-02-26 14:57:31',
    '2025-02-26 14:57:31'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230803000049',
    '黄柠檬',
    '启用',
    'GG230803000049',
    NULL,
    '黄柠檬',
    '15kg/箱',
    '启用',
    'SK230803000083',
    NULL,
    '箱',
    1.0,
    '启用',
    '12',
    '月',
    '保鲜',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2025-02-06 14:33:17',
    '2023-08-03 17:15:11'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241208000018',
    '制茶青年700磨砂塑杯（专属）',
    '启用',
    'GG241208000018',
    NULL,
    '制茶青年700磨砂塑杯（专属）',
    '25个*20袋/件',
    '启用',
    'SK241208000036',
    NULL,
    '件',
    20.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-08 10:03:57',
    '2024-12-08 10:02:22'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241208000018',
    '制茶青年700磨砂塑杯（专属）',
    '启用',
    'GG241208000018',
    NULL,
    '制茶青年700磨砂塑杯（专属）',
    '袋',
    '启用',
    'SK241208000035',
    NULL,
    '袋',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-08 10:03:57',
    '2024-12-08 10:02:22'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241207000002',
    '纸巾',
    '启用',
    'GG241207000002',
    NULL,
    '纸巾',
    '5提/箱',
    '启用',
    'SK241207000003',
    NULL,
    '箱',
    5.0,
    '启用',
    '3',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-07 08:11:02',
    '2024-12-07 08:11:02'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241207000002',
    '纸巾',
    '启用',
    'GG241207000002',
    NULL,
    '纸巾',
    '提',
    '启用',
    'SK241207000002',
    NULL,
    '提',
    1.0,
    '启用',
    '3',
    '年',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-07 08:11:02',
    '2024-12-07 08:11:02'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241202000062',
    '制茶素毛肚（零食）',
    '启用',
    'GG241202000062',
    NULL,
    '制茶素毛肚（零食）',
    '18g*360袋/箱',
    '启用',
    'SK241202000104',
    NULL,
    '箱',
    360.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-06 15:36:28',
    '2024-12-02 18:45:04'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241202000062',
    '制茶素毛肚（零食）',
    '启用',
    'GG241202000062',
    NULL,
    '制茶素毛肚（零食）',
    '袋',
    '启用',
    'SK241202000103',
    NULL,
    '袋',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-06 15:36:28',
    '2024-12-02 18:45:04'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241206000054',
    '红肉苹果果酱',
    '启用',
    'GG241206000054',
    NULL,
    '红肉苹果果酱',
    '1kg*16袋/件',
    '启用',
    'SK241206000093',
    NULL,
    '件',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-06 12:53:20',
    '2024-12-06 12:53:20'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241206000054',
    '红肉苹果果酱',
    '启用',
    'GG241206000054',
    NULL,
    '红肉苹果果酱',
    '袋',
    '启用',
    'SK241206000092',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-12-06 12:53:20',
    '2024-12-06 12:53:20'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241118000067',
    '冰勃朗基底乳(常温存储)',
    '启用',
    'GG241118000067',
    NULL,
    '冰勃朗基底乳(常温存储)',
    '1kg*12盒/箱',
    '启用',
    'SK241118000127',
    NULL,
    '箱',
    12.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-11-18 10:44:30',
    '2024-11-18 10:44:30'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241118000067',
    '冰勃朗基底乳(常温存储)',
    '启用',
    'GG241118000067',
    NULL,
    '冰勃朗基底乳(常温存储)',
    '盒',
    '启用',
    'SK241118000126',
    NULL,
    '盒',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-11-18 10:44:30',
    '2024-11-18 10:44:30'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241107000286',
    '外卖保温袋(四杯)含杯托',
    '启用',
    'GG241107000285',
    NULL,
    '外卖保温袋(四杯)含杯托',
    '150套/箱',
    '启用',
    'SK241107000317',
    NULL,
    '箱',
    150.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-11-07 18:24:45',
    '2024-11-07 18:24:45'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP241107000286',
    '外卖保温袋(四杯)含杯托',
    '启用',
    'GG241107000285',
    NULL,
    '外卖保温袋(四杯)含杯托',
    '套',
    '启用',
    'SK241107000316',
    NULL,
    '套',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-11-07 18:24:45',
    '2024-11-07 18:24:45'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240920000147',
    '制茶青年桂花乌龙茶',
    '启用',
    'GG240920000147',
    NULL,
    '制茶青年桂花乌龙茶',
    '100g*50包/箱',
    '启用',
    'SK240920000267',
    NULL,
    '箱',
    50.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-09-20 15:55:37',
    '2024-09-20 15:55:37'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240920000147',
    '制茶青年桂花乌龙茶',
    '启用',
    'GG240920000147',
    NULL,
    '制茶青年桂花乌龙茶',
    '包',
    '启用',
    'SK240920000266',
    NULL,
    '包',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-09-20 15:55:37',
    '2024-09-20 15:55:37'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240919000102',
    '制茶青年500桂花纸杯',
    '启用',
    'GG240919000102',
    NULL,
    '制茶青年500桂花纸杯',
    '500个/件',
    '启用',
    'SK240919000182',
    NULL,
    '件',
    500.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-09-19 17:51:04',
    '2024-09-19 17:51:04'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240919000102',
    '制茶青年500桂花纸杯',
    '启用',
    'GG240919000102',
    NULL,
    '制茶青年500桂花纸杯',
    '个',
    '启用',
    'SK240919000181',
    NULL,
    '个',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-09-19 17:51:04',
    '2024-09-19 17:51:04'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240826000029',
    '金萱乌龙/新',
    '启用',
    'GG240826000029',
    NULL,
    '金萱乌龙/新',
    '100g*100包/件',
    '启用',
    'SK240826000047',
    NULL,
    '件',
    100.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-26 10:51:25',
    '2024-08-26 10:47:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240826000029',
    '金萱乌龙/新',
    '启用',
    'GG240826000029',
    NULL,
    '金萱乌龙/新',
    '包',
    '启用',
    'SK240826000046',
    NULL,
    '包',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-26 10:51:25',
    '2024-08-26 10:47:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240812000022',
    '外卖保温袋（双杯）含杯托',
    '启用',
    'GG240812000022',
    NULL,
    '外卖保温袋（双杯）含杯托',
    '350套/箱',
    '启用',
    'SK240812000041',
    NULL,
    '箱',
    350.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-12 11:01:15',
    '2024-08-12 11:01:15'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240812000022',
    '外卖保温袋（双杯）含杯托',
    '启用',
    'GG240812000022',
    NULL,
    '外卖保温袋（双杯）含杯托',
    '套',
    '启用',
    'SK240812000040',
    NULL,
    '套',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-12 11:01:15',
    '2024-08-12 11:01:15'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240811000002',
    '青提果酱',
    '启用',
    'GG240811000002',
    NULL,
    '青提果酱',
    '1kg*16袋/箱',
    '启用',
    'SK240811000003',
    NULL,
    '箱',
    16.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-11 08:28:05',
    '2024-08-11 08:28:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240811000002',
    '青提果酱',
    '启用',
    'GG240811000002',
    NULL,
    '青提果酱',
    '1kg/袋',
    '启用',
    'SK240811000002',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-11 08:28:05',
    '2024-08-11 08:28:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240809000062',
    '速冻芒果浆',
    '启用',
    'GG240809000058',
    NULL,
    '速冻芒果浆',
    '1KG*12瓶/箱',
    '启用',
    'SK240809000101',
    NULL,
    '箱',
    12.0,
    '启用',
    '18',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-09 15:49:07',
    '2024-08-09 15:49:07'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240809000062',
    '速冻芒果浆',
    '启用',
    'GG240809000058',
    NULL,
    '速冻芒果浆',
    '1KG/瓶',
    '启用',
    'SK240809000100',
    NULL,
    '瓶',
    1.0,
    '启用',
    '18',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-09 15:49:07',
    '2024-08-09 15:49:07'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240801000005',
    '马蹄爆爆珠（果味酱）',
    '启用',
    'GG240801000005',
    NULL,
    '马蹄爆爆珠（果味酱）',
    '12袋/箱',
    '启用',
    'SK240801000008',
    NULL,
    '箱',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-01 08:56:46',
    '2024-08-01 08:56:46'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240801000005',
    '马蹄爆爆珠（果味酱）',
    '启用',
    'GG240801000005',
    NULL,
    '马蹄爆爆珠（果味酱）',
    '900g/袋',
    '启用',
    'SK240801000007',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-08-01 08:56:46',
    '2024-08-01 08:56:46'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240727000189',
    '咖啡盖（黑色）',
    '启用',
    'GG240727000183',
    NULL,
    '咖啡盖（黑色）',
    '1000套/箱',
    '启用',
    'SK240727000354',
    NULL,
    '箱',
    1000.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-27 13:28:24',
    '2024-07-27 13:28:24'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240727000189',
    '咖啡盖（黑色）',
    '启用',
    'GG240727000183',
    NULL,
    '咖啡盖（黑色）',
    '套',
    '启用',
    'SK240727000353',
    NULL,
    '套',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-27 13:28:24',
    '2024-07-27 13:28:24'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240726000018',
    '东方茉莉专用纸杯',
    '启用',
    'GG240726000018',
    NULL,
    '东方茉莉专用纸杯',
    '500只/箱',
    '启用',
    'SK240726000030',
    NULL,
    '箱',
    500.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-26 13:38:17',
    '2024-07-26 13:38:17'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240726000018',
    '东方茉莉专用纸杯',
    '启用',
    'GG240726000018',
    NULL,
    '东方茉莉专用纸杯',
    '只',
    '启用',
    'SK240726000029',
    NULL,
    '只',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-26 13:38:17',
    '2024-07-26 13:38:17'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240723000071',
    '1000Ml专用单杯袋',
    '启用',
    'GG240723000070',
    NULL,
    '1000Ml专用单杯袋',
    '2500支//箱',
    '启用',
    'SK240723000098',
    NULL,
    '箱',
    2500.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-23 11:49:40',
    '2024-07-23 11:49:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240723000071',
    '1000Ml专用单杯袋',
    '启用',
    'GG240723000070',
    NULL,
    '1000Ml专用单杯袋',
    '支',
    '启用',
    'SK240723000097',
    NULL,
    '支',
    1.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-23 11:49:40',
    '2024-07-23 11:49:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240719000077',
    '1000塑杯（大满贯）',
    '启用',
    'GG240719000076',
    NULL,
    '1000塑杯（大满贯）',
    '50个/箱',
    '启用',
    'SK240719000142',
    NULL,
    '箱',
    50.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-19 14:33:01',
    '2024-07-19 14:33:01'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240719000077',
    '1000塑杯（大满贯）',
    '启用',
    'GG240719000076',
    NULL,
    '1000塑杯（大满贯）',
    '个',
    '启用',
    'SK240719000141',
    NULL,
    '个',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-19 14:33:01',
    '2024-07-19 14:33:01'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240711000169',
    '制茶青年单杯袋（4000支装）',
    '启用',
    'GG240711000169',
    NULL,
    '制茶青年单杯袋（4000支装）',
    '箱',
    '启用',
    'SK240711000195',
    NULL,
    '箱',
    1.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-07-11 13:54:59',
    '2024-07-11 13:54:59'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076437',
    '厚椰乳',
    '启用',
    'SP2023062600048',
    NULL,
    '厚椰乳',
    '1L*12 盒',
    '启用',
    'OWL2023062600047',
    NULL,
    '箱',
    12.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2024-07-09 15:01:03',
    '2023-06-29 16:35:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076437',
    '厚椰乳',
    '启用',
    'SP2023062600048',
    NULL,
    '厚椰乳',
    '1L',
    '启用',
    'OSWL2023062600037',
    NULL,
    '盒',
    1.0,
    '启用',
    '9',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2024-07-09 15:01:03',
    '2023-06-29 16:35:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240619000057',
    '冰勃朗基底乳(保鲜存储)',
    '启用',
    'GG240619000057',
    NULL,
    '冰勃朗基底乳(保鲜存储)',
    '1kg*12盒/箱',
    '启用',
    'SK240619000099',
    NULL,
    '箱',
    12.0,
    '启用',
    '9',
    '月',
    '保鲜',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-06-19 17:02:08',
    '2024-06-19 14:10:47'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240619000057',
    '冰勃朗基底乳(保鲜存储)',
    '启用',
    'GG240619000057',
    NULL,
    '冰勃朗基底乳(保鲜存储)',
    '1kg/盒',
    '启用',
    'SK240619000098',
    NULL,
    '盒',
    1.0,
    '启用',
    '9',
    '月',
    '保鲜',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-06-19 17:02:08',
    '2024-06-19 14:10:47'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708074417',
    '制茶青年700磨砂塑杯',
    '启用',
    'SP2023062000041',
    NULL,
    '制茶青年700磨砂塑杯',
    '20袋*25个/箱',
    '启用',
    'OWL2023062000038',
    NULL,
    '箱',
    25.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2024-05-29 11:48:51',
    '2023-06-20 17:59:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708074417',
    '制茶青年700磨砂塑杯',
    '启用',
    'SP2023062000041',
    NULL,
    '制茶青年700磨砂塑杯',
    '25',
    '启用',
    'OSWL2023062000029',
    NULL,
    '个',
    1.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2024-05-29 11:48:51',
    '2023-06-20 17:59:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240403000035',
    '焙炒咖啡粉',
    '启用',
    'GG240403000035',
    NULL,
    '焙炒咖啡粉',
    '380g*20袋/箱',
    '启用',
    'SK240403000068',
    NULL,
    '箱',
    20.0,
    '启用',
    '120',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-04-03 10:38:13',
    '2024-04-03 10:36:35'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240403000035',
    '焙炒咖啡粉',
    '启用',
    'GG240403000035',
    NULL,
    '焙炒咖啡粉',
    '袋',
    '启用',
    'SK240403000067',
    NULL,
    '袋',
    1.0,
    '启用',
    '120',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-04-03 10:38:13',
    '2024-04-03 10:36:35'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240327000003',
    '原味水晶冻粉',
    '启用',
    'GG240327000003',
    NULL,
    '原味水晶冻粉',
    '1kg*12袋/箱',
    '启用',
    'SK240327000005',
    NULL,
    '箱',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2024-03-27 08:42:45',
    '2024-03-27 08:42:45'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240327000003',
    '原味水晶冻粉',
    '启用',
    'GG240327000003',
    NULL,
    '原味水晶冻粉',
    '袋',
    '启用',
    'SK240327000004',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2024-03-27 08:42:45',
    '2024-03-27 08:42:45'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240326000062',
    '速冻草莓',
    '启用',
    'GG240326000062',
    NULL,
    '速冻草莓',
    '10kg/件',
    '启用',
    'SK240326000099',
    NULL,
    '件',
    10.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2024-03-26 16:53:19',
    '2024-03-26 16:53:19'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240326000062',
    '速冻草莓',
    '启用',
    'GG240326000062',
    NULL,
    '速冻草莓',
    'kg',
    '启用',
    'SK240326000098',
    NULL,
    'kg',
    1.0,
    '启用',
    '24',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'OWT',
    '2024-03-26 16:53:19',
    '2024-03-26 16:53:19'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240319000056',
    '正山小种红',
    '启用',
    'GG240319000056',
    NULL,
    '正山小种红',
    '100包/箱',
    '启用',
    'SK240319000108',
    NULL,
    '箱',
    100.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-03-19 13:22:56',
    '2024-03-19 13:22:56'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP240319000056',
    '正山小种红',
    '启用',
    'GG240319000056',
    NULL,
    '正山小种红',
    '包',
    '启用',
    'SK240319000107',
    NULL,
    '包',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2024-03-19 13:22:56',
    '2024-03-19 13:22:56'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP231119000008',
    '圣代勺/新',
    '启用',
    'GG231119000008',
    NULL,
    '圣代勺/新',
    '100个*20包/件',
    '启用',
    'SK231119000014',
    NULL,
    '件',
    20.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-11-19 14:45:40',
    '2023-11-19 14:45:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP231119000008',
    '圣代勺/新',
    '启用',
    'GG231119000008',
    NULL,
    '圣代勺/新',
    '包',
    '启用',
    'SK231119000013',
    NULL,
    '包',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-11-19 14:45:40',
    '2023-11-19 14:45:40'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP231101000003',
    '青稞罐头',
    '启用',
    'GG231101000003',
    NULL,
    '青稞罐头',
    '900g*12罐/件',
    '启用',
    'SK231101000004',
    NULL,
    '件',
    12.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-11-01 09:07:47',
    '2023-11-01 09:07:47'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP231101000003',
    '青稞罐头',
    '启用',
    'GG231101000003',
    NULL,
    '青稞罐头',
    '罐',
    '启用',
    'SK231101000003',
    NULL,
    '罐',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-11-01 09:07:47',
    '2023-11-01 09:07:47'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075366',
    '牛奶冰淇淋粉',
    '启用',
    'SP2023062600042',
    NULL,
    '牛奶冰淇淋粉',
    '3KG*7袋/箱',
    '启用',
    'OWL2023062600041',
    NULL,
    '箱',
    7.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-10-04 10:56:25',
    '2023-06-27 13:57:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075366',
    '牛奶冰淇淋粉',
    '启用',
    'SP2023062600042',
    NULL,
    '牛奶冰淇淋粉',
    '3kg',
    '启用',
    'OSWL2023062600031',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-10-04 10:56:25',
    '2023-06-27 13:57:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075294',
    '制茶青年珍宝杯',
    '启用',
    'SP2023062700008',
    NULL,
    '制茶青年珍宝杯',
    '50个*20条/件',
    '启用',
    'OWL2023062700008',
    NULL,
    '件',
    20.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-10-04 08:24:19',
    '2023-06-27 09:49:01'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075294',
    '制茶青年珍宝杯',
    '启用',
    'SP2023062700008',
    NULL,
    '制茶青年珍宝杯',
    '50个',
    '启用',
    'OSWL2023062700007',
    NULL,
    '条',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-10-04 08:24:19',
    '2023-06-27 09:49:01'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230904000005',
    '果糖/新',
    '启用',
    'GG230904000005',
    NULL,
    '果糖/新',
    '6.25kg*4桶/箱',
    '启用',
    'SK230904000009',
    NULL,
    '箱',
    4.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-09-23 14:17:52',
    '2023-09-04 08:32:55'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230904000005',
    '果糖/新',
    '启用',
    'GG230904000005',
    NULL,
    '果糖/新',
    '桶',
    '启用',
    'SK230904000008',
    NULL,
    '桶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-09-23 14:17:52',
    '2023-09-04 08:32:55'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077973',
    '高档90球形盖',
    '启用',
    'SP2023070500080',
    NULL,
    '高档90球形盖',
    '1000个*1件',
    '启用',
    'OWL2023070500076',
    NULL,
    '件',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-16 14:32:36',
    '2023-07-05 15:06:02'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230718000002',
    '板栗罐头',
    '启用',
    'GG230718000002',
    NULL,
    '制茶青年板栗泥罐头',
    '12罐/箱',
    '启用',
    'SK230718000003',
    NULL,
    '箱',
    12.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-12 19:29:14',
    '2023-07-18 08:39:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230718000002',
    '板栗罐头',
    '启用',
    'GG230718000002',
    NULL,
    '制茶青年板栗泥罐头',
    '罐',
    '启用',
    'SK230718000002',
    NULL,
    '罐',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-12 19:29:14',
    '2023-07-18 08:39:05'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230718000004',
    '特制焦糖',
    '启用',
    'GG230718000004',
    NULL,
    '制茶青年特制焦糖',
    '15罐/箱',
    '启用',
    'SK230718000006',
    NULL,
    '箱',
    15.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-12 19:18:31',
    '2023-07-18 09:08:12'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230718000004',
    '特制焦糖',
    '启用',
    'GG230718000004',
    NULL,
    '制茶青年特制焦糖',
    '罐',
    '启用',
    'SK230718000005',
    NULL,
    '罐',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-12 19:18:31',
    '2023-07-18 09:08:12'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230720000002',
    '单色小芋圆（冷冻）',
    '启用',
    'GG230720000002',
    NULL,
    '单色小芋圆（冷冻）',
    '16袋/箱',
    '启用',
    'SK230720000003',
    NULL,
    '箱',
    16.0,
    '启用',
    '12',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-12 19:05:18',
    '2023-07-20 08:17:06'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230720000002',
    '单色小芋圆（冷冻）',
    '启用',
    'GG230720000002',
    NULL,
    '单色小芋圆（冷冻）',
    '袋',
    '启用',
    'SK230720000002',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-12 19:05:18',
    '2023-07-20 08:17:06'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230811000111',
    '制茶青年奶茶专用粉',
    '启用',
    'GG230811000110',
    NULL,
    '制茶青年奶茶专用粉',
    '8袋/箱',
    '启用',
    'SK230811000139',
    NULL,
    '箱',
    8.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-11 17:32:59',
    '2023-08-11 17:32:59'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230811000111',
    '制茶青年奶茶专用粉',
    '启用',
    'GG230811000110',
    NULL,
    '制茶青年奶茶专用粉',
    '2.5kg/袋',
    '启用',
    'SK230811000138',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-11 17:32:59',
    '2023-08-11 17:32:59'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077354',
    '制茶青年桃肉',
    '启用',
    'SP2023070400053',
    NULL,
    '制茶青年桃肉',
    '1.2kg*12瓶/箱',
    '启用',
    'OWL2023070400049',
    NULL,
    '箱',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-05 17:15:53',
    '2023-07-04 13:26:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077354',
    '制茶青年桃肉',
    '启用',
    'SP2023070400053',
    NULL,
    '制茶青年桃肉',
    '1.2kg',
    '启用',
    'OSWL2023070400042',
    NULL,
    '瓶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-05 17:15:53',
    '2023-07-04 13:26:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077302',
    '果蜜-',
    '启用',
    'SP2023070400040',
    NULL,
    '果蜜-',
    '6.25kg*4桶/箱',
    '启用',
    'OWL2023070400036',
    NULL,
    '箱',
    4.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-02 12:59:09',
    '2023-07-04 11:48:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077302',
    '果蜜-',
    '启用',
    'SP2023070400040',
    NULL,
    '果蜜-',
    '6.25kg',
    '启用',
    'OSWL2023070400029',
    NULL,
    '桶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-02 12:59:09',
    '2023-07-04 11:48:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230709000011',
    '制茶青年双杯袋',
    '启用',
    'GG230709000010',
    NULL,
    '制茶青年双杯袋',
    '3000个/件',
    '启用',
    'SK230709000018',
    NULL,
    '件',
    3000.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-02 12:58:11',
    '2023-07-09 11:15:08'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230709000011',
    '制茶青年双杯袋',
    '启用',
    'GG230709000010',
    NULL,
    '制茶青年双杯袋',
    '个',
    '启用',
    'SK230709000017',
    NULL,
    '个',
    1.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    '河南膳乡客餐饮管理有限公司',
    'OWT',
    '2023-08-02 12:58:11',
    '2023-07-09 11:15:08'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076846',
    '速冻香芋泥（冷冻）',
    '启用',
    'SP2023070100051',
    NULL,
    '速冻香芋泥（冷冻）',
    '1kg*22包/件',
    '启用',
    'OWL2023070100052',
    NULL,
    '件',
    22.0,
    '启用',
    '12',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-02 12:54:52',
    '2023-07-01 15:50:02'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076846',
    '速冻香芋泥（冷冻）',
    '启用',
    'SP2023070100051',
    NULL,
    '速冻香芋泥（冷冻）',
    '1kg',
    '启用',
    'OSWL2023070100036',
    NULL,
    '包',
    1.0,
    '启用',
    '12',
    '月',
    '冷冻',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-02 12:54:52',
    '2023-07-01 15:50:02'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075362',
    '森林玫果.',
    '启用',
    'SP2023062600046',
    NULL,
    '森林玫果.',
    '1.1KG*12瓶/箱',
    '启用',
    'OWL2023062600045',
    NULL,
    '箱',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-02 12:54:05',
    '2023-06-27 13:55:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075362',
    '森林玫果.',
    '启用',
    'SP2023062600046',
    NULL,
    '森林玫果.',
    '1.1KG',
    '启用',
    'OSWL2023062600035',
    NULL,
    '瓶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-08-02 12:54:05',
    '2023-06-27 13:55:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077985',
    '长粗吸管（230＊11）',
    '启用',
    'SP2023070500082',
    NULL,
    '长粗吸管（230＊11）',
    '4000个*1件',
    '启用',
    'OWL2023070500078',
    NULL,
    '件',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-26 09:51:39',
    '2023-07-05 15:09:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075061',
    '椰果果粒',
    '启用',
    'SP2023062600006',
    NULL,
    '椰果果粒',
    '2kg*10袋/件',
    '启用',
    'OWL2023062600006',
    NULL,
    '件',
    10.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-24 09:43:21',
    '2023-06-26 09:44:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075061',
    '椰果果粒',
    '启用',
    'SP2023062600006',
    NULL,
    '椰果果粒',
    '2kg',
    '启用',
    'OSWL2023062600004',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-24 09:43:21',
    '2023-06-26 09:44:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077981',
    '长细吸管（230＊6）',
    '启用',
    'SP2023070500081',
    NULL,
    '长细吸管（230＊6）',
    '4000个*1箱',
    '启用',
    'OWL2023070500077',
    NULL,
    '箱',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-24 09:42:06',
    '2023-07-05 15:08:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077989',
    '短粗吸管（190＊11）',
    '启用',
    'SP2023070500083',
    NULL,
    '短粗吸管（190＊11）',
    '4000个*1件',
    '启用',
    'OWL2023070500079',
    NULL,
    '件',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-24 09:41:39',
    '2023-07-05 15:10:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076693',
    '冰淇淋脆筒',
    '启用',
    'SP2023070100005',
    NULL,
    '冰淇淋脆筒',
    '20支*20条/箱',
    '启用',
    'OWL2023070100004',
    NULL,
    '箱',
    20.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-24 09:16:33',
    '2023-07-01 09:36:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076693',
    '冰淇淋脆筒',
    '启用',
    'SP2023070100005',
    NULL,
    '冰淇淋脆筒',
    '20支',
    '启用',
    'OSWL2023070100003',
    NULL,
    '条',
    1.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-24 09:16:33',
    '2023-07-01 09:36:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077306',
    '红西柚颗粒罐头',
    '启用',
    'SP2023070400041',
    NULL,
    '红西柚颗粒罐头',
    '800g*12罐/箱',
    '启用',
    'OWL2023070400037',
    NULL,
    '箱',
    12.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 10:03:06',
    '2023-07-04 11:49:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077306',
    '红西柚颗粒罐头',
    '启用',
    'SP2023070400041',
    NULL,
    '红西柚颗粒罐头',
    '800g',
    '启用',
    'OSWL2023070400030',
    NULL,
    '罐',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 10:03:06',
    '2023-07-04 11:49:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077330',
    '蓝莓果肉果酱',
    '启用',
    'SP2023070400047',
    NULL,
    '蓝莓果肉果酱',
    '2kg*6听/箱',
    '启用',
    'OWL2023070400043',
    NULL,
    '箱',
    6.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:59:52',
    '2023-07-04 12:16:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077330',
    '蓝莓果肉果酱',
    '启用',
    'SP2023070400047',
    NULL,
    '蓝莓果肉果酱',
    '2kg',
    '启用',
    'OSWL2023070400036',
    NULL,
    '听',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:59:52',
    '2023-07-04 12:16:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077338',
    '120口径封口膜',
    '启用',
    'SP2023070400049',
    NULL,
    '120口径封口膜',
    '6卷/箱',
    '启用',
    'OWL2023070400045',
    NULL,
    '箱',
    6.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:59:27',
    '2023-07-04 12:18:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077338',
    '120口径封口膜',
    '启用',
    'SP2023070400049',
    NULL,
    '120口径封口膜',
    '1',
    '启用',
    'OSWL2023070400038',
    NULL,
    '卷',
    1.0,
    '启用',
    '36',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:59:27',
    '2023-07-04 12:18:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077322',
    '可可饼干碎',
    '启用',
    'SP2023070400045',
    NULL,
    '可可饼干碎',
    '800g/袋*12袋/箱',
    '启用',
    'OWL2023070400041',
    NULL,
    '箱',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:58:47',
    '2023-07-04 12:12:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077322',
    '可可饼干碎',
    '启用',
    'SP2023070400045',
    NULL,
    '可可饼干碎',
    '800g/袋',
    '启用',
    'OSWL2023070400034',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:58:47',
    '2023-07-04 12:12:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077358',
    '红豆罐头',
    '启用',
    'SP2023070400054',
    NULL,
    '红豆罐头',
    '900g*12罐/箱',
    '启用',
    'OWL2023070400050',
    NULL,
    '箱',
    12.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:58:34',
    '2023-07-04 13:28:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077358',
    '红豆罐头',
    '启用',
    'SP2023070400054',
    NULL,
    '红豆罐头',
    '900g',
    '启用',
    'OSWL2023070400043',
    NULL,
    '罐',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:58:34',
    '2023-07-04 13:28:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077261',
    '90口径封口膜',
    '启用',
    'SP2023070400029',
    NULL,
    '90口径封口膜',
    '6卷/箱',
    '启用',
    'OWL2023070400028',
    NULL,
    '箱',
    6.0,
    '启用',
    '36',
    '月',
    '常温',
    FALSE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:54:35',
    '2023-07-04 11:23:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077261',
    '90口径封口膜',
    '启用',
    'SP2023070400029',
    NULL,
    '90口径封口膜',
    '1',
    '启用',
    'OSWL2023070400018',
    NULL,
    '卷',
    1.0,
    '启用',
    '36',
    '月',
    '常温',
    FALSE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:54:35',
    '2023-07-04 11:23:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075065',
    '原味晶球',
    '启用',
    'SP2023062600005',
    NULL,
    '原味晶球',
    '1kg*20袋/箱',
    '启用',
    'OWL2023062600005',
    NULL,
    '箱',
    20.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:51:57',
    '2023-06-26 09:44:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075065',
    '原味晶球',
    '启用',
    'SP2023062600005',
    NULL,
    '原味晶球',
    '1kg',
    '启用',
    'OSWL2023062600003',
    NULL,
    '袋',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:51:57',
    '2023-06-26 09:44:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077289',
    '巧克力风味糖浆',
    '启用',
    'SP2023070400036',
    NULL,
    '巧克力风味糖浆',
    '1.6L*6桶/箱',
    '启用',
    'OWL2023070400032',
    NULL,
    '箱',
    6.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:51:32',
    '2023-07-04 11:36:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077289',
    '巧克力风味糖浆',
    '启用',
    'SP2023070400036',
    NULL,
    '巧克力风味糖浆',
    '1.6',
    '启用',
    'OSWL2023070400025',
    NULL,
    '瓶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-23 09:51:32',
    '2023-07-04 11:36:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708078554',
    '橙柚果酱',
    '启用',
    'SP2023070800015',
    NULL,
    '橙柚果酱',
    '1.2kg*12罐/件',
    '启用',
    'OWL2023070800014',
    NULL,
    '件',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-08 13:03:38',
    '2023-07-08 11:05:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708078554',
    '橙柚果酱',
    '启用',
    'SP2023070800015',
    NULL,
    '橙柚果酱',
    '1.2kg',
    '启用',
    'OSWL2023070800010',
    NULL,
    '罐',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-08 13:03:38',
    '2023-07-08 11:05:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077949',
    '制茶青年（四杯）',
    '启用',
    'SP2023070500073',
    NULL,
    '制茶青年（四杯）',
    '200条*10捆/件',
    '启用',
    'OWL2023070500069',
    NULL,
    '件',
    10.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-05 14:30:12',
    '2023-07-05 13:38:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077949',
    '制茶青年（四杯）',
    '启用',
    'SP2023070500073',
    NULL,
    '制茶青年（四杯）',
    '200条',
    '启用',
    'OSWL2023070500060',
    NULL,
    '捆',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-05 14:30:12',
    '2023-07-05 13:38:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077040',
    '丝滑酸奶',
    '启用',
    'SP2023070300001',
    NULL,
    '丝滑酸奶',
    '1kg*12盒/件',
    '启用',
    'OWL2023070300001',
    NULL,
    '件',
    12.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-03 10:28:25',
    '2023-07-03 08:17:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077040',
    '丝滑酸奶',
    '启用',
    'SP2023070300001',
    NULL,
    '丝滑酸奶',
    '1kg',
    '启用',
    'OSWL2023070300001',
    NULL,
    '盒',
    1.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-03 10:28:25',
    '2023-07-03 08:17:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077060',
    '纯牛奶',
    '启用',
    'SP2023070300006',
    NULL,
    '纯牛奶',
    '1L*12盒/件',
    '启用',
    'OWL2023070300006',
    NULL,
    '件',
    12.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-03 10:28:19',
    '2023-07-03 09:07:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708077060',
    '纯牛奶',
    '启用',
    'SP2023070300006',
    NULL,
    '纯牛奶',
    '1L',
    '启用',
    'OSWL2023070300006',
    NULL,
    '盒',
    1.0,
    '启用',
    '6',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-03 10:28:19',
    '2023-07-03 09:07:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076701',
    '百香果果酱',
    '启用',
    'SP2023070100009',
    NULL,
    '百香果果酱',
    '2kg*6瓶/箱',
    '启用',
    'OWL2023070100009',
    NULL,
    '箱',
    6.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-01 13:17:14',
    '2023-07-01 09:52:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076701',
    '百香果果酱',
    '启用',
    'SP2023070100009',
    NULL,
    '百香果果酱',
    '2kg',
    '启用',
    'OSWL2023070100005',
    NULL,
    '瓶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-01 13:17:14',
    '2023-07-01 09:52:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076709',
    '凤梨味果酱',
    '启用',
    'SP2023070100011',
    NULL,
    '凤梨味果酱',
    '2kg*6瓶/件',
    '启用',
    'OWL2023070100011',
    NULL,
    '件',
    6.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-01 13:17:02',
    '2023-07-01 09:53:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708076709',
    '凤梨味果酱',
    '启用',
    'SP2023070100011',
    NULL,
    '凤梨味果酱',
    '2kg',
    '启用',
    'OSWL2023070100007',
    NULL,
    '瓶',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-07-01 13:17:02',
    '2023-07-01 09:53:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075446',
    '抹茶冰淇淋粉',
    '启用',
    'SP2023062700088',
    NULL,
    '抹茶冰淇淋粉',
    '3kg*8袋/箱',
    '启用',
    'OWL2023062700083',
    NULL,
    '箱',
    8.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 16:52:14',
    '2023-06-27 15:54:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075446',
    '抹茶冰淇淋粉',
    '启用',
    'SP2023062700088',
    NULL,
    '抹茶冰淇淋粉',
    '3kg',
    '启用',
    'OSWL2023062700083',
    NULL,
    '袋',
    1.0,
    '启用',
    '18',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 16:52:14',
    '2023-06-27 15:54:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075370',
    '葡萄果酱',
    '启用',
    'SP2023062600045',
    NULL,
    '葡萄果酱',
    '1.1KG*12罐/箱',
    '启用',
    'OWL2023062600044',
    NULL,
    '箱',
    12.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 14:23:15',
    '2023-06-27 13:57:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075370',
    '葡萄果酱',
    '启用',
    'SP2023062600045',
    NULL,
    '葡萄果酱',
    '1.IKG',
    '启用',
    'OSWL2023062600034',
    NULL,
    '罐',
    1.0,
    '启用',
    '12',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 14:23:15',
    '2023-06-27 13:57:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075282',
    '400U塑杯',
    '启用',
    'SP2023062700003',
    NULL,
    '400U塑杯',
    '50个*20袋/件',
    '启用',
    'OWL2023062700003',
    NULL,
    '件',
    20.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 10:19:01',
    '2023-06-27 09:49:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075282',
    '400U塑杯',
    '启用',
    'SP2023062700003',
    NULL,
    '400U塑杯',
    '50个',
    '启用',
    'OSWL2023062700002',
    NULL,
    '袋',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 10:19:01',
    '2023-06-27 09:49:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075286',
    '1000塑杯',
    '启用',
    'SP2023062700005',
    NULL,
    '1000塑杯',
    '50个*12袋/件',
    '启用',
    'OWL2023062700005',
    NULL,
    '件',
    12.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 10:18:52',
    '2023-06-27 09:49:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;

INSERT INTO product (
    product_code, product_name, product_status, spec_code, third_party_spec_code,
    spec_property, package_spec, spec_status, sku_code, third_party_sku_code,
    product_unit, unit_conversion_ratio, unit_status, shelf_life, shelf_life_type,
    storage_method, is_standard, shipper_name, shipper_level, supplier_name,
    product_source, updated_at, created_at
) VALUES (
    'SP230708075286',
    '1000塑杯',
    '启用',
    'SP2023062700005',
    NULL,
    '1000塑杯',
    '50个',
    '启用',
    'OSWL2023062700004',
    NULL,
    '袋',
    1.0,
    '启用',
    '24',
    '月',
    '常温',
    TRUE,
    '河南上黎供应链管理有限公司',
    'VIP客户',
    NULL,
    'WMS',
    '2023-06-27 10:18:52',
    '2023-06-27 09:49:00'
) ON CONFLICT (sku_code) DO UPDATE SET
    product_name = EXCLUDED.product_name,
    product_status = EXCLUDED.product_status,
    package_spec = EXCLUDED.package_spec;
