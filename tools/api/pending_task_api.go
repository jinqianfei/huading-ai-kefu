/**
 * pending_task_api.go - 待确认任务API
 * AI客服赋能系统 - B5任务API层
 *
 * 实现3个核心端点：
 * 1. GET /api/pending-tasks - 查询待确认任务列表
 * 2. POST /api/pending-tasks/{task_key}/confirm - 确认待确认任务
 * 3. POST /api/orders/process-row - 处理单行订单
 */

package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"ai-cs-support/tools/db"
)

// ==================== 数据结构定义 ====================

// PendingTask 待确认任务结构
type PendingTask struct {
	ID                int64           `json:"id"`
	TaskKey           string          `json:"task_key"`
	CustomerID        string          `json:"customer_id"`
	OrderID           string          `json:"order_id"`
	RowNumber         int             `json:"row_number"`
	MappingType       string          `json:"mapping_type"`
	SourceValue       string          `json:"source_value"`
	SourceField       string          `json:"source_field"`
	ExpectedFieldType string          `json:"expected_field_type"`
	Candidates        json.RawMessage `json:"candidates"`
	ManualValue       string          `json:"manual_value"`
	Status            string          `json:"status"`
	ConfirmedBy       string          `json:"confirmed_by"`
	ConfirmedAt       *time.Time      `json:"confirmed_at"`
	RejectionReason   string          `json:"rejection_reason"`
	Remark            string          `json:"remark"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

// Candidate 候选项结构
type Candidate struct {
	Value      string  `json:"value"`
	Label      string  `json:"label,omitempty"`
	Source     string  `json:"source"`
	Confidence float64 `json:"confidence"`
}

// PendingTaskListRequest 待确认任务列表查询请求
type PendingTaskListRequest struct {
	CustomerID   string `json:"customer_id"`
	Status       string `json:"status"`
	MappingType  string `json:"mapping_type"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
	OrderID      string `json:"order_id"`
	SourceValue  string `json:"source_value"`
}

// PendingTaskListResponse 待确认任务列表响应
type PendingTaskListResponse struct {
	Total       int64         `json:"total"`
	Page        int           `json:"page"`
	PageSize    int           `json:"page_size"`
	TotalPages  int           `json:"total_pages"`
	Tasks       []PendingTask `json:"tasks"`
}

// ConfirmTaskRequest 确认任务请求
type ConfirmTaskRequest struct {
	ManualValue string `json:"manual_value"`
	ConfirmedBy string `json:"confirmed_by"`
	Remark      string `json:"remark"`
}

// ConfirmTaskResponse 确认任务响应
type ConfirmTaskResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Task      PendingTask `json:"task"`
}

// ProcessRowRequest 处理单行订单请求
type ProcessRowRequest struct {
	CustomerID        string            `json:"customer_id"`
	OrderID           string            `json:"order_id"`
	RowNumber         int               `json:"row_number"`
	ConfirmedMappings map[string]string `json:"confirmed_mappings"` // field_name -> manual_value
}

// ProcessRowResponse 处理单行订单响应
type ProcessRowResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	OrderID       string `json:"order_id"`
	ExcelFilePath string `json:"excel_file_path"`
	GeneratedRows int    `json:"generated_rows"`
}

// APIResponse 通用API响应结构
type APIResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ==================== API处理器 ====================

// PendingTaskAPI 待确认任务API处理器
type PendingTaskAPI struct {
	db *sql.DB
}

// NewPendingTaskAPI 创建待确认任务API处理器
func NewPendingTaskAPI(database *sql.DB) *PendingTaskAPI {
	return &PendingTaskAPI{db: database}
}

