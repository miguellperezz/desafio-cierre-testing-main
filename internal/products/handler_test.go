package products_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/cmd/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProdcutsHandler_GetProducts(t *testing.T) {
	t.Run("should return a list of products", func(t *testing.T) {
		// arrange ->>>>Mis expectativas
		var (
			expectedStatusCode = 200
			expectedHeaders    = http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
			expectedBody       = `[{"ID":"mock","SellerID":"FEX112AC","Description":"generic product","Price":123.55}]`

			//algunos dirán generar el request
			request = httptest.NewRequest(http.MethodGet, "/api/v1/products?seller_id=1", nil)
			/*Si es un POST*/
			//request = httptest.NewRequest(http.MethodGet, "/api/v1/products?seller_id=1", strings.NewReader(`{"ID":"mock","SellerID":"FEX112AC","Description":"generic product","Price":123.55}`))
			response = httptest.NewRecorder()
			engine   = gin.New()
		)

		//repo := NewRepository()
		//service := NewService(repo)
		//handler := NewHandler(service)

		router.MapRoutes(engine)

		// act
		engine.ServeHTTP(response, request)

		// assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.JSONEq(t, expectedBody, response.Body.String())
	})

}
