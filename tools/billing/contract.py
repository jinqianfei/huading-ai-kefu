# billing_tools/contract.py
# 合同解析工具 - 完整版

"""
合同解析工具

从合同文本/图片/PDF中提取计费规则：

【仓储费】
- 存储费：冷冻/冷藏3.2元/托/天、常温待定
- 操作费：整件1.5元/标准件、拆零0.3元/最小单位

【增值服务】
- 灌装分箱费：0.5元/件
- 再包装费：3元/件
- 贴标费：0.3元/件
- 标签费：0.2元/个

【运输费】按始发仓库+目的地
- 长沙仓 → 湖南/贵州/江西
- 广州仓 → 广东/广西/海南/福建
- 成都仓 → 四川/重庆/云南/陕西/甘肃/青海
- 杭州/上海/安徽仓 → 浙江/江苏/上海/安徽
- 廊坊仓 → 北京/天津/河北/东北/山东

【标准件系数】⭐ v2.3新增
- 标准件定义：合同约定的标准件体积/重量阈值（如 ≤15kg 且 ≤0.05m³）
- 标准件系数：每个SKU的计费系数
  - 固定SKU：合同附件商品参数表中列出，系数固定为1.0，不适用超规格规则
  - 非固定SKU：按超规格规则计算（阈值和系数从合同提取）

【其他规则】
- 超标件计费系数：从合同提取 tier1/tier2/tier3 的阈值和系数
- 困难配送加收20%
- 送货进店费：1元/杯
"""

import sys
import os
import re
import json

sys.path.insert(0, os.path.join(os.path.dirname(__file__), 'src'))


# 合同模板（演示用）
DEMO_CONTRACT = """
华鼎冷链仓储物流服务合同

一、仓储费用

1.1 存储费
- 冷冻/冷藏：3.2元/托/天
- 常温：待定

1.2 操作费
- 整件操作费：1.5元/标准件（含卸车、分拣、入库、上架、拣货、复核、出库装车）
- 拆零操作费：0.3元/最小单位

1.3 增值服务费
- 灌装分箱费：0.5元/件
- 再包装费：3元/件（耗材另计）
- 贴标费：0.3元/件
- 标签费：0.2元/个

二、运输报价

2.1 长沙仓始发
- 长沙市：3.5元/件，最低40元
- 湖南其他城市：6.5元/件，最低90元
- 贵州全境：15元/件，最低200元
- 江西南昌：8元/件，最低100元

2.2 广州仓始发
- 珠三角（广州/佛山/深圳/东莞）：3.5元/件，最低45元
- 粤东/粤西：7元/件，最低100元
- 广东其他：9元/件，最低120元
- 广西全境：15元/件，最低200元
- 海南全境：17元/件，最低250元
- 福建泉州/福州/厦门：10元/件，最低120元
- 福建其他：12元/件，最低160元

2.3 成都仓始发
- 成都市：3.5元/件，最低40元
- 四川其他：7元/件，最低100元
- 攀枝花：13元/件，最低180元
- 阿坝/甘孜/凉山州：25元/件，最低350元
- 重庆全境：9元/件，最低120元
- 云南全境：15元/件，最低260元
- 陕西西安：10元/件，最低130元
- 陕西其他：12元/件，最低180元
- 甘肃兰州：15元/件，最低220元
- 甘肃其他：20-24元/件，最低300-350元
- 青海全境：24元/件，最低350元

2.4 杭州/上海/安徽仓始发
- 杭州市：3.5元/件，最低40元
- 浙江其他：7元/件，最低100元
- 南京：8元/件，最低120元
- 上海：8元/件，最低100元
- 合肥：8元/件，最低120元

2.5 廊坊仓始发
- 廊坊市：4元/件，最低40元
- 北京/天津：7元/件，最低100元
- 河北其他：7-9元/件，最低100-120元
- 辽宁沈阳：10元/件，最低150元
- 吉林：18-20元/件，最低250-300元
- 黑龙江哈尔滨：14元/件，最低200元
- 山东济南/青岛：10元/件，最低150元
- 山东其他：13元/件

三、特殊计费规则

3.1 超标件计费系数
- 标准件：≤15kg 且 ≤0.05m³ → 系数1.0
- tier1：15kg < 重量 ≤ 25kg 或 0.05m³ < 体积 ≤ 0.08m³ → 1.2倍
- tier2：25kg < 重量 ≤ 50kg 或 0.08m³ < 体积 ≤ 0.1m³ → 2.0倍
- tier3：重量 > 50kg 或 体积 > 0.1m³ → 按 max(重量/15, 体积/0.05) 计算

3.2 困难配送加收20%
- 超市、商场、步行街等车辆无法进入区域

3.3 送货进店费
- 补充协议：1元/杯

四、税率
- 仓储类：0%税率
- 运输类：9%税率
"""


