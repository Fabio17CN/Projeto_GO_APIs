package router

import (
	"github.com/Fabio17CN/Projeto_GO_APIs/controller"
	"github.com/Fabio17CN/Projeto_GO_APIs/db"
	"github.com/Fabio17CN/Projeto_GO_APIs/repository"
	"github.com/Fabio17CN/Projeto_GO_APIs/usecase"
	"github.com/gin-gonic/gin"
)


func Initialize()  {
		// inicializando o Router utilizando as configuração default do gin
		server := gin.Default()

		dbConnection, err := db.ConnectDB()
		if err != nil{
			panic(err)
		}
		// Camada de Repository
    ProductRepository := repository.NewProductRepository(	dbConnection)
		
		// camada usecase
		ProductUseCase := usecase.NewProductUseCase( ProductRepository)

		//Camada de controllers
		ProductController := controller.NewProductController(ProductUseCase)
	  server.GET("/ping", func(ctx *gin.Context){
			ctx.JSON(200,gin.H{
				"message": "pong",
			})
		})
		server.GET("/products",ProductController.GetProducts)
		server.POST("/product", ProductController.CreateProduct)
	  server.Run(":8080") // listen and serve on 0.0.0.0:8080
}