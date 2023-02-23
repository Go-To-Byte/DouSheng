-- @hexiaoming 

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE IF NOT EXISTS `dousheng` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `dousheng`;


DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(32) DEFAULT '' COMMENT '用户名',
    `password` varchar(60) NOT NULL DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_pk2` (`username`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

DROP TABLE IF EXISTS `video`;
CREATE TABLE IF NOT EXISTS `video` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `title` varchar(100) NOT NULL DEFAULT '' COMMENT '视频标题',
    `cover_url` varchar(100) NOT NULL DEFAULT '' COMMENT '视频封面地址',
    `play_url` varchar(256) NOT NULL DEFAULT '' COMMENT '视频播放地址',
    `author_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '作者ID',
    `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布时间',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

--
-- Table structure for table `user_follow`
--

DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow` (
  `user_id` bigint NOT NULL,
  `follow_id` bigint NOT NULL,
  `follow_flag` tinyint(1) NOT NULL,
  PRIMARY KEY (`user_id`,`follow_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;