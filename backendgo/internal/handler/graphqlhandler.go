package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/m2tx/neowaytc/backendgo/core/ports"
)

var (
	identificationNumberType *graphql.Object
	queryType                *graphql.Object
	schema                   graphql.Schema
)

type GraphQlHandler struct {
	service ports.IdentificationNumberService
}

func NewGraphQlHandler(service ports.IdentificationNumberService) *GraphQlHandler {
	return &GraphQlHandler{
		service: service,
	}
}

func (handler *GraphQlHandler) Handler(router *gin.Engine) {
	router.GET("/graphql", handler.GET)
	router.POST("/graphql", handler.POST)
	handler.configSchema()
}

func (handler *GraphQlHandler) configSchema() {
	identificationNumberType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "IdentificationNumber",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.ID,
				},
				"number": &graphql.Field{
					Type: graphql.String,
				},
				"blocked": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		},
	)
	queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"allIdentificationNumber": &graphql.Field{
					Type: graphql.NewList(identificationNumberType),
					Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
						return handler.service.GetAll(), nil
					},
				},
				"getIdentificationNumberByID": &graphql.Field{
					Type: identificationNumberType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
					},
					Resolve: func(rp graphql.ResolveParams) (interface{}, error) {
						id, ok := rp.Args["id"].(string)
						if ok {
							in, err := handler.service.Get(uuid.MustParse(id))
							if err != nil {
								return nil, err
							}
							return in, nil
						}
						return nil, ports.ErrorNotFoundIdentificationNumber
					},
				},
			},
		},
	)
	schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
}

func (handler *GraphQlHandler) execute(requestString string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestString,
	})
	if len(result.Errors) > 0 {
		log.Printf("GraphQl result with errors %v", result.Errors)
	}
	return result
}

func (handler *GraphQlHandler) GET(c *gin.Context) {
	query := c.Query("query")
	result := handler.execute(string(query))
	c.JSON(http.StatusOK, result)
}

func (handler *GraphQlHandler) POST(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, []string{err.Error()})
	}
	result := handler.execute(string(body))
	c.JSON(http.StatusOK, result)
}
