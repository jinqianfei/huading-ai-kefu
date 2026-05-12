"""
运输费完整匹配模块

根据仓库和门店名称匹配运输价格
"""

def match_transport_price(warehouse: str, store_name: str, transport_rules: dict = None) -> dict:
    """
    匹配运输价格
    
    Args:
        warehouse: 仓库名称（如"长沙仓"、"广州仓"等）
        store_name: 门店名称
        transport_rules: 运输规则字典，如果为None则使用默认规则
    
    Returns:
        dict: 包含 price(价格), min_fee(最低费用), destination(目的地区域) 的字典
    """
    if transport_rules is None:
        transport_rules = get_default_transport_rules()
    
    # 检查仓库是否存在
    if warehouse not in transport_rules:
        return {
            "price": 0,
            "min_fee": 0,
            "destination": "未知仓库",
            "matched": False,
            "error": f"未知仓库: {warehouse}"
        }
    
    warehouse_rules = transport_rules[warehouse]
    
    # 遍历仓库规则，匹配最具体的区域
    best_match = None
    best_priority = -1
    
    for destination, rule in warehouse_rules.items():
        if destination == "其他" or destination == "默认":
            continue
        
        # 检查门店是否匹配该区域
        if is_store_in_region(store_name, destination):
            # 计算优先级（精确城市 > 省会 > 省级 > 大区）
            priority = get_region_priority(destination)
            if priority > best_priority:
                best_priority = priority
                best_match = {
                    "price": rule["price"],
                    "min_fee": rule["min_fee"],
                    "destination": destination,
                    "matched": True
                }
    
    # 如果没有匹配到具体规则，使用默认规则
    if best_match is None:
        if "其他" in warehouse_rules:
            default_rule = warehouse_rules["其他"]
            best_match = {
                "price": default_rule["price"],
                "min_fee": default_rule["min_fee"],
                "destination": "其他区域",
                "matched": True
            }
        elif "默认" in warehouse_rules:
            default_rule = warehouse_rules["默认"]
            best_match = {
                "price": default_rule["price"],
                "min_fee": default_rule["min_fee"],
                "destination": "默认区域",
                "matched": True
            }
        else:
            best_match = {
                "price": 0,
                "min_fee": 0,
                "destination": "未匹配",
                "matched": False,
                "error": f"门店 '{store_name}' 无法匹配任何目的地区域"
            }
    
    return best_match


def _extract_location(store_name: str) -> str:
    """从门店名称中提取位置信息（括号内内容）"""
    if "（" in store_name and "）" in store_name:
        start = store_name.find("（") + 1
        end = store_name.find("）")
        return store_name[start:end]
    elif "(" in store_name and ")" in store_name:
        start = store_name.find("(") + 1
        end = store_name.find(")")
        return store_name[start:end]
    return store_name


