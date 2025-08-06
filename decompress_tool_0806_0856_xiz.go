// 代码生成时间: 2025-08-06 08:56:26
 * This tool provides an endpoint to decompress files. It supports different compression formats.
 *
 * @author Your Name
 * @date Today's Date
 */

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "archive/zip"
    "github.com/labstack/echo"
)

// DecompressFile decompresses a file to the specified directory.
func DecompressFile(src, dest string, format string) error {
    var err error
    switch format {
    case "zip":
        err = decompressZip(src, dest)
        break
    // Add other formats if needed.
    default:
        return fmt.Errorf("unsupported format: %s", format)
    }
    return err
}

// decompressZip decompresses a zip file to the specified directory.
func decompressZip(src, dest string) error {
    reader, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        filePath := filepath.Join(dest, file.Name)
        if file.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
        } else {
            if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
                return err
            }
            outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
            if err != nil {
                return err
            }
            defer outFile.Close()

            fileReader, err := file.Open()
            if err != nil {
                return err
            }
            defer fileReader.Close()

            _, err = outFile.Write(fileReader.Bytes())
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func main() {
    e := echo.New()
    e.POST("/decompress", decompressHandler)
    e.Logger.Fatal(e.Start(":8080"))
}

// decompressHandler handles the decompression of a file.
func decompressHandler(c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    dest := "./decompressed" // Destination directory for decompressed files.
    err = DecompressFile(file.Filename, dest, "zip")
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "File decompressed successfully.",
    })
}