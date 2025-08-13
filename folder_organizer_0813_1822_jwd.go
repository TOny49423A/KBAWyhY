// 代码生成时间: 2025-08-13 18:22:43
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/labstack/echo/v4"
)

// FolderOrganizer 结构体用于存储文件夹和文件的元数据
type FolderOrganizer struct {
    Path string
}

// NewFolderOrganizer 创建一个新的 FolderOrganizer 实例
func NewFolderOrganizer(path string) *FolderOrganizer {
    return &FolderOrganizer{Path: path}
}

// Organize 递归地组织文件夹中的文件，将文件按类型排序
func (fo *FolderOrganizer) Organize() error {
    // 获取文件夹中的所有文件和子文件夹
    files, err := os.ReadDir(fo.Path)
    if err != nil {
        return err
    }

    for _, file := range files {
        path := filepath.Join(fo.Path, file.Name())

        // 检查是否为文件夹
        if file.IsDir() {
            // 递归调用 Organize
            if err := fo.Organize(); err != nil {
                return err
            }
        } else {
            // 将文件按类型排序到不同的子文件夹
            ext := strings.TrimLeft(filepath.Ext(file.Name()), ".")
            destPath := filepath.Join(fo.Path, ext, file.Name())

            // 创建目标文件夹如果不存在
            if _, err := os.Stat(filepath.Dir(destPath)); os.IsNotExist(err) {
                if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
                    return err
                }
            }

            // 移动文件
            if err := os.Rename(path, destPath); err != nil {
                return err
            }
        }
    }

    return nil
}

// StartEchoServer 启动 ECHO 服务器
func StartEchoServer() *echo.Echo {
    e := echo.New()

    // 定义 API 路由
    e.GET("/organize", func(c echo.Context) error {
        // 示例路径，实际应用中应从请求参数获取
        path := "/path/to/your/folder"
        organizer := NewFolderOrganizer(path)
        if err := organizer.Organize(); err != nil {
            return c.JSON(500, echo.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(200, echo.Map{
            "message": "Folder organized successfully",
        })
    })

    return e
}

func main() {
    // 启动 ECHO 服务器
    e := StartEchoServer()
    e.Logger.Fatal(e.Start(":8080"))
}
