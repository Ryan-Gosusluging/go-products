package structs

import (
	"errors"
	"fmt"
	"bufio"
	"os"

	"github.com/Ryan-Gosusluging/go-products/model"
    "github.com/Ryan-Gosusluging/go-products/service"
)

func Display(){
	var choice int
    fmt.Println("Products Manager App")
    fmt.Println("Please choose the menu")
    for {
        fmt.Println("1. Get all products")
        fmt.Println("2. Get Product by ID")
        fmt.Println("3. Add new product")
        fmt.Println("4. Update product")
        fmt.Println("5. Delete product")
        fmt.Println("6. Quit")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            getAllProducts()
        case 2:
            getProductById()
        case 3:
            addProduct()
        case 4:
            updateProduct()
        case 5:
            deleteProduct()
        case 6:
            fmt.Println("Good bye..")
            return
        default:
            fmt.Println("Input invalid!")
            fmt.Println("Please choose from 1-6")
        }
    }
}

func getAllProducts() {
    var products []model.Product = service.GetProducts()
    if len(products) != 0 {
        fmt.Println("All products")
        for _, product := range products {
            showProductData(product)
        }
    } else {
        fmt.Println("No products found!")
    }
}

func getProductById() {
    var productId int64
    fmt.Println("Enter product ID: ")
    fmt.Scanln(&productId)
    _, product := service.GetProductById(productId)
    fmt.Println("Product data")
    showProductData(product)
}

func addProduct() {
    var reader *bufio.Reader = bufio.NewReader(os.Stdin)
    var name string
    var price int64
    var stock int

    fmt.Println("Enter product name: ")
    name, _ = reader.ReadString('\n')

    fmt.Println("Enter product price: ")
    fmt.Scanf("%f\n", &price)

    fmt.Println("Enter product stock: ")
    fmt.Scanf("%d\n", &stock)

    var newProduct model.ProductInput = model.ProductInput{
        Name:        name,
        Price:       price,
        Stock:       stock,
    }

    var insertedProduct = service.AddProduct(newProduct)
    fmt.Println("Your inserted product")
    showProductData(insertedProduct)
}

func updateProduct() {
    var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var id int64
    var name string
    var price int64
    var stock int

    fmt.Println("Enter product id: ")
    fmt.Scanf("%s\n", &productId)

    fmt.Println("Enter product name: ")
    name, _ = reader.ReadString('\n')

    fmt.Println("Enter product price: ")
    fmt.Scanf("%f\n", &price)

    fmt.Println("Enter product stock: ")
    fmt.Scanf("%d\n", &stock)

    var productData model.ProductInput = model.ProductInput{
        Name:        name,
        Price:       price,
        Stock:       stock,
    }

    var updatedProduct = service.UpdateProduct(productId, productData)
    fmt.Println("Your updated data")
    showProductData(updatedProduct)
}

func deleteProduct() {
    var productId int64
    fmt.Println("Enter product id: ")
    fmt.Scanf("%s\n", &productId)
    var isDeleted bool = service.DeleteProduct(productId)
    if isDeleted {
        errors.New("Data deleted!")
    }
}

func showProductData(product model.Product) {
    fmt.Println("ID: ", product.Id)
    fmt.Println("Name: ", product.Name)
    fmt.Println("Price: ", product.Price)
    fmt.Println("Stock: ", product.Stock)
}