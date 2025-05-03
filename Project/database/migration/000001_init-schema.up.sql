-- Desactivar restricciones temporales
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ENGINE_SUBSTITUTION';

-- Crear y usar base de datos
CREATE SCHEMA IF NOT EXISTS `coco_tours_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `coco_tours_db`;

-- Tabla: Administrador
CREATE TABLE IF NOT EXISTS `Administrador` (
  `idAdministrador` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `usuario` VARCHAR(40) NOT NULL,
  `contraseña` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`idAdministrador`)
) ENGINE=InnoDB;

-- Tabla: Cliente
CREATE TABLE IF NOT EXISTS `Cliente` (
  `idCliente` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `correo` VARCHAR(100) NOT NULL,
  `usuario` VARCHAR(40) NOT NULL,
  `contraseña` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`idCliente`)
) ENGINE=InnoDB;

-- Tabla: telefonoClientes
CREATE TABLE IF NOT EXISTS `telefonoClientes` (
  `idtelefonoClientes` INT NOT NULL AUTO_INCREMENT,
  `numero` VARCHAR(25) NOT NULL,
  `tipo` VARCHAR(40) NOT NULL,
  `idCliente` INT NOT NULL,
  PRIMARY KEY (`idtelefonoClientes`),
  INDEX `idx_cliente_telefono` (`idCliente`),
  CONSTRAINT `FK_telefono_cliente`
    FOREIGN KEY (`idCliente`)
    REFERENCES `Cliente` (`idCliente`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Tabla: Viaje
CREATE TABLE IF NOT EXISTS `Viaje` (
  `idViaje` INT NOT NULL AUTO_INCREMENT,
  `tipoViaje` VARCHAR(15) NOT NULL,
  PRIMARY KEY (`idViaje`)
) ENGINE=InnoDB;

-- Tabla: detalleViaje
CREATE TABLE IF NOT EXISTS `detalleViaje` (
  `iddetalleViaje` INT NOT NULL AUTO_INCREMENT,
  `tipoViaje` VARCHAR(100) NOT NULL,
  `fecha` DATE NOT NULL,
  `hora` TIME NOT NULL,
  `idProveedor` INT DEFAULT NULL,
  `idViaje` INT NOT NULL,
  PRIMARY KEY (`iddetalleViaje`),
  INDEX `idx_proveedor` (`idProveedor`),
  INDEX `idx_viaje` (`idViaje`),
  CONSTRAINT `FK_detalle_viaje_viaje`
    FOREIGN KEY (`idViaje`)
    REFERENCES `Viaje` (`idViaje`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_detalle_viaje_proveedor`
    FOREIGN KEY (`idProveedor`)
    REFERENCES `Proveedor` (`idProveedor`)
    ON DELETE SET NULL
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Tabla: Proveedor
CREATE TABLE IF NOT EXISTS `Proveedor` (
  `idProveedor` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `descrip` VARCHAR(70) NOT NULL,
  PRIMARY KEY (`idProveedor`)
) ENGINE=InnoDB;

-- Tabla: reservas
CREATE TABLE IF NOT EXISTS `reservas` (
  `idreservas` INT NOT NULL AUTO_INCREMENT,
  `idCliente` INT NOT NULL,
  `idAdministrador` INT NOT NULL,
  `idDetalle` INT NOT NULL,
  PRIMARY KEY (`idreservas`),
  INDEX `idx_reserva_cliente` (`idCliente`),
  INDEX `idx_reserva_admin` (`idAdministrador`),
  INDEX `idx_reserva_detalle` (`idDetalle`),
  CONSTRAINT `FK_reserva_cliente`
    FOREIGN KEY (`idCliente`)
    REFERENCES `Cliente` (`idCliente`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_reserva_admin`
    FOREIGN KEY (`idAdministrador`)
    REFERENCES `Administrador` (`idAdministrador`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_reserva_detalle`
    FOREIGN KEY (`idDetalle`)
    REFERENCES `detalleViaje` (`iddetalleViaje`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Tabla: Historial
CREATE TABLE IF NOT EXISTS `Historial` (
  `idHistorial` INT NOT NULL AUTO_INCREMENT,
  `descrip` VARCHAR(70) NOT NULL,
  `idCliente` INT NOT NULL,
  `idReserva` INT NOT NULL,
  PRIMARY KEY (`idHistorial`),
  INDEX `idx_historial_cliente` (`idCliente`),
  INDEX `idx_historial_reserva` (`idReserva`),
  CONSTRAINT `FK_historial_cliente`
    FOREIGN KEY (`idCliente`)
    REFERENCES `Cliente` (`idCliente`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_historial_reserva`
    FOREIGN KEY (`idReserva`)
    REFERENCES `reservas` (`idreservas`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Tabla: Pasajeros
CREATE TABLE IF NOT EXISTS `Pasajeros` (
  `idPasajeros` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `edad` INT NOT NULL,
  `idDetalle` INT NOT NULL,
  PRIMARY KEY (`idPasajeros`),
  INDEX `idx_pasajero_detalle` (`idDetalle`),
  CONSTRAINT `FK_pasajero_detalle`
    FOREIGN KEY (`idDetalle`)
    REFERENCES `detalleViaje` (`iddetalleViaje`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Restaurar configuración original
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
