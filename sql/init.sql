CREATE TABLE `tab_lost` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `avatar_url` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `gender` int(10) unsigned DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `subject` varchar(255) DEFAULT NULL,
  `characters` varchar(255) DEFAULT NULL,
  `details` varchar(255) DEFAULT NULL,
  `data_from` varchar(255) DEFAULT NULL,
  `birthed_province` varchar(255) DEFAULT NULL,
  `birthed_city` varchar(255) DEFAULT NULL,
  `birthed_country` varchar(255) DEFAULT NULL,
  `birthed_address` varchar(255) DEFAULT NULL,
  `birthed_at` datetime DEFAULT NULL,
  `missed_country` varchar(255) DEFAULT NULL,
  `missed_province` varchar(255) DEFAULT NULL,
  `missed_city` varchar(255) DEFAULT NULL,
  `missed_address` varchar(255) DEFAULT NULL,
  `missed_at` datetime DEFAULT NULL,
  `handler` varchar(255) DEFAULT NULL,
  `babyid` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL,
  `height` varchar(255) DEFAULT NULL,
  `syncstatus` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `data_from` (`data_from`),
  KEY `idx_articles_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `tab_rescue` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone1` varchar(255) DEFAULT NULL,
  `phone2` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `tab_lost_stat` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `lost_id` int(10) unsigned NOT NULL,
  `babyid` varchar(255) DEFAULT NULL,
  `share_count` int(10) unsigned NOT NULL DEFAULT 0,
  `show_count` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `lost_id` (`uniq_lost`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
