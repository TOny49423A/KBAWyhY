// 代码生成时间: 2025-08-30 01:42:56
// file_backup_sync.go
package main

import (
    "crypto/md5"
    "encoding/hex"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "time"
)

// FileSync contains source and destination paths
type FileSync struct {
    Source string
    Destination string
}

// NewFileSync initializes a new FileSync instance
func NewFileSync(source, destination string) *FileSync {
    return &FileSync{Source: source, Destination: destination}
}

// SyncFiles synchronizes files from source to destination
func (fs *FileSync) SyncFiles() error {
    srcFiles, err := ioutil.ReadDir(fs.Source)
    if err != nil {
        return err
    }

    for _, file := range srcFiles {
        srcPath := filepath.Join(fs.Source, file.Name())
        destPath := filepath.Join(fs.Destination, file.Name())

        if file.IsDir() {
            // Recursively copy directories
            err := fs.SyncFilesInDir(srcPath, destPath)
            if err != nil {
                return err
            }
        } else {
            // Copy individual files
            err := fs.CopyFile(srcPath, destPath)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// SyncFilesInDir synchronizes files in a directory
func (fs *FileSync) SyncFilesInDir(srcDir, destDir string) error {
    err := os.MkdirAll(destDir, 0755)
    if err != nil {
        return err
    }
    return fs.SyncFiles()
}

// CopyFile copies a single file from source to destination
func (fs *FileSync) CopyFile(srcPath, destPath string) error {
    srcFile, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    destFile, err := os.Create(destPath)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        return err
    }
    return destFile.Sync()
}

// CalculateMD5 calculates the MD5 checksum of a file
func CalculateMD5(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()
    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }
    return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
    srcPath := flag.String("source", "", "Source directory path")
    destPath := flag.String("destination", "", "Destination directory path")
    flag.Parse()

    if *srcPath == "" || *destPath == "" {
        log.Fatal("Both source and destination paths are required")
    }

    fs := NewFileSync(*srcPath, *destPath)
    err := fs.SyncFiles()
    if err != nil {
        log.Fatalf("Failed to sync files: %s", err)
    }
    fmt.Println("Files synced successfully")
}