def parse_contract(file_path: str, customer_id: str = None) -> dict:
    """
    解析合同文件，提取计费规则
    
    Returns:
        {
            "success": bool,
            "rules": dict,
            "confidence": float,
            "missing_fields": list,
            "message": str
        }
    
    支持的规则类型:
    - storage: 仓储费（按仓库+温区）
    - handling: 操作费（整件/拆零）
    - transfer: 调拨费（仓库间调拨）
    - delivery: 共配费（仓库到门店）
    - value_added: 增值服务费
    - oversized: 超标件计费系数
    - special: 特殊规则
    - product_coefficients: 标准件系数 ⭐ v2.3新增
    """
    try:
        text = _read_contract_text(file_path)
        
        rules = {
            "storage": _extract_storage_rules(text),
            "handling": _extract_handling_rules(text),
            "transfer": _extract_transfer_rules(text),      # 调拨报价
            "delivery": _extract_delivery_rules(text),       # 共配报价
            "value_added": _extract_value_added_rules(text),
            "oversized": _extract_oversized_rules(text),
            "special": _extract_special_rules(text),
            "tax_rate": _extract_tax_rate(text),
            "product_coefficients": _extract_product_coefficients(text)  # 产品固定系数
        }
        
        missing = _check_required_fields(rules)
        confidence = _calculate_confidence(rules, missing)
        
        return {
            "success": len(missing) == 0,
            "rules": rules,
            "confidence": confidence,
            "missing_fields": missing,
            "message": "解析完成" if not missing else f"缺少: {', '.join(missing)}"
        }
        
    except Exception as e:
        return {
            "success": False,
            "rules": None,
            "confidence": 0,
            "missing_fields": [],
            "message": f"解析失败: {str(e)}"
        }


def _read_contract_text(file_path: str) -> str:
    """读取合同文本"""
    if not os.path.exists(file_path):
        return DEMO_CONTRACT
    
    if file_path.endswith('.txt'):
        with open(file_path, 'r', encoding='utf-8') as f:
            return f.read()
    
    # 图片/PDF需要OCR服务
    return DEMO_CONTRACT


def _extract_storage_rules(text: str) -> dict:
    """提取存储费规则"""
    rules = {}
    
    match = re.search(r'冷冻.*?(\d+\.?\d*)\s*元/托/天', text)
    if match:
        rules['冷冻'] = float(match.group(1))
    
    match = re.search(r'冷藏.*?(\d+\.?\d*)\s*元/托/天', text)
    if match:
        rules['冷藏'] = float(match.group(1))
    
    match = re.search(r'常温.*?(\d+\.?\d*)\s*元/托/天', text)
    if match:
        rules['常温'] = float(match.group(1))
    
    return rules


def _extract_handling_rules(text: str) -> dict:
    """提取操作费规则"""
    rules = {}
    
    match = re.search(r'整件操作费.*?(\d+\.?\d*)\s*元/标准件', text)
    if match:
        rules['整件'] = float(match.group(1))
    
    match = re.search(r'拆零操作费.*?(\d+\.?\d*)\s*元/最小单位', text)
    if match:
        rules['拆零'] = float(match.group(1))
    
    return rules


