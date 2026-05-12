# billing_tools/tests/test_value_added.py
# 增值服务费单元测试

import unittest
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "..", "src"))
from calculator.value_added import (
    calc_value_added_fee,
    calc_all_value_added,
    get_value_added_summary,
    value_added_rules,
)


class TestValueAddedFee(unittest.TestCase):
    """增值服务费计算测试"""

    def test_fee_types(self):
        """验证各服务类型的费率"""
        self.assertEqual(value_added_rules["灌装分箱"], 0.5)
        self.assertEqual(value_added_rules["再包装"], 3.0)
        self.assertEqual(value_added_rules["贴标"], 0.3)
        self.assertEqual(value_added_rules["标签"], 0.2)

    def test_calc_value_added_fee(self):
        """单类增值服务费计算"""
        # 灌装分箱
        self.assertAlmostEqual(calc_value_added_fee("灌装分箱", 100), 50.0)
        self.assertAlmostEqual(calc_value_added_fee("灌装分箱", 0), 0.0)
        self.assertAlmostEqual(calc_value_added_fee("灌装分箱", -10), 0.0)

        # 再包装
        self.assertAlmostEqual(calc_value_added_fee("再包装", 50), 150.0)

        # 未知类型返回0
        self.assertAlmostEqual(calc_value_added_fee("未知服务", 10), 0.0)

    def test_calc_all_value_added(self):
        """批量计算所有增值服务费"""
        order = {
            "灌装分箱": 200,
            "再包装": 30,
            "贴标": 150,
            "标签": 100,
        }
        result = calc_all_value_added(order)

        self.assertAlmostEqual(result["灌装分箱"], 100.0)
        self.assertAlmostEqual(result["再包装"], 90.0)
        self.assertAlmostEqual(result["贴标"], 45.0)
        self.assertAlmostEqual(result["标签"], 20.0)

    def test_empty_order(self):
        """空订单"""
        result = calc_all_value_added({})
        self.assertEqual(len(result), 0)

        result = calc_all_value_added({"灌装分箱": 0, "再包装": 0})
        self.assertEqual(len(result), 0)

    def test_summary(self):
        """汇总信息"""
        order = {"灌装分箱": 100, "贴标": 50}
        summary = get_value_added_summary(order)

        self.assertEqual(len(summary["breakdown"]), 2)
        self.assertAlmostEqual(summary["total"], 65.0)
        self.assertTrue(summary["has_charges"])


if __name__ == "__main__":
    unittest.main()
