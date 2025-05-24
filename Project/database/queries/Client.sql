-- name: GetAllClients :many
SELECT * FROM Cliente;

-- name: GetClientById :one
SELECT * FROM Cliente WHERE idCliente = ? LIMIT 1;

-- name: GetClientByName :one
SELECT * FROM Cliente WHERE nombre = ? LIMIT 1;

-- name: CreateClient :execresult
INSERT INTO Cliente (nombre, correo)
VALUES (?, ?);

-- name: UpdateClient :execresult
UPDATE Cliente
SET nombre = ?, correo = ?
WHERE idCliente = ?;

-- name: DeleteClient :execresult
DELETE FROM Cliente WHERE idCliente = ?;