-- name: GetAllReservations :many
SELECT * FROM reservas;

-- name: GetReservationsById :one
SELECT * FROM reservas WHERE idreservas = ? LIMIT 1;

-- name: CreateReservation :execresult
INSERT INTO reservas (idUsuario, idDetalle, estado)
VALUES (?, ?, ?);

-- name: UpdateReservation :exec
UPDATE reservas
SET idUsuario = ?, idDetalle = ?
WHERE idreservas = ?;

-- name: UpdateStatus :exec
UPDATE reservas
SET estado = ?
WHERE idreservas = ?;

-- name: DeleteReservation :exec
DELETE FROM reservas WHERE idreservas = ?;