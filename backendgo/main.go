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
	db, err := repository.NewDb(os.Getenv("DB_URL"))
	if err != nil {
		panic("Connection database failed!")
	}
	repository := repository.NewIdentificationNumberRepository(db)
	service := services.NewIdentificationNumberService(repository)
	httpHandler := handler.NewHTTPHandler(service)
	graphQlHandler := handler.NewGraphQlHandler(service)

	gin.SetMode(os.Getenv("MODE"))
	router := gin.New()
	log.Println("Configuring CORSMiddleware")
	router.Use(CORSMiddleware())
	log.Println("Metrics Endpoint /metrics")
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.Use(router)
	log.Println("Health Endpoint /health")
	router.GET("/health", func(c *gin.Context) {
		log.Println(status)
		c.JSON(200, status)
	})
	log.Println("Configuring HttpHandler")
	httpHandler.Handler(router)
	log.Println("Configuring GraphQlHandler")
	graphQlHandler.Handler(router)
	log.Println("BackendGo - Running on :8081")
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
