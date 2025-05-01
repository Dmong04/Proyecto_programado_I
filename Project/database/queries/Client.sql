SELECT * FROM Cliente;

SELECT * FROM Cliente WHERE idCliente = ?;

INSERT INTO Cliente (nombre, correo, usuario, contraseña)
VALUES (?, ?, ?, ?);

UPDATE Cliente
SET nombre = ?, correo = ?, usuario = ?, contraseña = ?
WHERE idCliente = ?;

DELETE FROM Cliente WHERE idCliente = ?;