CREATE TABLE `message` (
                           `id` bigint(64) NOT NULL,
                           `to_user_id` bigint(64) NOT NULL,
                           `from_user_id` bigint(64) NOT NULL,
                           `content` varchar(128) NOT NULL,
                           `create_time` varchar(32) NOT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8