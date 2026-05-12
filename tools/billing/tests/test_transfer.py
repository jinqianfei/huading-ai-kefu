# billing_tools/tests/test_transfer.py
# 调拨计价单元测试

import unittest
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "..", "src"))
from calculator.transfer import (
    calc_transfer_fee,
    calc_combined_transport_fee,
)


class TestTransferFee(unittest.TestCase):
    """调拨计价测试"""

    def test_basic_calc(self):
        """基本调拨费计算"""
        result = calc_transfer_fee(50, 0.15)
        self.assertAlmostEqual(result.std_units, 50)
        self.assertAlmostEqual(result.transfer_price, 0.15)
        self.assertAlmostEqual(result.calculated_fee, 7.5)
        self.assertEqual(result.formula, "50 × 0.15 = 7.5")

    def test_zero_values(self):
        """零值处理"""
        result = calc_transfer_fee(0, 0.15)
        self.assertAlmostEqual(result.calculated_fee, 0.0)
        self.assertEqual(result.formula, "无调拨")

        result = calc_transfer_fee(50, 0)
        self.assertAlmostEqual(result.calculated_fee, 0.0)

        result = calc_transfer_fee(-10, 0.15)
        self.assertAlmostEqual(result.calculated_fee, 0.0)

    def test_combined_fee(self):
        """综合运输费计算"""
        result = calc_combined_transport_fee(
            std_units=50,
            shared_price=15,
            shared_min_fee=260,
            shared_fee=750,
            transfer_price=0.15,
            transfer_fee=7.5,
        )

        self.assertEqual(result["shared_fee"], 750)    # 750 > 260，不触发最低
        self.assertEqual(result["used_min_fee"], False)
        self.assertEqual(result["transfer_fee"], 7.5)
        self.assertEqual(result["total_fee"], 757.5)

    def test_combined_with_min_fee(self):
        """触发共配最低收费"""
        result = calc_combined_transport_fee(
            std_units=10,
            shared_price=15,
            shared_min_fee=260,
            shared_fee=100,   # 100 < 260，触发最低
            transfer_price=0.15,
            transfer_fee=1.5,
        )

        self.assertEqual(result["shared_fee"], 260)   # 触发最低
        self.assertEqual(result["used_min_fee"], True)
        self.assertEqual(result["transfer_fee"], 1.5)
        self.assertEqual(result["total_fee"], 261.5)


if __name__ == "__main__":
    unittest.main()
