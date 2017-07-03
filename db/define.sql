
-- change password
-- SET PASSWORD FOR 'root'@'localhost' = PASSWORD('newpass');

-- delete user
-- DROP USER 'poker'@'localhost';

-- create user
CREATE USER 'treehole'@'localhost' IDENTIFIED BY 'treehole1111';

-- create database
CREATE DATABASE treehole COLLATE 'utf8_general_ci';
GRANT ALL ON treehole.* TO 'treehole'@'localhost';

-- select database
USE treehole;

ALTER TABLE user AUTO_INCREMENT=10001;
-- create word table
CREATE TABLE IF NOT EXISTS word (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    nickName varchar(128) NOT NULL,
    word varchar(1024) NOT NULL,
    PRIMARY KEY(id)
) ENGINE = innoDB DEFAULT CHARACTER SET = utf8;
