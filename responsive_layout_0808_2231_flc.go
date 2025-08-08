// 代码生成时间: 2025-08-08 22:31:01
package main

import (
    "echo"
    "github.com/labstack/echo-contrib/static"
    "io"
    "net/http"
)

// define the root directory for static files
const staticDir = "public"

func main() {
    // Initialize an Echo instance
    e := echo.New()

    // Use the Static middleware to serve static files from the public directory
    e.Use(static.Static(staticDir))

    // Define the route for the home page
    e.GET("/", homeHandler)

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}

// homeHandler serves the home page with a responsive layout
func homeHandler(c echo.Context) error {
    // Open the index.html file from the public directory
    f, err := staticDir + "/index.html")
    if err != nil {
        // Return a 404 error if the file is not found
        return c.File(http.StatusNotFound, "404.html")
    }
    defer f.Close()

    // Read the contents of the index.html file
    contents, err := io.ReadAll(f)
    if err != nil {
        // Return a 500 error if there is an error reading the file
        return c.String(http.StatusInternalServerError, "Internal Server Error")
    }

    // Return the contents of the index.html file as the response
    return c.HTMLBlob(http.StatusOK, contents)
}
