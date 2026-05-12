"""
标准件计算模块

根据华鼎冷链合同规则计算标准件：

标准件判定规则（来自合同HDLL-CP-2025-286）：
- 公式：标准件 = 实际数量 × 系数
- 系数 = max(单件重量/15, 单件体积/0.05)

优先级：
1. 有商品资料库标准数据 → 用标准重量/体积计算系数
2. 无标准数据 → 用实际出库货物的重量/体积计算系数
"""

from typing import Union, Tuple, Optional, Dict
from dataclasses import dataclass
import json
import os


@dataclass
class CoefficientResult:
    """计费系数计算结果"""
    single_unit_weight_kg: float   # 单件重量(kg)
    single_unit_volume_m3: float   # 单件体积(m³)
    weight_ratio: float           # 重量比 (weight/15)
    volume_ratio: float           # 体积比 (volume/0.05)
    coefficient: float             # 计费系数
    data_source: str              # 数据来源：'standard'（标准数据）或 'actual'（实际数据）
    threshold_info: list          # 触发阈值详情
    
    def __str__(self) -> str:
        return (
            f"计费系数计算结果:\n"
            f"  单件重量: {self.single_unit_weight_kg:.4f} kg\n"
            f"  单件体积: {self.single_unit_volume_m3:.6f} m³\n"
            f"  重量比: {self.weight_ratio:.4f} (重量/15)\n"
            f"  体积比: {self.volume_ratio:.4f} (体积/0.05)\n"
            f"  系数: {self.coefficient:.4f}\n"
            f"  数据来源: {self.data_source}\n"
            f"  详情: {', '.join(self.threshold_info) if self.threshold_info else '无'}"
        )


@dataclass
class StdUnitsResult:
    """标准件计算结果"""
    actual_qty: float            # 实际数量
    total_weight_kg: float        # 总重量(kg)
    total_volume_m3: float        # 总体积(m³)
    single_unit_weight_kg: float  # 单件重量
    single_unit_volume_m3: float  # 单件体积
    coefficient: float             # 计费系数
    data_source: str              # 数据来源
    std_units: float              # 标准件数
    formula: str                 # 计算公式
    
    def __str__(self) -> str:
        return (
            f"标准件计算结果:\n"
            f"  实际数量: {self.actual_qty}\n"
            f"  单件重量: {self.single_unit_weight_kg:.4f} kg\n"
            f"  单件体积: {self.single_unit_volume_m3:.6f} m³\n"
            f"  系数: {self.coefficient:.4f} ({self.data_source}数据)\n"
            f"  标准件: {self.std_units:.4f}\n"
            f"  公式: {self.formula}"
        )


# 合同规定的档位阈值
TIER_THRESHOLDS = {
    "tier1": {"weight": 15, "volume": 0.05, "coef": 1.0},
    "tier2": {"weight": 25, "volume": 0.08, "coef": 1.2},
    "tier3": {"weight": 30, "volume": 0.10, "coef": 2.0},
}


