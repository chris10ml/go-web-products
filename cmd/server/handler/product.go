package handler

import (
	"fmt"
	"strconv"

	"github.com/chris10ml/go-web-products/internal/products"
	"github.com/chris10ml/go-web-products/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        string  `json:"code"`
	Posted      bool    `json:"posted"`
	DateCreated string  `json:"date_created"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		productsList, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, productsList, ""))
		// ctx.JSON(200, productsList)
	}
}

// AddProduct godoc
// @Summary Add product
// @Tags Products
// @Description add product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Validacion de token en Middleware

		// Cargo el body de la request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}

		// Validaciones
		requiredFieldsList := []string{"Name", "Color"}
		var errValidation error

		for _, fieldName := range requiredFieldsList {
			_, err := GetField(&req, fieldName)
			if err != nil {
				fmt.Printf("%s", err)
				errValidation = err
				break
			}
		}

		if errValidation != nil {
			ctx.JSON(422, web.NewResponse(422, nil, errValidation.Error()))
			return
		}

		// Hago el Post
		productCreated, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Posted, req.DateCreated)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, errValidation.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, productCreated, ""))
	}
}

// AddProduct godoc
// @Summary Update product
// @Tags Products
// @Description update product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Validacion de token en middleware

		//Obtener y Parsear id string a int
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}

		//Cargo el body de la request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		//Validaciones

		//Hago el Update
		productUpdated, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Posted, req.DateCreated)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, productUpdated, ""))
	}
}

// AddProduct godoc
// @Summary Update product
// @Tags Products
// @Description update name product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [patch]
func (p *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Obtener y Parsear id string a int
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}

		//Cargo el body de la request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		//Validaciones
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "name is required"))
			return
		}

		//Hago el Update del Name
		productUpdated, err := p.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, productUpdated, ""))
	}
}

// AddProduct godoc
// @Summary Delete product
// @Tags Products
// @Description delete product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Obtener y Parsear id string a int
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}

		//Hacer el delete y enviar respuestas
		err = p.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("Product %d deleted", id), ""))
	}
}
