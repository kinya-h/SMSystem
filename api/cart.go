package api

import (
	"context"
	"database/sql"
	"fmt"

	"encoding/json"
	// "macsinterapi/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kinya-h/northlead-store-management-system/db"
)


func (server *Server) createCart (w http.ResponseWriter, r *http.Request){
	result , err := server.db.CreateCart(context.Background())

	if err != nil {
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return
	}
	
	cartID, err := result.LastInsertId()

	if err != nil {
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return
	}

	
	cart , err := server.db.GetCart(context.Background(), cartID)
	
	if err != nil {
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return
	}

	

	type cartResponse struct{
		Cart db.Cart  `json:"cart"`
	}	

	createCart := cartResponse{Cart: cart}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createCart)

}


func (server *Server) addProductToCart(w http.ResponseWriter, r *http.Request){
	
	cartID := chi.URLParam(r, "id")
	

	var arg db.SaveCartItemsParams

	err := json.NewDecoder(r.Body).Decode(&arg)
	if err != nil{
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return 
	}

	arg.CartID = cartID;
	
	result , err := server.db.SaveCartItems(context.Background() , arg)
	if err!= nil{
		http.Error(w, fmt.Sprintf(" Error %s" , err) , http.StatusInternalServerError)
		return
	}

	prodId , err := result.LastInsertId()
	if err!=nil{
		fmt.Printf("An Error Occured while retrieving the index : %s" , err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Product Added to Cart successfully , id : %d " , prodId)))
}


func (server *Server) getCartItems(w http.ResponseWriter, r *http.Request){
	
	cartID := chi.URLParam(r, "id")
	

	cartItems , err := server.db.GetCartItems(context.Background() , cartID)


	if err!= nil{
		if err == sql.ErrNoRows{
			http.Error(w, "No cart was found with the given id" , http.StatusNotFound)
		return

		}else{
			http.Error(w, fmt.Sprintf(" Error %s" , err) , http.StatusInternalServerError)
		return

		}
	}

	type cartItemsResults struct{
		Items []db.GetCartItemsRow `json:"items"`
	}

	items := cartItemsResults{Items: cartItems}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	
}




func (server * Server) updateCartItems(w http.ResponseWriter, r *http.Request){

	cartID := chi.URLParam(r, "id")
	var arg db.UpdateCartItemsParams

	err := json.NewDecoder(r.Body).Decode(&arg)
	if err != nil{
		http.Error(w, fmt.Sprintf("An error occured %s" , err) , http.StatusBadRequest)
		return 
	}

	arg.CartID = cartID

	result, err := server.db.UpdateCartItems(context.Background() , arg);

	if err!= nil{
		http.Error(w, fmt.Sprintf(" An Error Occured %s" , err) , http.StatusInternalServerError)
		return
	}

	updatedCartItemID , _ := result.LastInsertId()
	

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Product quantity with id %d was updated succesfully for cart Item %d" , arg.ProductID ,updatedCartItemID )))


}



func (server *Server) deleteCartItems(w http.ResponseWriter, r *http.Request){
	
	cartID := chi.URLParam(r, "id")
	

	result , err := server.db.DeleteCartItems(context.Background() , cartID)


	if err!= nil{
		
		http.Error(w, fmt.Sprintf(" Error %s" , err) , http.StatusInternalServerError)
	
		return

	}
	
	deletedID ,err :=  result.LastInsertId()
	if err!= nil{
		
		http.Error(w, fmt.Sprintf(" Error %s" , err) , http.StatusInternalServerError)
	
		return

	}
	
	server.db.DeleteCart(r.Context() , cartID)

	
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("cart with id %d was deleted successfully ", deletedID)))
	
	
}