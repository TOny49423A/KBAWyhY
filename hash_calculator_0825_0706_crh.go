// 代码生成时间: 2025-08-25 07:06:30
 * This service provides endpoints to calculate the hash value of a given string.
 */

package main

import (
    "crypto/sha256"
    "encoding/hex"
    "net/http"
    "github.com/labstack/echo/v4"
    "log"
)

// HashCalculator defines a struct for the hash calculator service.
type HashCalculator struct{}

// CalculateSHA256 handles HTTP requests to calculate SHA-256 hash of a given string.
func (h *HashCalculator) CalculateSHA256(c echo.Context) error {
    // Get the input string from the query parameter.
    input := c.QueryParam("input")
    if input == "" {
        // Return a bad request error if no input is provided.
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Input parameter 'input' is required."
        })
    }
    
    // Calculate the SHA-256 hash of the input string.
    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])
    
    // Return the calculated hash as a JSON response.
    return c.JSON(http.StatusOK, map[string]string{
        "hash": hashString,
    })
}

func main() {
    // Create a new instance of the Echo framework.
    e := echo.New()
    
    // Define a new hash calculator service.
    hashService := &HashCalculator{}
    
    // Define the route for calculating SHA-256 hash.
    e.GET("/hash/sha256", hashService.CalculateSHA256)

    // Start the Echo server.
    log.Fatal(e.Start(":8080"))
}