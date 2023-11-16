CREATE TABLE users(
    id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(30) NOT NULL,
    phone VARCHAR(15) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE products (
            id INT  NOT NULL , 
            name VARCHAR(50) NOT NULL ,
            price FLOAT NOT NULL ,
            description VARCHAR(1000) NOT NULL,
            image VARCHAR(1000) NOT NULL,
            available BOOLEAN DEFAULT TRUE NOT NULL, 
            stock INT NOT NULL,
            category VARCHAR(20) NOT NULL,
            rating FLOAT NOT NULL,
            PRIMARY KEY (id) 
            );

CREATE TABLE cart_items(
    id INT AUTO_INCREMENT NOT NULL,
    product_id FLOAT NOT NULL ,
    quantity INT NOT NULL ,
    PRIMARY KEY (id),
    FOREIGN KEY(product_id) REFERENCES products(id)

);

CREATE TABLE carts (
    id INT AUTO_INCREMENT NOT NULL,
    items INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(items) REFERENCES cart_items(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)