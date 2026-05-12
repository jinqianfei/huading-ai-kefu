# -*- coding: utf-8 -*-
"""
Excel 读取模块

导出类和函数：
- ExcelReader: 基础Excel读取器
- OutboundReader: 出库单导出明细读取器
- StorageReader: 仓储透视表读取器
- OrderPivotReader: 订单明细透视读取器
- CustomerInfoReader: 客户信息导出读取器
- ProductInfoReader: 商品信息导出读取器

便捷函数：
- read_outbound_details()
- read_storage_pivot()
- read_order_pivot()
- read_customer_info()
- read_product_info()
- read_all_excel_data()
"""

from .reader import (
    ExcelReader,
    ExcelReadError,
    OutboundReader,
    StorageReader,
    OrderPivotReader,
    CustomerInfoReader,
    ProductInfoReader,
    read_outbound_details,
    read_storage_pivot,
    read_order_pivot,
    read_customer_info,
    read_product_info,
    read_all_excel_data,
)

__all__ = [
    "ExcelReader",
    "ExcelReadError",
    "OutboundReader",
    "StorageReader",
    "OrderPivotReader",
    "CustomerInfoReader",
    "ProductInfoReader",
    "read_outbound_details",
    "read_storage_pivot",
    "read_order_pivot",
    "read_customer_info",
    "read_product_info",
    "read_all_excel_data",
]
