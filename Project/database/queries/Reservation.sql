-- name: GetAllReservas :many
SELECT * FROM reservas;

-- name: GetReservaById :one
SELECT * FROM reservas WHERE idreservas = ? LIMIT 1;

-- name: CreateReserva :execresult
INSERT INTO reservas (idCliente, idAdministrador, idDetalle)
VALUES (?, ?, ?);

-- name: UpdateReserva :exec
UPDATE reservas
SET idCliente = ?, idAdministrador = ?, idDetalle = ?
WHERE idreservas = ?;

-- name: DeleteReserva :exec
DELETE FROM reservas WHERE idreservas = ?;