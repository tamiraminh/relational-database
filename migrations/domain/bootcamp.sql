-- MariaDB dump 10.18  Distrib 10.5.8-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: evermos_bootcamp
-- ------------------------------------------------------
-- Server version	10.5.8-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `brand`
--

DROP TABLE IF EXISTS `brand`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `brand` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `brand`
--

LOCK TABLES `brand` WRITE;
/*!40000 ALTER TABLE `brand` DISABLE KEYS */;
INSERT INTO `brand` VALUES ('1f9b7c1b-5e88-4a22-9d47-287328e94c8a','Nike','2022-09-01 08:00:00','1d50c7ab-2363-45a6-8e91-c86242f72d27','2022-09-01 08:00:00','1d50c7ab-2363-45a6-8e91-c86242f72d27',NULL,NULL),('2d1e3c2f-5bc2-4f2f-9f29-679d0df9160d','Adidas','2022-08-10 10:30:00','1d50c7ab-2363-45a6-8e91-c86242f72d27','2022-08-10 10:30:00','1d50c7ab-2363-45a6-8e91-c86242f72d27',NULL,NULL),('3eae1d35-76f7-4b2b-9071-4719ec299ee3','Puma','2022-07-20 12:45:00','1d50c7ab-2363-45a6-8e91-c86242f72d27','2022-07-20 12:45:00','1d50c7ab-2363-45a6-8e91-c86242f72d27',NULL,NULL);
/*!40000 ALTER TABLE `brand` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `image`
--

DROP TABLE IF EXISTS `image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `image` (
  `id` varchar(36) NOT NULL,
  `variantId` varchar(36) NOT NULL,
  `url` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `variantId` (`variantId`),
  CONSTRAINT `image_ibfk_2` FOREIGN KEY (`variantId`) REFERENCES `variant` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `image`
--

LOCK TABLES `image` WRITE;
/*!40000 ALTER TABLE `image` DISABLE KEYS */;
INSERT INTO `image` VALUES ('1a07e3d0-4801-44b3-9ee2-944e26e68b20','55a2eb7d-e135-405f-b8fd-682cc84ca6ee','https://example.com/image1.jpg','2023-07-27 12:00:00','8da6d159-6d75-4b19-8de7-08c79bdf5449','2023-07-27 12:00:00','8da6d159-6d75-4b19-8de7-08c79bdf5449',NULL,NULL),('2a3147a6-8b3d-4f57-8a3e-7201f26a92c5','59b503ac-d683-48a9-9be9-62313519f59f','https://example.com/image2.jpg','2023-07-27 13:30:00','e1b93559-3aa2-4c18-b5b0-91d3c35874e4','2023-07-27 13:30:00','e1b93559-3aa2-4c18-b5b0-91d3c35874e4',NULL,NULL),('3c043f2a-2c5f-4e54-9c28-38a6a3a998a9','90919957-8087-4f46-a35b-c9a578f8b95b','https://example.com/image3.jpg','2023-07-27 14:45:00','7c1a2eb2-e8e5-4262-b235-7828e9bb2de1','2023-07-27 14:45:00','7c1a2eb2-e8e5-4262-b235-7828e9bb2de1',NULL,NULL),('4eef3aa9-5d62-4d62-a59a-7a30f0a02f0d','cdec9878-f354-4a61-a421-c94797a11c78','https://example.com/image4.jpg','2023-07-27 16:00:00','8da6d159-6d75-4b19-8de7-08c79bdf5449','2023-07-27 16:00:00','8da6d159-6d75-4b19-8de7-08c79bdf5449',NULL,NULL),('5f9d64da-08f1-46f1-aa38-cda872b594b2','81b49b33-bfc3-4358-981c-058849a633eb','https://example.com/image5.jpg','2023-07-27 17:15:00','e1b93559-3aa2-4c18-b5b0-91d3c35874e4','2023-07-27 17:15:00','e1b93559-3aa2-4c18-b5b0-91d3c35874e4',NULL,NULL);
/*!40000 ALTER TABLE `image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product` (
  `id` varchar(36) NOT NULL,
  `brandId` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `stock` int(11) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `Product_index_0` (`id`) USING BTREE,
  KEY `brandId` (`brandId`),
  CONSTRAINT `product_ibfk_1` FOREIGN KEY (`brandId`) REFERENCES `brand` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES ('03aa6245-49b7-4050-98e4-5fd1f3769077','2d1e3c2f-5bc2-4f2f-9f29-679d0df9160d','Running Shoes',20,'2023-07-29 13:29:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 13:29:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('5908590b-0dc2-4190-98c3-d383766cfd5a','1f9b7c1b-5e88-4a22-9d47-287328e94c8a','Running Shoes',20,'2023-07-28 09:54:26','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-28 09:54:26','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('645bfa78-7529-4dfa-8c8c-8009ec3853d7','2d1e3c2f-5bc2-4f2f-9f29-679d0df9160d','Headband',30,'2023-07-29 13:30:43','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 15:36:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('8020c4d0-3eac-406a-b75e-ad8a68b1e691','1f9b7c1b-5e88-4a22-9d47-287328e94c8a','Running clothes',30,'2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL);
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = cp850 */ ;
/*!50003 SET character_set_results = cp850 */ ;
/*!50003 SET collation_connection  = cp850_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER update_variant_updatedAt
AFTER UPDATE ON Product
FOR EACH ROW
BEGIN
  UPDATE Variant
  SET updatedAt = NOW(),
      updatedBy = NEW.updatedBy
  WHERE Variant.productId = NEW.id;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `salesfact`
--

DROP TABLE IF EXISTS `salesfact`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `salesfact` (
  `id` varchar(36) NOT NULL,
  `productId` varchar(36) DEFAULT NULL,
  `variantId` varchar(36) DEFAULT NULL,
  `warehouseId` varchar(36) DEFAULT NULL,
  `saleDate` datetime DEFAULT NULL,
  `saleQuantity` int(11) DEFAULT NULL,
  `saleAmount` decimal(10,2) DEFAULT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `productId` (`productId`),
  KEY `variantId` (`variantId`),
  KEY `warehouseId` (`warehouseId`),
  CONSTRAINT `salesfact_ibfk_1` FOREIGN KEY (`productId`) REFERENCES `product` (`id`),
  CONSTRAINT `salesfact_ibfk_2` FOREIGN KEY (`variantId`) REFERENCES `variant` (`id`),
  CONSTRAINT `salesfact_ibfk_3` FOREIGN KEY (`warehouseId`) REFERENCES `warehouse` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `salesfact`
--

LOCK TABLES `salesfact` WRITE;
/*!40000 ALTER TABLE `salesfact` DISABLE KEYS */;
/*!40000 ALTER TABLE `salesfact` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `type` enum('admin','reguler') DEFAULT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('1d50c7ab-2363-45a6-8e91-c86242f72d27','John Doe','mypassword123','admin','2022-09-01 08:00:00','1d50c7ab-2363-45a6-8e91-c86242f72d27','2022-09-01 08:00:00','1d50c7ab-2363-45a6-8e91-c86242f72d27',NULL,NULL),('2e84a97e-6d18-4d71-9bf9-8e2074b6a882','Jane Smith','password456','reguler','2022-08-10 10:30:00','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2022-08-10 10:30:00','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('344f88df-90a0-4cc3-81cc-3188af33acdf','Michael Johnson','securepass789','reguler','2022-07-20 12:45:00','344f88df-90a0-4cc3-81cc-3188af33acdf','2022-07-20 12:45:00','344f88df-90a0-4cc3-81cc-3188af33acdf',NULL,NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userproduct`
--

DROP TABLE IF EXISTS `userproduct`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `userproduct` (
  `userId` varchar(36) NOT NULL,
  `productId` varchar(36) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`userId`,`productId`),
  KEY `productId` (`productId`),
  CONSTRAINT `userproduct_ibfk_1` FOREIGN KEY (`userId`) REFERENCES `user` (`id`),
  CONSTRAINT `userproduct_ibfk_2` FOREIGN KEY (`productId`) REFERENCES `product` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userproduct`
--

LOCK TABLES `userproduct` WRITE;
/*!40000 ALTER TABLE `userproduct` DISABLE KEYS */;
/*!40000 ALTER TABLE `userproduct` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `variant`
--

DROP TABLE IF EXISTS `variant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `variant` (
  `id` varchar(36) NOT NULL,
  `productId` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `stock` int(11) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `Variant_index_1` (`id`,`name`) USING BTREE,
  KEY `productId` (`productId`),
  CONSTRAINT `variant_ibfk_1` FOREIGN KEY (`productId`) REFERENCES `product` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `variant`
--

LOCK TABLES `variant` WRITE;
/*!40000 ALTER TABLE `variant` DISABLE KEYS */;
INSERT INTO `variant` VALUES ('55a2eb7d-e135-405f-b8fd-682cc84ca6ee','8020c4d0-3eac-406a-b75e-ad8a68b1e691','Black',149.99,10,'2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('59b503ac-d683-48a9-9be9-62313519f59f','8020c4d0-3eac-406a-b75e-ad8a68b1e691','Blue',129.99,10,'2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('81b49b33-bfc3-4358-981c-058849a633eb','8020c4d0-3eac-406a-b75e-ad8a68b1e691','Red',129.99,10,'2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 13:27:58','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('90919957-8087-4f46-a35b-c9a578f8b95b','645bfa78-7529-4dfa-8c8c-8009ec3853d7','Black',19.99,18,'2023-07-29 15:36:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 22:36:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('a42c9e80-8a8a-472a-ae1a-3c5529710429','03aa6245-49b7-4050-98e4-5fd1f3769077','Men Black',129.99,10,'2023-07-29 13:29:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 13:29:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('bcbefdb4-14a3-4335-a2dc-3bbecad75c0d','03aa6245-49b7-4050-98e4-5fd1f3769077','Woman Pink',129.99,10,'2023-07-29 13:29:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 13:29:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL),('cdec9878-f354-4a61-a421-c94797a11c78','645bfa78-7529-4dfa-8c8c-8009ec3853d7','Blue',19.99,12,'2023-07-29 15:36:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882','2023-07-29 22:36:12','2e84a97e-6d18-4d71-9bf9-8e2074b6a882',NULL,NULL);
/*!40000 ALTER TABLE `variant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `variantwarehouse`
--

DROP TABLE IF EXISTS `variantwarehouse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `variantwarehouse` (
  `variantId` varchar(36) NOT NULL,
  `warehouseId` varchar(36) NOT NULL,
  `stock` int(11) NOT NULL,
  `status` enum('ready','out of stock') DEFAULT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`variantId`,`warehouseId`),
  KEY `warehouseId` (`warehouseId`),
  CONSTRAINT `variantwarehouse_ibfk_1` FOREIGN KEY (`variantId`) REFERENCES `variant` (`id`),
  CONSTRAINT `variantwarehouse_ibfk_2` FOREIGN KEY (`warehouseId`) REFERENCES `warehouse` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `variantwarehouse`
--

LOCK TABLES `variantwarehouse` WRITE;
/*!40000 ALTER TABLE `variantwarehouse` DISABLE KEYS */;
/*!40000 ALTER TABLE `variantwarehouse` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `warehouse`
--

DROP TABLE IF EXISTS `warehouse`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `warehouse` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `deletedBy` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `warehouse`
--

LOCK TABLES `warehouse` WRITE;
/*!40000 ALTER TABLE `warehouse` DISABLE KEYS */;
/*!40000 ALTER TABLE `warehouse` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-30 22:10:55
