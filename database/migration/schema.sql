CREATE TABLE IF NOT EXISTS products (id INT PRIMARY KEY AUTO_INCREMENT,
                                    name VARCHAR(50) , description VARCHAR(1000) ,
                                    image VARCHAR(1000) , price DECIMAL(10,2),
                                    rating DECIMAL(3,2) , category VARCHAR(30) , 
                                    stock INT , available BOOLEAN DEFAULT TRUE);

