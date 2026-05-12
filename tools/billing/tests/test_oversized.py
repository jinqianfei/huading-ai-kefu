"""
超标件系数计算 - 单元测试
"""

import pytest
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'src', 'calculator'))
from oversized import calc_oversized_coefficient, get_oversized_tier


class TestOversizedCoefficient:
    """测试 calc_oversized_coefficient 函数"""
    
    def test_normal_item(self):
        """普通件：重量 ≤ 15kg 且 体积 ≤ 0.05m³"""
        assert calc_oversized_coefficient(10, 0.03) == 1.0
        assert calc_oversized_coefficient(15, 0.05) == 1.0
        assert calc_oversized_coefficient(5, 0.01) == 1.0
    
    def test_light_oversized_by_weight(self):
        """轻微超标（重量）: 15kg < 重量 ≤ 25kg"""
        assert calc_oversized_coefficient(20, 0.03) == 1.2
        assert calc_oversized_coefficient(25, 0.05) == 1.2
        assert calc_oversized_coefficient(16, 0.04) == 1.2
    
    def test_light_oversized_by_volume(self):
        """轻微超标（体积）: 0.05m³ < 体积 ≤ 0.08m³"""
        assert calc_oversized_coefficient(10, 0.06) == 1.2
        assert calc_oversized_coefficient(15, 0.08) == 1.2
        assert calc_oversized_coefficient(5, 0.07) == 1.2
    
    def test_medium_oversized_by_weight(self):
        """中度超标（重量）: 25kg < 重量 ≤ 50kg"""
        assert calc_oversized_coefficient(30, 0.03) == 2.0
        assert calc_oversized_coefficient(50, 0.05) == 2.0
        assert calc_oversized_coefficient(26, 0.04) == 2.0
    
    def test_medium_oversized_by_volume(self):
        """中度超标（体积）: 0.08m³ < 体积 ≤ 0.1m³"""
        assert calc_oversized_coefficient(10, 0.09) == 2.0
        assert calc_oversized_coefficient(15, 0.1) == 2.0
        assert calc_oversized_coefficient(5, 0.095) == 2.0
    
    def test_heavy_oversized_by_weight(self):
        """重度超标（重量）: 重量 > 50kg"""
        assert calc_oversized_coefficient(60, 0.03) == 4.0  # 60/15=4
        assert calc_oversized_coefficient(75, 0.03) == 5.0  # 75/15=5
    
    def test_heavy_oversized_by_volume(self):
        """重度超标（体积）: 体积 > 0.1m³"""
        result = calc_oversized_coefficient(10, 0.15)
        assert abs(result - 3.0) < 0.001  # 0.15/0.05=3, allow floating point tolerance
        assert calc_oversized_coefficient(15, 0.2) == 4.0   # 0.2/0.05=4
    
    def test_heavy_oversized_both(self):
        """重度超标（重量和体积同时超标）: 取较大值"""
        # weight: 60/15=4, volume: 0.15/0.05=3, 取较大值4
        assert calc_oversized_coefficient(60, 0.15) == 4.0
        # weight: 90/15=6, volume: 0.2/0.05=4, 取较大值6
        assert calc_oversized_coefficient(90, 0.2) == 6.0


class TestOversizedTier:
    """测试 get_oversized_tier 函数"""
    
    def test_normal_tier(self):
        assert get_oversized_tier(10, 0.03) == "普通件"
        assert get_oversized_tier(15, 0.05) == "普通件"
    
    def test_light_tier(self):
        assert get_oversized_tier(20, 0.03) == "轻微超标件"
        assert get_oversized_tier(10, 0.06) == "轻微超标件"
    
    def test_medium_tier(self):
        assert get_oversized_tier(30, 0.03) == "中度超标件"
        assert get_oversized_tier(10, 0.09) == "中度超标件"
    
    def test_heavy_tier(self):
        assert get_oversized_tier(60, 0.03) == "重度超标件"
        assert get_oversized_tier(10, 0.15) == "重度超标件"


if __name__ == "__main__":
    pytest.main([__file__, "-v"])