package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kinya-h/northlead-store-management-system/api"
	"github.com/kinya-h/northlead-store-management-system/db"
)



func createProduct(ID int64 ,name string, price float64, description string, image string, available bool, stock int64, category string, rating float64) (*db.Product, error) {

newProduct := &db.Product{
	    ID: ID,
        Name:        name,
        Price:       price,
        Description: description,
        Image:       image,
        Available:   available,
        Stock:      stock,
        Category:    category,
        Rating:      rating,
    }

 return newProduct, nil

}

func main(){

	
	
	ctx := context.Background()
	
	dbConnection, err := sql.Open("mysql", "root:root@/macs")
	
	if err != nil {
		fmt.Println("AN ERROR OCCURED ", err)
	}
	fmt.Print("CONNECTED SUCCESSFULLY ->" , dbConnection)
	
	queries := db.New(dbConnection)
	
	
	server := api.NewServer(queries)
	server.Start()
	

	// product, _ := createProduct() 
	// newProduct, _ := createProduct(2, "Product Name", 49.99, "Product Description", "product_image.jpg", true, 100, "Electronics", 4)
result, err:= queries.CreateProduct(ctx, db.CreateProductParams{
	Name: "Macadamia" , 
	Price: 18.5 ,
	Description: "Best Macadamia Nuts",
	Image: "my image",
	Stock: 120,
	Category: "Nuts",
	Rating: 3})



	if err != nil {
		fmt.Print("ERROR")
	}
	fmt.Print("RESULTS=" , result)

	insertedProductID, err := result.LastInsertId()
	// productID := int32(insertedProductID)

	if err != nil {
		fmt.Print("ERROR")
	}
	fmt.Println(insertedProductID)

	
	fetchedProduct, err := queries.GetProduct(ctx, insertedProductID)
	if err != nil {
		fmt.Print("ERROR")
	}
	fmt.Println("INSERTED PRODUCT ->"  , fetchedProduct)

}

