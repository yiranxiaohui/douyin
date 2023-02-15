create table users (
    id bigint(20) not null auto_increment comment 'ID',
    username varchar(64) not null comment '用户名',
    password varchar(30) not null comment '密码',
    follow_count int not null comment '关注总数',
    follower_count int not null comment '粉丝总数',
    is_follow int not null comment '是否关注',
    primary key (`id`),
    key `UK_USERNAME` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET =utf8;