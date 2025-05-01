SELECT * FROM Pasajeros;

SELECT * FROM Pasajeros WHERE idPasajeros = ?;

INSERT INTO Pasajeros (nombre, edad, idDetalle)
VALUES (?, ?, ?);

UPDATE Pasajeros
SET nombre = ?, edad = ?, idDetalle = ?
WHERE idPasajeros = ?;

DELETE FROM Pasajeros WHERE idPasajeros = ?;