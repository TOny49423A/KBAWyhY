// 代码生成时间: 2025-08-30 14:39:54
// network_status_checker.go

package main

import (
    "net"
    "time"
    "fmt"
    "strings"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// NetworkStatusChecker 结构体用于存储网络状态检查相关的属性
type NetworkStatusChecker struct {
    // 这里可以添加更多的属性，比如超时时间、检查间隔等
}

// NewNetworkStatusChecker 创建并初始化 NetworkStatusChecker 实例
func NewNetworkStatusChecker() *NetworkStatusChecker {
    return &NetworkStatusChecker{}
}

// Check 检查给定的网络地址是否可达
func (ns *NetworkStatusChecker) Check(address string) (bool, error) {
    // 定义网络连接的超时时间
    timeout := time.Duration(3 * time.Second)
    // 创建一个网络连接
    conn, err := net.DialTimeout("tcp", address, timeout)
    if err != nil {
        // 如果连接失败，返回错误
        return false, err
    }
    // 关闭连接
    defer conn.Close()
    // 如果连接成功，返回true
    return true, nil
}

// StartServer 启动 Echo 服务器
func StartServer() {
    // 创建 Echo 实例
    e := echo.New()
    // 使用默认中间件
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // 定义路由和处理函数
    e.GET("/check", func(c echo.Context) error {
        // 从请求中获取地址参数
        address := c.QueryParam("address")
        if address == "" {
            return c.JSON(400, echo.Map{
                "error": "Address parameter is required",
            })
        }

        // 创建 NetworkStatusChecker 实例
        ns := NewNetworkStatusChecker()

        // 检查网络状态
        reachable, err := ns.Check(address)
        if err != nil {
            // 如果检查失败，返回错误信息
            return c.JSON(500, echo.Map{
                "error": err.Error(),
            })
        }

        // 根据检查结果返回不同的响应
        if reachable {
            return c.JSON(200, echo.Map{
                "status": "reachable",
            })
        } else {
            return c.JSON(200, echo.Map{
                "status": "not reachable",
            })
        }
    })

    // 启动服务器
    e.Logger.Fatal(e.Start(":8080"))
}

func main() {
    // 启动服务器
    StartServer()
}
