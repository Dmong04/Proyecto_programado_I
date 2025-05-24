-- Desactivar restricciones temporales
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ENGINE_SUBSTITUTION';

-- Usar base de datos
USE `coco_tours_db_v2`;

-- Tabla: Administrador
CREATE TABLE IF NOT EXISTS `Administrador` (
  `idAdministrador` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `correo` VARCHAR(60) UNIQUE NOT NULL,
  PRIMARY KEY (`idAdministrador`)
) ENGINE=InnoDB;

-- Tabla: Cliente
CREATE TABLE IF NOT EXISTS `Cliente` (
  `idCliente` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `correo` VARCHAR(100) UNIQUE NOT NULL,
  PRIMARY KEY (`idCliente`)
) ENGINE=InnoDB;

-- Tabla: Proveedor
CREATE TABLE IF NOT EXISTS `Proveedor` (
  `idProveedor` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) UNIQUE NOT NULL,
  `descripcion` VARCHAR(70) NOT NULL,
  PRIMARY KEY (`idProveedor`)
) ENGINE=InnoDB;

-- Tabla: Viaje
CREATE TABLE IF NOT EXISTS `Viaje` (
  `idViaje` INT NOT NULL AUTO_INCREMENT,
  `tipoViaje` VARCHAR(15) UNIQUE NOT NULL,
  `descripcion` TEXT NOT NULL,
  `precio` DECIMAL NOT NULL,
  PRIMARY KEY (`idViaje`)
) ENGINE=InnoDB;

-- Tabla: Extras
CREATE TABLE IF NOT EXISTS `Extras` (
  `idExtra` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(20) UNIQUE NOT NULL,
  `descripcion` TEXT NOT NULL,
  `precio` DECIMAL NOT NULL,
  PRIMARY KEY (`idExtra`)
) ENGINE=InnoDB;

-- Tabla: Viaje_Extras
CREATE TABLE IF NOT EXISTS `Viaje_Extras` (
  `idViaje_Extra` INT NOT NULL AUTO_INCREMENT,
  `idViaje` INT NOT NULL,
  `idExtra` INT DEFAULT NULL,
  PRIMARY KEY (`idViaje_Extra`),
  INDEX `idx_Extra` (`idExtra`),
  INDEX `idx_viaje_Extra` (`idViaje`),
  CONSTRAINT `FK_Extra_Viaje`
    FOREIGN KEY (`idViaje`)
    REFERENCES `Viaje` (`idViaje`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_Viaje_Extra`
    FOREIGN KEY (`idExtra`)
    REFERENCES `Extras` (`idExtra`)
    ON DELETE SET NULL
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Tabla: detalleViaje
CREATE TABLE IF NOT EXISTS `detalleViaje` (
  `idDetalleViaje` INT NOT NULL AUTO_INCREMENT,
  `idProveedor` INT DEFAULT NULL,
  `idViaje` INT NOT NULL,
  PRIMARY KEY (`idDetalleViaje`),
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

-- Tabla: Usuario
CREATE TABLE IF NOT EXISTS `Usuario` (
  `idUsuario` INT NOT NULL AUTO_INCREMENT,
  `usuario` VARCHAR(40) UNIQUE NOT NULL,
  `contraseña` VARCHAR(100) UNIQUE NOT NULL,
  `admin` INT DEFAULT NULL,
  `cliente` INT DEFAULT NULL,
  PRIMARY KEY (`idUsuario`),
  INDEX `idx_admin` (`admin`),
  INDEX `idx_cliente` (`cliente`),
  CONSTRAINT `FK_usuario_admin`
    FOREIGN KEY (`admin`)
    REFERENCES `Administrador` (`idAdministrador`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_usuario_cliente`
    FOREIGN KEY (`cliente`)
    REFERENCES `Cliente` (`idCliente`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Tabla: telefonoClientes
CREATE TABLE IF NOT EXISTS `telefonoClientes` (
  `idtelefonoClientes` INT NOT NULL AUTO_INCREMENT,
  `numero` VARCHAR(25) UNIQUE NOT NULL,
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

-- Tabla: reservas
CREATE TABLE IF NOT EXISTS `reservas` (
  `idreservas` INT NOT NULL AUTO_INCREMENT,
  `idCliente` INT NOT NULL,
  `idAdministrador` INT NOT NULL,
  `idDetalle` INT NOT NULL,
  `fecha` VARCHAR(10) NOT NULL,
  `hora` VARCHAR(10) NOT NULL,
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
    REFERENCES `detalleViaje` (`idDetalleViaje`)
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
    REFERENCES `detalleViaje` (`idDetalleViaje`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB;

-- Restaurar configuración original
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
