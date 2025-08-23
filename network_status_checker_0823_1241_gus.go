// 代码生成时间: 2025-08-23 12:41:49
package main

import (
    "crypto/tls"
    "fmt"
    "net"
    "os"
    "time"

    "github.com/labstack/echo"
)

// NetworkStatusChecker represents a struct to hold the network status checker configurations.
type NetworkStatusChecker struct {
    // Timeout for network check in seconds.
    Timeout int
}

// NewNetworkStatusChecker creates a new instance of NetworkStatusChecker.
func NewNetworkStatusChecker(timeout int) *NetworkStatusChecker {
    return &NetworkStatusChecker{
        Timeout: timeout,
    }
}

// CheckNetworkStatus checks the network status by attempting to connect to a given host.
func (nsc *NetworkStatusChecker) CheckNetworkStatus(host string) (bool, error) {
    // Create a timeout context for network checking.
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(nsc.Timeout)*time.Second)
    defer cancel()

    dialer := net.Dialer{Timeout: time.Duration(nsc.Timeout) * time.Second}
    conn, err := tls.DialWithDialer(&dialer, "tcp", host, &tls.Config{InsecureSkipVerify: true})
    if err != nil {
        return false, err
    }
    defer conn.Close()

    return true, nil
}

// StartServer starts the Echo server with a route to check network status.
func StartServer() {
    e := echo.New()
    checker := NewNetworkStatusChecker(5) // 5 seconds timeout.

    e.GET("/check", func(c echo.Context) error {
        // Extract the host from the query parameter.
        host := c.QueryParam("host")
        if host == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "Host parameter is required")
        }

        status, err := checker.CheckNetworkStatus(host)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error checking network status: %s", err))
        }

        // Return the network status as JSON response.
        return c.JSON(http.StatusOK, map[string]bool{
            "status": status,
        })
    })

    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func main() {
    StartServer()
}