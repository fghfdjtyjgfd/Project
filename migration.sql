CREATE TABLE `Beer` (
  `ID` bigint(20)unsigned NOT NULL AUTO_INCREMENT,
  `deleted_at` datetime,
  `name` varchar(50),
  `type` longtext,
  `detail` longtext,
  `image_url` longtext,
  `company_id` bigint(20)unsigned DEFAULT NULL
);

CREATE TABLE `Users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(191)DEFAULT NULL,
  `password` longtext DEFAULT NULL
);

CREATE TABLE `Company` (
  `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL
);

CREATE TABLE `Distributer` (
  `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL
);

CREATE TABLE `DistributerBeer` (
  `beer_id` bigint(20) unsigned DEFAULT NULL,
  `distributer_id` bigint(20) unsigned DEFAULT NULL
);

ALTER TABLE `DistributerBeer` ADD FOREIGN KEY (`beer_id`) REFERENCES `Beer` (`ID`);

ALTER TABLE `DistributerBeer` ADD FOREIGN KEY (`distributer_id`) REFERENCES `Distributer` (`ID`);

ALTER TABLE `Beer` ADD FOREIGN KEY (`company_id`) REFERENCES `Company` (`ID`);
