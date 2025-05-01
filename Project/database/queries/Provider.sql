SELECT * FROM Proveedor;

SELECT * FROM Proveedor WHERE idProveedor = ?;

INSERT INTO Proveedor (nombre, descrip)
VALUES (?, ?);

UPDATE Proveedor
SET nombre = ?, descrip = ?
WHERE idProveedor = ?;

DELETE FROM Proveedor WHERE idProveedor = ?;