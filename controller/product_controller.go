package controller

import (
	"net/http"

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