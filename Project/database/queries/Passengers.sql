-- name: GetAllPassengers :many
SELECT * FROM Pasajeros;

-- name: GetPassengersById :one
SELECT * FROM Pasajeros WHERE idPasajeros = ? LIMIT 1;

-- name: CreatePassenger :execresult
INSERT INTO Pasajeros (nombre, edad, idDetalle)
VALUES (?, ?, ?);

-- name: UpdatePassenger :exec
UPDATE Pasajeros
SET nombre = ?, edad = ?, idDetalle = ?
WHERE idPasajeros = ?;

-- name: DeletePassenger :exec
DELETE FROM Pasajeros WHERE idPasajeros = ?;