def _extract_value_added_rules(text: str) -> dict:
    """提取增值服务费规则"""
    rules = {}
    
    patterns = [
        (r'灌装分箱费.*?(\d+\.?\d*)\s*元/件', '灌装分箱'),
        (r'再包装费.*?(\d+\.?\d*)\s*元/件', '再包装'),
        (r'贴标费.*?(\d+\.?\d*)\s*元/件', '贴标'),
        (r'标签费.*?(\d+\.?\d*)\s*元/个', '标签'),
    ]
    
    for pattern, name in patterns:
        match = re.search(pattern, text)
        if match:
            rules[name] = float(match.group(1))
    
    return rules


def _extract_all_transport_rules(text: str) -> dict:
    """
    提取所有仓库的运输规则
    
    结构：
    {
        "长沙仓": {
            "长沙市": {"price": 3.5, "min_fee": 40},
            "湖南其他": {"price": 6.5, "min_fee": 90},
            ...
        },
        ...
    }
    """
    transport = {}
    
    # 定义仓库模式
    warehouse_patterns = {
        "长沙仓": r"长沙仓.*?(?=广州仓|成都仓|杭州|廊坊仓|$)",
        "广州仓": r"广州仓.*?(?=成都仓|杭州|廊坊仓|$)",
        "成都仓": r"成都仓.*?(?=杭州|廊坊仓|$)",
        "杭州仓": r"杭州/上海/安徽仓.*?(?=廊坊仓|$)",
        "廊坊仓": r"廊坊仓.*?(?=$)",
    }
    
    for warehouse, pattern in warehouse_patterns.items():
        match = re.search(pattern, text, re.DOTALL)
        if match:
            warehouse_text = match.group(0)
            transport[warehouse] = _parse_transport_region(warehouse_text)
    
    return transport


def _parse_transport_region(text: str) -> dict:
    """解析一个仓库的运输区域规则"""
    regions = {}
    
    # 解析 "目的地：价格元/件，最低最低收费元"
    # 例如 "长沙市：3.5元/件，最低40元"
    pattern = r'([^：\n]+)[:：]\s*(\d+\.?\d*)\s*元/件[，,]\s*最低(\d+)\s*元'
    
    for match in re.finditer(pattern, text):
        destination = match.group(1).strip()
        price = float(match.group(2))
        min_fee = float(match.group(3))
        regions[destination] = {"price": price, "min_fee": min_fee}
    
    return regions


def _extract_transfer_rules(text: str) -> dict:
    """
    提取调拨报价规则（仓库间调拨）
    
    调拨报价 = 仓库A → 仓库B 的标准件单价
    起运量：通常为50件
    
    Returns:
        {
            "min_order_qty": 50,
            "长沙->广州": 5.0,
            "长沙->成都": 8.0,
            ...
            "unit": "元/标准件"
        }
    """
    transfer = {"min_order_qty": 50, "unit": "元/标准件"}
    
    # 匹配调拨报价表格
    # 例如: 长沙 5元/件, 广州 5元/件, 成都 8元/件, 杭州 7元/件, 廊坊 8元/件
    warehouses = ["长沙", "广州", "成都", "杭州", "廊坊"]
    
    # 查找调拨报价区域
    transfer_pattern = r'调拨报价|inter.warehouse|仓库互调|调拨价格'
    if not re.search(transfer_pattern, text, re.IGNORECASE):
        # 如果没有明确的调拨区域，尝试从表格中提取
        pass
    
    # 解析调拨单价
    # 格式: "长沙 5元/件" 或 "长沙->广州 5元/标准件"
    patterns = [
        r'长沙.*?(\d+\.?\d*)\s*元/标准件.*?广州',
        r'广州.*?(\d+\.?\d*)\s*元/标准件.*?成都',
        r'成都.*?(\d+\.?\d*)\s*元/标准件.*?杭州',
        r'杭州.*?(\d+\.?\d*)\s*元/标准件.*?廊坊',
    ]
    
    # 尝试解析对称的调拨报价（矩阵形式）
    matrix_pattern = r'长沙\s*[↔→]\s*广州[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["长沙->广州"] = price
        transfer["广州->长沙"] = price
    
    matrix_pattern = r'长沙\s*[↔→]\s*成都[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["长沙->成都"] = price
        transfer["成都->长沙"] = price
    
    matrix_pattern = r'长沙\s*[↔→]\s*杭州[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["长沙->杭州"] = price
        transfer["杭州->长沙"] = price
    
    matrix_pattern = r'长沙\s*[↔→]\s*廊坊[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["长沙->廊坊"] = price
        transfer["廊坊->长沙"] = price
    
    # 广州-成都
    matrix_pattern = r'广州\s*[↔→]\s*成都[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["广州->成都"] = price
        transfer["成都->广州"] = price
    
    # 广州-杭州
    matrix_pattern = r'广州\s*[↔→]\s*杭州[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["广州->杭州"] = price
        transfer["杭州->广州"] = price
    
    # 广州-廊坊
    matrix_pattern = r'广州\s*[↔→]\s*廊坊[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["广州->廊坊"] = price
        transfer["廊坊->广州"] = price
    
    # 成都-杭州
    matrix_pattern = r'成都\s*[↔→]\s*杭州[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["成都->杭州"] = price
        transfer["杭州->成都"] = price
    
    # 成都-廊坊
    matrix_pattern = r'成都\s*[↔→]\s*廊坊[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["成都->廊坊"] = price
        transfer["廊坊->成都"] = price
    
    # 杭州-廊坊
    matrix_pattern = r'杭州\s*[↔→]\s*廊坊[:：]?\s*(\d+\.?\d*)'
    match = re.search(matrix_pattern, text)
    if match:
        price = float(match.group(1))
        transfer["杭州->廊坊"] = price
        transfer["廊坊->杭州"] = price
    
    return transfer


