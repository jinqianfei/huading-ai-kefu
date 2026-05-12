# -*- coding: utf-8 -*-
"""
知识库模块

导出类和函数：
- RuleStore: 规则库存储管理器
- RuleFactory: 规则工厂
- RuleStoreError: 规则库异常
- RuleValidationError: 规则校验异常

便捷函数：
- save_rule()
- load_rule()
- delete_rule()
- list_rules()
- confirm_rule()
- is_confirmed()
- get_storage_price()
- get_transport_price()
- get_delivery_fee()
- get_special_product()
"""

from .rule_store import (
    RuleStore,
    RuleFactory,
    RuleStoreError,
    RuleValidationError,
    save_rule,
    load_rule,
    delete_rule,
    list_rules,
    confirm_rule,
    is_confirmed,
    get_storage_price,
    get_transport_price,
    get_delivery_fee,
    get_special_product,
)

__all__ = [
    "RuleStore",
    "RuleFactory",
    "RuleStoreError",
    "RuleValidationError",
    "save_rule",
    "load_rule",
    "delete_rule",
    "list_rules",
    "confirm_rule",
    "is_confirmed",
    "get_storage_price",
    "get_transport_price",
    "get_delivery_fee",
    "get_special_product",
]
