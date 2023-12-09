/*
项目建表语句
 */

CREATE TABLE `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '账号名称',
    `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '账号密码',
    `type` int unsigned NOT NULL DEFAULT '0' COMMENT '账号类型',
    `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `pictures` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '效果图名称',
    `user` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '账号名称',
    `picture` mediumblob NOT NULL COMMENT '效果图',
    `description` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '文案',
    `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `user_index` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
