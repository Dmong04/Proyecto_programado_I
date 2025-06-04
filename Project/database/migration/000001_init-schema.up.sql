-- Desactivar restricciones temporales
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ENGINE_SUBSTITUTION';

-- Usar base de datos
USE `coco_tours_db_v2`;

CREATE TABLE IF NOT EXISTS Cliente (
  idCliente INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS telefonosClientes (
  idtelefonosClientes INT AUTO_INCREMENT PRIMARY KEY,
  idCliente INT NOT NULL,
  telefono VARCHAR(20) NOT NULL,
  FOREIGN KEY (idCliente) REFERENCES Cliente(idCliente)
);

CREATE TABLE IF NOT EXISTS Administrador (
  idAdministrador INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS Usuario (
  idUsuario INT AUTO_INCREMENT PRIMARY KEY,
  correo VARCHAR(70) UNIQUE NOT NULL,
  usuario VARCHAR(30) UNIQUE NOT NULL,
  contraseña VARCHAR(15) NOT NULL,
  idCliente INT DEFAULT NULL,
  idAdmin INT DEFAULT NULL,
  FOREIGN KEY (idCliente) REFERENCES Cliente(idCliente),
  FOREIGN KEY (idAdmin) REFERENCES Administrador(idAdministrador)
);

CREATE TABLE IF NOT EXISTS Extra (
  idExtra INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(40) NOT NULL,
  descrip VARCHAR(50) NOT NULL,
  precioPersona DECIMAL(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS detalleExtra (
  idDetalleExtra INT AUTO_INCREMENT PRIMARY KEY,
  cantPersona INT NOT NULL,
  precio DECIMAL(10,2) GENERATED ALWAYS AS (cantPersona * 1.0) VIRTUAL, -- puedes ajustar esto si tienes acceso al precioPersona
  idExtra INT NOT NULL,
  FOREIGN KEY (idExtra) REFERENCES Extra(idExtra)
);

CREATE TABLE IF NOT EXISTS Proveedor (
  idProveedor INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  descrip VARCHAR(60) NOT NULL,
  correo VARCHAR(70) NOT NULL
);

CREATE TABLE IF NOT EXISTS telefonosProveedor (
  idTelefonoProveedor INT AUTO_INCREMENT PRIMARY KEY,
  telefono VARCHAR(20) NOT NULL,
  idProveedor INT NOT NULL,
  FOREIGN KEY (idProveedor) REFERENCES Proveedor(idProveedor)
);

CREATE TABLE IF NOT EXISTS Viaje (
  idViaje INT AUTO_INCREMENT PRIMARY KEY,
  tipo VARCHAR(30) NOT NULL,
  descrip VARCHAR(60) NOT NULL,
  origen VARCHAR(40) NOT NULL,
  destino VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS DetalleViaje (
  idDetalleViaje INT AUTO_INCREMENT PRIMARY KEY,
  numPasajeros VARCHAR(4) NOT NULL,
  precio DECIMAL(10,2) NOT NULL,
  idViaje INT NOT NULL,
  idProveedor INT NOT NULL,
  FOREIGN KEY (idViaje) REFERENCES Viaje(idViaje),
  FOREIGN KEY (idProveedor) REFERENCES Proveedor(idProveedor)
);

CREATE TABLE IF NOT EXISTS Reserva (
  idReserva INT AUTO_INCREMENT PRIMARY KEY,
  fecha DATE NOT NULL,
  hora TIME NOT NULL,
  descrip VARCHAR(60) NOT NULL,
  subtotalViaje DECIMAL(10,2) GENERATED ALWAYS AS (0) VIRTUAL, -- reemplazar por expresión real
  subtotalExtra DECIMAL(10,2) GENERATED ALWAYS AS (0) VIRTUAL, -- reemplazar por expresión real
  total DECIMAL(10,2) GENERATED ALWAYS AS (subtotalViaje + subtotalExtra) VIRTUAL,
  idExtra INT DEFAULT NULL,
  idViaje INT NOT NULL,
  idUsuario INT NOT NULL,
  FOREIGN KEY (idExtra) REFERENCES detalleExtra(idDetalleExtra),
  FOREIGN KEY (idViaje) REFERENCES DetalleViaje(idDetalleViaje),
  FOREIGN KEY (idUsuario) REFERENCES Usuario(idUsuario)
);

