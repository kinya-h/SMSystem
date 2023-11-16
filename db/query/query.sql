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


