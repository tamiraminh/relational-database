CREATE TABLE `User` (
  `id` varchar(36) PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `password` varchar(255),
  `type` enum('admin', 'reguler') NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36)
);

CREATE TABLE `Brand` (
  `id` varchar(36) PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36)
);

CREATE TABLE `Product` (
  `id` varchar(36) PRIMARY KEY,
  `brandId` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36)
);

CREATE TABLE `Variant` (
  `id` varchar(36) PRIMARY KEY,
  `productId` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` decimal(10, 2) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36)
);

CREATE TABLE `Image` (
  `id` varchar(36) PRIMARY KEY,
  `variantId` varchar(36) NOT NULL,
  `url` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36)
);

CREATE TABLE `Warehouse` (
  `id` varchar(36) PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36)
);

CREATE TABLE `VariantWarehouse` (
  `variantId` varchar(36) NOT NULL,
  `warehouseId` varchar(36) NOT NULL,
  `stock` int NOT NULL,
  `status` enum('ready', 'out of stock'),
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36),
  PRIMARY KEY (`variantId`, `warehouseId`)
);

CREATE TABLE `UserProduct` (
  `userId` varchar(36) NOT NULL,
  `productId` varchar(36) NOT NULL,
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36),
  PRIMARY KEY (`userId`, `productId`)
);

CREATE TABLE `SalesFact` (
  `id` varchar(36) PRIMARY KEY,
  `productId` varchar(36),
  `variantId` varchar(36),
  `warehouseId` varchar(36),
  `saleDate` datetime,
  `saleQuantity` int,
  `saleAmount` decimal(10, 2),
  `createdAt` datetime NOT NULL,
  `createdBy` varchar(36) NOT NULL,
  `updatedAt` datetime NOT NULL,
  `updatedBy` varchar(36) NOT NULL,
  `deletedAt` datetime,
  `deletedBy` varchar(36)
);

CREATE INDEX `Product_index_0` ON `Product` (`id`) USING BTREE;

CREATE INDEX `Variant_index_1` ON `Variant` (`id`, `name`) USING BTREE;

ALTER TABLE `Product` ADD FOREIGN KEY (`brandId`) REFERENCES `Brand` (`id`);

ALTER TABLE `Variant` ADD FOREIGN KEY (`productId`) REFERENCES `Product` (`id`);

ALTER TABLE `Image` ADD FOREIGN KEY (`variantId`) REFERENCES `Variant` (`id`);

ALTER TABLE `VariantWarehouse` ADD FOREIGN KEY (`variantId`) REFERENCES `Variant` (`id`);

ALTER TABLE `VariantWarehouse` ADD FOREIGN KEY (`warehouseId`) REFERENCES `Warehouse` (`id`);

ALTER TABLE `UserProduct` ADD FOREIGN KEY (`userId`) REFERENCES `User` (`id`);

ALTER TABLE `UserProduct` ADD FOREIGN KEY (`productId`) REFERENCES `Product` (`id`);

ALTER TABLE `SalesFact` ADD FOREIGN KEY (`productId`) REFERENCES `Product` (`id`);

ALTER TABLE `SalesFact` ADD FOREIGN KEY (`variantId`) REFERENCES `Variant` (`id`);

ALTER TABLE `SalesFact` ADD FOREIGN KEY (`warehouseId`) REFERENCES `Warehouse` (`id`);
