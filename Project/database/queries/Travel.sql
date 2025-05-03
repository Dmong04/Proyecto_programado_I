-- name: GetAllTravels :many
SELECT * FROM Viaje;

-- name: GetTravelById :one
SELECT * FROM Viaje WHERE idViaje = ? LIMIT 1;

-- name: CreateTravel :execresult
INSERT INTO Viaje (tipoViaje)
VALUES (?);

-- name: UpdateTravel :exec
UPDATE Viaje
SET tipoViaje = ?
WHERE idViaje = ?;

-- name: DeleteTravel :exec
DELETE FROM Viaje WHERE idViaje = ?;