-- name: GetAllUsers :many
SELECT idUsuario, correo, usuario, role, idAdministrador, idCliente FROM Usuario;

-- name: GetUserById :one
SELECT 
  idUsuario AS id, 
  usuario AS user, 
  correo AS email, 
  contraseña AS password, 
  role,
  img AS image,
  idCliente,
  idAdministrador
FROM Usuario
WHERE idUsuario = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO Usuario (correo, usuario, contraseña, img, idAdministrador, idCliente, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, now(), now());

-- name: UpdateUser :execresult
UPDATE Usuario
SET usuario = ?, correo = ?
WHERE idUsuario = ?;

-- name: GetUserByUserName :one
SELECT 
  idUsuario AS id,
  usuario AS user,
  correo AS email,
  contraseña AS password,
  role,
  img AS image,
  created_at,
  updated_at
FROM Usuario
WHERE usuario = ?;

-- name: UpdateUserPassword :execresult
UPDATE Usuario SET contraseña = ? 
WHERE idUsuario = ?;

-- name: GetUserByEmail :one
SELECT * FROM Usuario WHERE correo=? LIMIT 1;

-- name: DeleteUser :execresult
DELETE FROM Usuario WHERE idUsuario = ?;