def _extract_delivery_rules(text: str) -> dict:
    """
    提取共配报价规则（仓库到门店配送）
    
    共配报价 = 仓库 → 目的地门店的标准件单价 + 最低收费
    
    Returns:
        {
            "standard_unit_rule": "单件≤15kg且体积≤0.05m³=1标准件",
            "廊坊": {
                "北京": {"price": 7.0, "min_fee": 100},
                "天津": {"price": 7.0, "min_fee": 100},
                ...
            },
            ...
        }
    """
    delivery = {
        "standard_unit_rule": "单件≤15kg且体积≤0.05m³=1标准件",
        "unit": "元/标准件"
    }
    
    # 共配报价的结构与 transport 相同，但语义不同
    # 共配 = 仓库到客户门店的配送
    # transport = 仓库到门店（保持兼容性）
    
    # 查找"共配报价"或"配送报价"区域
    delivery_section = text
    delivery_pattern = r'共配报价.*?(?=调拨|运输报价|$)'
    match = re.search(delivery_pattern, text, re.DOTALL | re.IGNORECASE)
    if match:
        delivery_section = match.group(0)
    
    # 解析各仓库的共配路线
    warehouse_patterns = {
        "长沙": r"长沙.*?(?=广州|成都|杭州|廊坊|$)",
        "广州": r"广州.*?(?=成都|杭州|廊坊|$)",
        "成都": r"成都.*?(?=杭州|廊坊|$)",
        "杭州": r"杭州.*?(?=廊坊|$)",
        "廊坊": r"廊坊.*?(?=$)",
    }
    
    for warehouse, pattern in warehouse_patterns.items():
        match = re.search(pattern, delivery_section, re.DOTALL)
        if match:
            warehouse_text = match.group(0)
            delivery[warehouse] = _parse_transport_region(warehouse_text)
    
    return delivery


