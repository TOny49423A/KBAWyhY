// 代码生成时间: 2025-08-12 14:09:14
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

// PermissionManager is a struct that will handle user permissions.
type PermissionManager struct {
    // Permissions is a map where keys are user roles and values are permissions.
    Permissions map[string][]string
}

// NewPermissionManager creates a new instance of PermissionManager.
func NewPermissionManager() *PermissionManager {
    return &PermissionManager{
        Permissions: make(map[string][]string),
    }
}

// AddPermission adds a permission to a user role.
func (pm *PermissionManager) AddPermission(role string, permission string) {
    pm.Permissions[role] = append(pm.Permissions[role], permission)
}

// CheckPermission checks if a user has a specific permission.
func (pm *PermissionManager) CheckPermission(role string, permission string) bool {
    for _, perm := range pm.Permissions[role] {
        if perm == permission {
            return true
        }
    }
    return false
}

// Routes sets up the routes for the permission manager.
func (pm *PermissionManager) Routes(e *echo.Echo) {
    e.GET("/checkPermission", func(c echo.Context) error {
        role := c.QueryParam("role")
        permission := c.QueryParam("permission\)

        if hasPermission := pm.CheckPermission(role, permission); hasPermission {
            return c.JSON(http.StatusOK, map[string]bool{
                "hasPermission": true,
            })
        } else {
            return c.JSON(http.StatusForbidden, map[string]bool{
                "hasPermission": false,
            })
        }
    })
}

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    pm := NewPermissionManager()
    pm.AddPermission("admin", "create")
    pm.AddPermission("admin", "delete")
    pm.AddPermission("user", "read")

    pm.Routes(e)

    e.Logger.Fatal(e.Start(":8080\)