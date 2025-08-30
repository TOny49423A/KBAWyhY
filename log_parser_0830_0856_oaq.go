// 代码生成时间: 2025-08-30 08:56:22
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/labstack/echo"
)

// LogRecord represents a single log record with a timestamp and message.
type LogRecord struct {
    Timestamp string `json:"timestamp"`
    Message   string `json:"message"`
}

// ParseLogFile opens the specified log file and parses its content.
func ParseLogFile(filePath string) ([]LogRecord, error) {
    fileContent, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }

    var logRecords []LogRecord
    lines := strings.SplitAfter(string(fileContent), "
")
    for _, line := range lines {
        if line == "" {
            continue
        }
        // Assuming the log format is "<timestamp> <message>".
        parts := strings.Fields(line)
        if len(parts) < 2 {
            continue // Skip lines that do not contain a timestamp and message.
        }

        logRecord := LogRecord{
            Timestamp: parts[0],
            Message:   strings.Join(parts[1:], " "),
        }
        logRecords = append(logRecords, logRecord)
    }
    return logRecords, nil
}

// LogParserHandler is an Echo HTTP handler that parses a log file and returns its content in JSON format.
func LogParserHandler(c echo.Context) error {
    filePath := c.QueryParam("file")
    if filePath == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "file parameter is required")
    }

    logRecords, err := ParseLogFile(filePath)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error parsing log file: %s", err))
    }

    return c.JSON(http.StatusOK, logRecords)
}

func main() {
    e := echo.New()
    e.GET("/log/parse", LogParserHandler)

    // Start the Echo server on port 8080.
    if err := e.Start(":8080"); err != nil {
        log.Fatal(err)
    }
}
