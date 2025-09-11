// 代码生成时间: 2025-09-11 11:50:22
package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	
	"github.com/labstack/echo/v4"
)

// FolderOrganizer is a struct that holds the root directory path
type FolderOrganizer struct {
	RootDir string
}

// NewFolderOrganizer creates a new FolderOrganizer instance
func NewFolderOrganizer(rootDir string) *FolderOrganizer {
	return &FolderOrganizer{RootDir: rootDir}
}

// OrganizeFolderStructure scans the root directory and organizes its contents
func (f *FolderOrganizer) OrganizeFolderStructure() error {
	// Read the directory contents
	entries, err := os.ReadDir(f.RootDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// Sort the directory entries
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	// Organize each entry
	for _, entry := range entries {
		if entry.IsDir() {
			// Recursively organize subdirectories
			subFolderOrganizer := NewFolderOrganizer(filepath.Join(f.RootDir, entry.Name()))
			if err := subFolderOrganizer.OrganizeFolderStructure(); err != nil {
				return fmt.Errorf("failed to organize subdirectory %q: %w", entry.Name(), err)
			}
		} else {
			// Move files to a separate directory named 'files'
			filesDir := filepath.Join(f.RootDir, "files")
			if _, err := os.Stat(filesDir); os.IsNotExist(err) {
				if err := os.Mkdir(filesDir, fs.ModePerm); err != nil {
					return fmt.Errorf("failed to create files directory: %w", err)
				}
			}
			if err := os.Rename(filepath.Join(f.RootDir, entry.Name()), filepath.Join(filesDir, entry.Name())); err != nil {
				return fmt.Errorf("failed to move file %q: %w", entry.Name(), err)
			}
		}
	}
	return nil
}

// StartWebServer starts an Echo web server that can trigger folder organization
func StartWebServer(organizer *FolderOrganizer) {
	e := echo.New()
	e.GET("/organize", func(c echo.Context) error {
		if err := organizer.OrganizeFolderStructure(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Folder organized successfully"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func main() {
	// Create a new FolderOrganizer instance with the desired root directory
	organizer := NewFolderOrganizer("./")

	// Start the web server
	StartWebServer(organizer)
}
