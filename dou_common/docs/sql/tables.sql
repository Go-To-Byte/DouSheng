CREATE TABLE IF NOT EXISTS `user` (
    `id` BIGINT ( 20 ) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` VARCHAR ( 32 ) DEFAULT '' COMMENT '用户名',
    `password` VARCHAR ( 60 ) NOT NULL DEFAULT '' COMMENT '密码',
    PRIMARY KEY ( `id` ),
    UNIQUE KEY `user_pk2` ( `id`, `username` )
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

CREATE TABLE IF NOT EXISTS `video` (
                         `id` BIGINT ( 20 ) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '视频ID',
                         `title` VARCHAR ( 100 ) NOT NULL DEFAULT '' COMMENT '视频标题',
                         `cover_url` VARCHAR ( 100 ) NOT NULL DEFAULT '' COMMENT '视频封面地址',
                         `flay_url` VARCHAR ( 256 ) NOT NULL DEFAULT '' COMMENT '视频播放地址',
                         `user_id` BIGINT ( 20 ) UNSIGNED NOT NULL DEFAULT '0' COMMENT '作者ID',
                         PRIMARY KEY ( `id` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '视频表';
