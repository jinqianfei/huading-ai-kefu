# -*- coding: utf-8 -*-
"""
Dispatcher - 意图路由调度器

集成 intent_router.py，根据用户输入自动路由到对应 Skill
支持方式2：显式路由验证逻辑
"""

import sys
import os
from typing import Dict, Any, Optional

# 添加 skills 目录到路径
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from intent_router import resolve_intent, INTENT_KEYWORDS, match_intent


class Dispatcher:
    """意图路由调度器"""
    
    VERSION = "1.0"
    
    def __init__(self, db_config: Dict = None):
        """
        初始化调度器
        
        Args:
            db_config: 数据库配置
        """
        self.db_config = db_config or {
            "host": "localhost",
            "port": 5432,
            "database": "ai_cs_support",
            "user": "jinqianfei"
        }
        
        # Skill 路径映射
        self.skill_paths = {
            "skill-order-to-huading-template": "skills/skill_order_to_huading_template",
            "billing-workflow": "skills/billing-workflow",
        }
    
    def dispatch(self, user_input: str, context: Dict = None) -> Dict[str, Any]:
        """
        调度用户请求到对应 Skill
        
        Args:
            user_input: 用户输入
            context: 上下文数据（已上传的文件等）
        
        Returns:
            {
                "success": bool,
                "intent": str,
                "skill": str,
                "params": dict,
                "missing_params": list,
                "result": Any,  # Skill执行结果
                "message": str
            }
        """
        context = context or {}
        
        # 1. 解析意图
        route_info = resolve_intent(user_input, context)
        
        print(f"[Dispatcher] 意图解析: {route_info}")
        
        # 2. 检查意图是否识别
        if not route_info["intent"]:
            return {
                "success": False,
                "intent": None,
                "skill": None,
                "params": {},
                "missing_params": [],
                "result": None,
                "message": route_info["suggestion"]
            }
        
        # 3. 检查必填参数
        missing_params = route_info["missing_params"]
        if missing_params:
            return {
                "success": False,
                "intent": route_info["intent"],
                "skill": route_info["skill"],
                "params": route_info["params"],
                "missing_params": missing_params,
                "result": None,
                "message": f"缺少必填参数：{'、'.join(missing_params)}"
            }
        
        # 4. 执行 Skill
        skill_name = route_info["skill"]
        params = route_info["params"]
        
        try:
            result = self._execute_skill(skill_name, params)
            return {
                "success": True,
                "intent": route_info["intent"],
                "skill": skill_name,
                "params": params,
                "missing_params": [],
                "result": result,
                "message": "执行成功"
            }
        except Exception as e:
            return {
                "success": False,
                "intent": route_info["intent"],
                "skill": skill_name,
                "params": params,
                "missing_params": [],
                "result": None,
                "message": f"执行失败: {str(e)}"
            }
    
    def _execute_skill(self, skill_name: str, params: Dict) -> Any:
        """
        执行指定 Skill
        
        Args:
            skill_name: Skill名称
            params: 执行参数
        
        Returns:
            Skill执行结果
        """
        # skill-order-to-huading-template
        if skill_name == "skill-order-to-huading-template":
            return self._execute_order_template(params)
        
        # billing-workflow 的各个步骤
        elif skill_name.startswith("billing-workflow"):
            return self._execute_billing_workflow(skill_name, params)
        
        else:
            raise ValueError(f"未知的Skill: {skill_name}")
    
    def _execute_order_template(self, params: Dict) -> Dict:
        """执行订单转华鼎模板"""
        from skills.skill_order_to_huading_template import OrderToHuadingTemplate
        
        # 获取 shipper_id
        shipper_name = params.get("shipper_name")
        shipper_id = self._get_shipper_id(shipper_name)
        
        if not shipper_id:
            raise ValueError(f"未找到货主: {shipper_name}")
        
        # 获取订单文件
        order_file = params.get("order_file")
        if not order_file or not os.path.exists(order_file):
            raise ValueError(f"订单文件不存在: {order_file}")
        
        # 执行转换
        skill = OrderToHuadingTemplate(
            shipper_id=shipper_id,
            db_config=self.db_config
        )
        
        return skill.execute(order_file)
    
    def _execute_billing_workflow(self, skill_name: str, params: Dict) -> Dict:
        """执行账单工作流"""
        # 待实现
        return {"message": f"billing-workflow {skill_name} 待实现"}
    
    def _get_shipper_id(self, shipper_name: str) -> Optional[str]:
        """根据货主名称获取 shipper_id"""
        if not shipper_name:
            return None
        
        import psycopg2
        
        try:
            conn = psycopg2.connect(**self.db_config)
            cur = conn.cursor()
            
            cur.execute(
                "SELECT shipper_id FROM shipper WHERE shipper_name LIKE %s",
                (f"%{shipper_name}%",)
            )
            result = cur.fetchone()
            
            conn.close()
            
            return result[0] if result else None
            
        except Exception as e:
            print(f"查询货主失败: {e}")
            return None
    
    def get_available_intents(self) -> Dict[str, Dict]:
        """获取所有可用的意图及其配置"""
        return INTENT_KEYWORDS


def test_router():
    """测试路由逻辑"""
    dispatcher = Dispatcher()
    
    test_cases = [
        {
            "input": "帮我处理这个订单，生成华鼎模板",
            "context": {"order_file": "/path/to/order.xlsx"}
        },
        {
            "input": "生成长沙广承2026年2月账单",
            "context": {}
        },
        {
            "input": "上传合同，解析规则",
            "context": {"contract_file": "/path/to/contract.pdf"}
        },
        {
            "input": "处理订单",
            "context": {}
        },
    ]
    
    print("=" * 60)
    print("意图路由测试")
    print("=" * 60)
    
    for i, case in enumerate(test_cases):
        print(f"\n【测试 {i+1}】")
        print(f"输入: {case['input']}")
        print(f"上下文: {case['context']}")
        
        result = dispatcher.dispatch(case['input'], case['context'])
        
        print(f"结果:")
        print(f"  - success: {result['success']}")
        print(f"  - intent: {result['intent']}")
        print(f"  - skill: {result['skill']}")
        print(f"  - missing_params: {result['missing_params']}")
        print(f"  - message: {result['message']}")


if __name__ == "__main__":
    test_router()