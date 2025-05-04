-- name: GetAllTravelDetails :many
SELECT * FROM detalleViaje;

-- name: GetTravelDetailById :one
SELECT * FROM detalleViaje WHERE idDetalleViaje = ? LIMIT 1;

-- name: CreateTravelDetail :execresult
INSERT INTO detalleViaje (fecha, hora, idProveedor, idViaje)
VALUES (?, ?, ?, ?);

-- name: UpdateTravelDetail :exec
UPDATE detalleViaje
SET fecha = ?, hora = ?, idProveedor = ?, idViaje = ?
WHERE idDetalleViaje = ?;

-- name: DeleteTravelDetail :exec
DELETE FROM detalleViaje WHERE idDetalleViaje = ?;