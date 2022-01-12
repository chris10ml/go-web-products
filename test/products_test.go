package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chris10ml/go-web-products/cmd/server/handler"
	"github.com/chris10ml/go-web-products/internal/products"
	"github.com/chris10ml/go-web-products/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Función para crear el Server y definir las Rutas
func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "chris1234")

	db := store.New(store.FileType, "./products.json")

	productRepository := products.NewRepository(db)
	productService := products.NewService(productRepository)
	productHandler := handler.NewProduct(productService)

	router := gin.Default()

	productsRoutes := router.Group("/products")
	{
		productsRoutes.GET("/", productHandler.GetAll())      // ejemplo
		productsRoutes.POST("/", productHandler.Store())      // ejemplo
		productsRoutes.PATCH("/:id", productHandler.Update()) // EJERCICIO 1
	}

	return router
}

// Función para generar el Request y Response según nuestras necesidades
func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "chris1234")

	return req, httptest.NewRecorder()
}

// --------------- GETALL ---------------
// Se obtienen todos los productos y se valida la respuesta.
func Test_GetProduct_OK(t *testing.T) {
	// estructura response
	objRes := struct {
		Code string      `json:"code"`
		Data interface{} `json:"data"`
	}{}

	// crear el Server y definir las Rutas
	r := createServer()

	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	fmt.Println("respuesta: ", objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

// --------------- POST STORE ---------------
// Se da de alta un producto y se valida la creación exitosa.
func Test_SaveProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
        "nombre": "Tester","tipo": "Funcional","cantidad": 10,"precio": 99.99
    }`)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}
