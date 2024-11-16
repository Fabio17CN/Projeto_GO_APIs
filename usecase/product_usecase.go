package usecase

import (
	"github.com/Fabio17CN/Projeto_GO_APIs/model"
	"github.com/Fabio17CN/Projeto_GO_APIs/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product,error){
	return pu.repository.GetProducts()
}

func(pu *ProductUsecase) CreateProduct(product model.Product)(model.Product, error){

	productID,err := pu.repository.CreateProduct(product)
	if(err != nil){
		return model.Product{},err
	}

	product.ID = productID
	return product,nil
}

func(pu *ProductUsecase)GetProductById(id_product int)(*model.Product, error){
	product, err := pu.repository.GetProductById(id_product)
	if err != nil{
		return nil, err
	}
	return product, nil
}