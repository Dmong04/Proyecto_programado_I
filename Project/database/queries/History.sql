-- name: GetAllHistoriales :many
SELECT * FROM Historial;

-- name: GetHistorialById :one
SELECT * FROM Historial WHERE idHistorial = ? LIMIT 1;

-- name: CreateHistorial :execresult
INSERT INTO Historial (descrip, idCliente, idReserva)
VALUES (?, ?, ?);

-- name: UpdateHistorial :exec
UPDATE Historial
SET descrip = ?, idCliente = ?, idReserva = ?
WHERE idHistorial = ?;

-- name: DeleteHistorial :exec
DELETE FROM Historial WHERE idHistorial = ?;