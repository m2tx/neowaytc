package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/services"
	"github.com/m2tx/neowaytc/backendgo/internal/repository"
	"github.com/stretchr/testify/assert"
)

var (
	routerGraphQl = setUpRouterGraphQlHandler()
)

func setUpRouterGraphQlHandler() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	service := services.NewIdentificationNumberService(*repository.NewIdentificationNumberRepositoryTest([]domain.IdentificationNumber{
		{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", false},
		{uuid.MustParse("8ccf972c-6f24-4df3-ac65-b94853c10744"), "585.629.410-69", false},
		{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", true},
	}))
	NewGraphQlHandler(service).Handler(router)
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
