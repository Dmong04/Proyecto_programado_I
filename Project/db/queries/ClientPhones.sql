-- name: GetAllTelefonoClientes :many
SELECT * FROM telefonoClientes;

-- name: GetTelefonoClienteById :one
SELECT * FROM telefonoClientes WHERE idtelefonoClientes = ? LIMIT 1;

-- name: CreateTelefonoCliente :execresult
INSERT INTO telefonoClientes (numero, tipo, idCliente)
VALUES (?, ?, ?);

-- name: UpdateTelefonoCliente :exec
UPDATE telefonoClientes
SET numero = ?, tipo = ?, idCliente = ?
WHERE idtelefonoClientes = ?;

-- name: DeleteTelefonoCliente :exec
DELETE FROM telefonoClientes WHERE idtelefonoClientes = ?;