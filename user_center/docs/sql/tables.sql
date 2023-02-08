CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(32) DEFAULT '' COMMENT '用户名',
    `password` varchar(60) NOT NULL DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_pk2` (`username`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
