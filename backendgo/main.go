package main

import (
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
	repository := repository.NewIdentificationNumberPostgresRepository(os.Getenv("DB_URL"))
	service := services.NewIdentificationNumberService(repository)
	httpHandler := handler.NewHTTPHandler(service)

	router := gin.New()
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.Use(router)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, status)
	})
	router.GET("/api/identificationnumber/:id", httpHandler.GetById)
	router.POST("/api/identificationnumber/", httpHandler.New)
	router.PUT("/api/identificationnumber/:id", httpHandler.Update)
	router.Run(":8081")
}
