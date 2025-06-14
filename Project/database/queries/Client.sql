-- name: GetAllClients :many
SELECT idCliente, nombre, telefono FROM Cliente;

-- name: GetClientById :one
SELECT idCliente, nombre, telefono
FROM Cliente WHERE idCliente = ? LIMIT 1;

-- name: GetClientByName :one
SELECT idCliente, nombre, telefono 
FROM Cliente WHERE nombre = ? LIMIT 1;

-- name: CreateClient :execresult
INSERT INTO Cliente (nombre, telefono)
VALUES (?,?);

-- name: UpdateClient :execresult
UPDATE Cliente
SET nombre = ?, telefono = ?
WHERE idCliente = ?;

-- name: DeleteClient :execresult
DELETE FROM Cliente WHERE idCliente = ?;

-- name: DeleteClientByName :execresult
DELETE FROM Cliente WHERE nombre = ?;