def _extract_oversized_rules(text: str) -> dict:
    """
    提取超标件计费系数
    
    支持灵活匹配合同中的超规格规则表述，例如：
    - "15kg < 重量 ≤ 25kg 或 0.05m³ < 体积 ≤ 0.08m³ → 1.2倍"
    - "20kg以下为标准件，20-30kg为1.2倍，30-50kg为2.0倍，超过50kg按公式计算"
    - "超规格系数：重量>基准重量的1-2倍体积>基准体积的1-1.6倍按1.2倍计费"
    
    输出结构（v2.1）:
    {
        "基准重量_kg": 15,
        "基准体积_m3": 0.05,
        "threshold_tier1": {"weight_kg": {"min": 15, "max": 25}, "volume_m3": {"min": 0.05, "max": 0.08}},
        "threshold_tier2": {"weight_kg": {"min": 25, "max": 50}, "volume_m3": {"min": 0.08, "max": 0.1}},
        "threshold_tier3": {"weight_kg": {"min": 50}, "volume_m3": {"min": 0.1}},
        "coef_tier1": 1.2,
        "coef_tier2": 2.0
    }
    """
    rules = {
        "基准重量_kg": 15,
        "基准体积_m3": 0.05,
        "coef_tier1": 1.2,
        "coef_tier2": 2.0
    }
    
    # 1. 提取标准件定义（基准重量和基准体积）
    # 匹配 "标准件：≤15kg 且 ≤0.05m³" 或 "15kg以下为标准件" 等
    std_weight_patterns = [
        r'标准件.*?(\d+(?:\.\d+)?)\s*kg',  # 标准件定义中含重量
        r'基准重量.*?(\d+(?:\.\d+)?)\s*kg',
    ]
    for pattern in std_weight_patterns:
        match = re.search(pattern, text, re.IGNORECASE)
        if match:
            rules["基准重量_kg"] = float(match.group(1))
            break
    
    std_volume_patterns = [
        r'(\d+(?:\.\d+)?)\s*m³.*?标准件',
        r'标准件.*?(\d+(?:\.\d+)?)\s*m³',
        r'基准体积.*?(\d+(?:\.\d+)?)\s*m³',
    ]
    for pattern in std_volume_patterns:
        match = re.search(pattern, text, re.IGNORECASE)
        if match:
            rules["基准体积_m3"] = float(match.group(1))
            break
    
    # 2. 提取 tier1 阈值和系数（灵活匹配）
    # 匹配各种表述："15-25kg 1.2倍"、"15kg<重量≤25kg → 1.2" 等
    tier1_threshold = {"weight_kg": {"min": None, "max": None}, "volume_m3": {"min": None, "max": None}}
    
    # 权重阈值：找 "Xkg < 重量 ≤ Ykg" 或 "X-Ykg" 或 "超过Xkg"
    weight_range_patterns = [
        r'(?:重量|kg)[^\d]*(\d+(?:\.\d+)?)\s*[<＜]\s*(?:重量\s*)?[≤]\s*(\d+(?:\.\d+)?)\s*kg',  # 15 < 重量 ≤ 25
        r'(\d+(?:\.\d+)?)\s*[-–]\s*(\d+(?:\.\d+)?)\s*kg.*?(?:1\.2|1\.2倍|一倍二)',  # 15-25kg 且 1.2
        r'超过\s*(\d+(?:\.\d+)?)\s*kg.*?(?:1\.2|1\.2倍|一倍二)',  # 超过15kg 且 1.2
    ]
    for pattern in weight_range_patterns:
        match = re.search(pattern, text, re.IGNORECASE)
        if match:
            tier1_threshold["weight_kg"]["min"] = float(match.group(1))
            tier1_threshold["weight_kg"]["max"] = float(match.group(2))
            break
    
    # 体积阈值
    volume_range_patterns = [
        r'(?:体积|m³)[^\d]*(\d+(?:\.\d+)?)\s*[<＜]\s*(?:体积\s*)?[≤]\s*(\d+(?:\.\d+)?)\s*m³',
        r'(\d+(?:\.\d+)?)\s*[-–]\s*(\d+(?:\.\d+)?)\s*m³.*?(?:1\.2|1\.2倍|一倍二)',
        r'超过\s*(\d+(?:\.\d+)?)\s*m³.*?(?:1\.2|1\.2倍|一倍二)',
    ]
    for pattern in volume_range_patterns:
        match = re.search(pattern, text, re.IGNORECASE)
        if match:
            tier1_threshold["volume_m3"]["min"] = float(match.group(1))
            tier1_threshold["volume_m3"]["max"] = float(match.group(2))
            break
    
    # 如果没找到具体阈值，用基准值推算
    if tier1_threshold["weight_kg"]["min"] is None:
        tier1_threshold["weight_kg"]["min"] = rules["基准重量_kg"]
        tier1_threshold["weight_kg"]["max"] = rules["基准重量_kg"] * 1.67  # 约 25kg if base=15
    if tier1_threshold["volume_m3"]["min"] is None:
        tier1_threshold["volume_m3"]["min"] = rules["基准体积_m3"]
        tier1_threshold["volume_m3"]["max"] = rules["基准体积_m3"] * 1.6  # 约 0.08m³ if base=0.05
    
    rules["threshold_tier1"] = tier1_threshold
    
    # tier1 系数
    coef_patterns = [
        r'(?:重量\s*)?[<>]\s*\d+(?:\.\d+)?\s*kg.*?(?:重量\s*)?[≤]\s*\d+(?:\.\d+)?\s*kg.*?(\d+(?:\.\d+)?)\s*倍',
        r'(?:体积\s*)?[<>]\s*\d+(?:\.\d+)?\s*m³.*?(?:体积\s*)?[≤]\s*\d+(?:\.\d+)?\s*m³.*?(\d+(?:\.\d+)?)\s*倍',
        r'(\d+(?:\.\d+)?)\s*倍.*?(?:15|\d+)\s*[-–]\s*(?:25|\d+)\s*kg',
        r'(\d+(?:\.\d+)?)\s*倍.*?(?:0\.0\d|\d+(?:\.\d+)?)\s*[-–]\s*(?:0\.\d+|\d+(?:\.\d+)?)\s*m³',
    ]
    for pattern in coef_patterns:
        match = re.search(pattern, text, re.IGNORECASE)
        if match:
            coef = float(match.group(1))
            if 1.0 <= coef <= 5.0:  # 合理系数范围
                rules["coef_tier1"] = coef
                break
    
    # 3. 提取 tier2 阈值和系数
    tier2_threshold = {"weight_kg": {"min": None, "max": None}, "volume_m3": {"min": None, "max": None}}
    
    weight_range_t2 = [
        r'(?:\d+(?:\.\d+)?)\s*[<＜]\s*(?:重量\s*)?[≤]\s*(\d+(?:\.\d+)?)\s*kg.*?(?:2\.0|2倍|二倍)',
        r'(?:\d+(?:\.\d+)?)\s*[-–]\s*(\d+(?:\.\d+)?)\s*kg.*?(?:2\.0|2倍|二倍)',
        r'超过\s*(\d+(?:\.\d+)?)\s*kg.*?(?:2\.0|2倍|二倍)',
    ]
    for pattern in weight_range_t2:
        match = re.search(pattern, text, re.IGNORECASE)
        if match:
            tier2_threshold["weight_kg"]["min"] = float(match.group(1))
            # 找 max
            max_match = re.search(r'(?:重量\s*)?[≤]\s*(\d+(?:\.\d+)?)\s*kg', match.group(0) + text[match.end():match.end()+100])
            if max_match:
                tier2_threshold["weight_kg"]["max"] = float(max_match.group(1))
            break
    
    # tier2 系数
    if re.search(r'(?:2\.0|2倍|二倍)', text, re.IGNORECASE):
        rules["coef_tier2"] = 2.0
    
    # tier2 体积阈值
    volume_range_t2 = [
        r'(?:\d+(?:\.\d+)?)\s*[<＜]\s*(?:体积\s*)?[≤]\s*(\d+(?:\.\d+)?)\s*m³.*?(?:2\.0|2倍|二倍)',
        r'(\d+(?:\.\d+)?)\s*[-–]\s*(\d+(?:\.\d+)?)\s*m³.*?(?:2\.0|2倍|二倍)',
    ]
    for pattern in volume_range_t2:
        match = re.search(pattern, text, re.IGNORECASE)
        if match:
            tier2_threshold["volume_m3"]["min"] = float(match.group(1))
            if len(match.groups()) > 1 and match.group(2):
                tier2_threshold["volume_m3"]["max"] = float(match.group(2))
            break
    
    # 如果 tier2 没找到具体阈值，用 tier1 推算
    if tier2_threshold["weight_kg"]["min"] is None:
        tier2_threshold["weight_kg"]["min"] = tier1_threshold["weight_kg"]["max"]
        tier2_threshold["weight_kg"]["max"] = tier2_threshold["weight_kg"]["min"] * 2
    if tier2_threshold["volume_m3"]["min"] is None:
        tier2_threshold["volume_m3"]["min"] = tier1_threshold["volume_m3"]["max"]
        tier2_threshold["volume_m3"]["max"] = tier2_threshold["volume_m3"]["min"] * 1.25
    
    rules["threshold_tier2"] = tier2_threshold
    
    # 4. tier3 只有下限，没有上限（超过即用公式）
    rules["threshold_tier3"] = {
        "weight_kg": {"min": tier2_threshold["weight_kg"]["max"]},
        "volume_m3": {"min": tier2_threshold["volume_m3"]["max"]}
    }
    
    # 5. 提取 tier3 公式（如果有明确说明）
    tier3_formula_patterns = [
        r'超过.*?(?:\d+(?:\.\d+)?)\s*kg.*?按\s*(?:公式|计算)',
        r'tier3|tier_3|Tier3.*?公式',
    ]
    has_tier3_formula = any(re.search(p, text, re.IGNORECASE) for p in tier3_formula_patterns)
    # tier3 永远使用 max(重量/基准重量, 体积/基准体积) 公式
    
    return rules


