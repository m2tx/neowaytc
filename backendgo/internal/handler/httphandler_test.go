package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m2tx/neowaytc/backendgo/core/domain"
	"github.com/m2tx/neowaytc/backendgo/core/services"
	"github.com/m2tx/neowaytc/backendgo/internal/repository"
	"github.com/stretchr/testify/assert"
)

var (
	data = []domain.IdentificationNumber{
		{uuid.MustParse("789c728f-8fa2-494b-8db1-18808a5c61d8"), "046.847.189-80", false},
		{uuid.MustParse("8ccf972c-6f24-4df3-ac65-b94853c10744"), "585.629.410-69", false},
		{uuid.MustParse("35240f60-6a08-4774-becd-826bae221876"), "335.796.160-13", true},
	}
	rep         = repository.NewIdentificationNumberMemoryRepository(data)
	service     = services.NewIdentificationNumberService(rep)
	httpHandler = NewHTTPHandler(service)
	router      = setUpRouter()
)

func setUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	httpHandler.Handler(router)
	return router
}

func TestGetAllIdentificationNumberHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/identificationnumber/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var ins []domain.IdentificationNumber
	json.Unmarshal(w.Body.Bytes(), &ins)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, ins)

}

func TestGetIdentificationNumberHandler(t *testing.T) {
	type test struct {
		Name     string
		ID       string
		Expected int
	}
	numbers := []test{
		{"StatusOK", "789c728f-8fa2-494b-8db1-18808a5c61d8", http.StatusOK},
		{"StatusNotFound", "123e4567-e89b-12d3-a456-426614174000", http.StatusNotFound},
		{"StatusBadRequest", "10", http.StatusBadRequest},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/identificationnumber/"+nb.ID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, nb.Expected, w.Code)
			if w.Code == http.StatusOK {
				var in domain.IdentificationNumber
				json.Unmarshal(w.Body.Bytes(), &in)
				assert.Equal(t, uuid.MustParse(nb.ID), in.ID)
			}
		})
	}

}

func TestNewIdentificationNumberHandler(t *testing.T) {
	type test struct {
		Name     string
		Number   string
		Expected int
	}
	numbers := []test{
		{"StatusCreated", "103.742.240-64", http.StatusCreated},
		{"StatusInternalServerErrorWithExits", "046.847.189-80", http.StatusInternalServerError},
		{"StatusInternalServerErrorWithInvalidCPF", "046.847.189-81", http.StatusInternalServerError},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			in := &domain.IdentificationNumber{
				Number: nb.Number,
			}
			jsonValue, _ := json.Marshal(in)
			req, _ := http.NewRequest("POST", "/api/identificationnumber/", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, nb.Expected, w.Code)
			if w.Code == http.StatusCreated {
				var in domain.IdentificationNumber
				json.Unmarshal(w.Body.Bytes(), &in)
				assert.NotEqual(t, uuid.Nil, in.ID)
			}
		})
	}

}

func TestUpdateIdentificationNumberHandler(t *testing.T) {
	type test struct {
		Name     string
		ID       string
		Json     string
		Expected int
	}
	numbers := []test{
		{"StatusOK1", "789c728f-8fa2-494b-8db1-18808a5c61d8", "{\"id\":\"789c728f-8fa2-494b-8db1-18808a5c61d8\", \"number\":\"046.847.189-80\", \"blocked\":false}", http.StatusOK},
		{"StatusOK2", "8ccf972c-6f24-4df3-ac65-b94853c10744", "{\"id\":\"8ccf972c-6f24-4df3-ac65-b94853c10744\", \"number\":\"585.629.410-69\", \"blocked\":false}", http.StatusOK},
		{"StatusBadRequest", "8ccf972c-6f24-4df3-ac65-b94853c10744", "{x:'bad_request', number:null, blocked:false}", http.StatusBadRequest},
	}
	for _, nb := range numbers {
		t.Run(nb.Name, func(t *testing.T) {
			req, _ := http.NewRequest("PUT", "/api/identificationnumber/"+nb.ID, bytes.NewBuffer([]byte(nb.Json)))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, nb.Expected, w.Code)
			if w.Code == http.StatusOK {
				var in domain.IdentificationNumber
				json.Unmarshal(w.Body.Bytes(), &in)
				assert.Equal(t, uuid.MustParse(nb.ID), in.ID)
			}
		})
	}

}
