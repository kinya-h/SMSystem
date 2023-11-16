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
	router.Delete("/api/products/{id}" , server.deleteProduct)
	router.Get("/api/products/{id}" , server.getProduct)

	server.router = router
	return server



}

func (server *Server) Start(){
	http.ListenAndServe(":3000", server.router)

}