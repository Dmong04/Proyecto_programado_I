-- name: GetAllPasajeros :many
SELECT * FROM Pasajeros;

-- name: GetPasajeroById :one
SELECT * FROM Pasajeros WHERE idPasajeros = ? LIMIT 1;

-- name: CreatePasajero :execresult
INSERT INTO Pasajeros (nombre, edad, idDetalle)
VALUES (?, ?, ?);

-- name: UpdatePasajero :exec
UPDATE Pasajeros
SET nombre = ?, edad = ?, idDetalle = ?
WHERE idPasajeros = ?;

-- name: DeletePasajero :exec
DELETE FROM Pasajeros WHERE idPasajeros = ?;