
## 待开发需求 (2026-05-14 更新)

### skill_order_to_huading_template 多格式输入支持
**提出人**: 客服用户
**需求**: 订单入口应支持识别图片、文字、Excel等文件，识别后转为标准Excel再走现有skill流程
**优先级**: 高
**状态**: ✅ 已完成开发（集成到skill内部，agent调用image工具识别）
**完成时间**: 2026-05-14
**实现方式**:
- skill返回 `need_ocr=True` 时，由agent调用image工具识别
- 识别结果通过 `ocr_result` 参数回传给skill继续处理
- 复用现有门店匹配、货主确认、SKU映射流程