def is_store_in_region(store_name: str, region: str) -> bool:
    """
    判断门店是否在指定区域内
    
    Args:
        store_name: 门店名称
        region: 区域名称（如"长沙市"、"湖南省"、"贵州全境"等）
    
    Returns:
        bool: 是否匹配
    """
    if not store_name or not region:
        return False
    
    # 提取门店名称中的位置信息（括号内内容）
    location_info = _extract_location(store_name)
    # 清理门店名称，去除括号内的内容
    clean_name = store_name.split("（")[0].split("(")[0].strip()
    
    # 定义区域关键词映射
    region_keywords = {
        # 城市级别
        "长沙市": ["长沙", "长沙市"],
        "南昌市": ["南昌", "南昌市"],
        "贵阳市": ["贵阳", "贵阳市"],
        "昆明市": ["昆明", "昆明市"],
        "成都市": ["成都", "成都市"],
        "重庆市": ["重庆", "重庆市"],
        "广州市": ["广州", "广州市"],
        "深圳市": ["深圳", "深圳市"],
        "杭州市": ["杭州", "杭州市"],
        "南京市": ["南京", "南京市"],
        "上海市": ["上海", "上海市"],
        "北京市": ["北京", "北京市"],
        "天津市": ["天津", "天津市"],
        
        # 省级（添加简化key，如"云南"、"贵州"等，与运输规则key保持一致）
        "湖南省": ["湖南", "长沙", "株洲", "湘潭", "衡阳", "岳阳", "常德", "张家界", "益阳", "郴州", "永州", "怀化", "娄底", "湘西"],
        "贵州": ["贵州", "贵阳", "遵义", "六盘水", "安顺", "毕节", "铜仁", "黔东南", "黔南", "黔西南"],
        "贵州省": ["贵州", "贵阳", "遵义", "六盘水", "安顺", "毕节", "铜仁", "黔东南", "黔南", "黔西南"],
        "云南": ["云南", "昆明", "丽江", "大理", "曲靖", "玉溪", "保山", "昭通", "普洱", "临沧", "香格里拉", "德宏", "瑞丽", "芒市", "红河", "迪庆", "版纳", "西双版纳", "爱琴海"],
        "云南省": ["云南", "昆明", "丽江", "大理", "曲靖", "玉溪", "保山", "昭通", "普洱", "临沧", "香格里拉", "德宏", "瑞丽", "芒市", "红河", "迪庆", "版纳", "西双版纳"],
        "江西省": ["江西", "南昌", "景德镇", "萍乡", "九江", "新余", "鹰潭", "赣州", "吉安", "宜春", "抚州", "上饶"],
        "四川省": ["四川", "成都", "自贡", "攀枝花", "泸州", "德阳", "绵阳", "广元", "遂宁", "内江", "乐山", "南充", "眉山", "宜宾", "广安", "达州", "雅安", "巴中", "资阳", "阿坝", "甘孜", "凉山", "锦江", "万达", "犀浦", "龙泉", "郫县", "双流", "温江", "新都"],
        "重庆市": ["重庆", "渝中", "渝北", "南岸", "沙坪坝", "九龙坡", "大渡口", "江北", "北碚", "璧山", "铜梁", "潼南", "荣昌", "合川", "永川", "江津", "綦江", "万盛", "长寿", "涪陵", "垫江", "丰都", "武隆", "梁平", "开县", "云阳", "奉节", "巫山", "巫溪", "万州", "江北", "南坪", "杨家坪", "解放碑", "冉家坝", "石桥铺", "微电园", "新壹街", "爱琴海", "时代天街", "凤鸣", "沙坪坝", "钢城", "五里店"],
        "广东省": ["广东", "广州", "深圳", "珠海", "东莞", "佛山", "中山", "惠州", "汕头", "江门", "湛江", "茂名", "肇庆", "梅州", "汕尾", "河源", "阳江", "清远", "韶关", "揭阳", "潮州", "云浮"],
        "陕西": ["陕西", "西安", "咸阳", "宝鸡", "渭南", "铜川", "延安", "榆林", "汉中", "安康", "商洛"],
        "甘肃": ["甘肃", "兰州", "嘉峪关", "金昌", "白银", "天水", "武威", "张掖", "平凉", "酒泉", "庆阳", "定西", "陇南", "临夏", "甘南"],
        "青海": ["青海", "西宁", "海东", "海北", "黄南", "海南", "果洛", "玉树", "海西"],
        "宁夏": ["宁夏", "银川", "石嘴山", "吴忠", "固原", "中卫"],
        "新疆": ["新疆", "乌鲁木齐", "克拉玛依", "吐鲁番", "哈密"],
        "西藏": ["西藏", "拉萨", "日喀则", "昌都", "林芝", "山南", "那曲", "阿里"],
        
        # 大区
        "湖南其他城市": ["湖南", "长沙"],
        "四川其他城市": ["四川", "成都", "自贡", "攀枝花", "泸州", "德阳", "绵阳", "广元", "遂宁", "内江", "乐山", "南充", "眉山", "宜宾", "广安", "达州", "雅安", "巴中", "资阳", "阿坝", "甘孜", "凉山", "锦江", "万达", "犀浦", "龙泉", "郫县", "双流", "温江", "新都", "驷马桥", "时尚汇", "观音桥", "西城", "祥和", "西双十"],
        "贵州全境": ["贵州", "贵阳", "遵义", "六盘水", "安顺", "毕节", "铜仁", "黔东南", "黔南", "黔西南"],
        "江西全境": ["江西", "南昌", "景德镇", "萍乡", "九江", "新余", "鹰潭", "赣州", "吉安", "宜春", "抚州", "上饶"],
        "江西南昌": ["江西", "南昌"],
        "云南全境": ["云南", "昆明", "丽江", "大理", "曲靖", "玉溪", "保山", "昭通", "普洱", "临沧", "香格里拉", "德宏", "瑞丽", "芒市", "红河", "迪庆", "版纳", "西双版纳"],
        "四川全境": ["四川", "成都", "自贡", "攀枝花", "泸州", "德阳", "绵阳", "广元", "遂宁", "内江", "乐山", "南充", "眉山", "宜宾", "广安", "达州", "雅安", "巴中", "资阳", "阿坝", "甘孜", "凉山", "锦江", "万达", "犀浦"],
        "川渝全境": ["四川", "成都", "重庆", "渝中", "渝北", "南岸", "重庆", "锦江", "万达"],
    }
    
    # 获取区域的关键词
    keywords = region_keywords.get(region, [region])
    
    # 检查位置信息是否包含任何关键词
    for keyword in keywords:
        if keyword in location_info:
            return True
    
    # 如果位置信息没匹配到，尝试匹配清理后的名称
    for keyword in keywords:
        if keyword in clean_name:
            return True
    
    return False


