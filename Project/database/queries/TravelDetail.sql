-- name: GetAllTravelDetails :many
SELECT * FROM detalleViaje;

-- name: GetTravelDetailById :one
SELECT * FROM detalleViaje WHERE iddetalleViaje = ? LIMIT 1;

-- name: CreateTravelDetail :execresult
INSERT INTO detalleViaje (tipoViaje, idProveedor, idViaje)
VALUES (?, ?, ?);

-- name: UpdateTravelDetail :exec
UPDATE detalleViaje
SET tipoViaje = ?, idProveedor = ?, idViaje = ?
WHERE iddetalleViaje = ?;

-- name: DeleteTravelDetail :exec
DELETE FROM detalleViaje WHERE iddetalleViaje = ?;