// GetPendingTasksHandler GET /api/pending-tasks
// 查询待确认任务列表
func (api *PendingTaskAPI) GetPendingTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// 解析查询参数
	req := PendingTaskListRequest{
		Page:     1,
		PageSize: 20,
		Status:   "PENDING",
	}

	if page := r.URL.Query().Get("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			req.Page = p
		}
	}

	if pageSize := r.URL.Query().Get("page_size"); pageSize != "" {
		if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
			if ps > 100 {
				ps = 100 // 最大100条
			}
			req.PageSize = ps
		}
	}

	req.CustomerID = r.URL.Query().Get("customer_id")
	req.Status = r.URL.Query().Get("status")
	if req.Status == "" {
		req.Status = "PENDING"
	}
	req.MappingType = r.URL.Query().Get("mapping_type")
	req.OrderID = r.URL.Query().Get("order_id")
	req.SourceValue = r.URL.Query().Get("source_value")

	// 查询任务列表
	tasks, total, err := api.queryPendingTasks(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Query failed: %v", err))
		return
	}

	totalPages := int(total) / req.PageSize
	if int(total)%req.PageSize > 0 {
		totalPages++
	}

	response := PendingTaskListResponse{
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
		Tasks:      tasks,
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// queryPendingTasks 查询待确认任务列表（支持分页和筛选）
func (api *PendingTaskAPI) queryPendingTasks(req PendingTaskListRequest) ([]PendingTask, int64, error) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	if req.CustomerID != "" {
		conditions = append(conditions, fmt.Sprintf("customer_id = $%d", argIndex))
		args = append(args, req.CustomerID)
		argIndex++
	}

	if req.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, req.Status)
		argIndex++
	}

	if req.MappingType != "" {
		conditions = append(conditions, fmt.Sprintf("mapping_type = $%d", argIndex))
		args = append(args, req.MappingType)
		argIndex++
	}

	if req.OrderID != "" {
		conditions = append(conditions, fmt.Sprintf("order_id = $%d", argIndex))
		args = append(args, req.OrderID)
		argIndex++
	}

	if req.SourceValue != "" {
		conditions = append(conditions, fmt.Sprintf("source_value ILIKE $%d", argIndex))
		args = append(args, "%"+req.SourceValue+"%")
		argIndex++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// 查询总数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM pending_mapping_task %s", whereClause)
	var total int64
	err := api.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count tasks: %w", err)
	}

	// 查询数据
	query := fmt.Sprintf(`
		SELECT id, task_key, customer_id, order_id, row_number, mapping_type,
		       source_value, source_field, expected_field_type, candidates,
		       manual_value, status, confirmed_by, confirmed_at,
		       rejection_reason, remark, created_at, updated_at
		FROM pending_mapping_task
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)

	offset := (req.Page - 1) * req.PageSize
	args = append(args, req.PageSize, offset)

	rows, err := api.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []PendingTask
	for rows.Next() {
		var task PendingTask
		var confirmedAt sql.NullTime
		err := rows.Scan(
			&task.ID, &task.TaskKey, &task.CustomerID, &task.OrderID, &task.RowNumber,
			&task.MappingType, &task.SourceValue, &task.SourceField, &task.ExpectedFieldType,
			&task.Candidates, &task.ManualValue, &task.Status, &task.ConfirmedBy,
			&confirmedAt, &task.RejectionReason, &task.Remark, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan task: %w", err)
		}
		if confirmedAt.Valid {
			task.ConfirmedAt = &confirmedAt.Time
		}
		tasks = append(tasks, task)
	}

	return tasks, total, rows.Err()
}

// ConfirmTaskHandler POST /api/pending-tasks/{task_key}/confirm
// 确认待确认任务
func (api *PendingTaskAPI) ConfirmTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// 从URL路径中提取task_key
	path := strings.TrimPrefix(r.URL.Path, "/api/pending-tasks/")
	path = strings.TrimSuffix(path, "/confirm")
	taskKey := strings.TrimSpace(path)

	if taskKey == "" {
		respondWithError(w, http.StatusBadRequest, "Task key is required")
		return
	}

	// 解析请求体
	var req ConfirmTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	if req.ManualValue == "" {
		respondWithError(w, http.StatusBadRequest, "manual_value is required")
		return
	}

	if req.ConfirmedBy == "" {
		respondWithError(w, http.StatusBadRequest, "confirmed_by is required")
		return
	}

	// 更新任务状态
	task, err := api.confirmTask(taskKey, req)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Task not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to confirm task: %v", err))
		return
	}

	// 根据mapping_type执行相应的映射保存逻辑
	if err := api.saveMappingAfterConfirm(task); err != nil {
		// 记录错误但不影响确认结果
		fmt.Printf("Warning: failed to save mapping for task %s: %v\n", taskKey, err)
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Task confirmed successfully",
		Data: ConfirmTaskResponse{
			Success: true,
			Message: "Task confirmed",
			Task:    *task,
		},
	})
}

// confirmTask 确认任务的核心逻辑
func (api *PendingTaskAPI) confirmTask(taskKey string, req ConfirmTaskRequest) (*PendingTask, error) {
	now := time.Now()

	// 更新任务状态
	result, err := api.db.Exec(`
		UPDATE pending_mapping_task
		SET manual_value = $1,
		    status = 'CONFIRMED',
		    confirmed_by = $2,
		    confirmed_at = $3,
		    remark = $4,
		    updated_at = $5
		WHERE task_key = $6
		RETURNING id
	`, req.ManualValue, req.ConfirmedBy, now, req.Remark, now, taskKey)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	// 查询更新后的任务
	var task PendingTask
	var confirmedAt sql.NullTime
	err = api.db.QueryRow(`
		SELECT id, task_key, customer_id, order_id, row_number, mapping_type,
		       source_value, source_field, expected_field_type, candidates,
		       manual_value, status, confirmed_by, confirmed_at,
		       rejection_reason, remark, created_at, updated_at
		FROM pending_mapping_task
		WHERE task_key = $1
	`, taskKey).Scan(
		&task.ID, &task.TaskKey, &task.CustomerID, &task.OrderID, &task.RowNumber,
		&task.MappingType, &task.SourceValue, &task.SourceField, &task.ExpectedFieldType,
		&task.Candidates, &task.ManualValue, &task.Status, &task.ConfirmedBy,
		&confirmedAt, &task.RejectionReason, &task.Remark, &task.CreatedAt, &task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if confirmedAt.Valid {
		task.ConfirmedAt = &confirmedAt.Time
	}

	return &task, nil
}

// saveMappingAfterConfirm 确认后保存映射关系
func (api *PendingTaskAPI) saveMappingAfterConfirm(task *PendingTask) error {
	switch task.MappingType {
	case "SKU":
		return api.saveSkuMapping(task)
	case "STORE":
		return api.saveStoreMapping(task)
	case "WAREHOUSE":
		return api.saveWarehouseMapping(task)
	case "DELIVERY_METHOD":
		return api.saveDeliveryMethodMapping(task)
	case "BUSINESS_TYPE":
		return api.saveBusinessTypeMapping(task)
	default:
		return nil
	}
}

// saveSkuMapping 保存SKU映射
func (api *PendingTaskAPI) saveSkuMapping(task *PendingTask) error {
	// 解析candidates获取system_sku_code
	var candidates []Candidate
	if err := json.Unmarshal(task.Candidates, &candidates); err != nil {
		return err
	}

	var systemSkuCode, systemSkuName string
	for _, c := range candidates {
		if c.Value == task.ManualValue {
			systemSkuCode = c.Value
			systemSkuName = c.Label
			break
		}
	}

	if systemSkuCode == "" {
		return fmt.Errorf("manual value not found in candidates")
	}

	// 保存到customer_sku_mapping表
	_, err := api.db.Exec(`
		INSERT INTO customer_sku_mapping (
			customer_id, customer_sku_code, customer_sku_name,
			system_sku_code, system_sku_name, confidence, source, confirmed_by, confirmed_at
		) VALUES ($1, $2, $3, $4, $5, $6, 'MANUAL', $7, $8)
		ON CONFLICT (customer_id, customer_sku_code) DO UPDATE SET
			system_sku_code = EXCLUDED.system_sku_code,
			system_sku_name = EXCLUDED.system_sku_name,
			confidence = EXCLUDED.confidence,
			confirmed_by = EXCLUDED.confirmed_by,
			confirmed_at = EXCLUDED.confirmed_at,
			updated_at = NOW()
	`, task.CustomerID, task.SourceValue, task.SourceValue,
		systemSkuCode, systemSkuName, 1.0, task.ConfirmedBy, task.ConfirmedAt)

	return err
}

// saveStoreMapping 保存门店映射
func (api *PendingTaskAPI) saveStoreMapping(task *PendingTask) error {
	// 保存到customer_alias_mapping表
	_, err := api.db.Exec(`
		INSERT INTO customer_alias_mapping (
			customer_id, short_name, full_name, match_type, confidence, confirmed_by, confirmed_at
		) VALUES ($1, $2, $3, 'EXACT', 1.0, $4, $5)
		ON CONFLICT (short_name) DO UPDATE SET
			full_name = EXCLUDED.full_name,
			match_type = EXCLUDED.match_type,
			confidence = EXCLUDED.confidence,
			confirmed_by = EXCLUDED.confirmed_by,
			confirmed_at = EXCLUDED.confirmed_at
	`, task.CustomerID, task.SourceValue, task.ManualValue, task.ConfirmedBy, task.ConfirmedAt)

	return err
}

// saveWarehouseMapping 保存仓库映射
func (api *PendingTaskAPI) saveWarehouseMapping(task *PendingTask) error {
	// 保存到customer_warehouse_mapping表
	_, err := api.db.Exec(`
		INSERT INTO customer_warehouse_mapping (
			customer_id, warehouse_code, warehouse_name, enabled
		) VALUES ($1, $2, $3, true)
		ON CONFLICT (customer_id, warehouse_code) DO UPDATE SET
			warehouse_name = EXCLUDED.warehouse_name,
			enabled = true
	`, task.CustomerID, task.ManualValue, task.ManualValue)

	return err
}

// saveDeliveryMethodMapping 保存配送方式映射
func (api *PendingTaskAPI) saveDeliveryMethodMapping(task *PendingTask) error {
	_, err := api.db.Exec(`
		INSERT INTO delivery_method_mapping (
			customer_id, source_value, target_value, match_type, confidence
		) VALUES ($1, $2, $3, 'EXACT', 1.0)
		ON CONFLICT (customer_id, source_value) DO UPDATE SET
			target_value = EXCLUDED.target_value,
			match_type = EXCLUDED.match_type,
			confidence = EXCLUDED.confidence
	`, task.CustomerID, task.SourceValue, task.ManualValue)

	return err
}

// saveBusinessTypeMapping 保存业务类型映射
func (api *PendingTaskAPI) saveBusinessTypeMapping(task *PendingTask) error {
	_, err := api.db.Exec(`
		INSERT INTO business_type_mapping (
			customer_id, source_value, target_value
		) VALUES ($1, $2, $3)
		ON CONFLICT (customer_id, source_value) DO UPDATE SET
			target_value = EXCLUDED.target_value
	`, task.CustomerID, task.SourceValue, task.ManualValue)

	return err
}

// ProcessRowHandler POST /api/orders/process-row
// 处理单行订单（确认后重新生成）
func (api *PendingTaskAPI) ProcessRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// 解析请求体
	var req ProcessRowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	// 参数校验
	if req.CustomerID == "" {
		respondWithError(w, http.StatusBadRequest, "customer_id is required")
		return
	}
	if req.OrderID == "" {
		respondWithError(w, http.StatusBadRequest, "order_id is required")
		return
	}
	if req.RowNumber <= 0 {
		respondWithError(w, http.StatusBadRequest, "row_number must be greater than 0")
		return
	}

	// 处理订单行
	result, err := api.processOrderRow(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to process row: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Row processed successfully",
		Data:    result,
	})
}

// processOrderRow 处理订单行的核心逻辑
func (api *PendingTaskAPI) processOrderRow(req ProcessRowRequest) (*ProcessRowResponse, error) {
	// 1. 获取该行的所有待确认任务
	tasks, err := api.getPendingTasksForRow(req.CustomerID, req.OrderID, req.RowNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending tasks: %w", err)
	}

	// 2. 应用已确认的映射
	for fieldName, manualValue := range req.ConfirmedMappings {
		if err := api.updateTaskWithManualValue(req.CustomerID, req.OrderID, req.RowNumber, fieldName, manualValue); err != nil {
			return nil, fmt.Errorf("failed to update task for field %s: %w", fieldName, err)
		}
	}

	// 3. 检查是否还有未确认的任务
	pendingTasks, err := api.getPendingTasksForRow(req.CustomerID, req.OrderID, req.RowNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to check pending tasks: %w", err)
	}

	if len(pendingTasks) > 0 {
		// 还有未确认的任务，返回提示
		return &ProcessRowResponse{
			Success: false,
			Message: fmt.Sprintf("还有 %d 个字段待确认", len(pendingTasks)),
			OrderID: req.OrderID,
		}, nil
	}

	// 4. 生成Excel文件
	excelPath, generatedRows, err := api.generateExcelForRow(req.CustomerID, req.OrderID, req.RowNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to generate excel: %w", err)
	}

	return &ProcessRowResponse{
		Success:       true,
		Message:       "Row processed and Excel generated",
		OrderID:       req.OrderID,
		ExcelFilePath: excelPath,
		GeneratedRows: generatedRows,
	}, nil
}

// getPendingTasksForRow 获取指定行的待确认任务
func (api *PendingTaskAPI) getPendingTasksForRow(customerID, orderID string, rowNumber int) ([]PendingTask, error) {
	rows, err := api.db.Query(`
		SELECT id, task_key, customer_id, order_id, row_number, mapping_type,
		       source_value, source_field, expected_field_type, candidates,
		       manual_value, status, confirmed_by, confirmed_at,
		       rejection_reason, remark, created_at, updated_at
		FROM pending_mapping_task
		WHERE customer_id = $1 AND order_id = $2 AND row_number = $3 AND status = 'PENDING'
	`, customerID, orderID, rowNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []PendingTask
	for rows.Next() {
		var task PendingTask
		var confirmedAt sql.NullTime
		err := rows.Scan(
			&task.ID, &task.TaskKey, &task.CustomerID, &task.OrderID, &task.RowNumber,
			&task.MappingType, &task.SourceValue, &task.SourceField, &task.ExpectedFieldType,
			&task.Candidates, &task.ManualValue, &task.Status, &task.ConfirmedBy,
			&confirmedAt, &task.RejectionReason, &task.Remark, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if confirmedAt.Valid {
			task.ConfirmedAt = &confirmedAt.Time
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

// updateTaskWithManualValue 更新任务的manual_value
func (api *PendingTaskAPI) updateTaskWithManualValue(customerID, orderID string, rowNumber int, fieldName, manualValue string) error {
	taskKey := fmt.Sprintf("%s#%s#%d#%s", customerID, orderID, rowNumber, fieldName)

	_, err := api.db.Exec(`
		UPDATE pending_mapping_task
		SET manual_value = $1,
		    status = 'CONFIRMED',
		    updated_at = NOW()
		WHERE task_key = $2
	`, manualValue, taskKey)

	return err
}

// generateExcelForRow 为指定行生成Excel文件
func (api *PendingTaskAPI) generateExcelForRow(customerID, orderID string, rowNumber int) (string, int, error) {
	// 获取该行的所有已确认任务
	rows, err := api.db.Query(`
		SELECT mapping_type, source_field, manual_value, source_value
		FROM pending_mapping_task
		WHERE customer_id = $1 AND order_id = $2 AND row_number = $3 AND status = 'CONFIRMED'
	`, customerID, orderID, rowNumber)
	if err != nil {
		return "", 0, err
	}
	defer rows.Close()

	// 构建订单数据
	orderData := make(map[string]string)
	for rows.Next() {
		var mappingType, sourceField, manualValue, sourceValue string
		if err := rows.Scan(&mappingType, &sourceField, &manualValue, &sourceValue); err != nil {
			continue
		}

		// 使用manual_value，如果没有则使用source_value
		value := manualValue
		if value == "" {
			value = sourceValue
		}
		orderData[sourceField] = value
	}

	// 生成Excel文件路径
	timestamp := time.Now().Format("20060102_150405")
	excelPath := fmt.Sprintf("/tmp/huading_template_%s_%s_row%d_%s.xlsx",
		customerID, orderID, rowNumber, timestamp)

	// TODO: 调用excel工具生成华鼎标准出库单
	// 这里返回模拟结果，实际实现需要调用excel包中的生成函数
	// generatedRows, err := excel.GenerateHuadingTemplate(excelPath, orderData)

	return excelPath, 1, nil
}

// ==================== 辅助函数 ====================

// respondWithJSON 返回JSON响应
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// respondWithError 返回错误响应
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, APIResponse{
		Success: false,
		Code:    statusCode,
		Message: message,
	})
}

// ==================== 便捷函数 ====================

// CreatePendingTask 创建待确认任务
func CreatePendingTask(database *sql.DB, task *PendingTask) error {
	// 生成task_key
	task.TaskKey = fmt.Sprintf("%s#%s#%d#%s", task.CustomerID, task.OrderID, task.RowNumber, task.MappingType)

	candidatesJSON, err := json.Marshal(task.Candidates)
	if err != nil {
		return fmt.Errorf("failed to marshal candidates: %w", err)
	}

	_, err = database.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates,
			status, remark
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 'PENDING', $10)
		ON CONFLICT (task_key) DO UPDATE SET
			candidates = EXCLUDED.candidates,
			source_value = EXCLUDED.source_value,
			updated_at = NOW()
	`, task.TaskKey, task.CustomerID, task.OrderID, task.RowNumber, task.MappingType,
		task.SourceValue, task.SourceField, task.ExpectedFieldType, candidatesJSON, task.Remark)

	return err
}

