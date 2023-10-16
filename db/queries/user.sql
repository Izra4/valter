-- name: AddNewUser :execresult
INSERT INTO user(userName,email,number,address,password,token)
    VALUES (?,?,?,?,?,'');

-- name: GetAllUsers :many
SELECT * FROM user
    ORDER BY id asc;

-- name: GetUserbyId :one
SELECT * FROM user
    Where id = ?;

-- name: GetUserbyUsername :one
SELECT * FROM user
Where userName = ?;

-- name: GetUserbyEmail :one
SELECT * FROM user
Where email = ?;

-- name: UpdateUser :exec
UPDATE user
    SET userName = ?, email = ?, number = ?, address = ?
    WHERE id = ?;

-- name: ForgotPass :exec
UPDATE user
    SET password = ? WHERE id = ?;

-- name: SetToken :exec
UPDATE user
    SET token = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE from user
    WHERE id = ?;