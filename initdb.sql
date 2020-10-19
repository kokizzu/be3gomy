CREATE DATABASE db_apigo;
CREATE USER `user1`@`localhost` IDENTIFIED BY '1234567890';
USE db_apigo;
GRANT ALL PRIVILEGES ON db_apigo.* TO `user1`@`localhost`;
FLUSH PRIVILEGES;

CREATE TABLE `mahasiswa` (
	`id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	`nim` INT NOT NULL,
	`name` VARCHAR(64) NOT NULL,
	`semester` SMALLINT NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL
);
-- https://pastebin.com/dNaT8fHG 
