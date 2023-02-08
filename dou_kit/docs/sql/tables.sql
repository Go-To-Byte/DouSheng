CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(32) DEFAULT '' COMMENT '用户名',
    `password` varchar(60) NOT NULL DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_pk2` (`username`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE IF NOT EXISTS `video` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `title` varchar(100) NOT NULL DEFAULT '' COMMENT '视频标题',
    `cover_url` varchar(100) NOT NULL DEFAULT '' COMMENT '视频封面地址',
    `play_url` varchar(256) NOT NULL DEFAULT '' COMMENT '视频播放地址',
    `author_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '作者ID',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';
