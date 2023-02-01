use go_to_byte;

set FOREIGN_KEY_CHECKS = 0;

drop index user_id on comment;
drop index video_id on comment;
drop index video_id on favorite;
drop index user_id on favorite;
drop index user_id on message;
drop index to_user_id on relation;
drop index user_id on relation;

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
drop table if exists video;
drop table if exists relation;

set FOREIGN_KEY_CHECKS = 1;


/*==============================================================*/
/* Table: comment                                               */
/*==============================================================*/
create table comment
(
    id       bigint       not null,
    video_id bigint       not null,
    user_id  bigint       not null,
    content  varchar(128) not null,
    primary key (id)
);

/*==============================================================*/
/* Index: video_id                                              */
/*==============================================================*/
create index video_id on comment (video_id);

/*==============================================================*/
/* Index: user_id                                               */
/*==============================================================*/
create index user_id on comment (user_id);

/*==============================================================*/
/* Table: favorite                                              */
/*==============================================================*/
create table favorite
(
    id       bigint not null,
    user_id  bigint not null,
    video_id bigint not null,
    flag     bool   not null,
    primary key (id)
);

/*==============================================================*/
/* Index: user_id                                               */
/*==============================================================*/
create index user_id on favorite (user_id);

/*==============================================================*/
/* Index: video_id                                              */
/*==============================================================*/
create index video_id on favorite (video_id);

/*==============================================================*/
/* Table: message                                               */
/*==============================================================*/
create table message
(
    id         bigint       not null,
    user_id    bigint       not null,
    to_user_id bigint       not null,
    content    varchar(128) not null,
    primary key (id)
);

/*==============================================================*/
/* Index: user_id                                               */
/*==============================================================*/
create index user_id on message(user_id);

/*==============================================================*/
/* Table: relation                                              */
/*==============================================================*/
create table relation
(
    id         bigint not null,
    user_id    bigint not null,
    to_user_id bigint not null,
    flag       bool   not null,
    primary key (id)
);

/*==============================================================*/
/* Index: user_id                                               */
/*==============================================================*/
create index user_id on relation (user_id);

/*==============================================================*/
/* Index: to_user_id                                            */
/*==============================================================*/
create index to_user_id on relation (to_user_id);

/*==============================================================*/
/* Table: user                                                  */
/*==============================================================*/
create table user
(
    id       bigint    not null,
    username char(16)  not null,
    passwd   char(128) not null,
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

