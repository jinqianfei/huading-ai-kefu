# -*- coding: utf-8 -*-
"""
订单映射Skill处理器
整合SKU映射、门店映射、仓库映射、单位转换，实现订单全量映射

调用流程：
1. 读取订单 → OrderReader.read_order()
2. 识别客户 → 自动识别或使用传入的customer_id
3. SKU匹配 → sku_mapping.match_customer_sku()
4. 门店匹配 → store_mapping.get_huading_store_code()
5. 仓库匹配 → warehouse_mapping.get_warehouse_code()
6. 单位转换 → unit_converter.get_unit_type()
7. 返回结果
"""

import logging
from dataclasses import dataclass, field
from pathlib import Path
from typing import Any, Dict, List, Optional, Union

# 导入mapper模块
from . import sku_mapping
from . import store_mapping
from . import warehouse_mapping
from . import unit_converter

# 导入修复后的订单读取器
from .order_reader import read_order as read_order_func

# 配置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


# ============================================================
# 数据结构
# ============================================================

@dataclass
class OrderMappingResult:
    """订单映射结果"""
    success: bool
    order_id: str
    customer_id: str
    items: List[Dict] = field(default_factory=list)  # 包含sku匹配结果
    store: Dict = field(default_factory=dict)        # 门店匹配结果
    warehouse: Dict = field(default_factory=dict)    # 仓库匹配结果
    unmapped: List = field(default_factory=list)      # 未匹配项
    low_confidence: List = field(default_factory=list)  # 低置信度项
    message: str = ""


@dataclass
class OrderItemResult:
    """订单明细映射结果"""
    original_item: Dict[str, Any]
    sku_code: Optional[str] = None
    sku_name: Optional[str] = None
    sku_confidence: float = 0.0
    sku_match_method: str = "none"
    unit_type: Optional[str] = None
    unit_confidence: float = 0.0
    status: str = "pending"  # pending/mapped/low_confidence/unmapped


# ============================================================
# 订单读取器
# ============================================================

class OrderReader:
    """
    订单Excel读取器
    
    支持华鼎系统导出的标准订单模板
    """
    
    # 标准表头
    STANDARD_HEADERS = [
        "订单号", "出库单号", "单号",
        "门店名称", "门店", "目的门店",
        "商品名称", "商品", "SKU", "商品编码",
        "数量", "实际出库数量", "订货数量",
        "单位", "计量方式", "unit",
        "仓库", "出库仓库", "始发仓库",
        "客户", "客户名称", "客户编码",
    ]
    
    def __init__(self, file_path: Union[str, Path]):
        """
        初始化订单读取器
        
        Args:
            file_path: Excel文件路径
        """
        self.file_path = Path(file_path)
        if not self.file_path.exists():
            raise FileNotFoundError(f"文件不存在: {self.file_path}")
        
        self._reader: Optional[ExcelReader] = None
        self._data: List[Dict[str, Any]] = []
        self._order_id: Optional[str] = None
        self._customer_id: Optional[str] = None
    
    def _detect_headers(self, headers: Dict[int, str]) -> Dict[str, str]:
        """
        检测并标准化表头
        
        Args:
            headers: {列索引: 列名}
            
        Returns:
            {标准字段名: 列名}
        """
        header_map = {}
        reverse_map = {v: k for k, v in headers.items()}
        
        # 订单号
        for h in ["订单号", "出库单号", "单号"]:
            if h in reverse_map:
                header_map["order_id"] = h
                break
        
        # 门店
        for h in ["门店名称", "门店", "目的门店"]:
            if h in reverse_map:
                header_map["store"] = h
                break
        
        # 商品
        for h in ["商品名称", "商品", "SKU", "商品编码"]:
            if h in reverse_map:
                header_map["product"] = h
                break
        
        # 数量
        for h in ["数量", "实际出库数量", "订货数量"]:
            if h in reverse_map:
                header_map["quantity"] = h
                break
        
        # 单位
        for h in ["单位", "计量方式", "unit"]:
            if h in reverse_map:
                header_map["unit"] = h
                break
        
        # 仓库
        for h in ["仓库", "出库仓库", "始发仓库"]:
            if h in reverse_map:
                header_map["warehouse"] = h
                break
        
        # 客户
        for h in ["客户", "客户名称", "客户编码"]:
            if h in reverse_map:
                header_map["customer"] = h
                break
        
        return header_map
    
    def read_order(self) -> Dict[str, Any]:
        """
        读取订单数据 - 使用修复后的order_reader
        """
        result = read_order_func(str(self.file_path))
        
        if not result.get("success"):
            raise RuntimeError(result.get("error", "读取失败"))
        
        self._order_id = result.get("order_id", "")
        self._customer_id = result.get("customer_name", "")
        
        self._data = []
        for item in result.get("items", []):
            self._data.append({
                "store": result.get("store_name", ""),
                "product": item.get("sku_name", ""),
                "sku_code": item.get("sku_code", ""),
                "quantity": item.get("quantity", 0),
                "unit": item.get("unit", ""),
                "warehouse": "",
                "_raw": item,
            })
        
        return {
            "order_id": str(self._order_id),
            "customer_id": str(self._customer_id),
            "items": self._data,
            "raw_data": result.get("items", []),
        }

    def order_id(self) -> Optional[str]:
        return self._order_id
    
    @property
    def customer_id(self) -> Optional[str]:
        return self._customer_id


