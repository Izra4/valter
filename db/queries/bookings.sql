-- name: CreateBooking :execresult
INSERT INTO bookings(fname,lname,job,email,phone,country,address,message,productId)
    VALUES (?,?,?,?,?,?,?,?,?)