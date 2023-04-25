CREATE TABLE `image_route` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `user_id` char(64) NOT NULL COMMENT '用户账号',
     `route` char(128) NOT NULL COMMENT '图片路径',
     `enabled` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效 0无效 1有效',
     `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     KEY `utime` (`utime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8