create table IF not exists `articles`
(
 `id`               INT AUTO_INCREMENT,
 `title`            VARCHAR(100) NOT NULL,
 `body`             TEXT,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`ID`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
