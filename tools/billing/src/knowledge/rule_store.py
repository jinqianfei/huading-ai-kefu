# -*- coding: utf-8 -*-
"""
规则库模块

负责存储和管理客户计费规则：
- 规则库JSON结构设计（参考PRD 6.1）
- 规则存储和读取
- 规则版本管理
- 规则与客户关联

参考PRD文档：AI账单助手 v1.1
"""

import json
import logging
import shutil
from datetime import datetime
from pathlib import Path
from typing import Any, Dict, List, Optional, Union

# 配置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class RuleStoreError(Exception):
    """规则库异常"""
    pass


class RuleValidationError(RuleStoreError):
    """规则校验异常"""
    pass


# ============================================================
# 规则库数据结构
# ============================================================

class RuleStore:
    """
    规则库存储管理器
    
    规则库JSON结构（参考PRD 6.1）：
    
    ```json
    {
      "customer_id": "CUST_001",
      "customer_name": "长沙广承供应链有限公司",
      "contract_no": "HDLL-CP-2025-286",
      "effective_date": "2025-07-01",
      "expiry_date": null,
      "warehouses": ["成都仓", "长沙仓", "广州仓", "杭州仓", "廊坊仓"],
      "rules": {
        "storage": {
          "常温": {"price": 1.7, "unit": "元/托/天", "tax_rate": 0.06},
          "冷冻": {"price": 3.2, "unit": "元/托/天", "tax_rate": 0.06},
          "冷藏": {"price": 3.2, "unit": "元/托/天", "tax_rate": 0.06}
        },
        "transport": {
          "default": {"min_fee": 40, "unit_price": 3.5},
          "regions": {
            "成都->云南全境": {"min_fee": 260, "unit_price": 15},
            "成都->重庆全境": {"min_fee": 120, "unit_price": 9},
            "长沙->长沙": {"min_fee": 50, "unit_price": 8}
          }
        },
        "delivery_fee": {
          "price": 1,
          "unit": "元/件",
          "effective_date": "2025-07-16"
        },
        "outbound": {
          "托盘费": {"price": 1.5, "unit": "元/标准件", "tax_rate": 0.06}
        },
        "special_products": [
          {"sku": "YX02", "name": "鱼皮", "coefficient": 1.0},
          {"sku": "TSP01", "name": "300ml糖水瓶", "coefficient": 1.0}
        ]
      },
      "adjustments": [
        {
          "month": "2026-02",
          "type": "统配出库费",
          "amount": -495.02,
          "note": "人工优惠"
        }
      ],
      "confidence": {
        "overall": 0.95,
        "details": {
          "storage": 0.98,
          "transport": 0.92,
          "delivery_fee": 0.90
        }
      },
      "metadata": {
        "created_at": "2025-07-01T10:00:00",
        "updated_at": "2025-07-01T10:00:00",
        "created_by": "system",
        "confirmed": false,
        "confirmed_at": null,
        "confirmed_by": null
      }
    }
    ```
    """
    
    # 默认规则库目录
    DEFAULT_RULE_DIR = Path(__file__).parent.parent.parent / "data" / "rules"
    DEFAULT_RULE_FILE = "rules.json"
    
    # 规则版本文件后缀
    VERSION_SUFFIX = ".v{version}"
    VERSION_DIR = "versions"
    
    def __init__(
        self,
        rule_dir: Optional[Union[str, Path]] = None,
        auto_backup: bool = True
    ):
        """
        初始化规则库
        
        Args:
            rule_dir: 规则库目录，默认使用项目data/rules目录
            auto_backup: 是否自动备份旧版本
        """
        self.rule_dir = Path(rule_dir) if rule_dir else self.DEFAULT_RULE_DIR
        self.auto_backup = auto_backup
        
        # 确保目录存在
        self.rule_dir.mkdir(parents=True, exist_ok=True)
        
        # 版本目录
        self.version_dir = self.rule_dir / self.VERSION_DIR
        self.version_dir.mkdir(parents=True, exist_ok=True)
        
        logger.info(f"规则库初始化完成: {self.rule_dir}")
    
    # ============================================================
    # 规则文件路径
    # ============================================================
    
    def _get_rule_file(self, customer_id: str) -> Path:
        """获取客户规则文件路径"""
        return self.rule_dir / f"{customer_id}.json"
    
    def _get_version_file(self, customer_id: str, version: str) -> Path:
        """获取指定版本的规则文件路径"""
        filename = f"{customer_id}.v{version}.json"
        return self.version_dir / filename
    
    # ============================================================
    # 规则基础操作
    # ============================================================
    
    def save_rule(self, rule: Dict[str, Any], auto_backup: bool = True) -> str:
        """
        保存规则到规则库
        
        Args:
            rule: 规则数据（必须包含customer_id）
            auto_backup: 是否自动备份旧版本
            
        Returns:
            customer_id
            
        Raises:
            RuleValidationError: 规则校验失败
        """
        # 校验规则数据
        self._validate_rule(rule)
        
        customer_id = rule["customer_id"]
        rule_file = self._get_rule_file(customer_id)
        
        # 备份旧版本
        if auto_backup and self.auto_backup and rule_file.exists():
            self._backup_rule(customer_id)
        
        # 添加元数据
        rule = self._update_metadata(rule, is_new=not rule_file.exists())
        
        # 保存规则
        try:
            with open(rule_file, "w", encoding="utf-8") as f:
                json.dump(rule, f, ensure_ascii=False, indent=2)
            
            logger.info(f"规则保存成功: {customer_id}")
            return customer_id
            
        except Exception as e:
            raise RuleStoreError(f"保存规则失败: {e}")
    
    def load_rule(self, customer_id: str) -> Dict[str, Any]:
        """
        加载客户规则
        
        Args:
            customer_id: 客户ID
            
        Returns:
            规则数据
            
        Raises:
            RuleStoreError: 规则不存在或加载失败
        """
        rule_file = self._get_rule_file(customer_id)
        
        if not rule_file.exists():
            raise RuleStoreError(f"规则不存在: {customer_id}")
        
        try:
            with open(rule_file, "r", encoding="utf-8") as f:
                rule = json.load(f)
            
            logger.info(f"规则加载成功: {customer_id}")
            return rule
            
        except json.JSONDecodeError as e:
            raise RuleStoreError(f"规则文件格式错误: {e}")
        except Exception as e:
            raise RuleStoreError(f"加载规则失败: {e}")
    
    def delete_rule(self, customer_id: str, backup: bool = True) -> bool:
        """
        删除客户规则
        
        Args:
            customer_id: 客户ID
            backup: 是否备份后再删除
            
        Returns:
            是否删除成功
        """
        rule_file = self._get_rule_file(customer_id)
        
        if not rule_file.exists():
            return False
        
        try:
            # 备份
            if backup:
                self._backup_rule(customer_id)
            
            # 删除
            rule_file.unlink()
            logger.info(f"规则已删除: {customer_id}")
            return True
            
        except Exception as e:
            raise RuleStoreError(f"删除规则失败: {e}")
    
    def exists(self, customer_id: str) -> bool:
        """检查规则是否存在"""
        return self._get_rule_file(customer_id).exists()
    
    def list_rules(self) -> List[str]:
        """列出所有客户ID"""
        customer_ids = []
        for f in self.rule_dir.glob("*.json"):
            if f.name != self.DEFAULT_RULE_FILE:
                customer_ids.append(f.stem)
        return sorted(customer_ids)
    
    # ============================================================
    # 规则版本管理
    # ============================================================
    
    def _backup_rule(self, customer_id: str) -> Optional[str]:
        """
        备份当前规则版本
        
        Args:
            customer_id: 客户ID
            
        Returns:
            备份版本号
        """
        rule_file = self._get_rule_file(customer_id)
        
        if not rule_file.exists():
            return None
        
        try:
            # 读取当前规则
            with open(rule_file, "r", encoding="utf-8") as f:
                rule = json.load(f)
            
            # 生成版本号（时间戳）
            version = datetime.now().strftime("%Y%m%d%H%M%S")
            
            # 获取元数据中的版本信息
            metadata = rule.get("metadata", {})
            old_version = metadata.get("version", "1.0")
            
            # 更新元数据
            metadata["version"] = version
            metadata["previous_version"] = old_version
            metadata["backup_at"] = datetime.now().isoformat()
            rule["metadata"] = metadata
            
            # 保存版本文件
            version_file = self._get_version_file(customer_id, version)
            with open(version_file, "w", encoding="utf-8") as f:
                json.dump(rule, f, ensure_ascii=False, indent=2)
            
            logger.info(f"规则已备份: {customer_id} -> v{version}")
            return version
            
        except Exception as e:
            logger.warning(f"备份规则失败: {customer_id} - {e}")
            return None
    
    def get_version(self, customer_id: str, version: str) -> Dict[str, Any]:
        """
        获取指定版本的规则
        
        Args:
            customer_id: 客户ID
            version: 版本号
            
        Returns:
            规则数据
        """
        version_file = self._get_version_file(customer_id, version)
        
        if not version_file.exists():
            raise RuleStoreError(f"规则版本不存在: {customer_id} v{version}")
        
        try:
            with open(version_file, "r", encoding="utf-8") as f:
                return json.load(f)
        except Exception as e:
            raise RuleStoreError(f"读取规则版本失败: {e}")
    
    def list_versions(self, customer_id: str) -> List[Dict[str, str]]:
        """
        列出所有版本
        
        Args:
            customer_id: 客户ID
            
        Returns:
            版本列表 [{version, backup_at, ...}, ...]
        """
        versions = []
        pattern = f"{customer_id}.v*.json"
        
        for vf in self.version_dir.glob(pattern):
            try:
                with open(vf, "r", encoding="utf-8") as f:
                    rule = json.load(f)
                
                metadata = rule.get("metadata", {})
                versions.append({
                    "version": metadata.get("version", vf.stem.split(".v")[1]),
                    "backup_at": metadata.get("backup_at", ""),
                    "file": str(vf.name),
                })
            except Exception:
                continue
        
        return sorted(versions, key=lambda x: x.get("backup_at", ""), reverse=True)
    
    def restore_version(self, customer_id: str, version: str) -> Dict[str, Any]:
        """
        恢复指定版本
        
        Args:
            customer_id: 客户ID
            version: 版本号
            
        Returns:
            恢复后的规则数据
        """
        # 获取指定版本
        old_rule = self.get_version(customer_id, version)
        
        # 备份当前版本
        self._backup_rule(customer_id)
        
        # 恢复（更新元数据）
        old_rule["metadata"]["restored_from_version"] = version
        old_rule["metadata"]["restored_at"] = datetime.now().isoformat()
        
        # 保存为当前版本
        self.save_rule(old_rule, auto_backup=False)
        
        return old_rule
    
    # ============================================================
    # 规则校验
    # ============================================================
    
    def _validate_rule(self, rule: Dict[str, Any]) -> None:
        """
        校验规则数据
        
        Args:
            rule: 规则数据
            
        Raises:
            RuleValidationError: 校验失败
        """
        required_fields = ["customer_id", "customer_name", "rules"]
        
        for field in required_fields:
            if field not in rule:
                raise RuleValidationError(f"缺少必填字段: {field}")
        
        # 校验rules结构
        rules = rule.get("rules", {})
        
        # 仓储费规则
        if "storage" in rules:
            for zone, config in rules["storage"].items():
                if "price" not in config:
                    raise RuleValidationError(f"仓储费缺少price字段: {zone}")
        
        # 运输费规则
        if "transport" in rules:
            transport = rules["transport"]
            if "default" not in transport:
                raise RuleValidationError("运输费缺少default配置")
    
    @staticmethod
    def _update_metadata(rule: Dict[str, Any], is_new: bool = False) -> Dict[str, Any]:
        """更新元数据"""
        metadata = rule.get("metadata", {})
        
        if is_new:
            metadata["created_at"] = datetime.now().isoformat()
        
        metadata["updated_at"] = datetime.now().isoformat()
        rule["metadata"] = metadata
        
        return rule
    
    # ============================================================
    # 规则确认机制（必须环节）
    # ============================================================
    
    def confirm_rule(self, customer_id: str, confirmed_by: str = "system") -> Dict[str, Any]:
        """
        确认规则（人工确认环节）
        
        ⚠️ 重要：未经客服确认的规则，AI不得用于账单计算
        
        Args:
            customer_id: 客户ID
            confirmed_by: 确认人
            
        Returns:
            确认后的规则数据
        """
        rule = self.load_rule(customer_id)
        
        # 更新确认状态
        metadata = rule.get("metadata", {})
        metadata["confirmed"] = True
        metadata["confirmed_at"] = datetime.now().isoformat()
        metadata["confirmed_by"] = confirmed_by
        rule["metadata"] = metadata
        
        # 保存确认状态
        self.save_rule(rule, auto_backup=False)
        
        logger.info(f"规则已确认: {customer_id} by {confirmed_by}")
        return rule
    
    def is_confirmed(self, customer_id: str) -> bool:
        """检查规则是否已确认"""
        try:
            rule = self.load_rule(customer_id)
            return rule.get("metadata", {}).get("confirmed", False)
        except RuleStoreError:
            return False
    
    # ============================================================
    # 规则便捷查询
    # ============================================================
    
    def get_storage_price(
        self, 
        customer_id: str, 
        zone: str
    ) -> Optional[Dict[str, Any]]:
        """
        获取仓储价格
        
        Args:
            customer_id: 客户ID
            zone: 温区（常温/冷冻/冷藏）
            
        Returns:
            {price, unit, tax_rate}
        """
        rule = self.load_rule(customer_id)
        storage = rule.get("rules", {}).get("storage", {})
        return storage.get(zone)
    
    def get_transport_price(
        self,
        customer_id: str,
        warehouse: str,
        region: str
    ) -> Optional[Dict[str, Any]]:
        """
        获取运输价格
        
        Args:
            customer_id: 客户ID
            warehouse: 始发仓库
            region: 目的区域
            
        Returns:
            {min_fee, unit_price}
        """
        rule = self.load_rule(customer_id)
        transport = rule.get("rules", {}).get("transport", {})
        
        # 先尝试精确匹配
        route_key = f"{warehouse}->{region}"
        regions = transport.get("regions", {})
        
        if route_key in regions:
            return regions[route_key]
        
        # 尝试通配匹配（如"成都->云南全境"）
        for route_pattern, config in regions.items():
            if warehouse in route_pattern and region in route_pattern:
                return config
        
        # 返回默认配置
        return transport.get("default")
    
    def get_delivery_fee(self, customer_id: str) -> Optional[Dict[str, Any]]:
        """
        获取送货进店费
        
        Args:
            customer_id: 客户ID
            
        Returns:
            {price, unit, effective_date}
        """
        rule = self.load_rule(customer_id)
        return rule.get("rules", {}).get("delivery_fee")
    
    def get_special_product(
        self, 
        customer_id: str, 
        sku: str
    ) -> Optional[Dict[str, Any]]:
        """
        获取特殊产品计费系数
        
        Args:
            customer_id: 客户ID
            sku: 商品编码
            
        Returns:
            {sku, name, coefficient}
        """
        rule = self.load_rule(customer_id)
        special_products = rule.get("rules", {}).get("special_products", [])
        
        for product in special_products:
            if product.get("sku") == sku:
                return product
        
        return None
    
    def get_adjustments(
        self, 
        customer_id: str, 
        month: Optional[str] = None
    ) -> List[Dict[str, Any]]:
        """
        获取人工调整记录
        
        Args:
            customer_id: 客户ID
            month: 账期（YYYY-MM格式），不指定则返回全部
            
        Returns:
            调整记录列表
        """
        rule = self.load_rule(customer_id)
        adjustments = rule.get("adjustments", [])
        
        if month:
            return [adj for adj in adjustments if adj.get("month") == month]
        
        return adjustments
    
    # ============================================================
    # 规则导入导出
    # ============================================================
    
    def export_rule(self, customer_id: str, export_path: Optional[Path] = None) -> Path:
        """
        导出规则到指定路径
        
        Args:
            customer_id: 客户ID
            export_path: 导出路径，默认导出一份到rules目录
            
        Returns:
            导出文件路径
        """
        rule = self.load_rule(customer_id)
        
        if export_path is None:
            export_path = self.rule_dir / f"{customer_id}_export.json"
        
        with open(export_path, "w", encoding="utf-8") as f:
            json.dump(rule, f, ensure_ascii=False, indent=2)
        
        return export_path
    
    def import_rule(self, import_path: Union[str, Path], customer_id: Optional[str] = None) -> str:
        """
        从文件导入规则
        
        Args:
            import_path: 导入文件路径
            customer_id: 可选，指定客户ID覆盖原文件中的customer_id
            
        Returns:
            导入的客户ID
        """
        import_path = Path(import_path)
        
        if not import_path.exists():
            raise RuleStoreError(f"导入文件不存在: {import_path}")
        
        try:
            with open(import_path, "r", encoding="utf-8") as f:
                rule = json.load(f)
            
            # 覆盖customer_id
            if customer_id:
                rule["customer_id"] = customer_id
            
            # 保存
            return self.save_rule(rule)
            
        except json.JSONDecodeError as e:
            raise RuleStoreError(f"导入文件格式错误: {e}")


