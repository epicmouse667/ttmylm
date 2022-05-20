-- 创建数据库;
create database dou_sheng;
use dou_sheng;
-- 用户
create table user(
    id int auto_increment primary key,
    name varchar(64),
    hash varchar(128),
    salt varchar(128),
    follow_count int,
    followed_count int
);
-- 视频
create table video (
    id int auto_increment primary key ,
    author_id int,
    play_url varchar(128),
    favorite_count int,
    comment_count int,
    foreign key(author_id) references user(id)
);
-- 评论
create table comment(
    id int auto_increment primary key,
    user_id int,
    video_id int,
    content varchar(1024),
    create_date datetime,
    foreign key(user_id) references user(id),
    foreign key (video_id) references video(id)
);
-- 视频-评论
create table video_comment(
    video_id int,
    comment_id int,
    primary key(comment_id,video_id),
    foreign key(video_id) references video(id),
    foreign key(comment_id) references comment(id)
);
-- up-粉丝
create table user_follow(
    user_id int,
    subscribe_id int,
    primary key (user_id,subscribe_id),
    foreign key (user_id) references user(id),
    foreign key (subscribe_id) references user(id)
);
-- 用户-视频点赞
create table user_favorate(
    user_id int,
    video_id int,
    primary key (user_id,video_id),
    foreign key (user_id) references user(id),
    foreign key (video_id) references user(id)
);
