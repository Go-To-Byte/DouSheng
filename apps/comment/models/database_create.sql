use go_to_byte;

set FOREIGN_KEY_CHECKS = 0;

drop table if exists user;
drop table if exists user_follow;
drop table if exists user_follower;
drop table if exists user_info;
drop table if exists user_message;
drop table if exists video_comment;
drop table if exists video_favorite;
drop table if exists video_info;
drop table if exists comment;
drop table if exists favorite;
drop table if exists follow;
drop table if exists follower;
drop table if exists info;
drop table if exists message;
drop table if exists user;
drop table if exists video;

set FOREIGN_KEY_CHECKS = 1;

create table comment
(
    id       bigint       not null
        primary key,
    video_id bigint       not null,
    user_id  bigint       not null,
    content  varchar(128) not null
);

create table favorite
(
    user_id  bigint     not null,
    video_id bigint     not null,
    flag     tinyint(1) not null,
    primary key (user_id, video_id)
);

create table follow
(
    user_id    bigint     not null,
    to_user_id bigint     not null,
    flag       tinyint(1) not null,
    primary key (to_user_id, user_id)
);

create table follower
(
    user_id    bigint     not null,
    to_user_id bigint     not null,
    flag       tinyint(1) not null,
    primary key (user_id, to_user_id)
);

create table message
(
    id         bigint       not null
        primary key,
    user_id    bigint       not null,
    to_user_id bigint       not null,
    content    varchar(128) not null
);

create table user
(
    id       bigint      not null
        primary key,
    username varchar(16) not null,
    passwd   char(128)   not null
);

create table video
(
    id        bigint       not null
        primary key,
    auth_id   bigint       not null,
    titel     varchar(128) not null,
    cover_url longtext     not null,
    play_url  varchar(256) not null
);

