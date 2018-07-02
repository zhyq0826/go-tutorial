DROP TABLE IF EXISTS `app`;

CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_important` int(11) NOT NULL DEFAULT '0',
  `name` varchar(32) NOT NULL DEFAULT '',
  `url` varchar(256) NOT NULL DEFAULT '',
  `desc` varchar(512) NOT NULL DEFAULT '',
  `repository_url` varchar(512) NOT NULL DEFAULT '',
  `deploy_dir` varchar(512) NOT NULL DEFAULT '',
  `monitor_url` varchar(512) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


# Dump of table computer
# ------------------------------------------------------------

DROP TABLE IF EXISTS `computer`;

CREATE TABLE `computer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `cpu` int(11) NOT NULL DEFAULT '0',
  `ram` int(11) NOT NULL DEFAULT '0',
  `private_ip` varchar(32) NOT NULL DEFAULT '',
  `public_ip` varchar(32) NOT NULL DEFAULT '',
  `host_id` varchar(32) DEFAULT NULL,
  `service_id` int(11) NOT NULL DEFAULT '0',
  `app_id` int(11) NOT NULL DEFAULT '0',
  `tag` varchar(32) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table disk
# ------------------------------------------------------------

DROP TABLE IF EXISTS `disk`;

CREATE TABLE `disk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `size` int(11) NOT NULL DEFAULT '0',
  `left` int(11) NOT NULL DEFAULT '0',
  `computer_id` int(11) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table domain
# ------------------------------------------------------------

DROP TABLE IF EXISTS `domain`;

CREATE TABLE `domain` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `name` varchar(256) NOT NULL DEFAULT '',
  `private` tinyint(4) NOT NULL DEFAULT '0',
  `host` varchar(256) NOT NULL DEFAULT '',
  `ip` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;