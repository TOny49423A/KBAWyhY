// 代码生成时间: 2025-08-03 04:05:49
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

// AccessControlMiddleware is a middleware function that checks user access rights.
func AccessControlMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Check for user role in the request context
        role, exists := c.Get("role").(string)
        if !exists || role != "admin" {
            return echo.NewHTTPError(http.StatusForbidden, "Access denied")
        }
        return next(c)
    }
}

func main() {
    e := echo.New()
    
    // Define the route with access control middleware
    e.GET("/admin", func(c echo.Context) error {
        return c.String(http.StatusOK, "Welcome admin!")
    }, AccessControlMiddleware)

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