# ============================================================
# 规则工厂（用于创建新规则）
# ============================================================

class RuleFactory:
    """规则工厂"""
    
    @staticmethod
    def create_empty_rule(
        customer_id: str,
        customer_name: str,
        contract_no: Optional[str] = None,
        effective_date: Optional[str] = None
    ) -> Dict[str, Any]:
        """
        创建空白规则模板
        
        Args:
            customer_id: 客户ID
            customer_name: 客户名称
            contract_no: 合同编号
            effective_date: 生效日期
            
        Returns:
            空白规则模板
        """
        return {
            "customer_id": customer_id,
            "customer_name": customer_name,
            "contract_no": contract_no or "",
            "effective_date": effective_date or datetime.now().strftime("%Y-%m-%d"),
            "expiry_date": None,
            "warehouses": [],
            "rules": {
                "storage": {},
                "transport": {
                    "default": {"min_fee": 0, "unit_price": 0}
                },
                "delivery_fee": {},
                "outbound": {},
                "special_products": []
            },
            "adjustments": [],
            "confidence": {
                "overall": 0.0,
                "details": {}
            },
            "metadata": {
                "created_at": datetime.now().isoformat(),
                "updated_at": datetime.now().isoformat(),
                "created_by": "system",
                "confirmed": False,
                "confirmed_at": None,
                "confirmed_by": None
            }
        }
    
    @staticmethod
    def create_from_contract(
        customer_id: str,
        customer_name: str,
        contract_data: Dict[str, Any]
    ) -> Dict[str, Any]:
        """
        从合同数据创建规则
        
        Args:
            customer_id: 客户ID
            customer_name: 客户名称
            contract_data: 合同解析数据
            
        Returns:
            规则数据
        """
        rule = RuleFactory.create_empty_rule(customer_id, customer_name)
        
        # 填充仓储费
        if "storage" in contract_data:
            rule["rules"]["storage"] = contract_data["storage"]
        
        # 填充运输费
        if "transport" in contract_data:
            rule["rules"]["transport"] = contract_data["transport"]
        
        # 填充送货进店费
        if "delivery_fee" in contract_data:
            rule["rules"]["delivery_fee"] = contract_data["delivery_fee"]
        
        # 填充特殊产品
        if "special_products" in contract_data:
            rule["rules"]["special_products"] = contract_data["special_products"]
        
        # 填充置信度
        if "confidence" in contract_data:
            rule["confidence"] = contract_data["confidence"]
        
        return rule


