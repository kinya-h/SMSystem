package api

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"encoding/json"
	// "macsinterapi/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kinya-h/northlead-store-management-system/db"
)




type CreateProductParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Available   bool     `json:"available"`
	Rating      float64 `json:"rating"`
	Category    string  `json:"category"`
	Stock       int64   `json:"stock"`
}

func (server *Server) addProduct(w http.ResponseWriter, r *http.Request){
	// server.db


	var req CreateProductParams

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return 
	}

	arg := db.CreateProductParams{
		Name: req.Name , 
		Price: req.Price ,
		Description: req.Description,
		Image:req.Image,
		Stock: int64(req.Stock),
		Available: req.Available,
		Category: req.Category,
		Rating: req.Rating,
	}

	// var arg db.CreateProductParams
	
	result , err := server.db.CreateProduct(context.Background() , arg)
	if err!= nil{
		http.Error(w, fmt.Sprintf(" Error %s" , err) , http.StatusInternalServerError)
		return
	}

	prodId , err := result.LastInsertId()
	if err!=nil{
		fmt.Printf("An Error Occured while retrieving the index : %s" , err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Product created successfully , id : %d " , prodId)))
}



func (server *Server) getProducts(w http.ResponseWriter, r *http.Request){
	// server.db
	var err error
	var products []db.Product

	type productsResponse struct {
	Products []db.Product `json:"products"`
}

	products , err = server.db.ListProducts(context.Background())
	if err!= nil{
		http.Error(w, fmt.Sprintf(" Error %s" , err) , http.StatusInternalServerError)
		return
	}

	response := productsResponse{Products: products}
	json.NewEncoder(w).Encode(response)
	// w.WriteHeader(http.StatusOK)
}


func (server * Server) getProduct(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	productID, err := strconv.Atoi(id)

	if err != nil {
		
			http.Error(w, " id parameter Must be a number (id of the product)" , http.StatusBadRequest)
				
	}
	
	product , err := server.db.GetProduct(context.Background() ,int64(productID))
	if err != nil {

		if  err == sql.ErrNoRows{
			 emptyProducts:= make( []db.Product , 0) // Return [] instead of a struct with empty fields
			json.NewEncoder(w).Encode(emptyProducts)

			return
		}
		http.Error(w, fmt.Sprintf("An Error Occured %s" , err) , http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(product)
}


func (server * Server) updateProductPrice(w http.ResponseWriter, r *http.Request){
	var req  db.UpdateProductPriceParams

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return 
	}

	arg := db.UpdateProductPriceParams{
		Price: req.Price,
		ID: req.ID,
	}

	result, err := server.db.UpdateProductPrice(context.Background() , arg);

	if err!= nil{
		http.Error(w, fmt.Sprintf(" An Error Occured %s" , err) , http.StatusInternalServerError)
		return
	}

	updateProductId , _ := result.LastInsertId()
	

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Product with id %d was updated succesfully" , updateProductId )))


}

func (server * Server) updateProductRating(w http.ResponseWriter, r *http.Request){
	
	var arg  db.UpdateProductRatingParams

	err := json.NewDecoder(r.Body).Decode(&arg)
	if err != nil{
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return 
	}


	result, err := server.db.UpdateProductRating(context.Background() , arg);

	if err!= nil{
		http.Error(w, fmt.Sprintf(" An Error Occured %s" , err) , http.StatusInternalServerError)
		return
	}

	updatedProductId , _ := result.LastInsertId()
	

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Rating for Product with id %d was updated succesfully" , updatedProductId )))


}


func (server * Server) updateProductAvailability(w http.ResponseWriter, r *http.Request){
	
	var arg  db.UpdateProductAvailabilityParams

	err := json.NewDecoder(r.Body).Decode(&arg)
	if err != nil{
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return 
	}


	result, err := server.db.UpdateProductAvailability(context.Background() , arg);

	if err!= nil{
		http.Error(w, fmt.Sprintf(" An Error Occured %s" , err) , http.StatusInternalServerError)
		return
	}

	updatedProductId , _ := result.LastInsertId()
	

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Availability for Product with id %d was updated succesfully" , updatedProductId )))


}


func (server * Server) deleteProduct(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	productID, err := strconv.Atoi(id)

	if err != nil {
		
			http.Error(w, " id parameter Must be a number (id of the product)" , http.StatusBadRequest)
				
	}
	
	result , err := server.db.DeleteProduct(context.Background() , int64(productID))
	if err != nil {
		http.Error(w, "No Product with the given id was found" , http.StatusNotFound)
	}

	deleted, _ := result.LastInsertId()

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Product with id %d was deleted succesfully" , deleted )))

	
	// json.NewEncoder(w).Encode(product)
}