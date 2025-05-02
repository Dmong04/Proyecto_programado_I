-- name: GetAllClients :many
SELECT * FROM Cliente;

-- name: GetClientById :one
SELECT * FROM Cliente WHERE idCliente = ? LIMIT 1;

-- name: CreateClient :execresult
INSERT INTO Cliente (nombre, correo, usuario, contraseña)
VALUES (?, ?, ?, ?);

-- name: UpdateClient :exec
UPDATE Cliente
SET nombre = ?, correo = ?, usuario = ?
WHERE idCliente = ?;

-- name: updateClientPassword :exec
UPDATE Cliente SET contraseña = ? 
WHERE idCliente = ?;

-- name: DeleteClient :exec
DELETE FROM Cliente WHERE idCliente = ?;