package httpserver

import (
	"github.com/donnie4w/tlnet"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

var target = "/api/v1/users/123"

// 模拟文本响应处理
func textResponseHandler(c *gin.Context) {
	c.String(200, "Hello, Gin!")
}

// 测试Gin的复杂路由匹配
func setupGinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/api/v1/status", textResponseHandler)
	r.GET("/api/v1/users/:id", textResponseHandler)
	r.GET("/api/v1/users/:id/orders", textResponseHandler)
	r.GET("/api/v1/products", textResponseHandler)
	r.GET("/api/v1/products/:id", textResponseHandler)
	return r
}

// 基准测试Gin
func BenchmarkGin(b *testing.B) {
	router := setupGinRouter()
	req := httptest.NewRequest("GET", target, nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

// 基准测试Gin
func TestGin(t *testing.T) {
	router := setupGinRouter()
	req := httptest.NewRequest("GET", "/api/v1/users/123", nil)
	w := httptest.NewRecorder()
	for i := 0; i < 2; i++ {
		router.ServeHTTP(w, req)
		t.Logf("Response %d: %s", i+1, w.Body.String())
	}
}

// 模拟文本响应处理
func nativeTextResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, net/http!"))
}

// 设置原生HTTP的复杂路由匹配
func setupNativeHttpRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/status", nativeTextResponseHandler)
	mux.HandleFunc("/api/v1/users/", nativeTextResponseHandler) // 模拟路径参数处理
	mux.HandleFunc("/api/v1/users", nativeTextResponseHandler)
	mux.HandleFunc("/api/v1/products", nativeTextResponseHandler)
	mux.HandleFunc("/api/v1/products/", nativeTextResponseHandler)
	return mux
}

// 基准测试原生HTTP
func BenchmarkHttp(b *testing.B) {
	router := setupNativeHttpRouter()
	req := httptest.NewRequest("GET", target, nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

// 模拟文本响应处理
func tlnetTextResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, tlnet!"))
}

// 设置TLnet的复杂路由匹配
func setupTlnetRouter() tlnet.TlnetMux {
	mux := tlnet.NewTlnetMux()
	mux.HandleFunc("/api/v1/status", tlnetTextResponseHandler)
	mux.HandleFunc("/api/v1/users/", tlnetTextResponseHandler) // 模拟路径参数处理
	mux.HandleFunc("/api/v1/users", tlnetTextResponseHandler)
	mux.HandleFunc("/api/v1/products", tlnetTextResponseHandler)
	mux.HandleFunc("/api/v1/products/", tlnetTextResponseHandler)
	return mux
}

// 基准测试Tlnet
func BenchmarkTlnet(b *testing.B) {
	router := setupTlnetRouter()
	req := httptest.NewRequest("GET", target, nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

// 基准测试Tlnet
func TestTlnet(t *testing.T) {
	router := setupTlnetRouter()
	req := httptest.NewRequest("GET", "/api/v1/users/123", nil)
	w := httptest.NewRecorder()
	for i := 0; i < 2; i++ {
		router.ServeHTTP(w, req)
		t.Logf("Response %d: %s", i+1, w.Body.String())
	}
}

// 模拟文本响应处理
func textEchoResponseHandler(c echo.Context) error {
	return c.String(200, "Hello, Echo!")
}

// 测试Echo的复杂路由匹配
func setupEchoRouter() *echo.Echo {
	e := echo.New()
	e.GET("/api/v1/status", textEchoResponseHandler)
	e.GET("/api/v1/users/:id", textEchoResponseHandler)
	e.GET("/api/v1/users/:id/orders", textEchoResponseHandler)
	e.GET("/api/v1/products", textEchoResponseHandler)
	e.GET("/api/v1/products/:id", textEchoResponseHandler)
	return e
}

// 基准测试Echo
func BenchmarkEcho(b *testing.B) {
	router := setupEchoRouter()
	req := httptest.NewRequest("GET", target, nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}
