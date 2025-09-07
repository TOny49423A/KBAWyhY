// 代码生成时间: 2025-09-07 17:38:13
package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/labstack/echo/v4"
)

// FileSync is a structure that holds information about source and destination directories
type FileSync struct {
    SrcDir  string
    DstDir  string
    Verbose bool
}

// NewFileSync creates a new FileSync instance
func NewFileSync(src, dst string, verbose bool) *FileSync {
    return &FileSync{SrcDir: src, DstDir: dst, Verbose: verbose}
}

// Sync syncs the files from source to destination
func (fs *FileSync) Sync() error {
    // Get all files from source directory
    files, err := ioutil.ReadDir(fs.SrcDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        srcFilePath := filepath.Join(fs.SrcDir, file.Name())
        dstFilePath := filepath.Join(fs.DstDir, file.Name())

        // Skip directories
        if file.IsDir() {
            continue
        }

        // Check if the file exists in the destination
        if _, err := os.Stat(dstFilePath); os.IsNotExist(err) {
            // File does not exist in destination, copy it
            if fs.Verbose {
                fmt.Printf("Copying file: %s to %s
", srcFilePath, dstFilePath)
            }
            if err := fs.copyFile(srcFilePath, dstFilePath); err != nil {
                return fmt.Errorf("failed to copy file: %w", err)
            }
        } else {
            // File exists, check if it's different
            if different, err := fs.isFileDifferent(srcFilePath, dstFilePath); err != nil {
                return fmt.Errorf("failed to check if file is different: %w", err)
            } else if different {
                // File is different, copy it
                if fs.Verbose {
                    fmt.Printf("Updating file: %s to %s
", srcFilePath, dstFilePath)
                }
                if err := fs.copyFile(srcFilePath, dstFilePath); err != nil {
                    return fmt.Errorf("failed to copy file: %w", err)
                }
            }
        }
    }
    return nil
}

// copyFile copies a file from source to destination
func (fs *FileSync) copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("failed to open source file: %w", err)
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer destinationFile.Close()

    _, err = io.Copy(destinationFile, sourceFile)
    return err
}

// isFileDifferent checks if a file is different from the one in the destination
func (fs *FileSync) isFileDifferent(src, dst string) (bool, error) {
    srcHash, err := fs.getFileHash(src)
    if err != nil {
        return false, fmt.Errorf("failed to get hash of source file: %w", err)
    }
    dstHash, err := fs.getFileHash(dst)
    if err != nil {
        return false, fmt.Errorf("failed to get hash of destination file: %w", err)
    }
    return srcHash != dstHash, nil
}

// getFileHash gets the MD5 hash of a file
func (fs *FileSync) getFileHash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", fmt.Errorf("failed to compute hash: %w", err)
    }
    return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
    e := echo.New()
    sync := NewFileSync("./src", "./dst", true)

    e.GET("/sync", func(c echo.Context) error {
        if err := sync.Sync(); err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{
                "error": err.Error(),
            })
        }
        return c.JSON(http.StatusOK, map[string]string{
            "message": "Files synced successfully",
        })
    })

    e.Logger.Fatal(e.Start(":8080"))
}