class ProductDatabase:
    """
    商品资料库
    
    存储每个SKU的标准重量和标准体积
    """
    
    def __init__(self, db_path: Optional[str] = None):
        """
        初始化商品资料库
        
        Args:
            db_path: 商品资料库文件路径，如果为None则使用默认路径
        """
        self.db_path = db_path or self._get_default_db_path()
        self._db: Dict[str, Dict] = {}
        self._load_db()
    
    def _get_default_db_path(self) -> str:
        """获取默认资料库路径"""
        # 优先从knowledge/rules目录加载
        base_path = os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(__file__)))))
        return os.path.join(base_path, 'knowledge', 'rules', 'product_standards.json')
    
    def _load_db(self):
        """加载商品资料库"""
        if os.path.exists(self.db_path):
            try:
                with open(self.db_path, 'r', encoding='utf-8') as f:
                    self._db = json.load(f)
                print(f"[ProductDatabase] 已加载 {len(self._db)} 个商品的标准数据")
            except Exception as e:
                print(f"[ProductDatabase] 加载资料库失败: {e}")
                self._db = {}
        else:
            print(f"[ProductDatabase] 资料库文件不存在: {self.db_path}")
            self._db = {}
    
    def get_standard(self, product_name: str) -> Optional[Dict[str, float]]:
        """
        获取商品的标准重量和体积
        
        Args:
            product_name: 商品名称或SKU编码
        
        Returns:
            包含标准重量(kg)和标准体积(m³)的字典，如果没有则返回None
        """
        # 精确匹配
        if product_name in self._db:
            return self._db[product_name]
        
        # 模糊匹配（包含）
        for name, data in self._db.items():
            if product_name and name in product_name:
                return data
        
        # 反向模糊匹配
        for name, data in self._db.items():
            if name and name in product_name:
                return data
        
        return None
    
    def add_standard(self, product_name: str, std_weight_kg: float, std_volume_m3: float):
        """
        添加商品标准数据
        
        Args:
            product_name: 商品名称
            std_weight_kg: 标准重量(kg)
            std_volume_m3: 标准体积(m³)
        """
        self._db[product_name] = {
            "standard_weight_kg": std_weight_kg,
            "standard_volume_m3": std_volume_m3
        }
    
    def save(self):
        """保存资料库到文件"""
        os.makedirs(os.path.dirname(self.db_path), exist_ok=True)
        with open(self.db_path, 'w', encoding='utf-8') as f:
            json.dump(self._db, f, ensure_ascii=False, indent=2)
        print(f"[ProductDatabase] 已保存 {len(self._db)} 个商品到 {self.db_path}")


# 全局商品资料库实例
_product_db: Optional[ProductDatabase] = None


def get_product_db() -> ProductDatabase:
    """获取全局商品资料库实例"""
    global _product_db
    if _product_db is None:
        _product_db = ProductDatabase()
    return _product_db


def set_product_db(db_path: str):
    """设置商品资料库路径并重新加载"""
    global _product_db
    _product_db = ProductDatabase(db_path)


def calc_coefficient(
    single_unit_weight_kg: Union[int, float],
    single_unit_volume_m3: Union[int, float],
    validate: bool = True
) -> CoefficientResult:
    """
    根据合同规则计算单件计费系数
    
    规则：
    - 系数 = max(单件重量/15, 单件体积/0.05)
    
    Args:
        single_unit_weight_kg: 单件重量（公斤）
        single_unit_volume_m3: 单件体积（立方米）
        validate: 是否进行参数校验
        
    Returns:
        CoefficientResult: 计算结果
    """
    # 类型检查
    if not isinstance(single_unit_weight_kg, (int, float)):
        raise TypeError(f"单件重量必须是数字类型，当前类型: {type(single_unit_weight_kg).__name__}")
    if not isinstance(single_unit_volume_m3, (int, float)):
        raise TypeError(f"单件体积必须是数字类型，当前类型: {type(single_unit_volume_m3).__name__}")
    
    # 参数校验
    if validate:
        if single_unit_weight_kg < 0:
            raise ValueError(f"单件重量不能为负数: {single_unit_weight_kg}")
        if single_unit_volume_m3 < 0:
            raise ValueError(f"单件体积不能为负数: {single_unit_volume_m3}")
    
    weight = float(single_unit_weight_kg)
    volume = float(single_unit_volume_m3)
    
    # 计算重量比和体积比
    weight_ratio = weight / 15
    volume_ratio = volume / 0.05
    
    # 取高者作为系数
    coefficient = max(weight_ratio, volume_ratio)
    
    threshold_info = []
    if weight_ratio >= volume_ratio:
        threshold_info.append(f"重量比 {weight_ratio:.4f} >= 体积比 {volume_ratio:.4f}")
    else:
        threshold_info.append(f"体积比 {volume_ratio:.4f} > 重量比 {weight_ratio:.4f}")
    
    return CoefficientResult(
        single_unit_weight_kg=round(weight, 4),
        single_unit_volume_m3=round(volume, 6),
        weight_ratio=round(weight_ratio, 4),
        volume_ratio=round(volume_ratio, 4),
        coefficient=round(coefficient, 4),
        data_source="calculated",
        threshold_info=threshold_info
    )


