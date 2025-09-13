// 代码生成时间: 2025-09-14 06:24:22
package main

import (
    "testing"
    "github.com/labstack/echo"
    "net/http"
    "net/http/httptest"
)

// setupEchoServer 初始话Echo服务器
func setupEchoServer() *echo.Echo {
    e := echo.New()
    e.GET("/test", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    return e
}

// TestEchoServer 测试Echo服务器的GET请求
func TestEchoServer(t *testing.T) {
    e := setupEchoServer()
    rec := httptest.NewRecorder()
    req, err := http.NewRequest(http.MethodGet, "/test", nil)
    if err != nil {
        t.Fatal("Error creating request: \{err}")
    }
    e.ServeHTTP(rec, req)
    if rec.Code != http.StatusOK {
        t.Errorf("Expected status %v, got %v", http.StatusOK, rec.Code)
    }
    if rec.Body.String() != "Hello, World!" {
        t.Errorf("Expected body 'Hello, World!', got '%s'", rec.Body.String())
    }
}
