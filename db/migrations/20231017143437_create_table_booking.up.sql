CREATE TABLE bookings(
    id varchar(10) NOT NULL ,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deletedAt TIMESTAMP DEFAULT NULL,
    fname varchar(20) NOT NULL,
    lname varchar(20) NOT NULL,
    job varchar(20) NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(20) NOT NULL,
    country varchar(20) NOT NULL,
    address varchar(255) NOT NULL,
    message varchar(255) NOT NULL,
    productId int unsigned NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_booking_product
        FOREIGN KEY (productId) REFERENCES product(id)
)