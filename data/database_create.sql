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
/* Table: user                                                  */
/*==============================================================*/
create table user
(
    user_id        bigint   not null,
    phone          char(11) not null,
    name           char(16) not null,
    follow_count   int      not null,
    follower_count int      not null,
    primary key (user_id)
);


/*==============================================================*/
/* Table: user_follow                                           */
/*==============================================================*/
create table user_follow
(
    follow_id1  bigint not null,
    follow_id2  bigint not null,
    follow_flag bool   not null,
    primary key (follow_id2, follow_id1)
);


/*==============================================================*/
/* Table: user_follower                                         */
/*==============================================================*/
create table user_follower
(
    follower_id1  bigint not null,
    follower_id2  bigint not null,
    follower_flag bool   not null,
    primary key (follower_id1, follower_id2)
);


/*==============================================================*/
/* Table: user_info                                             */
/*==============================================================*/
create table user_info
(
    user_id  bigint      not null,
    username varchar(16) not null,
    passwd   char(128)   not null,
    primary key (user_id)
);


/*==============================================================*/
/* Table: user_message                                          */
/*==============================================================*/
create table user_message
(
    message_id      bigint       not null,
    user_id1        bigint       not null,
    user_id2        bigint       not null,
    message_content varchar(128) not null,
    primary key (message_id)
);


/*==============================================================*/
/* Table: video_comment                                         */
/*==============================================================*/
create table video_comment
(
    comment_id    bigint       not null,
    comment_video bigint       not null,
    comment_user  bigint       not null,
    comment       varchar(128) not null,
    primary key (comment_id)
);



/*==============================================================*/
/* Table: video_favorite                                        */
/*==============================================================*/
create table video_favorite
(
    favorite_user  bigint not null,
    favorite_video bigint not null,
    favorite_flag  bool   not null,
    primary key (favorite_user, favorite_video)
);


/*==============================================================*/
/* Table: video_info                                            */
/*==============================================================*/
create table video_info
(
    video_id       bigint       not null,
    auth_id        bigint       not null,
    titel          varchar(128) not null,
    comment_count  int          not null,
    favorite_count int          not null,
    cover_url      longtext     not null,
    play_url       varchar(256) not null,
    primary key (video_id)
);
