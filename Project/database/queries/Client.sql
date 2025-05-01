-- name: GetAllClientes :many
SELECT * FROM Cliente;

-- name: GetClienteById :one
SELECT * FROM Cliente WHERE idCliente = ? LIMIT 1;

-- name: CreateCliente :execresult
INSERT INTO Cliente (nombre, correo, usuario, contraseña)
VALUES (?, ?, ?, ?);

-- name: UpdateCliente :exec
UPDATE Cliente
SET nombre = ?, correo = ?, usuario = ?, contraseña = ?
WHERE idCliente = ?;

-- name: DeleteCliente :exec
DELETE FROM Cliente WHERE idCliente = ?;