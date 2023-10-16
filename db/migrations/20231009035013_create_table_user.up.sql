    CREATE TABLE user(
        id int unsigned NOT NULL AUTO_INCREMENT,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deletedAt TIMESTAMP DEFAULT NULL,
        userName varchar(20) NOT NULL UNIQUE ,
        email varchar(320) NOT NULL UNIQUE ,
        number varchar(20) NOT NULL,
        address varchar(255) NOT NULL,
        password varchar(100) NOT NULL,
        token varchar(10) NOT NULL,
        PRIMARY KEY(id)
    );