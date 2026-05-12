# billing_tools/config.py
# 账单助手工具配置

import os

# 知识库路径
KNOWLEDGE_BASE = os.path.expanduser("~/openclaw-workspaces/ai-kefu/knowledge")

RULES_PATH = os.path.join(KNOWLEDGE_BASE, "rules")
BILLS_PATH = os.path.join(KNOWLEDGE_BASE, "bills")
ADJUSTMENTS_PATH = os.path.join(KNOWLEDGE_BASE, "adjustments")

# 告警阈值
THRESHOLD = 0.01

# 确保目录存在
os.makedirs(RULES_PATH, exist_ok=True)
os.makedirs(BILLS_PATH, exist_ok=True)
os.makedirs(ADJUSTMENTS_PATH, exist_ok=True)
