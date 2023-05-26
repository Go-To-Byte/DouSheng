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

