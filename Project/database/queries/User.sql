-- name: GetAllUsers :many
SELECT idUsuario, correo, usuario, role FROM Usuario;

-- name: GetUserById :one
SELECT idUsuario, correo, usuario, role
 FROM Usuario WHERE idUsuario = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO Usuario (correo, usuario, contraseña, role, created_at, updated_at)
VALUES (?, ?, ?, ?, now(), now());

-- name: UpdateUser :execresult
UPDATE Usuario
SET correo = ?, usuario = ?, contraseña = ?, role = ?, updated_at = now()
WHERE idUsuario = ?;

-- name: GetUserByUserName :one
SELECT idUsuario AS id, usuario AS user, correo AS email, contraseña AS password, role
FROM Usuario
WHERE usuario = ? LIMIT 1;


-- name: UpdateUserRole :execresult
UPDATE Usuario set role=? WHERE idUsuario=?;

-- name: UpdateUserPassword :execresult
UPDATE Usuario SET contraseña = ? 
WHERE idUsuario = ?;

-- name: GetUserByEmail :one
SELECT * FROM Usuario WHERE correo=? LIMIT 1;

-- name: DeleteUser :execresult
DELETE FROM Usuario WHERE idUsuario = ?;