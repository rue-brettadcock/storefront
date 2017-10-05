CREATE DATABASE productInfo;
use productInfo;

CREATE TABLE `products` (
  `id` varchar(200) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `vendor` varchar(50) DEFAULT NULL,
  `quantity` int(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
)
