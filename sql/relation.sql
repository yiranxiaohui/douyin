create table follow_list(
      follower_id bigint(64) not null comment '关注人ID',
      user_id  bigint(64) not null comment '用户ID'
)ENGINE=InnoDB DEFAULT CHARSET =utf8;