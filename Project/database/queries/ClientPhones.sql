SELECT * FROM telefonoClientes;

SELECT * FROM telefonoClientes WHERE idtelefonoClientes = ?;

INSERT INTO telefonoClientes (numero, tipo, idCliente)
VALUES (?, ?, ?);

UPDATE telefonoClientes
SET numero = ?, tipo = ?, idCliente = ?
WHERE idtelefonoClientes = ?;

DELETE FROM telefonoClientes WHERE idtelefonoClientes = ?;