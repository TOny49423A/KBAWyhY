// 代码生成时间: 2025-08-28 11:26:56
package main

import (
    "archive/zip"
    "bytes"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/labstack/echo"
)

// DecompressHandler is the Echo handler for decompressing files.
func DecompressHandler(c echo.Context) error {
    // Get the uploaded file from the request form.
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    // Create a buffer to store the file data temporarily.
    dst := new(bytes.Buffer)
    _, err = io.Copy(dst, src)
    if err != nil {
        return err
    }

    // Decompress the file to the specified directory.
    err = decompress(dst.Bytes(), "./decompressed")
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, echo.Map{{"message": "File decompressed successfully"})
}

// decompress is a helper function to decompress zip files.
func decompress(data []byte, path string) error {
    // Create a reader for the zip data.
    r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
    if err != nil {
        return fmt.Errorf("failed to create zip reader: %w", err)
    }
    defer r.Close()

    // Iterate through the files in the zip archive.
    for _, f := range r.File {
        fr, err := f.Open()
        if err != nil {
            return fmt.Errorf("failed to open file %q in zip: %w", f.Name, err)
        }
        defer fr.Close()

        // Create the directory structure.
        dir := filepath.Join(path, filepath.Dir(f.Name))
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            if err := os.MkdirAll(dir, 0755); err != nil {
                return fmt.Errorf("failed to create directory %q: %w", dir, err)
            }
        }

        // Create the file.
        dest, err := os.Create(filepath.Join(path, f.Name))
        if err != nil {
            return fmt.Errorf("failed to create file %q: %w", f.Name, err)
        }
        defer dest.Close()

        // Copy the file data.
        if _, err := io.Copy(dest, fr); err != nil {
            return fmt.Errorf("failed to copy file %q data: %w", f.Name, err)
        }
    }
    return nil
}

func main() {
    e := echo.New()
    e.POST("/decompress", DecompressHandler)
    e.Start(":8080")
}
