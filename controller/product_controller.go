package controller

import (
	"net/http"
	"strconv"

	"github.com/Fabio17CN/Projeto_GO_APIs/model"
	"github.com/Fabio17CN/Projeto_GO_APIs/usecase"
	"github.com/gin-gonic/gin"
)

type productController struct {
	 ProductUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
	ProductUseCase: usecase,
}
}

func (p *productController) GetProducts(ctx *gin.Context){
	products, err := p.ProductUseCase.GetProducts()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
	}
	
 ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context){
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.ProductUseCase.CreateProduct(product)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context){
	
	id := ctx.Param("productId")
	if id == ""{
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil{
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
	}
	ctx.JSON(http.StatusBadRequest, response)
	return
}

	product, err := p.ProductUseCase.GetProductById(productId)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}
	if product == nil{
		response := model.Response{
			Message: "Produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
 ctx.JSON(http.StatusOK, product)
}