DROP DATABASE IF EXISTS `appdb`;
DROP USER IF EXISTS 'appuser'@'appdb';

CREATE DATABASE `appdb`;
--  @'%' means any host - this needs to be defined on both the create user and grant statements
CREATE USER 'appuser'@'%' IDENTIFIED BY 'pass';
GRANT ALL PRIVILEGES ON `appdb`.* TO 'appuser'@'%';

FLUSH PRIVILEGES;