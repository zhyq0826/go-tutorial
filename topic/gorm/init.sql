
CREATE TABLE IF NOT EXISTS `domain` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`private` int(10) unsigned NOT NULL default 0,
`name` varchar(200) NOT NULL default "",
`url` varchar(200) NOT NULL default "",
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;