// GetPendingTaskByKey 根据task_key获取待确认任务
func GetPendingTaskByKey(database *sql.DB, taskKey string) (*PendingTask, error) {
	var task PendingTask
	var confirmedAt sql.NullTime

	err := database.QueryRow(`
		SELECT id, task_key, customer_id, order_id, row_number, mapping_type,
		       source_value, source_field, expected_field_type, candidates,
		       manual_value, status, confirmed_by, confirmed_at,
		       rejection_reason, remark, created_at, updated_at
		FROM pending_mapping_task
		WHERE task_key = $1
	`, taskKey).Scan(
		&task.ID, &task.TaskKey, &task.CustomerID, &task.OrderID, &task.RowNumber,
		&task.MappingType, &task.SourceValue, &task.SourceField, &task.ExpectedFieldType,
		&task.Candidates, &task.ManualValue, &task.Status, &task.ConfirmedBy,
		&confirmedAt, &task.RejectionReason, &task.Remark, &task.CreatedAt, &task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if confirmedAt.Valid {
		task.ConfirmedAt = &confirmedAt.Time
	}

	return &task, nil
}

// BatchCreatePendingTasks 批量创建待确认任务
func BatchCreatePendingTasks(database *sql.DB, tasks []PendingTask) error {
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, task := range tasks {
		task.TaskKey = fmt.Sprintf("%s#%s#%d#%s", task.CustomerID, task.OrderID, task.RowNumber, task.MappingType)

		candidatesJSON, err := json.Marshal(task.Candidates)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`
			INSERT INTO pending_mapping_task (
				task_key, customer_id, order_id, row_number, mapping_type,
				source_value, source_field, expected_field_type, candidates,
				status, remark
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 'PENDING', $10)
			ON CONFLICT (task_key) DO UPDATE SET
				candidates = EXCLUDED.candidates,
				source_value = EXCLUDED.source_value,
				updated_at = NOW()
		`, task.TaskKey, task.CustomerID, task.OrderID, task.RowNumber, task.MappingType,
			task.SourceValue, task.SourceField, task.ExpectedFieldType, candidatesJSON, task.Remark)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