# ============================================================
# 订单映射Skill
# ============================================================

class OrderMappingSkill:
    """
    订单映射Skill处理器
    
    功能：
    - 读取订单文件
    - 自动识别客户
    - SKU匹配（4层匹配策略）
    - 门店匹配（3级匹配策略）
    - 仓库映射
    - 单位转换
    """
    
    # 低置信度阈值
    LOW_CONFIDENCE_SKU_THRESHOLD = 0.70
    LOW_CONFIDENCE_STORE_THRESHOLD = 0.70
    
    def __init__(self):
        """初始化订单映射Skill"""
        self._order_reader: Optional[OrderReader] = None
        self._current_order: Optional[Dict[str, Any]] = None
    
    def process(
        self,
        order_file: Union[str, Path],
        customer_id: Optional[str] = None
    ) -> OrderMappingResult:
        """
        处理订单映射
        
        Args:
            order_file: 订单文件路径（Excel）
            customer_id: 客户ID（可选，不传则自动识别）
            
        Returns:
            OrderMappingResult: 映射结果
        """
        logger.info(f"开始处理订单: {order_file}")
        
        # Step 1: 读取订单
        try:
            self._order_reader = OrderReader(order_file)
            order_data = self._order_reader.read_order()
            self._current_order = order_data
        except Exception as e:
            logger.error(f"读取订单失败: {e}")
            return OrderMappingResult(
                success=False,
                order_id="",
                customer_id=customer_id or "",
                message=f"读取订单失败: {e}"
            )
        
        # Step 2: 确定客户ID
        resolved_customer_id = customer_id or order_data.get("customer_id", "")
        if not resolved_customer_id:
            logger.warning("无法确定客户ID，映射结果可能不完整")
        
        logger.info(f"客户ID: {resolved_customer_id}, 订单ID: {order_data.get('order_id')}")
        
        # Step 3: 初始化结果
        result = OrderMappingResult(
            success=True,
            order_id=order_data.get("order_id", ""),
            customer_id=resolved_customer_id,
        )
        
        # Step 4: 处理每个订单明细
        items = order_data.get("items", [])
        for idx, item in enumerate(items):
            item_result = self._process_order_item(
                item,
                resolved_customer_id,
                idx
            )
            result.items.append({
                "index": idx,
                "original": item,
                "mapped": item_result,
            })
            
            # 分类处理
            if item_result.status == "unmapped":
                result.unmapped.append({
                    "index": idx,
                    "item": item,
                    "reason": "SKU未匹配",
                })
            elif item_result.status == "low_confidence":
                result.low_confidence.append({
                    "index": idx,
                    "item": item,
                    "sku_confidence": item_result.sku_confidence,
                })
        
        # Step 5: 门店匹配（取第一条明细的门店）
        if items:
            first_store = items[0].get("store", "")
            if first_store and resolved_customer_id:
                store_result = store_mapping.get_huading_store_code(
                    resolved_customer_id,
                    first_store
                )
                result.store = store_result
                
                # 检查门店置信度
                if store_result.get("success") and store_result.get("confidence", 0) < self.LOW_CONFIDENCE_STORE_THRESHOLD:
                    result.low_confidence.append({
                        "type": "store",
                        "store": first_store,
                        "confidence": store_result.get("confidence", 0),
                    })
            else:
                result.store = {"success": False, "error": "门店信息为空"}
        
        # Step 6: 仓库匹配（基于门店）
        if result.store.get("success") and resolved_customer_id:
            store_code = result.store.get("huading_store_code", "")
            if store_code:
                warehouse_result = warehouse_mapping.get_warehouse_code(
                    resolved_customer_id,
                    store_code
                )
                result.warehouse = warehouse_result
        else:
            result.warehouse = {"success": False, "error": "门店未匹配，无法确定仓库"}
        
        # Step 7: 汇总结果
        unmapped_count = len(result.unmapped)
        low_conf_count = len(result.low_confidence)
        
        if unmapped_count == 0 and low_conf_count == 0:
            result.message = f"订单映射完成，共{len(items)}项，全部匹配成功"
        elif unmapped_count == 0:
            result.message = f"订单映射完成，共{len(items)}项，{low_conf_count}项低置信度"
        else:
            result.message = f"订单映射完成，共{len(items)}项，{unmapped_count}项未匹配，{low_conf_count}项低置信度"
        
        logger.info(result.message)
        return result
    
    def _process_order_item(
        self,
        item: Dict[str, Any],
        customer_id: str,
        idx: int
    ) -> OrderItemResult:
        """
        处理单个订单明细
        
        Args:
            item: 订单明细原始数据
            customer_id: 客户ID
            idx: 明细索引
            
        Returns:
            OrderItemResult: 明细映射结果
        """
        result = OrderItemResult(original_item=item)
        
        # 获取商品名称
        product_name = item.get("product", "") or item.get("商品名称", "")
        if not product_name:
            result.status = "unmapped"
            return result
        
        # SKU匹配（尝试商品名称和SKU编码）
        sku_result = {"success": False}
        
        # 优先尝试商品名称匹配
        if product_name:
            sku_result = sku_mapping.match_customer_sku(product_name, customer_id or "")
        
        # 如果名称匹配失败，尝试SKU编码匹配
        if not sku_result.get("success") or not sku_result.get("huading_sku_code"):
            sku_code = item.get("sku_code", "")
            if sku_code:
                sku_result_by_code = sku_mapping.match_customer_sku(sku_code, customer_id or "")
                if sku_result_by_code.get("success"):
                    sku_result = sku_result_by_code
        
        if sku_result.get("success"):
            result.sku_code = sku_result.get("huading_sku_code")
            result.sku_confidence = sku_result.get("confidence", 0)
            result.sku_match_method = sku_result.get("match_method", "none")
            
            if result.sku_code is None:
                result.status = "unmapped"
            elif result.sku_confidence < self.LOW_CONFIDENCE_SKU_THRESHOLD:
                result.status = "low_confidence"
            else:
                result.status = "mapped"
        else:
            result.status = "unmapped"
        
        # 单位转换
        unit = item.get("unit", "") or item.get("单位", "")
        if unit:
            unit_result = unit_converter.get_unit_type(unit)
            if unit_result.get("success"):
                result.unit_type = unit_result.get("unit_type")
                result.unit_confidence = unit_result.get("confidence", 0)
        
        return result