# ============================================================
# 便捷函数
# ============================================================

# 全局默认规则库实例
_default_store: Optional[RuleStore] = None


def get_default_store() -> RuleStore:
    """获取默认规则库实例"""
    global _default_store
    if _default_store is None:
        _default_store = RuleStore()
    return _default_store


def save_rule(rule: Dict[str, Any]) -> str:
    """保存规则"""
    return get_default_store().save_rule(rule)


def load_rule(customer_id: str) -> Dict[str, Any]:
    """加载规则"""
    return get_default_store().load_rule(customer_id)


def delete_rule(customer_id: str) -> bool:
    """删除规则"""
    return get_default_store().delete_rule(customer_id)


def list_rules() -> List[str]:
    """列出所有规则"""
    return get_default_store().list_rules()


def confirm_rule(customer_id: str, confirmed_by: str = "system") -> Dict[str, Any]:
    """确认规则"""
    return get_default_store().confirm_rule(customer_id, confirmed_by)


def is_confirmed(customer_id: str) -> bool:
    """检查规则是否已确认"""
    return get_default_store().is_confirmed(customer_id)


def get_storage_price(customer_id: str, zone: str) -> Optional[Dict[str, Any]]:
    """获取仓储价格"""
    return get_default_store().get_storage_price(customer_id, zone)


