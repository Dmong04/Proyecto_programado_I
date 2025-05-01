SELECT * FROM Viaje;

SELECT * FROM Viaje WHERE idViaje = ?;

INSERT INTO Viaje (fecha, hora)
VALUES (?, ?);

UPDATE Viaje
SET fecha = ?, hora = ?
WHERE idViaje = ?;

DELETE FROM Viaje WHERE idViaje = ?;