# ============================================================
# 便捷函数
# ============================================================

def process_order(
    order_file: Union[str, Path],
    customer_id: Optional[str] = None
) -> OrderMappingResult:
    """
    处理订单映射（便捷函数）
    
    Args:
        order_file: 订单文件路径
        customer_id: 客户ID
        
    Returns:
        OrderMappingResult
    """
    skill = OrderMappingSkill()
    return skill.process(order_file, customer_id)


# ============================================================
# 测试
# ============================================================

if __name__ == "__main__":
    print("=== 订单映射Skill测试 ===")
    
    # 测试OrderMappingSkill
    print("\n1. OrderMappingSkill初始化:")
    skill = OrderMappingSkill()
    print(f"   ✅ 初始化成功")
    
    print("\n2. OrderMappingResult结构:")
    test_result = OrderMappingResult(
        success=True,
        order_id="TEST_ORDER_001",
        customer_id="CUST_001",
        items=[
            {
                "index": 0,
                "original": {"product": "1000塑杯", "quantity": 10},
                "mapped": OrderItemResult(
                    original_item={"product": "1000塑杯"},
                    sku_code="SKU001",
                    sku_name="1000塑杯",
                    sku_confidence=0.95,
                    sku_match_method="EXACT",
                    unit_type="BOX",
                    unit_confidence=1.0,
                    status="mapped"
                )
            }
        ],
        store={"success": True, "huading_store_code": "CHANGSHA001"},
        warehouse={"success": True, "warehouse_code": "WH_CHANGSHA"},
        unmapped=[],
        low_confidence=[],
        message="测试订单"
    )
    print(f"   success: {test_result.success}")
    print(f"   order_id: {test_result.order_id}")
    print(f"   items count: {len(test_result.items)}")
    print(f"   unmapped count: {len(test_result.unmapped)}")
    print(f"   low_confidence count: {len(test_result.low_confidence)}")
    
    print("\n✅ 订单映射Skill测试完成")
