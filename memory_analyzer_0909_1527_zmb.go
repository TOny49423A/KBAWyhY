// 代码生成时间: 2025-09-09 15:27:48
package main

import (
    "fmt"
    "log"
    "net/http"
    "runtime"
    "strings"

    "github.com/labstack/echo"
)

// MemoryStats represents the memory statistics
type MemoryStats struct {
    Alloc      uint64 `json:"alloc"`           // bytes allocated and not yet freed
    TotAlloc  uint64 `json:"totalAlloc"`       // bytes allocated (even if freed)
    Sys       uint64 `json:"sys"`             // bytes obtained from system (sum of XxxSys below)
    Lookups  uint64 `json:"lookups"`         // number of pointer lookups
    Mallocs   uint64 `json:"mallocs"`         // number of mallocs
    Frees     uint64 `json:"frees"`           // number of frees
    HeapAlloc uint64 `json:"heapAlloc"`       // bytes allocated in heap
    HeapSys   uint64 `json:"heapSys"`         // heap system bytes
    HeapIdle  uint64 `json:"heapIdle"`        // heap idle bytes
    HeapInuse uint64 `json:"heapInuse"`       // heap in-use bytes
    HeapReleased uint64 `json:"heapReleased"` // heap released bytes
    HeapObjects uint64 `json:"heapObjects"`    // number of allocated heap objects
   栈相关的统计信息可在需要时添加
}

// GetMemoryStats retrieves and returns the current memory statistics
func GetMemoryStats() (MemoryStats, error) {
    var stats runtime.MemStats
    runtime.ReadMemStats(&stats)

    return MemoryStats{
        Alloc:      stats.Alloc,
        TotAlloc:  stats.TotalAlloc,
        Sys:       stats.Sys,
        Lookups:   stats.Lookups,
        Mallocs:   stats.Mallocs,
        Frees:     stats.Frees,
        HeapAlloc: stats.HeapAlloc,
        HeapSys:   stats.HeapSys,
        HeapIdle:  stats.HeapIdle,
        HeapInuse: stats.HeapInuse,
        HeapReleased: stats.HeapReleased,
        HeapObjects: stats.HeapObjects,
    }, nil
}

// NewMemoryAnalyzer creates a new Echo HTTP server with a route for memory stats
func NewMemoryAnalyzer() *echo.Echo {
    e := echo.New()
    e.GET("/memory", func(c echo.Context) error {
        stats, err := GetMemoryStats()
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, stats)
    })
    return e
}

func main() {
    e := NewMemoryAnalyzer()
    fmt.Println("Starting memory analyzer on :8080")
    if err := e.Start(":8080"); err != nil {
        log.Fatal(err)
    }
}