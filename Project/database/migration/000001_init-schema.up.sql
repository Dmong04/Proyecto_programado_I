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
  PRIMARY KEY (`idAdministrador`)
) ENGINE=InnoDB;

-- Tabla: Cliente
CREATE TABLE IF NOT EXISTS `Cliente` (
  `idCliente` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `telefono` INT (12) NOT NULL,
  PRIMARY KEY (`idCliente`)
) ENGINE=InnoDB;

-- Tabla: Usuario
CREATE TABLE IF NOT EXISTS `Usuario` (
  `idUsuario` INT NOT NULL AUTO_INCREMENT,
  `correo` VARCHAR(60) NOT NULL,
  `usuario` VARCHAR(40) NOT NULL,
  `contraseña` VARCHAR(100) NOT NULL,
  `role` VARCHAR(20) NOT NULL,
  `idAdministrador` INT DEFAULT NULL,
  `idCliente` INT DEFAULT NULL,
  `created_at`	datetime DEFAULT NULL,
  `updated_at`	datetime DEFAULT NULL,
  `remember_token` VARCHAR(255),
  PRIMARY KEY (`idUsuario`),
  CONSTRAINT `FK_Usuario_Cliente`
    FOREIGN KEY (`idCliente`)
    REFERENCES `Cliente` (`idCliente`)
    ON DELETE SET NULL
    ON UPDATE CASCADE,
  CONSTRAINT `FK_Usuario_Administrador`
    FOREIGN KEY (`idAdministrador`)
    REFERENCES `Administrador` (`idAdministrador`)
    ON DELETE SET NULL
    ON UPDATE CASCADE
) ENGINE=InnoDB;
-- Tabla: Viaje
CREATE TABLE IF NOT EXISTS `Viaje` (
  `idViaje` INT NOT NULL AUTO_INCREMENT,
  `tipoViaje` VARCHAR(20) NOT NULL,
  `descripcion` TEXT NOT NULL,
  PRIMARY KEY (`idViaje`)
) ENGINE=InnoDB;

-- Tabla: Proveedor
CREATE TABLE IF NOT EXISTS `Proveedor` (
  `idProveedor` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(60) NOT NULL,
  `descrip` VARCHAR(70) NOT NULL,
  PRIMARY KEY (`idProveedor`)
) ENGINE=InnoDB;

-- Tabla: detalleViaje
CREATE TABLE IF NOT EXISTS `detalleViaje` (
  `idDetalleViaje` INT NOT NULL AUTO_INCREMENT,
  `fecha` DATE NOT NULL,
  `hora` TIME NOT NULL,
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

-- Tabla: Reservas
CREATE TABLE IF NOT EXISTS `reservas` (
  `idreservas` INT NOT NULL AUTO_INCREMENT,
  `idUsuario` INT NOT NULL,
  `idDetalle` INT NOT NULL,
  `estado` VARCHAR(10) NOT NULL,
  PRIMARY KEY (`idreservas`),
  INDEX `idx_reserva_usuario` (`idUsuario`),
  INDEX `idx_reserva_detalle` (`idDetalle`),
  CONSTRAINT `FK_reserva_usuario`
    FOREIGN KEY (`idUsuario`)
    REFERENCES `Usuario` (`idUsuario`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_reserva_detalle`
    FOREIGN KEY (`idDetalle`)
    REFERENCES `detalleViaje` (`idDetalleViaje`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB;
-- Restaurar configuración original
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;