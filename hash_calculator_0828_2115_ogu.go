// 代码生成时间: 2025-08-28 21:15:24
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "io"
    "log"
    "net/http"
    "github.com/labstack/echo"
)

// HashCalculator is a struct that holds the Echo instance.
type HashCalculator struct{}

// CalculateHash is a method that calculates the SHA-256 hash of a given string.
func (h *HashCalculator) CalculateHash(c echo.Context) error {
    // Get the input string from the query parameter.
    input := c.QueryParam("input")

    // Check if the input is empty.
    if input == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Input string is required")
    }

    // Calculate the hash.
    hash := sha256.Sum256([]byte(input))

    // Return the hash as a hexadecimal string.
    return c.JSON(http.StatusOK, map[string]string{
        "hash": hex.EncodeToString(hash[:]),
    })
}

func main() {
    e := echo.New()
    defer e.Close()

    // Create an instance of HashCalculator.
    hc := &HashCalculator{}

    // Define the route for calculating the hash.
    e.GET("/hash", hc.CalculateHash)

    // Start the Echo server.
    log.Printf("Hash calculator server is running on :8080")
    if err := e.Start(":8080"); err != nil {
        log.Fatal(err)
    }
}
