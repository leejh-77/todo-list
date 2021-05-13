CREATE DATABASE IF NOT EXISTS `todo` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
USE todo;

CREATE TABLE `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `emailAddress` VARCHAR(100) NOT NULL UNIQUE,
    `password` TEXT NOT NULL,
    `username` VARCHAR(50) NOT NULL,
    `registeredTime` INT NOT NULL
);