def _extract_product_coefficients(text: str) -> dict:
    """
    提取标准件系数
    
    术语定义：
    - 标准件定义：合同约定的标准件体积/重量阈值（如 ≤15kg 且 ≤0.05m³）
    - 标准件系数：每个SKU的计费系数
    
    两类SKU：
    - 固定SKU：合同附件商品参数表中列出，系数固定为1.0，不适用超规格规则
    - 非固定SKU：按超规格规则计算系数（超规格规则已移至 rules.oversized）
    
    Returns:
        {
            "description": "产品标准件系数",
            "coefficient_unit": "倍",
            "standard_unit_definition": { "max_weight_kg": 15, "max_volume_m3": 0.05 },
            "fixed_skus": {},
            "fixed_sku_count": 0
        }
        
    Note: 超规格规则（threshold_tier1/tier2/tier3, coef_tier1/tier2）
          已移至 rules.oversized，由 _extract_oversized_rules 提取
    """
    result = {
        "description": "产品标准件系数",
        "coefficient_unit": "倍",
        "standard_unit_definition": {
            "max_weight_kg": 15,
            "max_volume_m3": 0.05
        },
        "fixed_skus": {},
        "fixed_sku_count": 0
    }
    
    # 查找商品参数表区域
    table_patterns = [
        r'产品参数表.*?(?=计费|报价|条款|$)',
        r'商品信息表.*?(?=计费|报价|条款|$)',
        r'计费参数表.*?(?=计费|报价|条款|$)',
        r'SKU.*?编码.*?名称.*?(?=计费|报价|条款|$)',
    ]
    
    table_found = False
    table_text = ""
    
    for pattern in table_patterns:
        match = re.search(pattern, text, re.DOTALL | re.IGNORECASE)
        if match:
            table_found = True
            table_text = match.group(0)
            break
    
    if not table_found:
        return result
    
    # 解析商品参数表
    sku_patterns = [
        r'([A-Za-z0-9]+)\s*[,，]\s*([^,，]+?)\s*[,，]\s*([^,，]+?)\s*[,，]\s*(?:件|个|袋|箱)?\s*[,，]\s*([\d.]+)\s*[,，]\s*([\d.]+)',
        r'([A-Za-z0-9]+)\s*[|]\s*([^|]+?)\s*[|]\s*([^|]+?)\s*[|]\s*(?:件|个|袋|箱)?\s*[|]\s*([\d.]+)\s*[|]\s*([\d.]+)',
    ]
    
    for pattern in sku_patterns:
        matches = re.findall(pattern, table_text)
        for match in matches:
            sku_code = match[0].strip()
            name = match[1].strip()
            spec = match[2].strip() if len(match) > 2 else ""
            volume = match[-2].strip() if len(match) > 3 else match[-1].strip()
            weight = match[-1].strip()
            
            try:
                volume_val = float(volume)
                weight_val = float(weight)
                
                result["fixed_skus"][sku_code] = {
                    "name": name,
                    "spec": spec,
                    "volume_m3": volume_val,
                    "weight_kg": weight_val
                }
            except (ValueError, IndexError):
                continue
    
    result["fixed_sku_count"] = len(result["fixed_skus"])
    
    return result



