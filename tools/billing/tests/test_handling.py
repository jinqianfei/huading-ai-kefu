"""
操作费计算 - 单元测试
"""

import pytest
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'src', 'calculator'))
from handling import (
    calc_handling_fee, 
    calc_handling_fee_detail, 
    get_handling_fee_rates,
    HANDLING_FEE_STANDARD,
    HANDLING_FEE_SPLIT
)


class TestCalcHandlingFee:
    """测试 calc_handling_fee 函数"""
    
    def test_standard_handling_fee(self):
        """整件操作费：1.5元/标准件"""
        assert calc_handling_fee(10, is_split=False) == 15.0
        assert calc_handling_fee(1, is_split=False) == 1.5
        assert calc_handling_fee(5.5, is_split=False) == 8.25
    
    def test_split_handling_fee(self):
        """拆零操作费：0.3元/最小单位"""
        assert calc_handling_fee(10, is_split=True) == 3.0
        assert calc_handling_fee(1, is_split=True) == 0.3
        assert calc_handling_fee(5.5, is_split=True) == 1.65
    
    def test_zero_units(self):
        """0件操作费为0"""
        assert calc_handling_fee(0, is_split=False) == 0.0
        assert calc_handling_fee(0, is_split=True) == 0.0
    
    def test_negative_units(self):
        """负数件操作费为0"""
        assert calc_handling_fee(-5, is_split=False) == 0.0
        assert calc_handling_fee(-5, is_split=True) == 0.0
    
    def test_default_is_standard(self):
        """默认是整件操作"""
        assert calc_handling_fee(10) == 15.0
        assert calc_handling_fee(10, is_split=False) == 15.0
    
    def test_fractional_units(self):
        """支持小数件数"""
        assert calc_handling_fee(2.5, is_split=False) == 3.75
        assert calc_handling_fee(2.5, is_split=True) == 0.75
    
    def test_rounding(self):
        """四舍五入保留2位小数"""
        assert calc_handling_fee(1.33, is_split=False) == 2.0
        assert calc_handling_fee(1.33, is_split=True) == 0.4


class TestCalcHandlingFeeDetail:
    """测试 calc_handling_fee_detail 函数"""
    
    def test_standard_fee_detail(self):
        """整件操作费详情"""
        result = calc_handling_fee_detail(10, is_split=False)
        assert result["fee"] == 15.0
        assert result["std_units"] == 10
        assert result["fee_type"] == "整件"
        assert result["unit_price"] == 1.5
        assert "10 件 × 1.5 元 = 15.0 元" in result["calculation"]
    
    def test_split_fee_detail(self):
        """拆零操作费详情"""
        result = calc_handling_fee_detail(10, is_split=True)
        assert result["fee"] == 3.0
        assert result["std_units"] == 10
        assert result["fee_type"] == "拆零"
        assert result["unit_price"] == 0.3
        assert "10 件 × 0.3 元 = 3.0 元" in result["calculation"]
    
    def test_zero_units_detail(self):
        """0件的详情"""
        result = calc_handling_fee_detail(0, is_split=False)
        assert result["fee"] == 0.0
        assert result["calculation"] == "0 件 × 0 元 = 0 元"


class TestGetHandlingFeeRates:
    """测试 get_handling_fee_rates 函数"""
    
    def test_rates_structure(self):
        """费率结构正确"""
        rates = get_handling_fee_rates()
        assert "standard" in rates
        assert "split" in rates
    
    def test_rate_values(self):
        """费率数值正确"""
        rates = get_handling_fee_rates()
        assert rates["standard"]["rate"] == 1.5
        assert rates["split"]["rate"] == 0.3
    
    def test_rate_units(self):
        """费率单位正确"""
        rates = get_handling_fee_rates()
        assert rates["standard"]["unit"] == "元/标准件"
        assert rates["split"]["unit"] == "元/最小单位"


if __name__ == "__main__":
    pytest.main([__file__, "-v"])