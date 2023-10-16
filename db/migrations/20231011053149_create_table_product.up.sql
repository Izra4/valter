CREATE TABLE product(
    id int unsigned NOT NULL AUTO_INCREMENT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deletedAt TIMESTAMP DEFAULT NULL,
    name varchar(100) NOT NULL,
    description TEXT NOT NULL,
    link varchar(255) NOT NULL,
    book_pict varchar(255) NOT NULL,
    PRIMARY KEY (id)
);