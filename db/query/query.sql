-- name: GetProduct :one
SELECT * FROM products
WHERE id = ?;

-- name: GetProducts :many
SELECT * FROM products
ORDER BY name;

-- name: ListProducts :many
SELECT * FROM products
LIMIT 5;

-- name: CreateProduct :execresult
INSERT INTO products (
    name ,description ,image, price,rating , available , category, stock
) VALUES (
  ?,?,?,?,?,?,?,?
);


-- name: UpdateProductPrice :execresult
UPDATE products SET price = ? WHERE id  = ?;

-- name: UpdateProductRating :execresult
UPDATE products SET rating = ? WHERE id  = ?;

-- name: UpdateProductAvailability :execresult
UPDATE products SET available = ? WHERE id  = ?;

-- name: DeleteProduct :execresult
DELETE FROM products
WHERE id = ?;


-- name: CreateCart :execresult
INSERT INTO carts (cart_id) VALUES (UUID());

-- name: GetCart :one
SELECT * FROM carts WHERE id = ?;



-- name: GetCartItems :many
SELECT  p.*  , ci.quantity FROM cart_items ci JOIN products p ON ci.product_id = p.id WHERE cart_id = ?;

-- name: SaveCartItems :execresult
INSERT INTO cart_items (cart_id ,product_id , quantity) VALUES (?,?, ?);

-- name: GetCartItem :one
SELECT * FROM cart_items WHERE product_id = ? AND cart_id = ?;

-- name: UpdateCartItems :execresult
UPDATE cart_items SET quantity = quantity + 1 WHERE product_id = ? AND cart_id = ?;



-- name: DeleteCart :execresult
DELETE FROM carts
WHERE cart_id = ?;

-- name: DeleteCartItems :execresult
DELETE FROM cart_items
WHERE cart_id = ?;






