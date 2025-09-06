// 代码生成时间: 2025-09-06 13:02:10
main package is the entry point for the RESTful API server.
*/
package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

// Define a User struct to represent a user entity.
type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Age      int    `json:"age"`
}

// Initialize the Echo router and middleware.
func setupRouter() *echo.Echo {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.GET("/users", getUser)
    e.POST("/users", createUser)

    return e
}

// getUser handles GET requests for users.
func getUser(c echo.Context) error {
    // Example response
    return c.JSON(http.StatusOK, []User{
        {ID: "1", Name: "John Doe", Username: "john", Email: "john@example.com", Age: 30},
        {ID: "2", Name: "Jane Doe", Username: "jane", Email: "jane@example.com", Age: 25},
    })
}

// createUser handles POST requests for creating a new user.
func createUser(c echo.Context) error {
    // Parse the request body into a User struct.
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    // Here you would typically validate the user data and save it to a database.
    // For simplicity, we're just returning the user data as a response.
    return c.JSON(http.StatusCreated, u)
}

// Main function to start the server.
func main() {
    e := setupRouter()
    
    // Print the startup message.
    fmt.Println("Starting the API server...")
    
    // Start the server.
    if err := e.Start(":8080"); err != nil {
        e.Logger.Info("Shutting down the server...")
        e.Logger.Error(err)
    }
}