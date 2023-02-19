CREATE TABLE `comment` (
                           `user_id` bigint(64) NOT NULL COMMENT '用户id',
                           `id` bigint(64) NOT NULL COMMENT '评论id',
                           `text` varchar(128) NOT NULL COMMENT '评论内容',
                           `create_date` bigint(64) NOT NULL COMMENT '提交时间',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8