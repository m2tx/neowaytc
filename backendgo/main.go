package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/m2tx/neowaytc/backendgo/core/services"
	"github.com/m2tx/neowaytc/backendgo/internal/handler"
	"github.com/m2tx/neowaytc/backendgo/internal/repository"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

var (
	status = struct {
		Status string `json:"status"`
	}{
		Status: "UP",
	}
)

func main() {
	dbUrl := os.Getenv("DB_URL")
	mode := os.Getenv("MODE")

	repository := repository.NewIdentificationNumberPostgresRepository(dbUrl)
	service := services.NewIdentificationNumberService(repository)
	httpHandler := handler.NewHTTPHandler(service)

	gin.SetMode(mode)
	router := gin.New()
	router.Use(CORSMiddleware())
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.Use(router)
	router.GET("/health", func(c *gin.Context) {
		log.Println(status)
		c.JSON(200, status)
	})
	httpHandler.Handler(router)
	log.Println("BACKENDGO - RUNNING ON 8081")
	router.Run(":8081")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