def calc_std_units(
    actual_qty: Union[int, float],
    total_weight_kg: Union[int, float],
    total_volume_m3: Union[int, float],
    product_name: Optional[str] = None,
    validate: bool = True
) -> StdUnitsResult:
    """
    计算标准件数
    
    公式：标准件 = 实际数量 × 系数
    
    系数计算规则：
    1. 如果商品资料库中有该商品的标准数据，用标准数据计算系数
    2. 否则用实际出库货物的重量/体积计算系数
    
    Args:
        actual_qty: 实际出库数量
        total_weight_kg: 总重量（公斤）
        total_volume_m3: 总体积（立方米）
        product_name: 商品名称（用于查找资料库）
        validate: 是否进行参数校验
        
    Returns:
        StdUnitsResult: 计算结果
    """
    # 类型检查
    if not isinstance(actual_qty, (int, float)):
        raise TypeError(f"实际数量必须是数字类型，当前类型: {type(actual_qty).__name__}")
    if not isinstance(total_weight_kg, (int, float)):
        raise TypeError(f"总重量必须是数字类型，当前类型: {type(total_weight_kg).__name__}")
    if not isinstance(total_volume_m3, (int, float)):
        raise TypeError(f"总体积必须是数字类型，当前类型: {type(total_volume_m3).__name__}")
    
    # 参数校验
    if validate:
        if actual_qty <= 0:
            raise ValueError(f"实际数量必须大于0: {actual_qty}")
        if total_weight_kg < 0:
            raise ValueError(f"总重量不能为负数: {total_weight_kg}")
        if total_volume_m3 < 0:
            raise ValueError(f"总体积不能为负数: {total_volume_m3}")
    
    actual_qty = float(actual_qty)
    total_weight = float(total_weight_kg)
    total_volume = float(total_volume_m3)
    
    # 计算单件重量和单件体积
    single_unit_weight = total_weight / actual_qty
    single_unit_volume = total_volume / actual_qty
    
    # 尝试从商品资料库获取标准数据
    data_source = "actual"
    std_weight = single_unit_weight
    std_volume = single_unit_volume
    
    if product_name:
        db = get_product_db()
        standard = db.get_standard(product_name)
        if standard:
            std_weight = standard.get("standard_weight_kg", single_unit_weight)
            std_volume = standard.get("standard_volume_m3", single_unit_volume)
            data_source = "standard"
    
    # 计算系数
    coef_result = calc_coefficient(std_weight, std_volume, validate)
    coefficient = coef_result.coefficient
    
    # 计算标准件
    std_units = actual_qty * coefficient
    
    if data_source == "standard":
        formula = f"{actual_qty} × max({std_weight:.4f}/15, {std_volume:.6f}/0.05) = {actual_qty} × {coefficient:.4f} = {std_units:.4f}"
    else:
        formula = f"{actual_qty} × max({single_unit_weight:.4f}/15, {single_unit_volume:.6f}/0.05) = {actual_qty} × {coefficient:.4f} = {std_units:.4f}"
    
    return StdUnitsResult(
        actual_qty=actual_qty,
        total_weight_kg=round(total_weight, 4),
        total_volume_m3=round(total_volume, 6),
        single_unit_weight_kg=round(single_unit_weight, 4),
        single_unit_volume_m3=round(single_unit_volume, 6),
        coefficient=coefficient,
        data_source=data_source,
        std_units=round(std_units, 4),
        formula=formula
    )


