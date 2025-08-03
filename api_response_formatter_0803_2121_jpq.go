// 代码生成时间: 2025-08-03 21:21:54
package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo"
)

// ErrorResponse is a struct to define the error response format
type ErrorResponse struct {
    Error string `json:"error"`
}

// SuccessResponse is a struct to define the success response format
type SuccessResponse struct {
    Message string `json:"message"`
    Data    interface{} `json:"data"`
}

// ResponseFormatter is a function to format API responses
func ResponseFormatter(c echo.Context, message string, data interface{}) error {
    return c.JSON(http.StatusOK, SuccessResponse{Message: message, Data: data})
}

// ErrorResponseFormatter is a function to format API error responses
func ErrorResponseFormatter(c echo.Context, message string, code int) error {
    return c.JSON(code, ErrorResponse{Error: message})
}

func main() {
    // Initialize Echo instance
    e := echo.New()

    // Define route for a sample API
    e.GET("/api/sample", func(c echo.Context) error {
        // Simulate a response
        return ResponseFormatter(c, "Sample API called successfully", map[string]string{"key": "value"})
    })

    // Define route for a sample error API
    e.GET("/api/error", func(c echo.Context) error {
        // Simulate an error response
        return ErrorResponseFormatter(c, "An error occurred", http.StatusInternalServerError)
    })

    // Start the Echo server
    e.Start(":8080")
}
