// 代码生成时间: 2025-08-04 23:55:37
package main

import (
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/labstack/echo"
)

// ScrapeWebPage defines the structure for the scraped data
type ScrapeWebPage struct {
    URL   string `json:"url"`
    Title string `json:"title"`
    Body  string `json:"body"`
}

// scraperService handles the scraping logic
type scraperService struct {
    // No additional fields needed for this simple example
}

// NewScraperService creates a new instance of scraperService
func NewScraperService() *scraperService {
    return &scraperService{}
}

// Scrape performs the actual scraping of the webpage
func (svc *scraperService) Scrape(url string) (*ScrapeWebPage, error) {
    // Create an HTTP client without checking for SSL certificates
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }
    
    response, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()
    
    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch page, status: %d", response.StatusCode)
    }
    
    bodyBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }
    
    body := string(bodyBytes)
    title := strings.Title(strings.TrimSpace(extractTitle(body)))
    
    return &ScrapeWebPage{URL: url, Title: title, Body: body}, nil
}

// extractTitle attempts to extract the title from the HTML body
func extractTitle(body string) string {
    // This is a very basic implementation and may not work for all websites
    start := strings.Index(body, "<title>")
    end := strings.Index(body, "</title>")
    if start == -1 || end == -1 {
        return ""
    }
    return strings.TrimSpace(body[start+len("<title>") : end])
}

func main() {
    e := echo.New()
    scraper := NewScraperService()
    
    // Define the route for scraping a webpage
    e.GET("/scrape", func(c echo.Context) error {
        url := c.QueryParam("url")
        if url == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "URL parameter is required")
        }
        
        result, err := scraper.Scrape(url)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }
        
        return c.JSON(http.StatusOK, result)
    })
    
    // Start the Echo server
    e.Start(":8080")
}
