CREATE TABLE IF NOT EXISTS `user` (
                        `id` BIGINT ( 20 ) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                        `username` VARCHAR ( 32 ) DEFAULT '' COMMENT '用户名',
                        `password` VARCHAR ( 60 ) NOT NULL DEFAULT '' COMMENT '密码',
                        PRIMARY KEY ( `id` ),
                        UNIQUE KEY `user_pk2` ( `id`, `username` )
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';
