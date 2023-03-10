create table videos(
    id bigint(20) not null auto_increment comment 'ID',
    use_id varchar(64) not null comment '作者ID',
    play_url varchar(128) not null comment '播放地址',
    cover_url varchar(128) not null comment '封面地址',
    comment_count int not null comment '评论数',
    title nvarchar(32) not null comment '视频标题',
    release_time bigint(64) not null comment '提交时间',
    primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET =utf8;