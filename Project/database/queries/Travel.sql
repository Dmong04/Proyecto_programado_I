-- name: GetAllViajes :many
SELECT * FROM Viaje;

-- name: GetViajeById :one
SELECT * FROM Viaje WHERE idViaje = ? LIMIT 1;

-- name: CreateViaje :execresult
INSERT INTO Viaje (fecha, hora)
VALUES (?, ?);

-- name: UpdateViaje :exec
UPDATE Viaje
SET fecha = ?, hora = ?
WHERE idViaje = ?;

-- name: DeleteViaje :exec
DELETE FROM Viaje WHERE idViaje = ?;