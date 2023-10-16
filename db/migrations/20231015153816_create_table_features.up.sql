CREATE TABLE features(
    id int unsigned NOT NULL AUTO_INCREMENT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deletedAt TIMESTAMP DEFAULT NULL,
    title varchar(20) NOT NULL,
    description TEXT NOT NULL,
    productId int unsigned NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_product_features
        FOREIGN KEY (productId) REFERENCES product(id)
)