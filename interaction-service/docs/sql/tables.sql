CREATE TABLE IF NOT EXISTS `comment`  (
    `id` bigint NOT NULL COMMENT '评论ID',
    `user_id` bigint NOT NULL COMMENT '用户ID',
    `video_id` bigint NOT NULL COMMENT '视频ID',
    `content` varchar(256) NOT NULL DEFAULT '' COMMENT '评论内容',
    `create_date` varchar(64) NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `message` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `to_user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
    `from_user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
    `content` text NOT NULL DEFAULT '',
    `created_at` bigint(20) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
