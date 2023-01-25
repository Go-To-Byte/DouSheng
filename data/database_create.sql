use go_to_byte;

set FOREIGN_KEY_CHECKS=0;

drop table if exists user;

drop table if exists user_follow;

drop table if exists user_follower;

drop table if exists user_info;

drop table if exists user_message;

drop table if exists video_comment;

drop table if exists video_favorite;

drop table if exists video_info;

set FOREIGN_KEY_CHECKS=1;


/*==============================================================*/
/* Table: user                                                  */
/*==============================================================*/
create table user
(
    user                 bigint not null,
    username             varchar(16),
    phone                char(11) not null,
    name                 char(16) not null,
    follow_count         int not null,
    follower_count       int not null,
    primary key (user)
);

/*==============================================================*/
/* Table: user_follow                                           */
/*==============================================================*/
create table user_follow
(
    follow_id1           bigint not null,
    follow_id2           bigint not null,
    follow_flag          bool not null,
    primary key (follow_id2, follow_id1)
);

/*==============================================================*/
/* Table: user_follower                                         */
/*==============================================================*/
create table user_follower
(
    follower_id1         bigint not null,
    follower_id2         bigint not null,
    follower_flag        bool not null,
    primary key (follower_id1, follower_id2)
);

/*==============================================================*/
/* Table: user_info                                             */
/*==============================================================*/
create table user_info
(
    username             varchar(16) not null,
    user                 bigint,
    passwd               char(128) not null,
    user_id              bigint not null,
    token                char(128),
    token_time           bigint,
    primary key (username)
);

/*==============================================================*/
/* Table: user_message                                          */
/*==============================================================*/
create table user_message
(
    message_id           bigint not null,
    user_id1             bigint not null,
    user_id2             bigint not null,
    message_content      varchar(128) not null,
    primary key (message_id)
);

/*==============================================================*/
/* Table: video_comment                                         */
/*==============================================================*/
create table video_comment
(
    comment_id           bigint not null,
    comment_video        bigint not null,
    comment_user         bigint not null,
    comment              varchar(128) not null,
    primary key (comment_id)
);

/*==============================================================*/
/* Table: video_favorite                                        */
/*==============================================================*/
create table video_favorite
(
    favorite_user        bigint not null,
    favorite_video       bigint not null,
    favorite_flag        bool not null,
    primary key (favorite_user, favorite_video)
);

/*==============================================================*/
/* Table: video_info                                            */
/*==============================================================*/
create table video_info
(
    video_id             bigint not null,
    auth_id              bigint not null,
    titel                varchar(128) not null,
    comment_count        int not null,
    favorite_count       int not null,
    cover_url            longtext not null,
    play_url             varchar(256) not null,
    primary key (video_id)
);

alter table user add constraint FK_id foreign key (username)
    references user_info (username) on delete restrict on update restrict;

alter table user_follow add constraint FK_follow foreign key (follow_id1)
    references user (user) on delete restrict on update restrict;

alter table user_follower add constraint FK_follower foreign key (follower_id1)
    references user (user) on delete restrict on update restrict;

alter table user_info add constraint FK_id2 foreign key (user)
    references user (user) on delete restrict on update restrict;

alter table user_message add constraint FK_message1 foreign key (user_id1)
    references user (user) on delete restrict on update restrict;

alter table user_message add constraint FK_message2 foreign key (user_id2)
    references user (user) on delete restrict on update restrict;

alter table video_comment add constraint FK_comment foreign key (comment_video)
    references video_info (video_id) on delete restrict on update restrict;

alter table video_comment add constraint FK_commenter foreign key (comment_user)
    references user (user) on delete restrict on update restrict;

alter table video_favorite add constraint FK_favorite foreign key (favorite_user)
    references video_info (video_id) on delete restrict on update restrict;

alter table video_info add constraint FK_video foreign key (auth_id)
    references user (user) on delete restrict on update restrict;