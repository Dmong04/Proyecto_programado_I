-- name: GetAllReservations :many
SELECT * FROM reservas;

-- name: GetReservationsById :one
SELECT * FROM reservas WHERE idreservas = ? LIMIT 1;

-- name: CreateReservation :execresult
INSERT INTO reservas (idCliente, idAdministrador, idDetalle)
VALUES (?, ?, ?, ?,);

-- name: UpdateReservation :exec
UPDATE reservas
SET idCliente = ?, idAdministrador = ?, idDetalle = ?
WHERE idreservas = ?;

-- name: DeleteReservation :exec
DELETE FROM reservas WHERE idreservas = ?;