def get_transport_price(
    customer_id: str, 
    warehouse: str, 
    region: str
) -> Optional[Dict[str, Any]]:
    """获取运输价格"""
    return get_default_store().get_transport_price(customer_id, warehouse, region)


def get_delivery_fee(customer_id: str) -> Optional[Dict[str, Any]]:
    """获取送货进店费"""
    return get_default_store().get_delivery_fee(customer_id)


def get_special_product(customer_id: str, sku: str) -> Optional[Dict[str, Any]]:
    """获取特殊产品计费系数"""
    return get_default_store().get_special_product(customer_id, sku)


if __name__ == "__main__":
    # 测试代码
    import sys
    
    print("=" * 50)
    print("规则库模块测试")
    print("=" * 50)
    
    # 创建测试规则
    test_rule = {
        "customer_id": "TEST_001",
        "customer_name": "测试客户",
        "contract_no": "TEST-CONTRACT-001",
        "effective_date": "2025-01-01",
        "warehouses": ["成都仓", "长沙仓"],
        "rules": {
            "storage": {
                "常温": {"price": 1.7, "unit": "元/托/天", "tax_rate": 0.06},
                "冷冻": {"price": 3.2, "unit": "元/托/天", "tax_rate": 0.06}
            },
            "transport": {
                "default": {"min_fee": 40, "unit_price": 3.5},
                "regions": {
                    "成都->云南全境": {"min_fee": 260, "unit_price": 15}
                }
            },
            "delivery_fee": {
                "price": 1,
                "unit": "元/件",
                "effective_date": "2025-07-16"
            },
            "outbound": {
                "托盘费": {"price": 1.5, "unit": "元/标准件", "tax_rate": 0.06}
            },
            "special_products": [
                {"sku": "YX02", "name": "鱼皮", "coefficient": 1.0}
            ]
        },
        "adjustments": [],
        "confidence": {
            "overall": 0.95,
            "details": {
                "storage": 0.98,
                "transport": 0.92
            }
        }
    }
    
    # 测试规则工厂
    print("\n1. 测试规则工厂...")
    empty_rule = RuleFactory.create_empty_rule("CUST_TEST", "测试客户")
    print(f"   创建空白规则: {empty_rule['customer_id']}")
    
    # 测试规则库
    print("\n2. 测试规则库...")
    store = RuleStore()
    
    try:
        # 保存规则
        print("   保存规则...")
        saved_id = store.save_rule(test_rule)
        print(f"   保存成功: {saved_id}")
        
        # 加载规则
        print("   加载规则...")
        loaded_rule = store.load_rule("TEST_001")
        print(f"   加载成功: {loaded_rule['customer_name']}")
        
        # 查询价格
        print("   查询仓储价格...")
        storage_price = store.get_storage_price("TEST_001", "常温")
        print(f"   常温价格: {storage_price}")
        
        # 列出所有规则
        print("   列出所有规则...")
        all_rules = store.list_rules()
        print(f"   规则列表: {all_rules}")
        
        # 确认规则
        print("   确认规则...")
        confirmed_rule = store.confirm_rule("TEST_001", "测试用户")
        print(f"   确认状态: {confirmed_rule['metadata']['confirmed']}")
        
        # 删除测试规则
        print("   删除测试规则...")
        store.delete_rule("TEST_001")
        print("   删除成功")
        
    except Exception as e:
        print(f"   错误: {e}")
    
    print("\n" + "=" * 50)
    print("测试完成")
    print("=" * 50)
