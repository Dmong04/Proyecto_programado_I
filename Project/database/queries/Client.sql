-- name: GetAllClients :many
SELECT idCliente, nombre FROM Cliente;

-- name: GetClientById :one
SELECT idCliente, nombre
FROM Cliente WHERE idCliente = ? LIMIT 1;

-- name: GetClientByName :one
SELECT idCliente, nombre 
FROM Cliente WHERE nombre = ? LIMIT 1;

-- name: CreateClient :execresult
INSERT INTO Cliente (nombre)
VALUES (?);

-- name: UpdateClient :execresult
UPDATE Cliente
SET nombre = ?
WHERE idCliente = ?;

-- name: DeleteClient :execresult
DELETE FROM Cliente WHERE idCliente = ?;

-- name: DeleteClientByName :execresult
DELETE FROM Cliente WHERE nombre = ?;