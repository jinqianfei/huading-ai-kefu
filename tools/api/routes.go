/**
 * routes.go - API路由注册
 * AI客服赋能系统 - B5任务API层
 *
 * 集中管理所有API路由的注册和配置
 */

package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Router API路由管理器
type Router struct {
	mux              *http.ServeMux
	pendingTaskAPI   *PendingTaskAPI
	orderMappingAPI  *OrderMappingAPI
	middlewares      []Middleware
}

// Middleware 中间件函数类型
type Middleware func(http.HandlerFunc) http.HandlerFunc

// NewRouter 创建API路由管理器
func NewRouter(db *sql.DB) *Router {
	router := &Router{
		mux:             http.NewServeMux(),
		pendingTaskAPI:  NewPendingTaskAPI(db),
		orderMappingAPI: NewOrderMappingAPI(db),
		middlewares:     []Middleware{},
	}

	// 注册默认中间件
	router.Use(LoggingMiddleware)
	router.Use(CORSMiddleware)
	router.Use(RecoveryMiddleware)

	// 注册路由
	router.registerRoutes()

	return router
}

// Use 添加中间件
func (r *Router) Use(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

// registerRoutes 注册所有API路由
func (r *Router) registerRoutes() {
	// ==================== 待确认任务API (B5核心) ====================

	// GET /api/pending-tasks - 查询待确认任务列表
	// 支持查询参数: customer_id, status, mapping_type, page, page_size, order_id, source_value
	r.HandleFunc("/api/pending-tasks", r.pendingTaskAPI.GetPendingTasksHandler)

	// POST /api/pending-tasks/{task_key}/confirm - 确认待确认任务
	// 请求体: {manual_value, confirmed_by, remark}
	r.HandleFunc("/api/pending-tasks/", r.handleConfirmTask)

	// POST /api/orders/process-row - 处理单行订单（确认后重新生成）
	// 请求体: {customer_id, order_id, row_number, confirmed_mappings}
	r.HandleFunc("/api/orders/process-row", r.pendingTaskAPI.ProcessRowHandler)

	// ==================== 订单映射API ====================

	// SKU映射API
	// GET /api/mappings/sku - 查询SKU映射
	// POST /api/mappings/sku - 创建SKU映射
	r.HandleFunc("/api/mappings/sku", r.handleSkuMapping)

	// GET /api/mappings/sku/list - 查询SKU映射列表（支持分页）
	r.HandleFunc("/api/mappings/sku/list", r.orderMappingAPI.ListSkuMappingsHandler)

	// 门店映射API
	// GET /api/mappings/store - 查询门店映射
	// POST /api/mappings/store - 创建门店映射
	r.HandleFunc("/api/mappings/store", r.handleStoreMapping)

	// 仓库映射API
	// GET /api/mappings/warehouse - 查询仓库映射
	r.HandleFunc("/api/mappings/warehouse", r.orderMappingAPI.GetWarehouseMappingHandler)

	// 配送方式映射API
	// GET /api/mappings/delivery-method - 查询配送方式映射
	r.HandleFunc("/api/mappings/delivery-method", r.orderMappingAPI.GetDeliveryMethodMappingsHandler)

	// 业务类型映射API
	// GET /api/mappings/business-type - 查询业务类型映射
	r.HandleFunc("/api/mappings/business-type", r.orderMappingAPI.GetBusinessTypeMappingsHandler)

	// 通用映射匹配API
	// POST /api/mappings/match - 通用映射匹配
	r.HandleFunc("/api/mappings/match", r.orderMappingAPI.MatchMappingHandler)

	// ==================== 健康检查API ====================

	// GET /api/health - 健康检查
	r.HandleFunc("/api/health", r.handleHealthCheck)

	// GET /api/version - 版本信息
	r.HandleFunc("/api/version", r.handleVersion)
}

// HandleFunc 包装路由处理器，应用中间件
func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	// 应用所有中间件
	wrappedHandler := handler
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		wrappedHandler = r.middlewares[i](wrappedHandler)
	}
	r.mux.HandleFunc(pattern, wrappedHandler)
}

// ServeHTTP 实现http.Handler接口
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

// ==================== 路由辅助函数 ====================

// handleConfirmTask 处理确认任务路由
// 路由模式: /api/pending-tasks/{task_key}/confirm
func (r *Router) handleConfirmTask(w http.ResponseWriter, req *http.Request) {
	// 检查路径是否匹配确认任务模式
	if strings.HasSuffix(req.URL.Path, "/confirm") && req.Method == http.MethodPost {
		r.pendingTaskAPI.ConfirmTaskHandler(w, req)
		return
	}

	// 不匹配，返回404
	http.NotFound(w, req)
}

// handleSkuMapping 处理SKU映射路由
// GET /api/mappings/sku - 查询
// POST /api/mappings/sku - 创建
func (r *Router) handleSkuMapping(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.orderMappingAPI.GetSkuMappingHandler(w, req)
	case http.MethodPost:
		r.orderMappingAPI.CreateSkuMappingHandler(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleStoreMapping 处理门店映射路由
// GET /api/mappings/store - 查询
// POST /api/mappings/store - 创建
func (r *Router) handleStoreMapping(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.orderMappingAPI.GetStoreMappingHandler(w, req)
	case http.MethodPost:
		r.orderMappingAPI.CreateStoreMappingHandler(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleHealthCheck 健康检查处理器
func (r *Router) handleHealthCheck(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Service is healthy",
		Data: map[string]interface{}{
			"status": "UP",
		},
	})
}

// handleVersion 版本信息处理器
func (r *Router) handleVersion(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Code:    200,
		Message: "Success",
		Data: map[string]interface{}{
			"version":   "1.0.0",
			"name":      "AI客服赋能系统API",
			"build_time": "2026-05-07",
		},
	})
}

// ==================== 中间件 ====================

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 记录请求信息
		start := time.Now()

		// 包装ResponseWriter以捕获状态码
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next(wrapped, r)

		// 打印日志
		fmt.Printf("[%s] %s %s - %d - %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			wrapped.statusCode,
			time.Since(start),
		)
	}
}

// CORSMiddleware CORS中间件
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置CORS头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// 处理预检请求
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// RecoveryMiddleware 恢复中间件（捕获panic）
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[PANIC] %v\n", err)
				respondWithError(w, http.StatusInternalServerError, "Internal server error")
			}
		}()

		next(w, r)
	}
}

// AuthMiddleware 认证中间件（示例）
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头获取token
		token := r.Header.Get("Authorization")
		if token == "" {
			respondWithError(w, http.StatusUnauthorized, "Authorization required")
			return
		}

		// TODO: 验证token
		// 这里可以调用认证服务验证token的有效性

		next(w, r)
	}
}

// ==================== 辅助类型 ====================

// responseWriter 包装http.ResponseWriter以捕获状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader 捕获状态码
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// ==================== 便捷函数 ====================

// SetupServer 配置并返回HTTP服务器
func SetupServer(db *sql.DB, port string) *http.Server {
	if port == "" {
		port = "8080"
	}

	router := NewRouter(db)

	return &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
}

// SetupServerWithMiddleware 配置带自定义中间件的HTTP服务器
func SetupServerWithMiddleware(db *sql.DB, port string, middlewares []Middleware) *http.Server {
	if port == "" {
		port = "8080"
	}

	router := NewRouter(db)

	// 添加自定义中间件
	for _, mw := range middlewares {
		router.Use(mw)
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
}
