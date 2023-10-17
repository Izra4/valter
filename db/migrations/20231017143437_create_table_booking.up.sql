CREATE TABLE bookings(
    id varchar(10) NOT NULL AUTO_INCREMENT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deletedAt TIMESTAMP DEFAULT NULL,
    fname varchar(20) NOT NULL,
    lname varchar(20) NOT NULL,
    job varchar(20) NOT NULL,
    country varchar(20) NOT NULL,
    address varchar(255) NOT NULL,
    message varchar(255) NOT NULL,
)