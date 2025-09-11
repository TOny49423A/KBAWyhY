// 代码生成时间: 2025-09-12 00:50:52
package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/labstack/echo"
    "github.com/robfig/cron/v3"
)

// Scheduler is the main struct for the application.
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler creates a new instance of Scheduler.
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(),
    }
}

// Start starts the scheduler and begins processing the cron jobs.
func (s *Scheduler) Start() {
    s.cron.Start()
    defer s.cron.Stop()

    // Handle graceful shutdown.
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    sig := <-sigChan
    log.Printf("Received signal: %s", sig)
}

// AddJob adds a new job to the scheduler.
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    if _, err := s.cron.AddFunc(spec, cmd); err != nil {
        return err
    }
    return nil
}

// RunCronJob is a sample function that can be scheduled to run at a specific time.
func RunCronJob() {
    log.Println("Cron job executed...")
}

func main() {
    // Create a new scheduler instance.
    scheduler := NewScheduler()

    // Add a job to the scheduler that runs every minute.
    _ = scheduler.AddJob("* * * * *", RunCronJob)

    // Start the scheduler.
    scheduler.Start()

    // Start the Echo server on a separate goroutine.
    go func() {
        e := echo.New()
        // Define routes and middleware here.
        e.GET("/", func(c echo.Context) error {
            return c.String(http.StatusOK, "Scheduler is running...")
        })
        e.Logger.Fatal(e.Start(":8080"))
    }()
}
