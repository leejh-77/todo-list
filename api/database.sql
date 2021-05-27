CREATE DATABASE IF NOT EXISTS `todo` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
USE todo;

CREATE TABLE IF NOT EXISTS `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `emailAddress` VARCHAR(100) NOT NULL UNIQUE,
    `password` TEXT NOT NULL,
    `username` VARCHAR(50) NOT NULL,
    `registeredTime` INT NOT NULL
);

CREATE TABLE IF NOT EXISTS `todos` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `folderId` INT NOT NULL,
    `userId` INT NOT NULL,
    `subject` VARCHAR(50),
    `body` TEXT,
    `status` INT NOT NULL,
    `completedTime` INT,
    `position` INT
);

CREATE TABLE IF NOT EXISTS `workspaces` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(50),
    `createdTime` INT
);

CREATE TABLE IF NOT EXISTS `workspaceMembers` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `type` TINYINT(1) NOT NULL,
    `workspaceId` INT NOT NULL,
    `userId` INT NOT NULL
);

CREATE TABLE IF NOT EXISTS `folders` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(50),
    `workspaceId` INT NOT NULL
);