def get_region_priority(region: str) -> int:
    """
    获取区域匹配的优先级
    数值越大优先级越高
    """
    # 城市级别 - 最高优先级
    city_level = ["长沙市", "南昌市", "贵阳市", "昆明市", "成都市", "重庆市"]
    if region in city_level:
        return 100
    
    # 省级（包括简化key如"贵州"、"云南"等）
    province_level = ["湖南省", "贵州省", "云南省", "江西省", "四川省", 
                      "贵州", "云南", "陕西", "甘肃", "青海", "宁夏", "新疆", "西藏",
                      "陕西", "甘肃", "青海", "宁夏", "新疆", "西藏",
                      "广西", "广东"]
    if region in province_level:
        return 50
    
    # 大区级别 - 较低优先级
    region_level = ["湖南其他城市", "贵州全境", "江西全境", "江西南昌", "云南全境", "四川全境", "川渝全境"]
    if region in region_level:
        return 30
    
    return 10


def get_default_transport_rules() -> dict:
    """
    获取默认运输规则（5仓库×多目的地）
    
    Returns:
        dict: 运输规则字典
    """
    return {
        "长沙仓": {
            "长沙市": {"price": 3.5, "min_fee": 40},
            "湖南其他城市": {"price": 6.5, "min_fee": 90},
            "贵州全境": {"price": 15, "min_fee": 200},
            "江西南昌": {"price": 8, "min_fee": 100},
            "贵州": {"price": 15, "min_fee": 200},
            "云南": {"price": 18, "min_fee": 220},
            "江西": {"price": 10, "min_fee": 120},
            "广西": {"price": 12, "min_fee": 150},
            "其他": {"price": 20, "min_fee": 250},
        },
        "广州仓": {
            "广州市": {"price": 3.0, "min_fee": 35},
            "广东省其他城市": {"price": 5.5, "min_fee": 80},
            "广西": {"price": 8, "min_fee": 100},
            "海南": {"price": 12, "min_fee": 150},
            "湖南": {"price": 6, "min_fee": 85},
            "江西": {"price": 7, "min_fee": 95},
            "福建": {"price": 9, "min_fee": 110},
            "其他": {"price": 18, "min_fee": 220},
        },
        "成都仓": {
            "成都市": {"price": 3.5, "min_fee": 40},
            "四川其他城市": {"price": 3.5, "min_fee": 40},
            "重庆市": {"price": 9, "min_fee": 120},
            "贵州": {"price": 8, "min_fee": 100},
            "云南": {"price": 15, "min_fee": 260},
            "陕西": {"price": 12, "min_fee": 140},
            "甘肃": {"price": 15, "min_fee": 180},
            "青海": {"price": 18, "min_fee": 200},
            "宁夏": {"price": 18, "min_fee": 200},
            "新疆": {"price": 22, "min_fee": 280},
            "西藏": {"price": 25, "min_fee": 300},
            "其他": {"price": 20, "min_fee": 250},
        },
        "杭州仓": {
            "杭州市": {"price": 3.0, "min_fee": 35},
            "浙江其他城市": {"price": 5, "min_fee": 70},
            "上海": {"price": 4, "min_fee": 50},
            "江苏": {"price": 5.5, "min_fee": 75},
            "安徽": {"price": 6, "min_fee": 85},
            "福建": {"price": 8, "min_fee": 100},
            "江西": {"price": 7, "min_fee": 95},
            "其他": {"price": 15, "min_fee": 180},
        },
        "廊坊仓": {
            "廊坊市": {"price": 3.0, "min_fee": 35},
            "北京": {"price": 4, "min_fee": 50},
            "天津": {"price": 4.5, "min_fee": 55},
            "河北其他城市": {"price": 6, "min_fee": 80},
            "山西": {"price": 8, "min_fee": 100},
            "山东": {"price": 9, "min_fee": 110},
            "河南": {"price": 10, "min_fee": 120},
            "内蒙古": {"price": 12, "min_fee": 150},
            "辽宁": {"price": 10, "min_fee": 120},
            "吉林": {"price": 12, "min_fee": 150},
            "黑龙江": {"price": 15, "min_fee": 180},
            "其他": {"price": 18, "min_fee": 220},
        }
    }


def list_available_warehouses() -> list:
    """列出所有可用仓库"""
    return list(get_default_transport_rules().keys())


def list_destinations_by_warehouse(warehouse: str) -> list:
    """列出指定仓库的所有目的地"""
    rules = get_default_transport_rules()
    if warehouse in rules:
        return list(rules[warehouse].keys())
    return []