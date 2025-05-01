SELECT * FROM reservas;

SELECT * FROM reservas WHERE idreservas = ?;

INSERT INTO reservas (idCliente, idAdministrador, idDetalle)
VALUES (?, ?, ?);

UPDATE reservas
SET idCliente = ?, idAdministrador = ?, idDetalle = ?
WHERE idreservas = ?;

DELETE FROM reservas WHERE idreservas = ?;