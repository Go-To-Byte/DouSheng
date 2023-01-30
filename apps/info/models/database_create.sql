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

set FOREIGN_KEY_CHECKS = 1;

/*==============================================================*/
/* Table: comment                                               */
/*==============================================================*/
create table comment
(
    id       bigint       not null,
    video_id bigint       not null,
    user_id  bigint       not null,
    comment  varchar(128) not null,
    primary key (id)
);



/*==============================================================*/
/* Table: favorite                                              */
/*==============================================================*/
create table favorite
(
    user_id  bigint not null,
    video_id bigint not null,
    flag     bool   not null,
    primary key (user_id, video_id)
);


/*==============================================================*/
/* Table: follow                                                */
/*==============================================================*/
create table follow
(
    follow    bigint not null,
    to_follow bigint not null,
    flag      bool   not null,
    primary key (to_follow, follow)
);


/*==============================================================*/
/* Table: follower                                              */
/*==============================================================*/
create table follower
(
    follower    bigint not null,
    to_follower bigint not null,
    flag        bool   not null,
    primary key (follower, to_follower)
);


/*==============================================================*/
/* Table: message                                               */
/*==============================================================*/
create table message
(
    id      bigint       not null,
    send    bigint       not null,
    receive bigint       not null,
    content varchar(128) not null,
    primary key (id)
);


/*==============================================================*/
/* Table: user                                                  */
/*==============================================================*/
create table user
(
    id             bigint   not null,
    phone          char(11) not null,
    name           char(16) not null,
    follow_count   int      not null,
    follower_count int      not null,
    primary key (id)
);


/*==============================================================*/
/* Table: user_info                                             */
/*==============================================================*/
create table user_info
(
    id       bigint      not null,
    username varchar(16) not null,
    passwd   char(128)   not null,
    primary key (id)
);


/*==============================================================*/
/* Table: video                                                 */
/*==============================================================*/
create table video
(
    id             bigint       not null,
    auth_id        bigint       not null,
    titel          varchar(128) not null,
    comment_count  int          not null,
    favorite_count int          not null,
    cover_url      longtext     not null,
    play_url       varchar(256) not null,
    primary key (id)
);
