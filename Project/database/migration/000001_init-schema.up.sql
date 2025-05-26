-- Desactivar restricciones temporales
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ENGINE_SUBSTITUTION';

-- Usar base de datos
USE `coco_tours_db_v2`;

-- Tabla Cliente
CREATE TABLE Cliente (
  idCliente INT AUTO_INCREMENT PRIMARY KEY
  nombre VARCHAR(50) NOT NULL
);

-- Tabla telefonosClientes
CREATE TABLE telefonosClientes (
  idtelefonosClientes INT AUTO_INCREMENT PRIMARY KEY,
  idCliente INT NOT NULL,
  telefono VARCHAR(20) NOT NULL,
  FOREIGN KEY (idCliente) REFERENCES Cliente(idCliente)
);

-- Tabla Administrador
CREATE TABLE Administrador (
  idAdministrador INT AUTO_INCREMENT PRIMARY KEY
  nombre VARCHAR(50) NOT NULL
);

-- Tabla Usuario
CREATE TABLE Usuario (
  idUsuario INT AUTO_INCREMENT PRIMARY KEY,
  correo VARCHAR(70) UNIQUE NOT NULL,
  usuario VARCHAR(30) UNIQUE NOT NULL,
  contraseña VARCHAR(15) NOT NULL,
  idCliente INT DEFAULT NULL,
  idAdmin INT DEFAULT NULL,
  FOREIGN KEY (idCliente) REFERENCES Cliente(idCliente),
  FOREIGN KEY (idAdmin) REFERENCES Administrador(idAdministrador)
);

-- Tabla Extra
CREATE TABLE Extra (
  idExtra INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(40) NOT NULL,
  descrip VARCHAR(50) NOT NULL,
  precioPersona DECIMAL NOT NULL
);

-- Tabla detalleExtra
CREATE TABLE detalleExtra (
  idDetalleExtra INT AUTO_INCREMENT PRIMARY KEY,
  cantPersona INT NOT NULL,
  -- precio se omitirá si no se define la expresión
  idExtra INT NOT NULL,
  FOREIGN KEY (idExtra) REFERENCES Extra(idExtra)
);

-- Tabla Proveedor
CREATE TABLE Proveedor (
  idProveedor INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  descrip VARCHAR(60) NOT NULL,
  correo VARCHAR(70) NOT NULL
);

-- Tabla telefonosProveedor
CREATE TABLE telefonosProveedor (
  idTelefonoProovedor INT AUTO_INCREMENT PRIMARY KEY,
  telefono VARCHAR(20) NOT NULL,
  idProveedor INT NOT NULL,
  FOREIGN KEY (idProveedor) REFERENCES Proveedor(idProveedor)
);

-- Tabla Viaje
CREATE TABLE Viaje (
  idViaje INT AUTO_INCREMENT PRIMARY KEY,
  tipo VARCHAR(30) NOT NULL,
  descrip VARCHAR(60) NOT NULL,
  origen VARCHAR(40) NOT NULL,
  destino VARCHAR(40) NOT NULL
);

-- Tabla DetalleViaje
CREATE TABLE DetalleViaje (
  idDetalleViaje INT AUTO_INCREMENT PRIMARY KEY,
  numPasajeros VARCHAR(4) NOT NULL,
  precio DECIMAL NOT NULL,
  idViaje INT NOT NULL,
  idProveedor INT NOT NULL,
  FOREIGN KEY (idViaje) REFERENCES Viaje(idViaje),
  FOREIGN KEY (idProveedor) REFERENCES Proveedor(idProveedor)
);

-- Tabla Reserva
CREATE TABLE Reserva (
  idReserva INT AUTO_INCREMENT PRIMARY KEY,
  fecha VARCHAR(20) NOT NULL,
  hora VARCHAR(20) NOT NULL,
  descrip VARCHAR(60) NOT NULL,
  idExtra INT,
  idViaje INT NOT NULL,
  idUsuario INT NOT NULL,
  FOREIGN KEY (idExtra) REFERENCES detalleExtra(idDetalleExtra),
  FOREIGN KEY (idViaje) REFERENCES DetalleViaje(idDetalleViaje),
  FOREIGN KEY (idUsuario) REFERENCES Usuario(idUsuario)
);

-- Restaurar configuración original
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
