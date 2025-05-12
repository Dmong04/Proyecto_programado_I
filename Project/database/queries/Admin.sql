-- name: GetAllAdmins :many
SELECT idAdministrador, nombre, correo, usuario FROM Administrador;

-- name: GetAdminById :one
SELECT idAdministrador, nombre, correo, usuario
 FROM Administrador WHERE idAdministrador = ? LIMIT 1;

-- name: GetAdminByName :one
SELECT idAdministrador, nombre, correo, usuario 
FROM Administrador WHERE nombre = ? LIMIT 1;

-- name: CreateAdmin :execresult
INSERT INTO Administrador (nombre, correo, usuario, contraseña)
VALUES (?, ?, ?, ?);

-- name: UpdateAdmin :execresult
UPDATE Administrador
SET nombre = ?, correo = ?, usuario = ? WHERE idAdministrador = ?;

-- name: UpdateAdminPassword :execresult
UPDATE Administrador SET contraseña = ? 
WHERE idAdministrador = ?;

-- name: DeleteAdmin :execresult
DELETE FROM Administrador WHERE idAdministrador = ?;

-- name: DeleteAdminByName :execresult
DELETE FROM Administrador WHERE nombre = ?;

-- name: GetAdminByUser :one
SELECT * FROM Administrador
WHERE usuario = ? LIMIT 1;