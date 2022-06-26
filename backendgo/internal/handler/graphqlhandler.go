package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
)

type GraphQlHandler struct {
	service ports.IdentificationNumberService
}

func NewGraphQlHandler(service ports.IdentificationNumberService) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (handler *GraphQlHandler) Handler(router *gin.Engine) {
	router.GET("/graphql/", handler.GET)
	router.POST("/graphql/", handler.POST)
}

func (handler *GraphQlHandler) GET(c *gin.Context) {

}

func (handler *GraphQlHandler) POST(c *gin.Context) {

}
