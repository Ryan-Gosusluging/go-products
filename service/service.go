package service

import (
	"errors"
	"github.com/Ryan-Gosusluging/go-products/model"
    "github.com/Ryan-Gosusluging/go-products/repository"
)

type Service interface {
	GetAllProducts() []model.Product
	GetProductByID(id int64) (model.Product, error)
	AddProduct(name string, price int64) error
	UpdateProduct(id int64, name string, price int64) error
	DeleteProduct(id int64) error
}

func GetAllProducts() []model.Product{
	return repository.GetAllProducts()
}

func GetProductByID(id int64) (int64, model.Product, error){
	productIndex, product := repository.GetProductById(id)
    return productIndex, product
}

func AddProduct(productInput model.ProductInput) (model.Product, error){
	_, err := repository.GetProductByName(name)
	if err == nil{
		return errors.New("Product with this name already exists!")
	}
	if !errors.Is(err, errors.New("Product with this name isn't found!")){
		return err
	}
	return repository.AddProduct(productInput)
}

func UpdateProduct(id int64, productInput model.ProductInput) (model.Product, error){
	_, err := repository.GetProductByID(id)
	if err != nil{
		return err
	}
	existingProduct, err := repository.GetProductByName(name)
	if err == nil && existingProduct.ID != id {
		return errors.New("product with this id already exists")
	}
	if !errors.Is(err, errors.New("Product with this id isn't found!")){
		return err
	}
	return repository.UpdateProduct(id, productInput)
}

func DeleteProduct(id int64) bool{
	return repository.DeleteProduct(id)
}