package handler

import (
	"net/http"

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

func (handler *HTTPHandler) Handler(router *gin.Engine) {
	router.GET("/api/identificationnumber/", handler.GetAll)
	router.GET("/api/identificationnumber/:id", handler.GetById)
	router.POST("/api/identificationnumber/", handler.New)
	router.PUT("/api/identificationnumber/:id", handler.Update)
	router.DELETE("/api/identificationnumber/:id", handler.Delete)
	router.POST("/api/identificationnumber/query/", handler.Query)
}

func (handler *HTTPHandler) GetAll(c *gin.Context) {
	ins := handler.service.GetAll()
	c.JSON(http.StatusOK, ins)
}

func (handler *HTTPHandler) GetById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, []string{err.Error()})
		return
	}
	in, err := handler.service.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, []string{err.Error()})
		return
	}
	c.JSON(http.StatusOK, in)
}

func (handler *HTTPHandler) New(c *gin.Context) {
	var body struct {
		Number string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, []string{err.Error()})
		return
	}
	in, err := handler.service.New(body.Number)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, []string{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, in)
}

func (handler *HTTPHandler) Update(c *gin.Context) {
	_, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, []string{err.Error()})
		return
	}
	var body domain.IdentificationNumber
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, []string{err.Error()})
		return
	}
	err = handler.service.Update(body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, []string{err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

func (handler *HTTPHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, []string{err.Error()})
		return
	}
	identificationnumber := domain.IdentificationNumber{ID: id}
	err = handler.service.Delete(identificationnumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, []string{err.Error()})
		return
	}

	c.JSON(http.StatusOK, identificationnumber)
}

func (handler *HTTPHandler) Query(c *gin.Context) {
	params := make(map[string]any)
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, []string{err.Error()})
		return
	}
	pageIndex := utils.StringToInt(c.Query("page"))
	size := utils.StringToInt(c.Query("size"))
	sort, _ := domain.ParseSort(c.Query("sort"))
	pageable := domain.Pageable{
		Page:     pageIndex,
		PageSize: size,
		Sort:     sort,
	}
	page, err := handler.service.Query(params, pageable)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, []string{err.Error()})
		return
	}
	c.JSON(http.StatusOK, page)
}
