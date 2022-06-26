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
	repository := repository.NewIdentificationNumberPostgresRepository(dbUrl)
	service := services.NewIdentificationNumberService(repository)
	httpHandler := handler.NewHTTPHandler(service)

	router := gin.New()
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.Use(router)
	router.GET("/health", func(c *gin.Context) {
		log.Println(status)
		c.JSON(200, status)
	})
	router.GET("/api/identificationnumber/", httpHandler.GetAll)
	router.GET("/api/identificationnumber/:id", httpHandler.GetById)
	router.POST("/api/identificationnumber/", httpHandler.New)
	router.PUT("/api/identificationnumber/:id", httpHandler.Update)
	router.POST("/api/identificationnumber/query/", httpHandler.Query)
	router.Run(":8081")
}