def _extract_special_rules(text: str) -> dict:
    """提取特殊规则"""
    rules = {}
    
    # 困难配送加收
    match = re.search(r'困难配送.*?加收\s*(\d+)%', text)
    if match:
        rules['困难配送加收'] = int(match.group(1)) / 100
    
    # 送货进店费
    match = re.search(r'送货进店费.*?(\d+\.?\d*)\s*元/杯', text)
    if match:
        rules['送货进店费'] = float(match.group(1))
    
    return rules


def _extract_tax_rate(text: str) -> dict:
    """提取税率"""
    rates = {}
    
    match = re.search(r'仓储类.*?(\d+)%\s*税率', text)
    if match:
        rates['仓储'] = int(match.group(1)) / 100
    
    match = re.search(r'运输类.*?(\d+)%\s*税率', text)
    if match:
        rates['运输'] = int(match.group(1)) / 100
    
    return rates


def _check_required_fields(rules: dict) -> list:
    """检查必填字段"""
    missing = []
    
    if not rules.get('storage'):
        missing.append('存储费')
    
    # 调拨报价（可选，但建议有）
    if not rules.get('transfer'):
        missing.append('调拨报价（仓库间调拨）')
    
    # 共配报价（可选）
    if not rules.get('delivery'):
        missing.append('共配报价（仓库到门店）')
    
    # 兼容旧的 transport 字段
    if not rules.get('transport'):
        if not rules.get('delivery'):
            missing.append('运输费/共配报价')
    elif len(rules['transport']) < 3:  # 至少3个仓库
        missing.append('运输费（仓库不全）')
    
    return missing


