-- name: GetAllUsers :many
SELECT idUsuario, nombre, correo, usuario FROM Usuario;

-- name: GetUserById :one
SELECT idUsuario, nombre, correo, usuario
 FROM Usuario WHERE idUsuario = ? LIMIT 1;

-- name: GetUserByName :one
SELECT idUsuario, nombre, correo, usuario 
FROM Usuario WHERE nombre = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO Usuario (nombre, correo, usuario, contraseña)
VALUES (?, ?, ?, ?);

-- name: UpdateUser :execresult
UPDATE Usuario
SET correo = ?, SET usuario = ?, SET contraseña = ?,
WHERE idUsuario = ?;

-- name: UpdateUserPassword :execresult
UPDATE Usuario SET contraseña = ? 
WHERE idUsuario = ?;

-- name: DeleteUser :execresult
DELETE FROM Usuario WHERE idUsuario = ?;

-- name: DeleteUserByName :execresult
DELETE FROM Usuario WHERE nombre = ?;