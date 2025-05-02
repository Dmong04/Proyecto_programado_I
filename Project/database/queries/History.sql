-- name: GetAllHistories :many
SELECT * FROM Historial;

-- name: GetHistoryById :one
SELECT * FROM Historial WHERE idHistorial = ? LIMIT 1;

-- name: CreateHistory :execresult
INSERT INTO Historial (descrip, idCliente, idReserva)
VALUES (?, ?, ?);

-- name: UpdateHistory :exec
UPDATE Historial
SET descrip = ?, idCliente = ?, idReserva = ?
WHERE idHistorial = ?;

-- name: DeleteHistory :exec
DELETE FROM Historial WHERE idHistorial = ?;