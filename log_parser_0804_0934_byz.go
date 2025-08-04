// 代码生成时间: 2025-08-04 09:34:26
@author Your Name
@date 2023-10-01
*/

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
    "github.com/labstack/echo"
)

// LogEntry represents a single log entry with its details
type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
}

// parseLogLine takes a line from the log file and attempts to parse it into a LogEntry
func parseLogLine(line string) (*LogEntry, error) {
    parts := strings.Split(line, " ")
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid log line format: %s", line)
    }
    timestamp, err := time.Parse("2006-01-02 15:04:05", parts[0]+" "+parts[1])
    if err != nil {
        return nil, fmt.Errorf("failed to parse timestamp: %s", err)
    }
    level := parts[2]
    message := strings.Join(parts[3:], " ")
    return &LogEntry{
        Timestamp: timestamp,
        Level:     level,
        Message:   message,
    }, nil
}

// parseLogFile reads the log file and parses each line
func parseLogFile(filePath string) ([]LogEntry, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %s", err)
    }
    defer file.Close()

    var entries []LogEntry
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        entry, err := parseLogLine(line)
        if err != nil {
            log.Printf("skipping invalid log entry: %s", err)
            continue
        }
        entries = append(entries, *entry)
    }
    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("failed to read file: %s", err)
    }
    return entries, nil
}

// setupRoutes sets up the Echo routes for the log parser tool
func setupRoutes(e *echo.Echo) {
    e.GET("/parse", func(c echo.Context) error {
        filePath := c.QueryParam("file")
        if filePath == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "log file path is required")
        }
        entries, err := parseLogFile(filePath)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }
        return c.JSON(http.StatusOK, entries)
    })
}

func main() {
    e := echo.New()
    setupRoutes(e)
    e.Logger.Fatal(e.Start(":8080"))
}
