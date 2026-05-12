/**
 * OpenClaw Agent Core 数据库模块
 * 负责所有与数据库相关的操作
 */

package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// 全局数据库连接（由 main.go 初始化）
var db *sql.DB

// InitDB 初始化数据库连接
func InitDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// ==================== 客户相关 ====================

// Customer 客户结构
type Customer struct {
	ID        string    `json:"customer_id"`
	Name      string    `json:"customer_name"`
	Code      string    `json:"customer_code"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllCustomers 获取所有客户
func GetAllCustomers() ([]Customer, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT customer_id, customer_name, customer_code, status, created_at
		FROM customer WHERE status = 'active' ORDER BY customer_name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.Status, &c.CreatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, rows.Err()
}

// GetCustomer 获取单个客户
func GetCustomer(customerID string) (*Customer, error) {
	var c Customer
	err := db.QueryRowContext(context.Background(), `
		SELECT customer_id, customer_name, customer_code, status, created_at
		FROM customer WHERE customer_id = $1
	`, customerID).Scan(&c.ID, &c.Name, &c.Code, &c.Status, &c.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// ==================== SKU映射相关 ====================

// CustomerSkuMapping 客户SKU映射结构
type CustomerSkuMapping struct {
	ID                  string    `json:"id"`
	CustomerID           string    `json:"customer_id"`
	CustomerSkuCode     string    `json:"customer_sku_code"`
	CustomerSkuName     string    `json:"customer_sku_name"`
	SystemSkuCode       string    `json:"system_sku_code"`
	SystemSkuName       string    `json:"system_sku_name"`
	Confidence          float64   `json:"confidence"`
	Source              string    `json:"source"`
	UnitConversionRule  string    `json:"unit_conversion_rule"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// GetCustomerSkuMapping 获取客户商品编码→系统SKU的精确映射
func GetCustomerSkuMapping(customerID, customerSkuCode string) (*CustomerSkuMapping, error) {
	var m CustomerSkuMapping
	err := db.QueryRowContext(context.Background(), `
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
			   system_sku_code, system_sku_name, confidence, source,
			   unit_conversion_rule, status, created_at, updated_at
		FROM customer_sku_mapping
		WHERE customer_id = $1 AND customer_sku_code = $2 AND status = 'active'
	`, customerID, customerSkuCode).Scan(
		&m.ID, &m.CustomerID, &m.CustomerSkuCode, &m.CustomerSkuName,
		&m.SystemSkuCode, &m.SystemSkuName, &m.Confidence, &m.Source,
		&m.UnitConversionRule, &m.Status, &m.CreatedAt, &m.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// SaveCustomerSkuMapping 保存客户SKU映射
func SaveCustomerSkuMapping(m *CustomerSkuMapping) error {
	_, err := db.ExecContext(context.Background(), `
		INSERT INTO customer_sku_mapping (
			customer_id, customer_sku_code, customer_sku_name,
			system_sku_code, system_sku_name, confidence, source,
			unit_conversion_rule, status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (customer_id, customer_sku_code) DO UPDATE SET
			system_sku_code = EXCLUDED.system_sku_code,
			system_sku_name = EXCLUDED.system_sku_name,
			confidence = EXCLUDED.confidence,
			updated_at = NOW()
	`, m.CustomerID, m.CustomerSkuCode, m.CustomerSkuName,
		m.SystemSkuCode, m.SystemSkuName, m.Confidence, m.Source,
		m.UnitConversionRule, m.Status)
	return err
}

// ==================== 门店别名相关 ====================

// StoreAliasMapping 门店别名映射结构
type StoreAliasMapping struct {
	ID          string  `json:"id"`
	CustomerID  string  `json:"customer_id"`
	Alias       string  `json:"alias"`
	FullName    string  `json:"full_name"`
	MatchType   string  `json:"match_type"`
	Confidence  float64 `json:"confidence"`
	Address     string  `json:"address"`
	Contact     string  `json:"contact"`
	Phone       string  `json:"phone"`
}

// GetStoreAliasMapping 获取门店简称映射
func GetStoreAliasMapping(alias string) (*StoreAliasMapping, error) {
	var m StoreAliasMapping
	err := db.QueryRowContext(context.Background(), `
		SELECT id, customer_id, alias, full_name, match_type,
			   confidence, address, contact, phone
		FROM customer_alias_mapping
		WHERE alias = $1 AND status = 'active'
		LIMIT 1
	`, alias).Scan(&m.ID, &m.CustomerID, &m.Alias, &m.FullName,
		&m.MatchType, &m.Confidence, &m.Address, &m.Contact, &m.Phone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// ==================== 匹配历史 ====================

// MatchHistory 匹配历史结构
type MatchHistory struct {
	ID              string    `json:"id"`
	CustomerID      string    `json:"customer_id"`
	CustomerSkuCode string    `json:"customer_sku_code"`
	CustomerSkuName string    `json:"customer_sku_name"`
	SystemSkuCode   string    `json:"system_sku_code"`
	SystemSkuName   string    `json:"system_sku_name"`
	MatchType       string    `json:"match_type"`
	Confidence      float64   `json:"confidence"`
	Corrected       bool      `json:"corrected"`
	CorrectedBy     string    `json:"corrected_by"`
	CreatedAt       time.Time `json:"created_at"`
}

// GetMatchHistory 获取匹配历史
func GetMatchHistory(customerID, customerSkuCode string) (*MatchHistory, error) {
	var h MatchHistory
	err := db.QueryRowContext(context.Background(), `
		SELECT id, customer_id, customer_sku_code, customer_sku_name,
			   system_sku_code, system_sku_name, match_type, confidence,
			   corrected, corrected_by, created_at
		FROM match_history
		WHERE customer_id = $1 AND customer_sku_code = $2
		ORDER BY created_at DESC LIMIT 1
	`, customerID, customerSkuCode).Scan(
		&h.ID, &h.CustomerID, &h.CustomerSkuCode, &h.CustomerSkuName,
		&h.SystemSkuCode, &h.SystemSkuName, &h.MatchType, &h.Confidence,
		&h.Corrected, &h.CorrectedBy, &h.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &h, nil
}

// SaveMatchHistory 保存匹配历史
func SaveMatchHistory(h *MatchHistory) error {
	_, err := db.ExecContext(context.Background(), `
		INSERT INTO match_history (
			customer_id, customer_sku_code, customer_sku_name,
			system_sku_code, system_sku_name, match_type, confidence
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, h.CustomerID, h.CustomerSkuCode, h.CustomerSkuName,
		h.SystemSkuCode, h.SystemSkuName, h.MatchType, h.Confidence)
	return err
}

// ==================== 系统SKU ====================

// SystemSku 系统SKU结构
type SystemSku struct {
	ID       string    `json:"id"`
	SkuCode  string    `json:"sku_code"`
	SkuName  string    `json:"sku_name"`
	Category string    `json:"category"`
	Unit     string    `json:"unit"`
	Status   string    `json:"status"`
}

// GetSystemSkuCandidates 获取所有可用的系统SKU候选
func GetSystemSkuCandidates() ([]SystemSku, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT id, sku_code, sku_name, category, unit, status
		FROM system_sku WHERE status = 'active'
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skus []SystemSku
	for rows.Next() {
		var s SystemSku
		if err := rows.Scan(&s.ID, &s.SkuCode, &s.SkuName, &s.Category, &s.Unit, &s.Status); err != nil {
			return nil, err
		}
		skus = append(skus, s)
	}
	return skus, rows.Err()
}

// ==================== 合同规则相关 ====================

// ContractRule 合同规则结构
type ContractRule struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customer_id"`
	ContractNo   string    `json:"contract_no"`
	RuleCategory string    `json:"rule_category"`
	FieldName    string    `json:"field_name"`
	FieldValue   string    `json:"field_value"`
	FieldUnit    string    `json:"field_unit"`
	Source       string    `json:"source"`
	Confidence   float64   `json:"confidence"`
	Status       string    `json:"status"`
	EffDate      time.Time `json:"effective_date"`
	ExpDate      time.Time `json:"expiry_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GetContractRules 获取客户的合同规则
func GetContractRules(customerID string) ([]ContractRule, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT id, customer_id, contract_no, rule_category, field_name,
			   field_value, field_unit, source, confidence, status,
			   effective_date, expiry_date, created_at, updated_at
		FROM customer_contract_rules
		WHERE customer_id = $1
			AND status = 'active'
			AND effective_date <= NOW()
			AND expiry_date >= NOW()
		ORDER BY rule_category, field_name
	`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []ContractRule
	for rows.Next() {
		var r ContractRule
		if err := rows.Scan(&r.ID, &r.CustomerID, &r.ContractNo, &r.RuleCategory,
			&r.FieldName, &r.FieldValue, &r.FieldUnit, &r.Source, &r.Confidence,
			&r.Status, &r.EffDate, &r.ExpDate, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		rules = append(rules, r)
	}
	return rules, rows.Err()
}

// SaveContractRule 保存合同规则
func SaveContractRule(r *ContractRule) error {
	_, err := db.ExecContext(context.Background(), `
		INSERT INTO customer_contract_rules (
			customer_id, contract_no, rule_category, field_name,
			field_value, field_unit, source, confidence, status,
			effective_date, expiry_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (customer_id, contract_no, field_name) DO UPDATE SET
			field_value = EXCLUDED.field_value,
			confidence = EXCLUDED.confidence,
			updated_at = NOW()
	`, r.CustomerID, r.ContractNo, r.RuleCategory, r.FieldName,
		r.FieldValue, r.FieldUnit, r.Source, r.Confidence, r.Status,
		r.EffDate, r.ExpDate)
	return err
}

// ==================== 出库记录相关 ====================

// OutboundRecord 出库记录结构
type OutboundRecord struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customer_id"`
	OrderNo      string    `json:"order_no"`
	OutboundDate time.Time `json:"outbound_date"`
	StoreName    string    `json:"store_name"`
	SkuCode      string    `json:"sku_code"`
	SkuName      string    `json:"sku_name"`
	Quantity     float64   `json:"quantity"`
	Unit         string    `json:"unit"`
	Weight       float64   `json:"weight"`
	Region       string    `json:"region"`
}

// GetOutboundRecords 获取出库记录
func GetOutboundRecords(customerID, startDate, endDate string) ([]OutboundRecord, error) {
	rows, err := db.QueryContext(context.Background(), `
		SELECT id, customer_id, order_no, outbound_date, store_name,
			   sku_code, sku_name, quantity, unit, weight, region
		FROM outbound_records
		WHERE customer_id = $1
			AND outbound_date >= $2
			AND outbound_date <= $3
		ORDER BY outbound_date DESC
	`, customerID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []OutboundRecord
	for rows.Next() {
		var r OutboundRecord
		if err := rows.Scan(&r.ID, &r.CustomerID, &r.OrderNo, &r.OutboundDate,
			&r.StoreName, &r.SkuCode, &r.SkuName, &r.Quantity, &r.Unit,
			&r.Weight, &r.Region); err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, rows.Err()
}
