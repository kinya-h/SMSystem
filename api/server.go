package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kinya-h/northlead-store-management-system/db"
)




type Server struct{
	db  *db.Queries
	router *chi.Mux
}

func NewServer(db  *db.Queries) *Server{
	server := &Server{db:db}
	router := chi.NewRouter()


	router.Post("/api/products" , server.addProduct)
	router.Get("/api/products" , server.getProducts)
	router.Patch("/api/products/{id}/price" , server.updateProductPrice)
	router.Patch("/api/products/{id}/rating" , server.updateProductRating)
	router.Patch("/api/products/{id}/available" , server.updateProductAvailability)
	router.Delete("/api/products/{id}" , server.deleteProduct)
	router.Get("/api/products/{id}" , server.getProduct)
	router.Post("/api/carts" , server.createCart)
	router.Post("/api/carts/{id}/items" , server.addProductToCart)
	router.Get("/api/carts/{id}/items" , server.getCartItems)
	router.Delete("/api/carts/{id}/items" , server.deleteCartItems)
	router.Patch("/api/carts/{id}/items" , server.updateCartItems)

	server.router = router
	return server



}

func (server *Server) Start(){
	http.ListenAndServe(":3000", server.router)

}