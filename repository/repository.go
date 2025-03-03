package repository

import (
	"errors"
	"math/rand"
	"time"
	"github.com/google/uuid"
	"github.com/Ryan-Gosusluging/go-products/model"
)

type Repository interface{
	GetAllProducts() []model.Product
	GetProductByID(id int64) (model.Product, error)
	AddProduct(product model.Product) error
	UpdateProduct(product model.Product) error
	DeleteProduct(id int64) error
	GetProductByName(name string) (model.Product, error) 
}

var database []model.Product = []model.Product{}

func GetAllProducts() []model.Product{
	return database
}

func GetProductByID(id int64) (model.Product, error){
	for i := 0; i < len(database); i++{
		if database[i].ID == id{
			return database[i], nil
		}
	}
	return model.Product{}, errors.New("Product with this id isn't found!")
}

func GetProductByName(name string) (model.Product, error){
	for i := 0; i < len(database); i++{
		if database[i].name == name{
			return database[i], nil
		}
	}
	return model.Product{}, errors.New("Product with this name isn't found!")
}

func AddProduct(productInput model.ProductInput) (model.Product, error){
	var newProduct model.Product = model.Product{
        Id:          rand.Int63(),
        Name:        productInput.Name,
        Price:       productInput.Price,
        Stock:       productInput.Stock,
    }
	database = append(database, newProduct)
	return newProduct
}

func UpdateProduct(id int64, productInput model.ProductInput) (model.Product){
	index, product := GetProductById(id)
    product.Name = productInput.Name
    product.Price = productInput.Price
    product.Stock = productInput.Stock
    database[index] = product
    return product
}

func DeleteProduct(id int64) bool{
	var afterDeleted []model.Product = []model.Product{}
    for _, value := range database {
        if value.Id != id {
            afterDeleted = append(afterDeleted, value)
        }
    }
    database = afterDeleted
    return true
}