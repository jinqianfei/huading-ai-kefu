/**
 * api_test_helpers.go - API测试辅助方法
 * AI客服赋能系统 - B6任务集成测试
 *
 * 为集成测试提供必要的辅助方法
 */

package api

import (
	"database/sql"
)

// ==================== OrderMappingAPI 测试辅助方法 ====================

// GetSkuMappingForTest 测试用：获取SKU映射
func (api *OrderMappingAPI) GetSkuMappingForTest(customerID, customerSkuCode string) (*SkuMappingResponse, error) {
	return api.getSkuMapping(customerID, customerSkuCode)
}

// GetWarehouseMappingsForTest 测试用：获取仓库映射列表
func (api *OrderMappingAPI) GetWarehouseMappingsForTest(customerID string) ([]WarehouseMappingResponse, error) {
	return api.getWarehouseMappings(customerID)
}

// GetDeliveryMethodMappingsForTest 测试用：获取配送方式映射
func (api *OrderMappingAPI) GetDeliveryMethodMappingsForTest(customerID, sourceValue string) ([]DeliveryMethodMappingResponse, error) {
	return api.getDeliveryMethodMappings(customerID, sourceValue)
}

// GetBusinessTypeMappingsForTest 测试用：获取业务类型映射
func (api *OrderMappingAPI) GetBusinessTypeMappingsForTest(customerID, sourceValue string) ([]BusinessTypeMappingResponse, error) {
	return api.getBusinessTypeMappings(customerID, sourceValue)
}

// ==================== PendingTaskAPI 测试辅助方法 ====================

// QueryPendingTasksForTest 测试用：查询待确认任务
func (api *PendingTaskAPI) QueryPendingTasksForTest(req PendingTaskListRequest) ([]PendingTask, int64, error) {
	return api.queryPendingTasks(req)
}

// ConfirmTaskForTest 测试用：确认任务
func (api *PendingTaskAPI) ConfirmTaskForTest(taskKey string, req ConfirmTaskRequest) (*PendingTask, error) {
	return api.confirmTask(taskKey, req)
}

// GetPendingTasksForRowForTest 测试用：获取指定行的待确认任务
func (api *PendingTaskAPI) GetPendingTasksForRowForTest(customerID, orderID string, rowNumber int) ([]PendingTask, error) {
	return api.getPendingTasksForRow(customerID, orderID, rowNumber)
}

// ==================== 数据库测试辅助函数 ====================

// TestDBHelper 数据库测试辅助结构
type TestDBHelper struct {
	db *sql.DB
}

// NewTestDBHelper 创建测试辅助实例
func NewTestDBHelper(db *sql.DB) *TestDBHelper {
	return &TestDBHelper{db: db}
}

// CleanupTestData 清理测试数据
func (h *TestDBHelper) CleanupTestData(customerIDPattern string) error {
	tables := []string{
		"pending_mapping_task",
		"customer_sku_mapping",
		"customer_alias_mapping",
		"huading_store_mapping",
		"customer_warehouse_mapping",
		"order_task",
		"match_history",
		"delivery_method_mapping",
		"business_type_mapping",
	}

	for _, table := range tables {
		_, err := h.db.Exec("DELETE FROM "+table+" WHERE customer_id LIKE $1", customerIDPattern+"%")
		if err != nil {
			return err
		}
	}
	return nil
}

// CountTableRows 统计表行数
func (h *TestDBHelper) CountTableRows(table, whereClause string, args ...interface{}) (int64, error) {
	query := "SELECT COUNT(*) FROM " + table
	if whereClause != "" {
		query += " WHERE " + whereClause
	}

	var count int64
	err := h.db.QueryRow(query, args...).Scan(&count)
	return count, err
}

// InsertTestSkuMapping 插入测试SKU映射
func (h *TestDBHelper) InsertTestSkuMapping(customerID, customerSkuCode, customerSkuName, systemSkuCode, systemSkuName string, confidence float64) error {
	_, err := h.db.Exec(`
		INSERT INTO customer_sku_mapping (
			customer_id, customer_sku_code, customer_sku_name,
			system_sku_code, system_sku_name, confidence, source, status
		) VALUES ($1, $2, $3, $4, $5, $6, 'AUTO', 'ACTIVE')
		ON CONFLICT (customer_id, customer_sku_code) DO UPDATE SET
			system_sku_code = EXCLUDED.system_sku_code,
			confidence = EXCLUDED.confidence
	`, customerID, customerSkuCode, customerSkuName, systemSkuCode, systemSkuName, confidence)
	return err
}

// InsertTestStoreMapping 插入测试门店映射
func (h *TestDBHelper) InsertTestStoreMapping(customerID, shortName, fullName, matchType string, confidence float64) error {
	_, err := h.db.Exec(`
		INSERT INTO customer_alias_mapping (
			customer_id, short_name, full_name, match_type, confidence
		) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (short_name) DO UPDATE SET
			full_name = EXCLUDED.full_name,
			confidence = EXCLUDED.confidence
	`, customerID, shortName, fullName, matchType, confidence)
	return err
}

// InsertTestWarehouseMapping 插入测试仓库映射
func (h *TestDBHelper) InsertTestWarehouseMapping(customerID, warehouseCode, warehouseName string, priority int) error {
	_, err := h.db.Exec(`
		INSERT INTO customer_warehouse_mapping (
			customer_id, warehouse_code, warehouse_name, priority, enabled
		) VALUES ($1, $2, $3, $4, true)
		ON CONFLICT (customer_id, warehouse_code) DO UPDATE SET
			warehouse_name = EXCLUDED.warehouse_name,
			priority = EXCLUDED.priority
	`, customerID, warehouseCode, warehouseName, priority)
	return err
}

// InsertTestPendingTask 插入测试待确认任务
func (h *TestDBHelper) InsertTestPendingTask(task *PendingTask) error {
	return CreatePendingTask(h.db, task)
}

// GetPendingTaskStatus 获取待确认任务状态
func (h *TestDBHelper) GetPendingTaskStatus(taskKey string) (string, error) {
	var status string
	err := h.db.QueryRow(`
		SELECT status FROM pending_mapping_task WHERE task_key = $1
	`, taskKey).Scan(&status)
	return status, err
}

// CountPendingTasks 统计待确认任务数量
func (h *TestDBHelper) CountPendingTasks(customerID, status string) (int64, error) {
	var count int64
	err := h.db.QueryRow(`
		SELECT COUNT(*) FROM pending_mapping_task
		WHERE customer_id = $1 AND status = $2
	`, customerID, status).Scan(&count)
	return count, err
}
