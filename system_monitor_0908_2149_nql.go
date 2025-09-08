// 代码生成时间: 2025-09-08 21:49:48
package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "strings"
    "time"

    "github.com/labstack/echo/v4"
)

// SystemInfo contains system performance data.
type SystemInfo struct {
    CpuUsage float64 `json:"cpu_usage"`
    MemoryUsage float64 `json:"memory_usage"`
    DiskUsage float64 `json:"disk_usage"`
}

// GetSystemInfo retrieves system performance data.
func GetSystemInfo() (SystemInfo, error) {
    var sysInfo SystemInfo
    var err error

    // Get CPU usage
    sysInfo.CpuUsage, err = getCPUUsage()
    if err != nil {
        return sysInfo, err
    }

    // Get Memory usage
    sysInfo.MemoryUsage, err = getMemoryUsage()
    if err != nil {
        return sysInfo, err
    }

    // Get Disk usage
    sysInfo.DiskUsage, err = getDiskUsage()
    if err != nil {
        return sysInfo, err
    }

    return sysInfo, nil
}

// getCPUUsage retrieves the CPU usage.
func getCPUUsage() (float64, error) {
    // Command to get CPU usage
    cmd := exec.Command("top", "-b", "-n1")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return 0, err
    }

    // Parse CPU usage from the output
    lines := strings.Split(string(output), "
")
    for _, line := range lines {
        if strings.Contains(line, "Cpu(s)") {
            fields := strings.Fields(line)
            usage, err := strconv.ParseFloat(fields[len(fields)-3], 64)
            if err != nil {
                return 0, err
            }
            return usage, nil
        }
    }
    return 0, fmt.Errorf("could not parse CPU usage")
}

// getMemoryUsage retrieves the memory usage.
func getMemoryUsage() (float64, error) {
    // Command to get memory usage
    cmd := exec.Command("free", "-m")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return 0, err
    }

    // Parse memory usage from the output
    lines := strings.Split(string(output), "
")
    for i, line := range lines {
        if i == 1 {
            fields := strings.Fields(line)
            usage, err := strconv.ParseFloat(fields[2], 64)
            if err != nil {
                return 0, err
            }
            return usage, nil
        }
    }
    return 0, fmt.Errorf("could not parse memory usage")
}

// getDiskUsage retrieves the disk usage.
func getDiskUsage() (float64, error) {
    // Command to get disk usage
    cmd := exec.Command("df", "-h")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return 0, err
    }

    // Parse disk usage from the output
    lines := strings.Split(string(output), "
")
    for _, line := range lines {
        if !strings.HasPrefix(line, "Filesystem") && !strings.HasPrefix(line, "tmpfs") {
            fields := strings.Fields(line)
            usage, err := strconv.ParseFloat(fields[3], 64)
            if err != nil {
                return 0, err
            }
            return usage, nil
        }
    }
    return 0, fmt.Errorf("could not parse disk usage")
}

func main() {
    e := echo.New()
    e.GET("/monitor", func(c echo.Context) error {
        sysInfo, err := GetSystemInfo()
        if err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(http.StatusOK, sysInfo)
    })

    e.Logger.Fatal(e.Start(":8080"))
}