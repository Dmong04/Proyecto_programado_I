-- name: GetAllTravels :many
SELECT * FROM Viaje;

-- name: GetTravelById :one
SELECT * FROM Viaje WHERE idViaje = ? LIMIT 1;

-- name: CreateTravel :execresult
INSERT INTO Viaje (tipoViaje, descripcion)
VALUES (?, ?);

-- name: UpdateTravel :exec
UPDATE Viaje
SET tipoViaje = ?, descripcion = ?
WHERE idViaje = ?;

-- name: DeleteTravel :exec
DELETE FROM Viaje WHERE idViaje = ?;