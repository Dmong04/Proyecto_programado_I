-- name: GetAllAdmins :many
SELECT * FROM Administrador;

-- name: GetAdminById :one
SELECT * FROM Administrador WHERE idAdministrador = ? LIMIT 1;

-- name: CreateAdmin :execresult
INSERT INTO Administrador (nombre, usuario, contraseña)
VALUES (?, ?, ?);

-- name: UpdateAdmin :exec
UPDATE Administrador
SET nombre = ?, usuario = ? WHERE idAdministrador = ?;

-- name: UpdateAdminPassword :exec
UPDATE Administrador SET contraseña = ? 
WHERE idAdministrador = ?;

-- name: DeleteAdmin :exec
DELETE FROM Administrador WHERE idAdministrador = ?;