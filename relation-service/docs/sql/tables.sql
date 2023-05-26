-- @hexiaoming 

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE TABLE IF NOT EXISTS  `user_follower` (
  `user_id` bigint NOT NULL,
  `follower_id` bigint NOT NULL,
  `follower_flag` tinyint(1) NOT NULL,
  PRIMARY KEY (`user_id`,`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS  `user_follow` (
                               `user_id` bigint NOT NULL,
                               `follow_id` bigint NOT NULL,
                               `follow_flag` tinyint(1) NOT NULL,
                               PRIMARY KEY (`user_id`,`follow_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
