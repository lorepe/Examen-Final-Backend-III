-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: localhost    Database: clinica
-- ------------------------------------------------------
-- Server version	8.0.30
create database clinica;
use clinica;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `odontologo`
--

DROP TABLE IF EXISTS `odontologo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `odontologo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) NOT NULL,
  `apellido` varchar(45) NOT NULL,
  `matricula` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `odontologo`
--

LOCK TABLES `odontologo` WRITE;
/*!40000 ALTER TABLE `odontologo` DISABLE KEYS */;
/*!40000 ALTER TABLE `odontologo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paciente`
--

DROP TABLE IF EXISTS `paciente`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paciente` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) NOT NULL,
  `apellido` varchar(45) NOT NULL,
  `domicilio` varchar(45) NOT NULL,
  `dni` varchar(45) NOT NULL,
  `fecha_alta` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paciente`
--

LOCK TABLES `paciente` WRITE;
/*!40000 ALTER TABLE `paciente` DISABLE KEYS */;
/*!40000 ALTER TABLE `paciente` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `turno`
--

DROP TABLE IF EXISTS `turno`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `turno` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_paciente` int NOT NULL,
  `id_odontologo` int NOT NULL,
  `fecha_hora` datetime DEFAULT NULL,
  `descripcion` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_paciente` (`id_paciente`),
  KEY `id_odontologo` (`id_odontologo`),
  CONSTRAINT `turno_ibfk_1` FOREIGN KEY (`id_paciente`) REFERENCES `paciente` (`id`),
  CONSTRAINT `turno_ibfk_2` FOREIGN KEY (`id_odontologo`) REFERENCES `odontologo` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `turno`
--

LOCK TABLES `turno` WRITE;
/*!40000 ALTER TABLE `turno` DISABLE KEYS */;
/*!40000 ALTER TABLE `turno` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

insert into odontologo values(default, "Gustavo", "Aguirre", "MN20017");
insert into odontologo values(default, "Alejandra", "Torres", "MN89512");
insert into odontologo values(default, "Marta", "Guitierrez", "MN48947");
insert into odontologo values(default, "Ricardo", "Perez", "MN94578");
insert into odontologo values(default, "Pedro", "Costas", "MN256462");
insert into odontologo values(default, "Sandra", "Milani", "MN20547");

insert into paciente values(default, "Tamara", "Gomez", "Las Violetas 435", "1111", "2020-10-30");
insert into paciente values(default, "Luciano", "Lopez", "San Martin 2014", "2222", "2021-01-09");
insert into paciente values(default, "Teresa", "Calles", "Avellaneda 1078", "3333", "2022-09-01");
insert into paciente values(default, "Damián", "Brown", "Los Cardos 4735", "4444", "2021-11-08");
insert into paciente values(default, "Gonzalo", "Estevez", "Las Violetas 435", "5555", "2022-07-24");
insert into paciente values(default, "Ezequiel", "Avalo", "Pintos 3535", "6666", "2023-04-25");
insert into paciente values(default, "Jesús", "Luna", "Callenas 1457", "7777", "2021-09-13");
insert into paciente values(default, "Luaciana", "Safoni", "Del Valle 2689", "8888", "2022-12-20");
insert into paciente values(default, "Romina", "Crivaro", "Sarmiento 3451", "9999", "2021-03-10");
insert into paciente values(default, "Esteban", "Cuenca", "Trabajadores 1235", "1010", "2023-02-07");

insert into turno values(default, 1, 6, "2023-06-27 08:30:00", "Dolor de Muela");
insert into turno values(default, 2, 6, "2023-06-27 08:15:00", "Extracción");
insert into turno values(default, 3, 6, "2023-06-27 09:00:00", "Control cirugía");
insert into turno values(default, 4, 3, "2023-06-27 08:30:00", "Prótesis");
insert into turno values(default, 5, 3, "2023-06-27 08:45:00", "Limpieza");
insert into turno values(default, 6, 3, "2023-06-27 09:30:00", "Colocación ortodoncia");

-- Dump completed on 2023-06-26 21:18:47
