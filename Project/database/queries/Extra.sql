-- name: GetAllExtras :many
SELECT * FROM Extras;

-- name: GetExtraByName :one
SELECT * FROM Extras WHERE nombre = ? LIMIT 1;

-- name: CreateExtra :execresult
INSERT INTO Extras (nombre, descripcion, precio) 
VALUES (?, ?, ?);

-- name: UpdateExtra :execresult
Update Extras 
SET nombre = ?, descripcion = ?, precio
WHERE idExtra = ?;

-- name: DeleteExtra :execresult
DELETE * FROM Extras WHERE idExtra = ?;