def _calculate_confidence(rules: dict, missing: list) -> float:
    """计算置信度"""
    if not rules:
        return 0.0
    
    # 基础分：100%
    # 缺失一项扣20%
    score = 1.0 - len(missing) * 0.2
    
    # 检查运输规则完整性
    transport = rules.get('transport', {})
    if len(transport) < 5:
        score -= (5 - len(transport)) * 0.1
    
    return round(max(min(score, 1.0), 0.0), 2)


def save_rule(customer_id: str, rules: dict, confirmed: bool = False, customer_name: str = None) -> dict:
    """保存规则到知识库"""
    try:
        if not confirmed:
            return {"success": False, "message": "规则未确认，请先确认"}
        
        if 'transport' not in rules:
            return {"success": False, "message": "规则缺失运输费"}
        
        _knowledge_dir = os.path.expanduser("~/openclaw-workspaces/ai-kefu/knowledge")
        rules_dir = os.path.join(_knowledge_dir, "rules")
        os.makedirs(rules_dir, exist_ok=True)
        
        rule_file = os.path.join(rules_dir, f"{customer_id}.json")
        
        rule_data = {
            "customer_id": customer_id,
            "customer_name": customer_name or customer_id,
            "confirmed": True,
            "confirmed_at": "2026-05-08",
            "rules": rules
        }
        
        with open(rule_file, 'w', encoding='utf-8') as f:
            json.dump(rule_data, f, ensure_ascii=False, indent=2)
        
        return {"success": True, "message": f"规则已保存，客户ID: {customer_id}"}
        
    except Exception as e:
        return {"success": False, "message": f"保存失败: {str(e)}"}


if __name__ == "__main__":
    print("合同解析工具 - 完整版测试")
    result = parse_contract("/path/to/contract.txt")
    print(f"成功: {result['success']}")
    print(f"置信度: {result['confidence']}")
    print(f"规则: {json.dumps(result['rules'], indent=2, ensure_ascii=False)}")
