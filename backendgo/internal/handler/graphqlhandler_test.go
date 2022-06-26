package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

var (
	graphQlHandler = NewGraphQlHandler(service)
	routerGraphQl  = setUpRouterGraphQlHandler()
)

func setUpRouterGraphQlHandler() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	graphQlHandler.Handler(router)
	return router
}

func TestGraphQlHandler(t *testing.T) {
	type test struct {
		Name          string
		RequestString string
		Expected      int
	}
	tests := []test{
		{"AllIdentificationNumberStatusOK", "query { allIdentificationNumber { id } }", http.StatusOK},
		{"getIdentificationNumberByIDStatusOK1", "query { getIdentificationNumberByID (id:\"98b72201-446f-461f-bed4-d8193eded5ea\") { id,number } }", http.StatusOK},
		{"getIdentificationNumberByIDStatusOK2", "query { getIdentificationNumberByID (id:\"123e4567-e89b-12d3-a456-426614174000\") { id } }", http.StatusOK},
	}
	for _, nb := range tests {
		t.Run(nb.Name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/graphql", bytes.NewBuffer([]byte(nb.RequestString)))
			w := httptest.NewRecorder()
			routerGraphQl.ServeHTTP(w, req)
			assert.Equal(t, nb.Expected, w.Code)
			if w.Code == http.StatusOK {
				var in graphql.Result
				json.Unmarshal(w.Body.Bytes(), &in)
				assert.NotEmpty(t, in.Data)
			}
		})
	}
}
