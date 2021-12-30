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
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Validacion de token en Middleware
		// token := ctx.GetHeader("token") //token := ctx.Request.Header.Get("token")
		// if token != "chris1234" {
		// 	ctx.JSON(401, gin.H{"error": "invalid token"})
		// }

		// Cargo el body de la request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
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
			ctx.JSON(422, gin.H{"error": errValidation.Error()})
			return
		}

		// Hago el Post
		productCreated, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Posted, req.DateCreated)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, productCreated)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Validar token
		token := ctx.GetHeader("token")
		if token != "chris1234" {
			ctx.JSON(401, gin.H{"error": "invalid token"})
		}

		//Obtener y Parsear id string a int
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
		}

		//Cargo el body de la request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Validaciones

		//Hago el Update
		p, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Posted, req.DateCreated)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Validar token
		token := ctx.GetHeader("token")
		if token != "chris1234" {
			ctx.JSON(401, gin.H{"error": "invalid token"})
		}

		//Obtener y Parsear id string a int
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
		}

		//Cargo el body de la request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Validaciones
		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "name is required"})
			return
		}

		//Hago el Update del Name
		p, err := p.service.UpdateName(int(id), req.Name)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Validar token
		token := ctx.GetHeader("token")
		if token != "chris1234" {
			ctx.JSON(401, gin.H{"error": "invalid token"})
		}

		//Obtener y Parsear id string a int
		idParam := ctx.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
		}

		//Hacer el delete y enviar respuestas
		err = p.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{
			"data": fmt.Sprintf("Product %d deleted", id),
		})
	}
}
