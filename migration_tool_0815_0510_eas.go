// 代码生成时间: 2025-08-15 05:10:54
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/lib/pq" // PostgreSQL driver
    "github.com/labstack/echo/v4"
)

const (
    migrationDir = "./migrations"
)

// Main function
func main() {
    setupMigration()
}

// setupMigration sets up the migration tool
func setupMigration() {
    // Create a new Echo instance
    e := echo.New()

    // Define a route for applying migrations
    e.GET("/migrate", applyMigrations)

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}

// applyMigrations applies the database migrations
func applyMigrations(c echo.Context) error {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        return &echo.HTTPError{
            Code:     500,
            Message:  "DATABASE_URL environment variable is not set",
        }
    }

    // Initialize the migrator
    driver, err := postgres.WithInstance(dbURL, &sql.DB{})
    if err != nil {
        return &echo.HTTPError{
            Code:     500,
            Message:  err.Error(),
        }
    }
    defer driver.Close()

    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s", migrationDir),
        "postgres", driver,
    )
    if err != nil {
        return &echo.HTTPError{
            Code:     500,
            Message:  err.Error(),
        }
    }

    // Migrate up to the latest version
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return &echo.HTTPError{
            Code:     500,
            Message:  err.Error(),
        }
    }

    return c.JSON(200, echo.Map{
        "message": "Migrations applied successfully",
    })
}

// readMigrationFiles reads migration files from the directory
func readMigrationFiles() ([]string, error) {
    // Read all files from the migration directory
    files, err := os.ReadDir(migrationDir)
    if err != nil {
        return nil, err
    }

    var migrations []string
    for _, f := range files {
        if f.IsDir() {
            continue
        }
        // Ensure the file has the correct extension
        if strings.HasSuffix(f.Name(), ".sql") {
            filePath := filepath.Join(migrationDir, f.Name())
            migrations = append(migrations, filePath)
        }
    }
    return migrations, nil
}
