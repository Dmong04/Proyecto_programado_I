-- name: GetAllPassengers :many
SELECT * FROM Pasajeros;

-- name: GetPassengersById :one
SELECT * FROM Pasajeros WHERE idPasajeros = ? LIMIT 1;

-- name: GetPassengersByName :one
SELECT * FROM Pasajeros WHERE nombre = ? LIMIT 1;

-- name: GetPassengersByDetailID :many
SELECT * FROM Pasajeros WHERE idDetalle = ?;

-- name: CreatePassenger :execresult
INSERT INTO Pasajeros (nombre, edad, idDetalle)
VALUES (?, ?, ?);

-- name: UpdatePassenger :exec
UPDATE Pasajeros
SET nombre = ?, edad = ?, idDetalle = ?
WHERE idPasajeros = ?;

-- name: DeletePassenger :exec
DELETE FROM Pasajeros WHERE idPasajeros = ?;

-- name: DeletePassengerByName :exec
DELETE FROM Pasajeros WHERE nombre = ?;