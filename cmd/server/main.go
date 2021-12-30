package main

import (
	"log"
	"os"

	"github.com/chris10ml/go-web-products/cmd/server/handler"
	"github.com/chris10ml/go-web-products/docs"
	"github.com/chris10ml/go-web-products/internal/products"
	"github.com/chris10ml/go-web-products/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error trying to load file .env")
	}

	db := store.New(store.FileType, "./products.json")
	productRepository := products.NewRepository(db)
	productService := products.NewService(productRepository)
	productHandler := handler.NewProduct(productService)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(handler.AuthTokenHandler)

	productsRoutes := router.Group("/products")
	{
		productsRoutes.GET("/", productHandler.GetAll())
		productsRoutes.POST("/", productHandler.Store())
		productsRoutes.PUT("/:id", productHandler.Update())
		productsRoutes.PATCH("/:id", productHandler.UpdateName())
		productsRoutes.DELETE("/:id", productHandler.Delete())
	}

	router.Run()
}
