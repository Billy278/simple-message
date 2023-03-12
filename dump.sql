-- MySQL dump 10.16  Distrib 10.1.22-MariaDB, for Win32 (AMD64)
--
-- Host: localhost    Database: simple_message
-- ------------------------------------------------------
-- Server version	10.1.22-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `last_read`
--

DROP TABLE IF EXISTS `last_read`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `last_read` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message_id` int(11) DEFAULT NULL,
  `sender_id` int(16) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sender_id` (`sender_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `last_read`
--

LOCK TABLES `last_read` WRITE;
/*!40000 ALTER TABLE `last_read` DISABLE KEYS */;
INSERT INTO `last_read` VALUES (6,8,1),(7,4,3),(10,10,4);
/*!40000 ALTER TABLE `last_read` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sender` varchar(200) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `message` varchar(500) DEFAULT NULL,
  `receiver` varchar(200) DEFAULT NULL,
  `sender_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_receiver` (`receiver`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
INSERT INTO `message` VALUES (1,'hallo test','2023-02-06','hallo test','bima',3),(2,'hallo test 2','2023-03-06','hallo test 2','bima',3),(3,'hallo test 3','2023-03-07','hallo test 3','bima',3),(4,'hallo test 4','2023-03-08','hallo test 4','bima',3),(5,'billy','2023-03-11','hallo test 11','bima',1),(6,'billy','2023-03-12','hallo test 12','bima',1),(7,'billy','2023-03-14','hallo test 13','bima',1),(8,'billy','2023-03-14','hallo test 14','bima',1),(10,'andi','2023-03-14','hallo test 15','bima',4),(11,'andi','2023-03-14','hallo test 16','bima',4),(12,'eva','2023-03-14','hallo test 17','bima',5),(13,'eva','2023-03-14','hallo test 18','bima',5);
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_test`
--

DROP TABLE IF EXISTS `message_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message_test` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sender` varchar(200) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `message` varchar(500) DEFAULT NULL,
  `receiver` varchar(200) DEFAULT NULL,
  `sender_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_test`
--

LOCK TABLES `message_test` WRITE;
/*!40000 ALTER TABLE `message_test` DISABLE KEYS */;
INSERT INTO `message_test` VALUES (1,'billy','2023-03-14','hallo test 1','bima',1),(2,'billy','2023-03-14','hallo test 1','bima',1),(3,'billy','2023-03-14','hallo test 1','bima',1);
/*!40000 ALTER TABLE `message_test` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `username` varchar(200) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'Billy Ginting','billy','$2a$10$66c9r1Opl4i7ezgfhXL3mumOMoK58ibgGFdxSqKie1goCz0bP/gcC'),(2,'Bima Ginting','bima','$2a$10$vrMPsR9CKIX5OBFT5R7aqelMb5KkApBKvOxMu7X7T07HnWdSpOvQC'),(3,'Aldo','aldo','$2a$10$5g9CP00dM6Uc4HXfIQhmj.iiWQwZyb1l18z6DVODuKo25A9bB3GwS'),(4,'Andi','andi','$2a$10$cTVj.APQ73o/2yZMF/B6cudKqXTDe8IxPvGOhhQaM5XI252ccjQ1G'),(5,'eva','eva','$2a$10$CNSgn1zS0FeXSybq4.vWzeKLDu/84s7m1kAfn1tgrfBBKHcq3Zqq.');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-12 21:08:48
