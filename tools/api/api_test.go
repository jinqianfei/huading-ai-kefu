/**
 * api_test.go - API测试用例
 * AI客服/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    =/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T)/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	d/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1",/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		)/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2,/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3,/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1",/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t */**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := New/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", test/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", test/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expected/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt :=/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr :=/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expected/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response);/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data !=/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data,/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expected/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTask/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNot/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody:/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				Confirmed/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.Status/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setup/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanup/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRow/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
			/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
		/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
				OrderID:   testOrder/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
				OrderID:   testOrderID,
				RowNumber: 1,
			},
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
				OrderID:   testOrderID,
				RowNumber: 1,
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
	/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
				OrderID:   testOrderID,
				RowNumber: 1,
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name: "缺少order_id",
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
				OrderID:   testOrderID,
				RowNumber: 1,
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name: "缺少order_id",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomer/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
				OrderID:   testOrderID,
				RowNumber: 1,
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name: "缺少order_id",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				RowNumber:  1,
			},
/**
 * api_test.go - API测试用例
 * AI客服赋能系统 - B5任务API层
 *
 * 包含所有API端点的单元测试和集成测试
 */

package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// ==================== 测试配置 ====================

const (
	testDBConnStr = "host=localhost port=5432 user=test password=test dbname=ai_cs_support_test sslmode=disable"
	testCustomerID = "CUST001"
	testOrderID    = "ORD202605070001"
)

// setupTestDB 设置测试数据库连接
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConnStr)
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		t.Skipf("Skipping test: cannot connect to database: %v", err)
		return nil
	}

	return db
}

// setupTestData 设置测试数据
func setupTestData(t *testing.T, db *sql.DB) {
	// 清理旧数据
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)

	// 插入测试用的待确认任务
	candidates, _ := json.Marshal([]Candidate{
		{Value: "SKU001", Label: "测试商品A", Source: "SYSTEM", Confidence: 0.95},
		{Value: "SKU002", Label: "测试商品B", Source: "HISTORY", Confidence: 0.85},
	})

	_, err := db.Exec(`
		INSERT INTO pending_mapping_task (
			task_key, customer_id, order_id, row_number, mapping_type,
			source_value, source_field, expected_field_type, candidates, status
		) VALUES 
		($1, $2, $3, 1, 'SKU', 'P1556689351', 'product_code', 'string', $4, 'PENDING'),
		($5, $2, $3, 1, 'STORE', '王口镇店', 'store_name', 'string', $6, 'PENDING'),
		($7, $2, $3, 2, 'SKU', 'P1556689352', 'product_code', 'string', $4, 'PENDING')
	`,
		fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID),
		testCustomerID,
		testOrderID,
		candidates,
		fmt.Sprintf("%s#%s#1#STORE", testCustomerID, testOrderID),
		`[{"value":"天津王口镇店","source":"SYSTEM","confidence":0.9}]`,
		fmt.Sprintf("%s#%s#2#SKU", testCustomerID, testOrderID),
	)

	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
}

// cleanupTestData 清理测试数据
func cleanupTestData(t *testing.T, db *sql.DB) {
	db.Exec("DELETE FROM pending_mapping_task WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_sku_mapping WHERE customer_id = $1", testCustomerID)
	db.Exec("DELETE FROM customer_alias_mapping WHERE customer_id = $1", testCustomerID)
}

// ==================== PendingTaskAPI测试 ====================

func TestGetPendingTasksHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		expectedCount  int
	}{
		{
			name:           "查询所有待确认任务",
			queryParams:    fmt.Sprintf("customer_id=%s&status=PENDING", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "按mapping_type筛选",
			queryParams:    fmt.Sprintf("customer_id=%s&mapping_type=SKU", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "分页查询",
			queryParams:    fmt.Sprintf("customer_id=%s&page=1&page_size=2", testCustomerID),
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name:           "无参数查询",
			queryParams:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "/api/pending-tasks"
			if tt.queryParams != "" {
				url = url + "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			api.GetPendingTasksHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedCount > 0 {
				var response APIResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if response.Data != nil {
					data, _ := json.Marshal(response.Data)
					var listResp PendingTaskListResponse
					json.Unmarshal(data, &listResp)

					if len(listResp.Tasks) != tt.expectedCount {
						t.Errorf("Expected %d tasks, got %d", tt.expectedCount, len(listResp.Tasks))
					}
				}
			}
		})
	}
}

func TestConfirmTaskHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	taskKey := fmt.Sprintf("%s#%s#1#SKU", testCustomerID, testOrderID)

	tests := []struct {
		name           string
		taskKey        string
		requestBody    ConfirmTaskRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name:    "成功确认任务",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
				Remark:      "测试确认",
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name:    "任务不存在",
			taskKey: "NON_EXISTENT_KEY",
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusNotFound,
			expectSuccess:  false,
		},
		{
			name:    "缺少manual_value",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ConfirmedBy: "test_user",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name:    "缺少confirmed_by",
			taskKey: taskKey,
			requestBody: ConfirmTaskRequest{
				ManualValue: "SKU001",
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重新设置测试数据（因为前面的测试可能已经更新了状态）
			setupTestData(t, db)

			body, _ := json.Marshal(tt.requestBody)
			url := fmt.Sprintf("/api/pending-tasks/%s/confirm", tt.taskKey)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			api.ConfirmTaskHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var response APIResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if response.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, response.Success)
			}
		})
	}
}

func TestProcessRowHandler(t *testing.T) {
	db := setupTestDB(t)
	if db == nil {
		return
	}
	defer db.Close()

	setupTestData(t, db)
	defer cleanupTestData(t, db)

	api := NewPendingTaskAPI(db)

	tests := []struct {
		name           string
		requestBody    ProcessRowRequest
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "成功处理行（所有字段已确认）",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				OrderID:    testOrderID,
				RowNumber:  1,
				ConfirmedMappings: map[string]string{
					"SKU":   "SKU001",
					"STORE": "天津王口镇店",
				},
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "缺少customer_id",
			requestBody: ProcessRowRequest{
				OrderID:   testOrderID,
				RowNumber: 1,
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false,
		},
		{
			name: "缺少order_id",
			requestBody: ProcessRowRequest{
				CustomerID: testCustomerID,
				RowNumber:  1,
			},
			expectedStatus: http.StatusBadRequest,
			expectSuccess:  false