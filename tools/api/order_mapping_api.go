/**
 * order_mapping_api.go - 订单映射相关API
 * AI客服赋能系统 - B5任务API层
 *
 * 提供订单映射相关的查询和管理API：
 * - SKU映射查询和创建
 * - 门店映射查询和创建
 * - 仓库映射查询
 * - 配送方式映射查询
 * - 业务类型映射查询
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
)

// ==================== 数据结构定义 ====================

// SkuMappingRequest SKU映射请求
type SkuMappingRequest struct {
	CustomerID       string  `json:"customer_id"`
	CustomerSkuCode  string  `json:"customer_sku_code"`
	CustomerSkuName  string  `json:"customer_sku_name"`
	SystemSkuCode    string  `json:"system_sku_code"`
	SystemSkuName    string  `json:"system_sku_name"`
	UnitConversion   string  `json:"unit_conversion,omitempty"`
	ConfirmedBy      string  `json:"confirmed_by"`
}

// SkuMappingResponse SKU映射响应
type SkuMappingResponse struct {
	ID               int64     `json:"id"`
	CustomerID       string    `json:"customer_id"`
	CustomerSkuCode  string    `json:"customer_sku_code"`
	CustomerSkuName  string    `json:"customer_sku_name"`
	SystemSkuCode    string    `json:"system_sku_code"`
	SystemSkuName    string    `json:"system_sku_name"`
	UnitConversion   string    `json:"unit_conversion,omitempty"`
	Confidence       float64   `json:"confidence"`
	Source           string    `json:"source"`
	Status           string    `json:"status"`
	ConfirmedBy      string    `json:"confirmed_by,omitempty"`
	ConfirmedAt      *time.Time `json:"confirmed_at,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
}

// StoreMappingRequest 门店映射请求
type StoreMappingRequest struct {
	CustomerID  string  `json:"customer_id"`
	ShortName   string  `json:"short_name"`
	FullName    string  `json:"full_name"`
	Address     string  `json:"address,omitempty"`
	Contact     string  `json:"contact,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	ConfirmedBy string  `json:"confirmed_by"`
}

// StoreMappingResponse 门店映射响应
type StoreMappingResponse struct {
	ID         int64      `json:"id"`
	CustomerID string     `json:"customer_id"`
	ShortName  string     `json:"short_name"`
	FullName   string     `json:"full_name"`
	Address    string     `json:"address,omitempty"`
	Contact    string     `json:"contact,omitempty"`
	Phone      string     `json:"phone,omitempty"`
	MatchType  string     `json:"match_type"`
	Confidence float64    `json:"confidence"`
	ConfirmedBy string    `json:"confirmed_by,omitempty"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}

// WarehouseMappingResponse 仓库映射响应
type WarehouseMappingResponse struct {
	ID               int64     `json:"id"`
	CustomerID       string    `json:"customer_id"`
	WarehouseCode    string    `json:"warehouse_code"`
	WarehouseName    string    `json:"warehouse_name"`
	WarehouseDictCode string   `json:"warehouse_dict_code,omitempty"`
	DeliveryType     string    `json:"delivery_type,omitempty"`
	DefaultRegion    string    `json:"default_region,omitempty"`
	Priority         int       `json:"priority"`
	Enabled          bool      `json:"enabled"`
	CreatedAt        time.Time `json:"created_at"`
}

// DeliveryMethodMappingResponse 配送方式映射响应
type DeliveryMethodMappingResponse struct {
	ID         int64     `json:"id"`
	CustomerID string    `json:"customer_id,omitempty"`
	SourceValue string   `json:"source_value"`
	TargetValue string   `json:"target_value"`
	MatchType  string    `json:"match_type"`
	Confidence float64   `json:"confidence"`
	Priority   int       `json:"priority"`
	CreatedAt  time.Time `json:"created_at"`
}

// BusinessTypeMappingResponse 业务类型映射响应
type BusinessTypeMappingResponse struct {
	ID          int64     `json:"id"`
	CustomerID  string    `json:"customer_id,omitempty"`
	SourceValue string    `json:"source_value"`
	TargetValue string    `json:"target_value"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
}

// MappingMatchRequest 映射匹配请求
type MappingMatchRequest struct {
	CustomerID string `json:"customer_id"`
	Value      string `json:"value"`
	MappingType string `json:"mapping_type"` // SKU, STORE, WAREHOUSE, DELIVERY_METHOD, BUSINESS_TYPE
}

// MappingMatchResponse 映射匹配响应
type MappingMatchResponse struct {
	Success     bool        `json:"success"`
	MappingType string      `json:"mapping_type"`
	SourceValue string      `json:"source_value"`
	Matched     bool        `json:"matched"`
	Result      interface{} `json:"result,omitempty"`
	Candidates  interface{} `json:"candidates,omitempty"`
	Confidence  float64     `json:"confidence"`
}

// MappingListRequest 映射列表请求
type MappingListRequest struct {
	CustomerID  string `json:"customer_id"`
	MappingType string `json:"mapping_type"`
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
	Keyword     string `json:"keyword"`
}

// MappingListResponse 映射列表响应
type MappingListResponse struct {
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

// ==================== API处理器 ====================

// OrderMappingAPI 订单映射API处理器
type OrderMappingAPI struct {
	db *sql.DB
}

// NewOrderMappingAPI 创建订单映射API处理器
func NewOrderMappingAPI(database *sql.DB) *OrderMappingAPI {
	return &OrderMappingAPI{db: database}
}

// ==================== SKU映射API ====================

// GetSkuMappingHandler GET /api/mappings/sku
// 查询SKU映射
func (api *OrderMappingAPI) GetSkuMappingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	customerID := r.URL.Query().Get("customer_id")
	customerSkuCode := r.URL.Query().Get("customer_sku_code")

	if customerID == "" || customerSkuCode == "" {
		respondWithError(w, http.StatusBadRequest, "customer_id and customer_sku_code are required")
		return
	}

	mapping, err := api.getSkuMapping(customerID, customerSkuCode)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Code:    200,
				Message: "Mapping not found",
				Data:    nil,
			})
			return
		}
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Query failed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    mapping,
	})
}

// getSkuMapping 查询SKU映射
func (api *OrderMappingAPI) getSkuMapping(customerID, customerSkuCode string) (*SkuMappingResponse, error) {
	var mapping SkuMappingResponse
	var confirmedAt sql.NullTime
	var unitConversion sql.NullString

	err := api.db.QueryRow(`
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
		       system_sku_code, system_sku_name, unit_conversion_rule,
		       confidence, source, status, confirmed_by, confirmed_at, created_at
		FROM customer_sku_mapping
		WHERE customer_id = $1 AND customer_sku_code = $2 AND status = 'ACTIVE'
	`, customerID, customerSkuCode).Scan(
		&mapping.ID, &mapping.CustomerID, &mapping.CustomerSkuCode, &mapping.CustomerSkuName,
		&mapping.SystemSkuCode, &mapping.SystemSkuName, &unitConversion,
		&mapping.Confidence, &mapping.Source, &mapping.Status, &mapping.ConfirmedBy,
		&confirmedAt, &mapping.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	if confirmedAt.Valid {
		mapping.ConfirmedAt = &confirmedAt.Time
	}
	if unitConversion.Valid {
		mapping.UnitConversion = unitConversion.String
	}

	return &mapping, nil
}

// CreateSkuMappingHandler POST /api/mappings/sku
// 创建SKU映射
func (api *OrderMappingAPI) CreateSkuMappingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req SkuMappingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	// 参数校验
	if req.CustomerID == "" || req.CustomerSkuCode == "" || req.SystemSkuCode == "" {
		respondWithError(w, http.StatusBadRequest, "customer_id, customer_sku_code and system_sku_code are required")
		return
	}

	mapping, err := api.createSkuMapping(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create mapping: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Code:    201,
		Message: "SKU mapping created successfully",
		Data:    mapping,
	})
}

// createSkuMapping 创建SKU映射
func (api *OrderMappingAPI) createSkuMapping(req SkuMappingRequest) (*SkuMappingResponse, error) {
	now := time.Now()

	_, err := api.db.Exec(`
		INSERT INTO customer_sku_mapping (
			customer_id, customer_sku_code, customer_sku_name,
			system_sku_code, system_sku_name, unit_conversion_rule,
			confidence, source, status, confirmed_by, confirmed_at
		) VALUES ($1, $2, $3, $4, $5, $6, 1.0, 'MANUAL', 'ACTIVE', $7, $8)
		ON CONFLICT (customer_id, customer_sku_code) DO UPDATE SET
			system_sku_code = EXCLUDED.system_sku_code,
			system_sku_name = EXCLUDED.system_sku_name,
			unit_conversion_rule = EXCLUDED.unit_conversion_rule,
			confidence = EXCLUDED.confidence,
			confirmed_by = EXCLUDED.confirmed_by,
			confirmed_at = EXCLUDED.confirmed_at,
			updated_at = NOW()
	`, req.CustomerID, req.CustomerSkuCode, req.CustomerSkuName,
		req.SystemSkuCode, req.SystemSkuName, req.UnitConversion,
		req.ConfirmedBy, now)

	if err != nil {
		return nil, err
	}

	return api.getSkuMapping(req.CustomerID, req.CustomerSkuCode)
}

// ListSkuMappingsHandler GET /api/mappings/sku/list
// 查询SKU映射列表
func (api *OrderMappingAPI) ListSkuMappingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	req := MappingListRequest{
		Page:     1,
		PageSize: 20,
	}

	if page := r.URL.Query().Get("page"); page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			req.Page = p
		}
	}

	if pageSize := r.URL.Query().Get("page_size"); pageSize != "" {
		if ps, err := strconv.Atoi(pageSize); err == nil && ps > 0 {
			if ps > 100 {
				ps = 100
			}
			req.PageSize = ps
		}
	}

	req.CustomerID = r.URL.Query().Get("customer_id")
	req.Keyword = r.URL.Query().Get("keyword")

	mappings, total, err := api.listSkuMappings(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Query failed: %v", err))
		return
	}

	totalPages := int(total) / req.PageSize
	if int(total)%req.PageSize > 0 {
		totalPages++
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data: MappingListResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
			Data:       mappings,
		},
	})
}

// listSkuMappings 查询SKU映射列表
func (api *OrderMappingAPI) listSkuMappings(req MappingListRequest) ([]SkuMappingResponse, int64, error) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	if req.CustomerID != "" {
		conditions = append(conditions, fmt.Sprintf("customer_id = $%d", argIndex))
		args = append(args, req.CustomerID)
		argIndex++
	}

	if req.Keyword != "" {
		conditions = append(conditions, fmt.Sprintf("(customer_sku_code ILIKE $%d OR customer_sku_name ILIKE $%d OR system_sku_name ILIKE $%d)", argIndex, argIndex, argIndex))
		args = append(args, "%"+req.Keyword+"%")
		argIndex++
	}

	conditions = append(conditions, "status = 'ACTIVE'")

	whereClause := "WHERE " + strings.Join(conditions, " AND ")

	// 查询总数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM customer_sku_mapping %s", whereClause)
	var total int64
	err := api.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 查询数据
	query := fmt.Sprintf(`
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
		       system_sku_code, system_sku_name, unit_conversion_rule,
		       confidence, source, status, confirmed_by, confirmed_at, created_at
		FROM customer_sku_mapping
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)

	offset := (req.Page - 1) * req.PageSize
	args = append(args, req.PageSize, offset)

	rows, err := api.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var mappings []SkuMappingResponse
	for rows.Next() {
		var mapping SkuMappingResponse
		var confirmedAt sql.NullTime
		var unitConversion sql.NullString

		err := rows.Scan(
			&mapping.ID, &mapping.CustomerID, &mapping.CustomerSkuCode, &mapping.CustomerSkuName,
			&mapping.SystemSkuCode, &mapping.SystemSkuName, &unitConversion,
			&mapping.Confidence, &mapping.Source, &mapping.Status, &mapping.ConfirmedBy,
			&confirmedAt, &mapping.CreatedAt,
		)
		if err != nil {
			continue
		}

		if confirmedAt.Valid {
			mapping.ConfirmedAt = &confirmedAt.Time
		}
		if unitConversion.Valid {
			mapping.UnitConversion = unitConversion.String
		}

		mappings = append(mappings, mapping)
	}

	return mappings, total, rows.Err()
}

// ==================== 门店映射API ====================

// GetStoreMappingHandler GET /api/mappings/store
// 查询门店映射
func (api *OrderMappingAPI) GetStoreMappingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	shortName := r.URL.Query().Get("short_name")
	if shortName == "" {
		respondWithError(w, http.StatusBadRequest, "short_name is required")
		return
	}

	mapping, err := api.getStoreMapping(shortName)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Code:    200,
				Message: "Mapping not found",
				Data:    nil,
			})
			return
		}
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Query failed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    mapping,
	})
}

// getStoreMapping 查询门店映射
func (api *OrderMappingAPI) getStoreMapping(shortName string) (*StoreMappingResponse, error) {
	var mapping StoreMappingResponse
	var customerID sql.NullString
	var address, contact, phone sql.NullString
	var confirmedAt sql.NullTime
	var confirmedBy sql.NullString

	err := api.db.QueryRow(`
		SELECT id, customer_id, short_name, full_name, match_type,
		       confidence, address, contact, phone, confirmed_by, confirmed_at, created_at
		FROM customer_alias_mapping
		WHERE short_name = $1
	`, shortName).Scan(
		&mapping.ID, &customerID, &mapping.ShortName, &mapping.FullName, &mapping.MatchType,
		&mapping.Confidence, &address, &contact, &phone, &confirmedBy,
		&confirmedAt, &mapping.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	if customerID.Valid {
		mapping.CustomerID = customerID.String
	}
	if address.Valid {
		mapping.Address = address.String
	}
	if contact.Valid {
		mapping.Contact = contact.String
	}
	if phone.Valid {
		mapping.Phone = phone.String
	}
	if confirmedBy.Valid {
		mapping.ConfirmedBy = confirmedBy.String
	}
	if confirmedAt.Valid {
		mapping.ConfirmedAt = &confirmedAt.Time
	}

	return &mapping, nil
}

// CreateStoreMappingHandler POST /api/mappings/store
// 创建门店映射
func (api *OrderMappingAPI) CreateStoreMappingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req StoreMappingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	if req.ShortName == "" || req.FullName == "" {
		respondWithError(w, http.StatusBadRequest, "short_name and full_name are required")
		return
	}

	mapping, err := api.createStoreMapping(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create mapping: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Code:    201,
		Message: "Store mapping created successfully",
		Data:    mapping,
	})
}

// createStoreMapping 创建门店映射
func (api *OrderMappingAPI) createStoreMapping(req StoreMappingRequest) (*StoreMappingResponse, error) {
	now := time.Now()

	_, err := api.db.Exec(`
		INSERT INTO customer_alias_mapping (
			customer_id, short_name, full_name, match_type, confidence,
			address, contact, phone, confirmed_by, confirmed_at
		) VALUES ($1, $2, $3, 'EXACT', 1.0, $4, $5, $6, $7, $8)
		ON CONFLICT (short_name) DO UPDATE SET
			full_name = EXCLUDED.full_name,
			match_type = EXCLUDED.match_type,
			confidence = EXCLUDED.confidence,
			address = EXCLUDED.address,
			contact = EXCLUDED.contact,
			phone = EXCLUDED.phone,
			confirmed_by = EXCLUDED.confirmed_by,
			confirmed_at = EXCLUDED.confirmed_at
	`, req.CustomerID, req.ShortName, req.FullName,
		req.Address, req.Contact, req.Phone, req.ConfirmedBy, now)

	if err != nil {
		return nil, err
	}

	return api.getStoreMapping(req.ShortName)
}

// ==================== 仓库映射API ====================

// GetWarehouseMappingHandler GET /api/mappings/warehouse
// 查询仓库映射
func (api *OrderMappingAPI) GetWarehouseMappingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	customerID := r.URL.Query().Get("customer_id")
	if customerID == "" {
		respondWithError(w, http.StatusBadRequest, "customer_id is required")
		return
	}

	mappings, err := api.getWarehouseMappings(customerID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Query failed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    mappings,
	})
}

// getWarehouseMappings 查询仓库映射列表
func (api *OrderMappingAPI) getWarehouseMappings(customerID string) ([]WarehouseMappingResponse, error) {
	rows, err := api.db.Query(`
		SELECT id, customer_id, warehouse_code, warehouse_name, warehouse_dict_code,
		       delivery_type, default_region, priority, enabled, created_at
		FROM customer_warehouse_mapping
		WHERE customer_id = $1 AND enabled = true
		ORDER BY priority ASC
	`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mappings []WarehouseMappingResponse
	for rows.Next() {
		var mapping WarehouseMappingResponse
		var deliveryType, defaultRegion, warehouseDictCode sql.NullString

		err := rows.Scan(
			&mapping.ID, &mapping.CustomerID, &mapping.WarehouseCode, &mapping.WarehouseName,
			&warehouseDictCode, &deliveryType, &defaultRegion, &mapping.Priority,
			&mapping.Enabled, &mapping.CreatedAt,
		)
		if err != nil {
			continue
		}

		if deliveryType.Valid {
			mapping.DeliveryType = deliveryType.String
		}
		if defaultRegion.Valid {
			mapping.DefaultRegion = defaultRegion.String
		}
		if warehouseDictCode.Valid {
			mapping.WarehouseDictCode = warehouseDictCode.String
		}

		mappings = append(mappings, mapping)
	}

	return mappings, rows.Err()
}

// ==================== 配送方式映射API ====================

// GetDeliveryMethodMappingsHandler GET /api/mappings/delivery-method
// 查询配送方式映射
func (api *OrderMappingAPI) GetDeliveryMethodMappingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	customerID := r.URL.Query().Get("customer_id")
	sourceValue := r.URL.Query().Get("source_value")

	mappings, err := api.getDeliveryMethodMappings(customerID, sourceValue)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Query failed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    mappings,
	})
}

// getDeliveryMethodMappings 查询配送方式映射
func (api *OrderMappingAPI) getDeliveryMethodMappings(customerID, sourceValue string) ([]DeliveryMethodMappingResponse, error) {
	var query string
	var args []interface{}

	if customerID != "" && sourceValue != "" {
		query = `
			SELECT id, customer_id, source_value, target_value, match_type, confidence, priority, created_at
			FROM delivery_method_mapping
			WHERE (customer_id = $1 OR customer_id IS NULL) AND source_value = $2
			ORDER BY priority ASC
		`
		args = []interface{}{customerID, sourceValue}
	} else if customerID != "" {
		query = `
			SELECT id, customer_id, source_value, target_value, match_type, confidence, priority, created_at
			FROM delivery_method_mapping
			WHERE customer_id = $1 OR customer_id IS NULL
			ORDER BY priority ASC
		`
		args = []interface{}{customerID}
	} else {
		query = `
			SELECT id, customer_id, source_value, target_value, match_type, confidence, priority, created_at
			FROM delivery_method_mapping
			WHERE customer_id IS NULL
			ORDER BY priority ASC
		`
	}

	rows, err := api.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mappings []DeliveryMethodMappingResponse
	for rows.Next() {
		var mapping DeliveryMethodMappingResponse
		var customerID sql.NullString

		err := rows.Scan(
			&mapping.ID, &customerID, &mapping.SourceValue, &mapping.TargetValue,
			&mapping.MatchType, &mapping.Confidence, &mapping.Priority, &mapping.CreatedAt,
		)
		if err != nil {
			continue
		}

		if customerID.Valid {
			mapping.CustomerID = customerID.String
		}

		mappings = append(mappings, mapping)
	}

	return mappings, rows.Err()
}

// ==================== 业务类型映射API ====================

// GetBusinessTypeMappingsHandler GET /api/mappings/business-type
// 查询业务类型映射
func (api *OrderMappingAPI) GetBusinessTypeMappingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	customerID := r.URL.Query().Get("customer_id")
	sourceValue := r.URL.Query().Get("source_value")

	mappings, err := api.getBusinessTypeMappings(customerID, sourceValue)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Query failed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    mappings,
	})
}

// getBusinessTypeMappings 查询业务类型映射
func (api *OrderMappingAPI) getBusinessTypeMappings(customerID, sourceValue string) ([]BusinessTypeMappingResponse, error) {
	var query string
	var args []interface{}

	if customerID != "" && sourceValue != "" {
		query = `
			SELECT id, customer_id, source_value, target_value, priority, created_at
			FROM business_type_mapping
			WHERE (customer_id = $1 OR customer_id IS NULL) AND source_value = $2
			ORDER BY priority ASC
		`
		args = []interface{}{customerID, sourceValue}
	} else if customerID != "" {
		query = `
			SELECT id, customer_id, source_value, target_value, priority, created_at
			FROM business_type_mapping
			WHERE customer_id = $1 OR customer_id IS NULL
			ORDER BY priority ASC
		`
		args = []interface{}{customerID}
	} else {
		query = `
			SELECT id, customer_id, source_value, target_value, priority, created_at
			FROM business_type_mapping
			WHERE customer_id IS NULL
			ORDER BY priority ASC
		`
	}

	rows, err := api.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mappings []BusinessTypeMappingResponse
	for rows.Next() {
		var mapping BusinessTypeMappingResponse
		var customerID sql.NullString

		err := rows.Scan(
			&mapping.ID, &customerID, &mapping.SourceValue, &mapping.TargetValue,
			&mapping.Priority, &mapping.CreatedAt,
		)
		if err != nil {
			continue
		}

		if customerID.Valid {
			mapping.CustomerID = customerID.String
		}

		mappings = append(mappings, mapping)
	}

	return mappings, rows.Err()
}

// ==================== 通用映射匹配API ====================

// MatchMappingHandler POST /api/mappings/match
// 通用映射匹配接口
func (api *OrderMappingAPI) MatchMappingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req MappingMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	if req.Value == "" || req.MappingType == "" {
		respondWithError(w, http.StatusBadRequest, "value and mapping_type are required")
		return
	}

	result, err := api.matchMapping(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Match failed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// matchMapping 执行映射匹配
func (api *OrderMappingAPI) matchMapping(req MappingMatchRequest) (*MappingMatchResponse, error) {
	switch strings.ToUpper(req.MappingType) {
	case "SKU":
		return api.matchSkuMapping(req)
	case "STORE":
		return api.matchStoreMapping(req)
	case "WAREHOUSE":
		return api.matchWarehouseMapping(req)
	case "DELIVERY_METHOD":
		return api.matchDeliveryMethodMapping(req)
	case "BUSINESS_TYPE":
		return api.matchBusinessTypeMapping(req)
	default:
		return nil, fmt.Errorf("unsupported mapping type: %s", req.MappingType)
	}
}

// matchSkuMapping 匹配SKU映射
func (api *OrderMappingAPI) matchSkuMapping(req MappingMatchRequest) (*MappingMatchResponse, error) {
	mapping, err := api.getSkuMapping(req.CustomerID, req.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return &MappingMatchResponse{
				Success:     true,
				MappingType: req.MappingType,
				SourceValue: req.Value,
				Matched:     false,
				Confidence:  0,
			}, nil
		}
		return nil, err
	}

	return &MappingMatchResponse{
		Success:     true,
		MappingType: req.MappingType,
		SourceValue: req.Value,
		Matched:     true,
		Result:      mapping,
		Confidence:  mapping.Confidence,
	}, nil
}

// matchStoreMapping 匹配门店映射
func (api *OrderMappingAPI) matchStoreMapping(req MappingMatchRequest) (*MappingMatchResponse, error) {
	mapping, err := api.getStoreMapping(req.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return &MappingMatchResponse{
				Success:     true,
				MappingType: req.MappingType,
				SourceValue: req.Value,
				Matched:     false,
				Confidence:  0,
			}, nil
		}
		return nil, err
	}

	return &MappingMatchResponse{
		Success:     true,
		MappingType: req.MappingType,
		SourceValue: req.Value,
		Matched:     true,
		Result:      mapping,
		Confidence:  mapping.Confidence,
	}, nil
}

// matchWarehouseMapping 匹配仓库映射
func (api *OrderMappingAPI) matchWarehouseMapping(req MappingMatchRequest) (*MappingMatchResponse, error) {
	mappings, err := api.getWarehouseMappings(req.CustomerID)
	if err != nil {
		return nil, err
	}

	if len(mappings) == 0 {
		return &MappingMatchResponse{
			Success:     true,
			MappingType: req.MappingType,
			SourceValue: req.Value,
			Matched:     false,
			Confidence:  0,
		}, nil
	}

	// 返回第一个（优先级最高）的仓库
	return &MappingMatchResponse{
		Success:     true,
		MappingType: req.MappingType,
		SourceValue: req.Value,
		Matched:     true,
		Result:      mappings[0],
		Confidence:  1.0,
	}, nil
}

// matchDeliveryMethodMapping 匹配配送方式映射
func (api *OrderMappingAPI) matchDeliveryMethodMapping(req MappingMatchRequest) (*MappingMatchResponse, error) {
	mappings, err := api.getDeliveryMethodMappings(req.CustomerID, req.Value)
	if err != nil {
		return nil, err
	}

	if len(mappings) == 0 {
		return &MappingMatchResponse{
			Success:     true,
			MappingType: req.MappingType,
			SourceValue: req.Value,
			Matched:     false,
			Confidence:  0,
		}, nil
	}

	return &MappingMatchResponse{
		Success:     true,
		MappingType: req.MappingType,
		SourceValue: req.Value,
		Matched:     true,
		Result:      mappings[0],
		Confidence:  mappings[0].Confidence,
	}, nil
}

// matchBusinessTypeMapping 匹配业务类型映射
func (api *OrderMappingAPI) matchBusinessTypeMapping(req MappingMatchRequest) (*MappingMatchResponse, error) {
	mappings, err := api.getBusinessTypeMappings(req.CustomerID, req.Value)
	if err != nil {
		return nil, err
	}

	if len(mappings) == 0 {
		return &MappingMatchResponse{
			Success:     true,
			MappingType: req.MappingType,
			SourceValue: req.Value,
			Matched:     false,
			Confidence:  0,
		}, nil
	}

	return &MappingMatchResponse{
		Success:     true,
		MappingType: req.MappingType,
		SourceValue: req.Value,
		Matched:     true,
		Result:      mappings[0],
		Confidence:  1.0,
	}, nil
}
