package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
	"github.com/m2tx/neowaytc/backendgo/core/utils"
)

type HTTPHandler struct {
	service ports.IdentificationNumberService
}

func NewHTTPHandler(service ports.IdentificationNumberService) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (handler *HTTPHandler) GetAll(c *gin.Context) {
	ins := handler.service.GetAll()
	c.JSON(200, ins)
}

func (handler *HTTPHandler) GetById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, []string{err.Error()})
		return
	}
	in, err := handler.service.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(400, []string{err.Error()})
		return
	}
	c.JSON(200, in)
}

func (handler *HTTPHandler) New(c *gin.Context) {
	body := domain.IdentificationNumber{}
	c.BindJSON(&body)

	in, err := handler.service.New(body.Number)
	if err != nil {
		c.AbortWithStatusJSON(500, []string{err.Error()})
		return
	}

	c.JSON(201, in)
}

func (handler *HTTPHandler) Update(c *gin.Context) {
	_, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, []string{err.Error()})
		return
	}
	body := domain.IdentificationNumber{}
	c.BindJSON(&body)

	err = handler.service.Update(body)
	if err != nil {
		c.AbortWithStatusJSON(500, []string{err.Error()})
		return
	}

	c.JSON(200, body)
}

func (handler *HTTPHandler) Query(c *gin.Context) {
	params := make(map[string]string)
	c.BindJSON(&params)
	sort, err := domain.ParseSort(c.Param("sort"))
	if err != nil {
		log.Println(err.Error())
	}
	pageable := domain.Pageable{
		Page:     utils.StringToInt(c.Param("page")),
		PageSize: utils.StringToInt(c.Param("size")),
		Sort:     sort,
	}

	page, err := handler.service.Query(params, pageable)
	if err != nil {
		c.AbortWithStatusJSON(500, []string{err.Error()})
		return
	}

	c.JSON(200, page)
}
