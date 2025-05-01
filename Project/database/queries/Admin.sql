-- name: GetAll :many
SELECT * FROM Administrador;

-- name: GetAdministradorById :one
SELECT * FROM Administrador WHERE idAdministrador = ?;

-- name: CreateAdministrador :execresult
INSERT INTO Administrador (nombre, usuario, contraseña)
VALUES (?, ?, ?);

-- name: UpdateAdministrador :exec
UPDATE Administrador
SET nombre = ?, usuario = ?, contraseña = ?
WHERE idAdministrador = ?;

-- name: DeleteAdministrador :exec
DELETE FROM Administrador WHERE idAdministrador = ?;