CREATE TABLE `foods` (
  `food_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `carbohydrates` decimal(6,2) DEFAULT NULL,
  `fat` decimal(6,2) DEFAULT NULL,
  `category` varchar(100) DEFAULT NULL,
  `nutrition` json DEFAULT NULL,
  `image` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`food_id`)
);