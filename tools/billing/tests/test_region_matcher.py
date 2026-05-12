"""
目的地区域匹配 - 单元测试
"""

import pytest
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'src', 'calculator'))
from region_matcher import (
    match_region, 
    is_hunan_store,
    get_all_regions,
    get_province_by_city,
    PROVINCE_KEYWORDS,
    CITY_KEYWORDS
)


class TestMatchRegion:
    """测试 match_region 函数"""
    
    def test_kunming_store(self):
        """昆明门店 → 云南省"""
        result = match_region("宫小燕（昆明世纪城店）")
        assert result == "云南省"
    
    def test_chengdu_store(self):
        """成都门店 → 四川省"""
        result = match_region("宫小燕（成都星河店）")
        assert result == "四川省"
    
    def test_changsha_store(self):
        """长沙门店 → 湖南省"""
        result = match_region("宫小燕（长沙五一广场店）")
        assert result == "湖南省"
    
    def test_changsha_hunan_store(self):
        """长沙门店（无省份标注）→ 湖南省"""
        result = match_region("宫小燕（长沙店）")
        assert result == "湖南省"
    
    def test_guiyang_store(self):
        """贵阳门店 → 贵州省"""
        result = match_region("宫小燕（贵阳世纪城店）")
        assert result == "贵州省"
    
    def test_nanchang_store(self):
        """南昌门店 → 江西省"""
        result = match_region("宫小燕（南昌红谷滩店）")
        assert result == "江西省"
    
    def test_guangzhou_store(self):
        """广州门店 → 广东省"""
        result = match_region("宫小燕（广州天河城店）")
        assert result == "广东省"
    
    def test_shanghai_store(self):
        """上海门店 → 上海市"""
        result = match_region("宫小燕（上海陆家嘴店）")
        assert result == "上海市"
    
    def test_beijing_store(self):
        """北京门店 → 北京市"""
        result = match_region("宫小燕（北京国贸店）")
        assert result == "北京市"
    
    def test_chongqing_store(self):
        """重庆门店 → 重庆市"""
        result = match_region("宫小燕（重庆解放碑店）")
        assert result == "重庆市"
    
    def test_hangzhou_store(self):
        """杭州门店 → 浙江省"""
        result = match_region("宫小燕（杭州西湖店）")
        assert result == "浙江省"
    
    def test_nanjing_store(self):
        """南京门店 → 江苏省"""
        result = match_region("宫小燕（南京新街口店）")
        assert result == "江苏省"
    
    def test_empty_store_name(self):
        """空门店名称"""
        result = match_region("")
        assert result == "未知区域"
    
    def test_unknown_store(self):
        """无法识别的门店"""
        result = match_region("宫小燕（某地店）")
        assert result == "未匹配区域"
    
    def test_store_name_without_parentheses(self):
        """无括号的门店名称"""
        result = match_region("宫小燕长沙店")
        assert result == "湖南省"


class TestIsHunanStore:
    """测试 is_hunan_store 函数"""
    
    def test_changsha_is_hunan(self):
        assert is_hunan_store("宫小燕（长沙店）") == True
        assert is_hunan_store("宫小燕（长沙五一广场店）") == True
    
    def test_guizhou_changsha_not_hunan(self):
        """贵州的长沙不是湖南"""
        assert is_hunan_store("宫小燕（贵州长沙店）") == False
    
    def test_other_provinces(self):
        """其他省份门店"""
        assert is_hunan_store("宫小燕（广州店）") == False
        assert is_hunan_store("宫小燕（贵阳店）") == False


class TestGetAllRegions:
    """测试 get_all_regions 函数"""
    
    def test_has_provinces(self):
        """包含省份"""
        regions = get_all_regions()
        assert "湖南省" in regions
        assert "贵州省" in regions
        assert "云南省" in regions
    
    def test_has_cities(self):
        """包含城市"""
        regions = get_all_regions()
        assert "长沙市" in regions
        assert "成都市" in regions
        assert "重庆市" in regions
    
    def test_returns_sorted_list(self):
        """返回排序列表"""
        regions = get_all_regions()
        assert isinstance(regions, list)
        assert regions == sorted(regions)


class TestGetProvinceByCity:
    """测试 get_province_by_city 函数"""
    
    def test_changsha(self):
        assert get_province_by_city("长沙") == "湖南省"
        assert get_province_by_city("长沙市") == "湖南省"
    
    def test_guiyang(self):
        assert get_province_by_city("贵阳") == "贵州省"
        assert get_province_by_city("贵阳市") == "贵州省"
    
    def test_kunming(self):
        assert get_province_by_city("昆明") == "云南省"
        assert get_province_by_city("昆明市") == "云南省"
    
    def test_unknown_city(self):
        assert get_province_by_city("某城市") == ""


class TestProvinceKeywords:
    """测试省份关键词配置"""
    
    def test_all_provinces_have_keywords(self):
        """所有省份都有关键词"""
        for province, keywords in PROVINCE_KEYWORDS.items():
            assert len(keywords) > 0, f"{province} 没有关键词"
    
    def test_province_name_in_keywords(self):
        """省份名称在关键词中"""
        for province, keywords in PROVINCE_KEYWORDS.items():
            province_base = province.replace("省", "").replace("市", "")
            assert province_base in keywords or province in keywords, \
                f"{province} 不在自身关键词列表中"


if __name__ == "__main__":
    pytest.main([__file__, "-v"])