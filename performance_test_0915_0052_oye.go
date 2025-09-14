// 代码生成时间: 2025-09-15 00:52:37
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/labstack/echo/v4"
)

// RunTest runs a performance test on the provided URL
func RunTest(url string, requests int, duration time.Duration) error {
    start := time.Now()
    defer func() {
        fmt.Printf("Test completed in %v
", time.Since(start))
    }()

    for i := 0; i < requests; i++ {
        if time.Since(start) > duration {
            break
        }
        _, err := http.Get(url)
        if err != nil {
            return fmt.Errorf("error during request: %w", err)
        }
    }
    return nil
}

func main() {
    e := echo.New()
    defer e.Close()

    // Define the performance test endpoint
    e.GET("/test", func(c echo.Context) error {
        url := c.QueryParam("url")
        requests := c.QueryParam("requests")
        durationStr := c.QueryParam("duration")
        duration, err := time.ParseDuration(durationStr)
        if err != nil {
            return err
        }

        // Parse the requests parameter
        numRequests, err := strconv.Atoi(requests)
        if err != nil {
            return err
        }

        // Run the performance test
        if err := RunTest(url, numRequests, duration); err != nil {
            return err
        }

        return c.JSON(http.StatusOK, map[string]string{
            "message": "Performance test completed successfully"
        })
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}
