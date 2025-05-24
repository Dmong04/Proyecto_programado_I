-- name: GetAllAdmins :many
SELECT * FROM Administrador;

-- name: GetAdminById :one
SELECT * FROM Administrador WHERE idAdministrador = ? LIMIT 1;

-- name: GetAdminByName :one
SELECT * FROM Administrador WHERE nombre = ? LIMIT 1;

-- name: CreateAdmin :execresult
INSERT INTO Administrador (nombre, correo)
VALUES (?, ?);

-- name: UpdateAdmin :execresult
UPDATE Administrador
SET nombre = ?, correo = ? 
WHERE idAdministrador = ?;

-- name: DeleteAdmin :execresult
DELETE FROM Administrador WHERE idAdministrador = ?;