-- name: GetAllAdmins :many
SELECT idAdministrador, nombre FROM Administrador;

-- name: GetAdminById :one
SELECT idAdministrador, nombre
 FROM Administrador WHERE idAdministrador = ? LIMIT 1;

-- name: GetAdminByName :one
SELECT idAdministrador, nombre 
FROM Administrador WHERE nombre = ? LIMIT 1;

-- name: CreateAdmin :execresult
INSERT INTO Administrador (nombre)
VALUES (?);

-- name: UpdateAdmin :execresult
UPDATE Administrador
SET nombre = ?
WHERE idAdministrador = ?;

-- name: DeleteAdmin :execresult
DELETE FROM Administrador WHERE idAdministrador = ?;

-- name: DeleteAdminByName :execresult
DELETE FROM Administrador WHERE nombre = ?;