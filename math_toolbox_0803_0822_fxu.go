// 代码生成时间: 2025-08-03 08:22:26
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "strconv"
    "fmt"
)

// MathToolbox is a structure that represents the math toolbox.
type MathToolbox struct{}

// Add handles the addition operation.
func (mt *MathToolbox) Add(c echo.Context) error {
    // Extract parameters from the query string.
    param1, err := strconv.ParseFloat(c.QueryParam("a"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for a")
    }
    param2, err := strconv.ParseFloat(c.QueryParam("b"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for b")
    }
    
    // Perform addition and return the result.
    result := param1 + param2
    return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
}

// Subtract handles the subtraction operation.
func (mt *MathToolbox) Subtract(c echo.Context) error {
    param1, err := strconv.ParseFloat(c.QueryParam("a"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for a")
    }
    param2, err := strconv.ParseFloat(c.QueryParam("b"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for b")
    }
    
    result := param1 - param2
    return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
}

// Multiply handles the multiplication operation.
func (mt *MathToolbox) Multiply(c echo.Context) error {
    param1, err := strconv.ParseFloat(c.QueryParam("a"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for a")
    }
    param2, err := strconv.ParseFloat(c.QueryParam("b"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for b")
    }
    
    result := param1 * param2
    return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
}

// Divide handles the division operation.
func (mt *MathToolbox) Divide(c echo.Context) error {
    param1, err := strconv.ParseFloat(c.QueryParam("a"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for a")
    }
    param2, err := strconv.ParseFloat(c.QueryParam("b"), 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input for b")
    }
    
    if param2 == 0 {
        return c.JSON(http.StatusBadRequest, "Cannot divide by zero")
    }
    
    result := param1 / param2
    return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
}

func main() {
    // Initialize Echo instance.
    e := echo.New()
    
    // Define routes.
    toolbox := MathToolbox{}
    e.GET("/add", toolbox.Add)
    e.GET("/subtract", toolbox.Subtract)
    e.GET("/multiply", toolbox.Multiply)
    e.GET("/divide", toolbox.Divide)
    
    // Start the server.
    e.Logger.Fatal(e.Start(":8080"))
}