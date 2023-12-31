// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import ()

type Cart struct {
	ID        int64       `json:"id"`
	CartID    string      `json:"cart_id"`
	CreatedAt interface{} `json:"created_at"`
}

type CartItem struct {
	ID        int64  `json:"id"`
	CartID    string `json:"cart_id"`
	ProductID int64  `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Available   bool    `json:"available"`
	Stock       int64   `json:"stock"`
	Category    string  `json:"category"`
	Rating      float64 `json:"rating"`
}

type User struct {
	ID        interface{} `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Password  string      `json:"password"`
}
