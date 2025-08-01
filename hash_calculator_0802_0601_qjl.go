// 代码生成时间: 2025-08-02 06:01:45
package main

import (
    "crypto/sha1"
    "encoding/hex"
    "net/http"
    "github.com/labstack/echo"
    "strings"
)

// HashCalculator contains the necessary methods to calculate hash values.
type HashCalculator struct{}

// CalculateSHA1 takes a string input and returns its SHA1 hash.
func (h *HashCalculator) CalculateSHA1(input string) (string, error) {
    if input == "" {
        return "", echo.NewHTTPError(http.StatusBadRequest, "Input cannot be empty")
    }

    sha1Hash := sha1.New()
    sha1Hash.Write([]byte(input))
    return hex.EncodeToString(sha1Hash.Sum(nil)), nil
}

// HashHandler is the HTTP handler for the hash calculator.
func HashHandler(hashCalculator *HashCalculator) echo.HandlerFunc {
    return func(c echo.Context) error {
        input := c.QueryParam("input")
        if input == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "Input parameter is required")
        }

        hash, err := hashCalculator.CalculateSHA1(input)
        if err != nil {
            return err
        }

        return c.JSON(http.StatusOK, map[string]string{
            "input": input,
            "hash": hash,
        })
    }
}

func main() {
    e := echo.New()
    hashCalculator := &HashCalculator{}

    // Define the route for the hash calculator.
    e.GET("/hash", HashHandler(hashCalculator))

    // Start the Echo server.
    e.Logger.Fatal(e.Start(":8080"))
}
