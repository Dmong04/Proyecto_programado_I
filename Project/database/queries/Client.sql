-- name: GetAllClients :many
SELECT idCliente, nombre, correo, usuario FROM Cliente;

-- name: GetClientById :one
SELECT idCliente, nombre, correo, usuario 
FROM Cliente WHERE idCliente = ? LIMIT 1;

-- name: GetClientByName :one
SELECT idCliente, nombre, correo, usuario 
FROM Cliente WHERE nombre = ? LIMIT 1;

-- name: CreateClient :execresult
INSERT INTO Cliente (nombre, correo, usuario, contraseña)
VALUES (?, ?, ?, ?);

-- name: UpdateClient :execresult
UPDATE Cliente
SET nombre = ?, correo = ?, usuario = ?
WHERE idCliente = ?;

-- name: UpdateClientPassword :execresult
UPDATE Cliente SET contraseña = ? 
WHERE idCliente = ?;

-- name: DeleteClient :execresult
DELETE FROM Cliente WHERE idCliente = ?;

-- name: DeleteClientByName :execresult
DELETE FROM Cliente WHERE nombre = ?;

-- name: GetClientByUser :one
SELECT * FROM Cliente
WHERE usuario = ? LIMIT 1;