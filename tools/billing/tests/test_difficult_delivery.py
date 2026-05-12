# billing_tools/tests/test_difficult_delivery.py
# 困难配送加收单元测试

import unittest
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "..", "src"))
from calculator.difficult_delivery import (
    is_difficult_delivery,
    calc_difficult_surcharge,
    get_difficult_delivery_info,
    difficult_keywords,
)


class TestDifficultDelivery(unittest.TestCase):
    """困难配送判定测试"""

    def test_is_difficult_delivery(self):
        """困难配送判定"""
        # 困难配送
        difficult, _ = is_difficult_delivery("永辉超市")
        self.assertTrue(difficult)

        difficult, _ = is_difficult_delivery("万达广场")
        self.assertTrue(difficult)

        difficult, _ = is_difficult_delivery("成都春熙路步行街")
        self.assertTrue(difficult)

        difficult, kw = is_difficult_delivery("家乐福超市成都高新店")
        self.assertTrue(difficult)
        self.assertEqual(kw, "超市")

        difficult, _ = is_difficult_delivery("某城市 mall")
        self.assertTrue(difficult)

        # 普通配送（不含困难配送关键词）
        difficult, _ = is_difficult_delivery("普通便利店")
        self.assertFalse(difficult)

        difficult, _ = is_difficult_delivery("街角杂货铺")
        self.assertFalse(difficult)

        # 空字符串
        difficult, _ = is_difficult_delivery("")
        self.assertFalse(difficult)

    def test_calc_difficult_surcharge(self):
        """困难配送加收计算"""
        # 困难配送：加收20%
        fee = calc_difficult_surcharge(100.0, "永辉超市")
        self.assertAlmostEqual(fee, 20.0)

        fee = calc_difficult_surcharge(500.0, "万达广场")
        self.assertAlmostEqual(fee, 100.0)

        # 普通配送：无加收
        fee = calc_difficult_surcharge(100.0, "普通便利店")
        self.assertAlmostEqual(fee, 0.0)

        fee = calc_difficult_surcharge(100.0, "")
        self.assertAlmostEqual(fee, 0.0)

    def test_get_difficult_delivery_info(self):
        """完整信息返回"""
        info = get_difficult_delivery_info(100.0, "永辉超市")
        self.assertTrue(info.is_difficult)
        self.assertEqual(info.matched_keyword, "超市")
        self.assertAlmostEqual(info.surcharge_rate, 0.20)
        self.assertAlmostEqual(info.surcharge_amount, 20.0)

        info = get_difficult_delivery_info(100.0, "普通便利店")
        self.assertFalse(info.is_difficult)
        self.assertEqual(info.matched_keyword, "")
        self.assertAlmostEqual(info.surcharge_rate, 0.0)
        self.assertAlmostEqual(info.surcharge_amount, 0.0)


if __name__ == "__main__":
    unittest.main()
