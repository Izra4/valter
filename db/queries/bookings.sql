-- name: CreateBooking :execresult
INSERT INTO bookings(id,fname,lname,job,email,phone,country,address,message,productId)
    VALUES (?,?,?,?,?,?,?,?,?,?);

-- name: ShowAllBookings :many
select * from bookings;
