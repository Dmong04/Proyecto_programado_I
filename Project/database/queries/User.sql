-- name: GetAllUsers :many
SELECT u.idUsuario, a.idAdministrador as ID, a.nombre as nombre, u.usuario as usuario, a.correo as correo, 'Admin' AS tipo
FROM Usuario as u 
inner join Administrador as a on u.admin = a.idAdministrador
UNION ALL
(SELECT u.idUsuario, c.idCliente as ID, c.nombre as nombre, u.usuario as usuario, c.correo as correo, 'Cliente' AS tipo
FROM Usuario as u 
inner join Cliente as c on u.cliente = c.idCliente);

-- name: GetUserById :one
SELECT * FROM 
(SELECT u.idUsuario, a.nombre, a.correo, u.usuario, 'Admin' AS tipo FROM Usuario as u inner join 
Administrador as a on u.admin = a.idAdministrador
UNION
(SELECT u.idUsuario, c.nombre, c.correo, u.usuario, 'Cliente' AS tipo FROM Usuario as u inner join 
Cliente as c on u.cliente = c.idCliente)) AS all_users WHERE idUsuario = ? LIMIT 1;

-- name: GetAdminByusername :one
SELECT * FROM 
(SELECT u.idUsuario, u.nombre, u.usuario, u.correo as correo, 'Admin' AS tipo FROM Usuario as u inner join 
Administrador as a on u.admin = a.idAdministrador
UNION
(SELECT u.idUsuario, c.nombre, u.usuario, c.correo as correo, 'Cliente' AS tipo FROM Usuario as u inner join 
Cliente as c on u.cliente = c.idCliente)) AS all_users WHERE usuario = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO Usuario (usuario, contraseña, admin, cliente)
VALUES (?, ?, ?, ?);

-- name: UpdateUsername :execresult
UPDATE Usuario
SET usuario = ? WHERE idUsuario = ?;

-- name: UpdatePassword :execresult
UPDATE Usuario SET contraseña = ? 
WHERE idUsuario = ?;

-- name: DeleteUser :execresult
DELETE FROM Usuario WHERE idUsuario = ?;

-- name: DeleteByUsername :execresult
DELETE FROM Usuario WHERE usuario = ?;