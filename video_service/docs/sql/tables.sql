CREATE TABLE IF NOT EXISTS `video` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `title` varchar(100) NOT NULL DEFAULT '' COMMENT '视频标题',
    `cover_url` varchar(100) NOT NULL DEFAULT '' COMMENT '视频封面地址',
    `play_url` varchar(256) NOT NULL DEFAULT '' COMMENT '视频播放地址',
    `author_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '作者ID',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';
