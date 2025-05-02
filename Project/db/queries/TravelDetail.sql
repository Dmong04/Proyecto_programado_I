-- name: GetAllDetalleViajes :many
SELECT * FROM detalleViaje;

-- name: GetDetalleViajeById :one
SELECT * FROM detalleViaje WHERE iddetalleViaje = ? LIMIT 1;

-- name: CreateDetalleViaje :execresult
INSERT INTO detalleViaje (tipoViaje, idProveedor, idViaje)
VALUES (?, ?, ?);

-- name: UpdateDetalleViaje :exec
UPDATE detalleViaje
SET tipoViaje = ?, idProveedor = ?, idViaje = ?
WHERE iddetalleViaje = ?;

-- name: DeleteDetalleViaje :exec
DELETE FROM detalleViaje WHERE iddetalleViaje = ?;