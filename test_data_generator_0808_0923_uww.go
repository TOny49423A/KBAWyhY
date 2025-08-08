// 代码生成时间: 2025-08-08 09:23:41
 * It provides a RESTful API to generate and return test data.
 */

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "time"

    "github.com/labstack/echo"
)

// testData represents the structure of the test data.
type testData struct {
    ID       uint   "json:\u0069d"
    Name     string "json:\"name\""
    Email    string "json:\"email\""
    Age      int    "json:\"age\""
    Birthday string "json:\"birthday\"" // in the format YYYY-MM-DD
}

// generateTestData generates random test data.
func generateTestData() testData {
    rand.Seed(time.Now().UnixNano())
    return testData{
        ID:      (uint)(rand.Intn(10000)),
        Name:    fmt.Sprintf("Name%d", rand.Intn(100)),
        Email:   fmt.Sprintf("%d@example.com", rand.Intn(100)),
        Age:     (rand.Intn(50) + 18), // age between 18 and 68
        Birthday: time.Now().AddDate(-68, 0, 0).Format("2006-01-02"),
    }
}

// getTestData handles the HTTP GET request to generate and return test data.
func getTestData(c echo.Context) error {
    data := generateTestData()
    return c.JSON(http.StatusOK, data)
}

func main() {
    e := echo.New()

    // Define a route for the test data endpoint.
    e.GET("/test-data", getTestData)

    // Start the Echo server.
    log.Fatal(e.Start(":8080"))
}