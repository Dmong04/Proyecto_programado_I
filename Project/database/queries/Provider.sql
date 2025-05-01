-- name: GetAllProviders :many
SELECT * FROM Proveedor;

-- name: GetProviderById :one
SELECT * FROM Proveedor WHERE idProveedor = ? LIMIT 1;

-- name: CreateProvider :execresult
INSERT INTO Proveedor (nombre, descrip)
VALUES (?, ?);

-- name: UpdateProvider :exec
UPDATE Proveedor
SET nombre = ?, descrip = ?
WHERE idProveedor = ?;

-- name: DeleteProvider :exec
DELETE FROM Proveedor WHERE idProveedor = ?;