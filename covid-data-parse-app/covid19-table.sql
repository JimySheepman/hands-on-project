CREATE DATABASE challenge;
USE challenge;
CREATE TABLE `covid19` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date_time` varchar(128) COLLATE utf8_general_ci NOT NULL,
  `country` varchar(128) COLLATE utf8_general_ci NOT NULL,
  `cases` varchar(128) COLLATE utf8_general_ci NOT NULL,
  `deaths` varchar(128) COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
