-- name: GetAllClientPhones :many
SELECT * FROM telefonoClientes;

-- name: GetClientPhonesById :one
SELECT * FROM telefonoClientes WHERE idtelefonoClientes = ? LIMIT 1;

-- name: CreateClientPhones :execresult
INSERT INTO telefonoClientes (numero, tipo, idCliente)
VALUES (?, ?, ?);

-- name: UpdateClientPhones :exec
UPDATE telefonoClientes
SET numero = ?, tipo = ?
WHERE idtelefonoClientes = ?;

-- name: DeleteClientPhones :exec
DELETE FROM telefonoClientes WHERE idtelefonoClientes = ?;