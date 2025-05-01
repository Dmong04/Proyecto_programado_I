SELECT * FROM Historial;

SELECT * FROM Historial WHERE idHistorial = ?;

INSERT INTO Historial (descrip, idCliente, idReserva)
VALUES (?, ?, ?);

UPDATE Historial
SET descrip = ?, idCliente = ?, idReserva = ?
WHERE idHistorial = ?;

DELETE FROM Historial WHERE idHistorial = ?;