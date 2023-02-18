create table favorites(
    video_id bigint(64) not null comment '喜欢的视频ID',
    user_id bigint(64) not null comment '用户ID'
)ENGINE=InnoDB DEFAULT CHARSET =utf8;