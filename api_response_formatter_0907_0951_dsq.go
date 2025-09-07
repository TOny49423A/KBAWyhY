// 代码生成时间: 2025-09-07 09:51:06
package main

import (
    "echo"
    "net/http"
    "time"
)

// ApiResponse defines the structure of a formatted API response.
type ApiResponse struct {
    Success bool        `json:"success"`
    Message string     `json:"message"`
    Data    interface{} `json:"data"`
    Time    string     `json:"time"`
}

// ErrorResponse defines the structure of an error response.
type ErrorResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Time    string `json:"time"`
}

// NewErrorResponse creates a new ErrorResponse with a formatted message and current timestamp.
func NewErrorResponse(message string) ErrorResponse {
    return ErrorResponse{
        Success: false,
        Message: message,
        Time:    time.Now().Format(time.RFC3339),
    }
}

// NewApiResponse creates a new ApiResponse with a success flag, message, data, and current timestamp.
func NewApiResponse(message string, data interface{}) ApiResponse {
    return ApiResponse{
        Success: true,
        Message: message,
        Data:    data,
        Time:    time.Now().Format(time.RFC3339),
    }
}

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        // Example data to be sent in the response
        data := map[string]string{
            "example": "data",
        }
        // Create a new API response
        response := NewApiResponse("Success", data)
        return c.JSON(http.StatusOK, response)
    })

    e.GET("/error", func(c echo.Context) error {
        // Create a new error response
        errorResponse := NewErrorResponse("Something went wrong")
        return c.JSON(http.StatusInternalServerError, errorResponse)
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}