def calc_all(
    actual_qty: Union[int, float],
    total_weight_kg: Union[int, float],
    total_volume_m3: Union[int, float],
    product_name: Optional[str] = None,
    validate: bool = True
) -> Tuple[CoefficientResult, StdUnitsResult]:
    """
    一次性计算计费系数和标准件
    
    Args:
        actual_qty: 实际出库数量
        total_weight_kg: 总重量（公斤）
        total_volume_m3: 总体积（立方米）
        product_name: 商品名称（用于查找资料库）
        validate: 是否进行参数校验
        
    Returns:
        Tuple[CoefficientResult, StdUnitsResult]: (计费系数结果, 标准件结果)
    """
    std_result = calc_std_units(
        actual_qty=actual_qty,
        total_weight_kg=total_weight_kg,
        total_volume_m3=total_volume_m3,
        product_name=product_name,
        validate=validate
    )
    
    coef_result = CoefficientResult(
        single_unit_weight_kg=std_result.single_unit_weight_kg,
        single_unit_volume_m3=std_result.single_unit_volume_m3,
        weight_ratio=std_result.single_unit_weight_kg / 15,
        volume_ratio=std_result.single_unit_volume_m3 / 0.05,
        coefficient=std_result.coefficient,
        data_source=std_result.data_source,
        threshold_info=[]
    )
    
    return coef_result, std_result


# ============================================================
# 兼容旧接口的函数
# ============================================================

def calculate_coefficient(weight_kg: float, volume_m3: float) -> float:
    """
    计算计费系数（兼容旧接口）
    
    注意：这是按单件计算的，weight_kg和volume_m3应该是单件的值
    """
    result = calc_coefficient(weight_kg, volume_m3)
    return result.coefficient


def calculate_std_units(
    actual_qty: float,
    total_weight_kg: float,
    total_volume_m3: float
) -> float:
    """
    计算标准件数（兼容旧接口）
    """
    result = calc_std_units(actual_qty, total_weight_kg, total_volume_m3)
    return result.std_units


if __name__ == "__main__":
    print("=" * 60)
    print("标准件计算示例（按合同规则 + 商品资料库）")
    print("=" * 60)
    
    # 示例1：300ml糖水瓶（有商品资料库标准数据）
    print("\n1. 300ml糖水瓶（96个/箱）20箱（使用标准数据）")
    print("   假设标准重量: 18kg/箱，标准体积: 0.06m³/箱")
    coef1, std1 = calc_all(
        actual_qty=20, 
        total_weight_kg=402, 
        total_volume_m3=1.326,
        product_name="300ml糖水瓶（96个/箱）"
    )
    print(f"   数据来源: {std1.data_source}")
    print(f"   系数: {coef1.coefficient:.4f}")
    print(f"   标准件: {std1.std_units:.4f}")
    print(f"   公式: {std1.formula}")
    
    # 示例2：无纺布铝箔袋（无标准数据，使用实际数据）
    print("\n2. 无纺布铝箔袋（500个/包）4包（使用实际数据）")
    print("   总重量: 67.2kg，体积: 0.4416m³")
    coef2, std2 = calc_all(
        actual_qty=4, 
        total_weight_kg=67.2, 
        total_volume_m3=0.4416,
        product_name="无纺布铝箔袋（500个/包）"
    )
    print(f"   数据来源: {std2.data_source}")
    print(f"   单件重量: {coef2.single_unit_weight_kg:.2f} kg")
    print(f"   单件体积: {coef2.single_unit_volume_m3:.4f} m³")
    print(f"   系数: {coef2.coefficient:.4f}")
    print(f"   标准件: {std2.std_units:.4f}")
    print(f"   官方数据: 8.84")
    
    # 示例3：直接使用实际数据（不指定商品名称）
    print("\n3. 300ml糖水瓶20箱（不使用资料库，直接用实际数据）")
    coef3, std3 = calc_all(
        actual_qty=20, 
        total_weight_kg=402, 
        total_volume_m3=1.326
    )
    print(f"   数据来源: {std3.data_source}")
    print(f"   系数: {coef3.coefficient:.4f}")
    print(f"   标准件: {std3.std_units:.4f}")