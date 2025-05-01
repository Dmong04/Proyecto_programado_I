SELECT * FROM detalleViaje;

SELECT * FROM detalleViaje WHERE iddetalleViaje = ?;

INSERT INTO detalleViaje (tipoViaje, idProveedor, idViaje)
VALUES (?, ?, ?);

UPDATE detalleViaje
SET tipoViaje = ?, idProveedor = ?, idViaje = ?
WHERE iddetalleViaje = ?;

DELETE FROM detalleViaje WHERE iddetalleViaje = ?;