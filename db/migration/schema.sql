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





CREATE TABLE carts (
    id INT NOT NULL,
    cart_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
    PRIMARY KEY (id)
    
);



CREATE TABLE cart_items (
    id INT NOT NULL,
    cart_id VARCHAR(36) NOT NULL,
    product_id  INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (cart_id) REFERENCES carts(cart_id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
