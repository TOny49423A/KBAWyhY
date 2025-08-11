// 代码生成时间: 2025-08-12 06:20:37
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

// RenameFunc 是一个函数类型，用于定义重命名操作
type RenameFunc func(oldName, newName string) error

// BulkRenamer 结构体包含重命名操作的函数和文件路径
type BulkRenamer struct {
    renFunc RenameFunc
    baseDir string
}

// NewBulkRenamer 创建并返回一个新的BulkRenamer实例
func NewBulkRenamer(baseDir string) *BulkRenamer {
    return &BulkRenamer{
        renFunc: os.Rename,
        baseDir: baseDir,
    }
}

// RenameAll 重命名指定目录下所有匹配旧模式的文件到新模式
func (br *BulkRenamer) RenameAll(oldPattern, newPattern string) error {
    // 遍历目录
    err := filepath.WalkDir(br.baseDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }
        // 检查文件名是否匹配旧模式
        if strings.Contains(filepath.Base(path), oldPattern) {
            // 构建新文件名
            newBase := strings.ReplaceAll(filepath.Base(path), oldPattern, newPattern)
            newPath := filepath.Join(filepath.Dir(path), newBase)
            // 执行重命名操作
            if err := br.renFunc(path, newPath); err != nil {
                return err
            }
            fmt.Printf("Renamed '%s' to '%s'
", path, newPath)
        }
        return nil
    })
    return err
}

func main() {
    baseDir := "./" // 指定目录
    oldPattern := "old" // 旧文件名模式
    newPattern := "new" // 新文件名模式

    // 创建BulkRenamer实例
    renamer := NewBulkRenamer(baseDir)

    // 执行批量重命名
    if err := renamer.RenameAll(oldPattern, newPattern); err != nil {
        fmt.Printf("Error occurred: %s
", err)
    }
}
