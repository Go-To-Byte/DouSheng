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
    `created_at` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发布时间',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

CREATE TABLE IF NOT EXISTS `comment`  (
  `id` bigint NOT NULL COMMENT '评论ID',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `video_id` bigint NOT NULL COMMENT '视频ID',
  `content` varchar(256) NOT NULL DEFAULT '' COMMENT '评论内容',
  `create_date` varchar(64) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `favorite`  (
    `id` bigint NOT NULL COMMENT '喜欢ID',
    `user_id` bigint NOT NULL COMMENT '用户ID',
    `video_id` bigint NOT NULL COMMENT '视频ID',
    PRIMARY KEY (`id`)
);

