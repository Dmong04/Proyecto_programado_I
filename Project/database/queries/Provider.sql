-- name: GetAllProveedores :many
SELECT * FROM Proveedor;

-- name: GetProveedorById :one
SELECT * FROM Proveedor WHERE idProveedor = ? LIMIT 1;

-- name: CreateProveedor :execresult
INSERT INTO Proveedor (nombre, descrip)
VALUES (?, ?);

-- name: UpdateProveedor :exec
UPDATE Proveedor
SET nombre = ?, descrip = ?
WHERE idProveedor = ?;

-- name: DeleteProveedor :exec
DELETE FROM Proveedor WHERE idProveedor = ?;