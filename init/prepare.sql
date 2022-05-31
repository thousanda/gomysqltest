CREATE DATABASE IF NOT EXISTS dbname;
use dbname;

CREATE TABLE IF NOT EXISTS dbname.t_user
(
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64),
    PRIMARY KEY (`id`)
);

INSERT INTO dbname.t_user (name) VALUES ("佐藤");
INSERT INTO dbname.t_user (name) VALUES ("鈴木");
INSERT INTO dbname.t_user (name) VALUES ("山田");
INSERT INTO dbname.t_user (name) VALUES ("加藤");
INSERT INTO dbname.t_user (name) VALUES ("武田");



