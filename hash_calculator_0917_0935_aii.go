// 代码生成时间: 2025-09-17 09:35:12
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "echo"
    "fmt"
    "net/http"
    "strings"
)

// HashCalculator 结构体用于封装哈希值计算逻辑
type HashCalculator struct{}

// CalculateSHA256 计算给定字符串的SHA256哈希值
func (h *HashCalculator) CalculateSHA256(c echo.Context) error {
    // 从请求中提取字符串参数
    input := c.QueryParam("input")
    if input == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Input string is required")
    }

    // 计算SHA256哈希值
    hash := sha256.Sum256([]byte(input))
    hexHash := hex.EncodeToString(hash[:])

    // 返回哈希值结果
    return c.JSON(http.StatusOK, map[string]string{
        "hash": hexHash,
    })
}

func main() {
    e := echo.New()

    // 设置Echo中间件
    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.Gzip())

    // 创建HashCalculator实例
    calculator := &HashCalculator{}

    // 为Echo框架注册路由
    e.GET("/hash", calculator.CalculateSHA256)

    // 启动Echo服务器
    e.Logger.Fatal(e.Start(":8080"))
}

// middleware 是一个包含Echo框架中间件的包
// 这里假设我们有一个名为middleware的自定义包
// 包含日志记录、错误恢复和Gzip压缩的中间件
package middleware

import (
    "echo"
    "github.com/labstack/echo/middleware"
)

// Logger 中间件用于记录请求日志
func Logger() echo.MiddlewareFunc {
    return middleware.Logger()
}

// Recover 中间件用于捕获和处理panic
func Recover() echo.MiddlewareFunc {
    return middleware.Recover()
}

// Gzip 中间件用于启用Gzip压缩
func Gzip() echo.MiddlewareFunc {
    return middleware.GzipWithConfig(middleware.GzipConfig{Level: 5})
}