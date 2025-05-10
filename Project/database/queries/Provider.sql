-- name: GetAllProviders :many
SELECT * FROM Proveedor;

-- name: GetProviderById :one
SELECT * FROM Proveedor WHERE idProveedor = ? LIMIT 1;

-- name: GetProviderByName :one
SELECT * FROM Proveedor WHERE nombre = ? LIMIT 1;

-- name: CreateProvider :execresult
INSERT INTO Proveedor (nombre, descrip)
VALUES (?, ?);

-- name: UpdateProvider :exec
UPDATE Proveedor
SET nombre = ?, descrip = ?
WHERE idProveedor = ?;

-- name: UpdateProviderByName :exec
UPDATE Proveedor
SET nombre = ?, descrip = ?
WHERE nombre = ?;

-- name: DeleteProvider :exec
DELETE FROM Proveedor WHERE idProveedor = ?;

-- name: DeleteProviderByName :exec
DELETE FROM Proveedor WHERE nombre = ?;