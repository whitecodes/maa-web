package main

import (
	"fmt"
	"log"
	"maa-web/internal/router"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create or open a log file
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing log file: %v\n", err)
		}
	}()

	// Initialize the Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))
	e.Use(middleware.Recover())

	// Serve static files from the "web/public" directory
	e.Static("/", "web/dist")

	// 设置路由
	router.SetupRoutes(e)

	// 启动服务器
	err = e.Start("127.0.0.1:8080")
	if err != nil {
		fmt.Println("服务器启动失败：", err)
	}
}
