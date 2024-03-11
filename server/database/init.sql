CREATE DATABASE IF NOT EXISTS puzzle DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

USE puzzle;

DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `username` VARCHAR(100) NOT NULL COMMENT '用户名',
  `password` VARCHAR(255) NOT NULL COMMENT '密码',
  `nickname` VARCHAR(15) NOT NULL COMMENT '昵称',
  `accolade_id` INT(3) DEFAULT 0 COMMENT '头衔ID',
  `email` VARCHAR(50) COMMENT '邮箱',
  `phone` VARCHAR(11) COMMENT '手机号',
  `avatar` VARCHAR(255) COMMENT '头像',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';

-- 为`user`表添加唯一索引，以提高用户名的唯一性查询效率
ALTER TABLE `user` ADD UNIQUE INDEX `idx_user_username` (`username`);
ALTER TABLE `user` ADD UNIQUE INDEX `idx_user_nickname` (`nickname`);
ALTER TABLE `user` ADD INDEX `idx_user_status` (`status`);
-- ----------------------------
-- 新增admin账号
-- ----------------------------
-- BEGIN;
-- INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `phone`, `avatar`, `status`, `created_at`, `updated_at`) VALUES (271060023935762432, 'admin1', 'eb7f6a2e87831026c14971228cb27b37', '黄叶', '', '', '', 1, '2024-02-04 23:11:32', '2024-02-04 23:11:32');
-- COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `accolade`;
CREATE TABLE IF NOT EXISTS `accolade` (
  `id` INT(3) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` VARCHAR(20) NOT NULL COMMENT '头衔名称',
  `icon` VARCHAR(255) NOT NULL COMMENT '头衔图标',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '头衔表';

-- 为`accolade`表添加唯一索引，以提高头衔名称的唯一性查询效率
ALTER TABLE `accolade` ADD UNIQUE INDEX `idx_accolade_name` (`name`);
ALTER TABLE `accolade` ADD INDEX `idx_accolade_status` (`status`);

DROP TABLE IF EXISTS `record`;
CREATE TABLE IF NOT EXISTS `record` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `type` TINYINT(1) NOT NULL COMMENT '类型 1:练习 2:排行榜 3:对战',
  `duration` INT NOT NULL COMMENT '耗时',
  `step` INT NOT NULL COMMENT '步数',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `scramble` VARCHAR(255) NOT NULL COMMENT '打乱公式',
  `solution` TEXT NOT NULL COMMENT '还原公式',
  `idx` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱随机数',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '记录表';

-- 为`record`表添加索引，以提高查询效率
ALTER TABLE `record` ADD INDEX `idx_record_user_id` (`user_id`);
ALTER TABLE `record` ADD INDEX `idx_record_dimension` (`dimension`);
ALTER TABLE `record` ADD INDEX `idx_record_type` (`type`);
ALTER TABLE `record` ADD INDEX `idx_record_status` (`status`);

DROP TABLE IF EXISTS `record_best_single`;
CREATE TABLE IF NOT EXISTS `record_best_single` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `record_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '记录ID',
  `record_duration` INT NOT NULL COMMENT '记录耗时',
  `record_step` INT NOT NULL COMMENT '记录步数',
  `record_break_count` INT NOT NULL DEFAULT 1 COMMENT '打破最佳单次记录的次数',
  `ranked` INT UNSIGNED COMMENT '排名',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳单次记录表';


-- 为`record_best_single`表添加索引，以提高按事件类型进行的查询效率
ALTER TABLE `record_best_single` ADD INDEX `idx_record_best_single_user_id` (`user_id`);
ALTER TABLE `record_best_single` ADD INDEX `idx_record_best_single_dimension` (`dimension`);
ALTER TABLE `record_best_single` ADD INDEX `idx_record_best_single_record_duration` (`record_duration`);
ALTER TABLE `record_best_single` ADD INDEX `idx_record_best_single_ranked` (`ranked`);


DROP TABLE IF EXISTS `record_best_average`;
CREATE TABLE IF NOT EXISTS `record_best_average` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `type` TINYINT(5) NOT NULL COMMENT '类型 5:5次平均 12:12次平均',
  `record_ids` TEXT NOT NULL COMMENT '记录ID',
  `record_average_duration` INT NOT NULL COMMENT '记录平均用时',
  `record_break_count` INT NOT NULL DEFAULT 1 COMMENT '打破最佳平均记录的次数',
  `ranked` INT UNSIGNED COMMENT '排名',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳平均记录表';

-- 为`record_best_average`表添加索引，以提高按事件类型进行的查询效率
ALTER TABLE `record_best_average` ADD INDEX `idx_record_best_average_user_id` (`user_id`);
ALTER TABLE `record_best_average` ADD INDEX `idx_record_best_average_dimension` (`dimension`);
ALTER TABLE `record_best_average` ADD INDEX `idx_record_best_average_type` (`type`);
ALTER TABLE `record_best_average` ADD INDEX `idx_record_best_average_record_average_duration` (`record_average_duration`);
ALTER TABLE `record_best_average` ADD INDEX `idx_record_best_average_ranked` (`ranked`);


DROP TABLE IF EXISTS `record_best_step`;
CREATE TABLE IF NOT EXISTS `record_best_step` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `record_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '记录ID',
  `record_step` INT NOT NULL COMMENT '记录步数',
  `record_break_count` INT NOT NULL DEFAULT 1 COMMENT '打破最佳步数记录的次数',
  `ranked` INT UNSIGNED COMMENT '排名',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳步数记录表';

-- 为`record_best_step`表添加索引，以提高按事件类型进行的查询效率
ALTER TABLE `record_best_step` ADD INDEX `idx_record_best_step_user_id` (`user_id`);
ALTER TABLE `record_best_step` ADD INDEX `idx_record_best_step_dimension` (`dimension`);
ALTER TABLE `record_best_step` ADD INDEX `idx_record_best_step_record_step` (`record_step`);
ALTER TABLE `record_best_step` ADD INDEX `idx_record_best_step_ranked` (`ranked`);


DROP TABLE IF EXISTS `scramble`;
CREATE TABLE IF NOT EXISTS `scramble` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `idx` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱随机数',
  `scramble` VARCHAR(255) NOT NULL COMMENT '打乱公式',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '打乱表';

-- 为`scramble`表添加索引，以提高按阶数进行的查询效率
ALTER TABLE `scramble` ADD INDEX `idx_scramble_dimension` (`dimension`);
ALTER TABLE `scramble` ADD INDEX `idx_scramble_idx` (`idx`);
ALTER TABLE `scramble` ADD INDEX `idx_scramble_idx_status` (`status`);

DROP TABLE IF EXISTS `scrambled_user_status`;
CREATE TABLE IF NOT EXISTS `scrambled_user_status`(
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `scramble_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱ID',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:未完成 2:已完成',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户打乱状态表';

-- 为`scrambled_user_status`表添加索引，以提高按用户ID进行的查询效率
ALTER TABLE `scrambled_user_status` ADD INDEX `idx_scrambled_user_status_user_id` (`user_id`);
ALTER TABLE `scrambled_user_status` ADD INDEX `idx_scrambled_user_status_dimension` (`dimension`);
ALTER TABLE `scrambled_user_status` ADD INDEX `idx_scrambled_user_status_scramble_id` (`scramble_id`);
ALTER TABLE `scrambled_user_status` ADD INDEX `idx_scrambled_user_status_user_status` (`status`);

DROP TABLE IF EXISTS `notification`;
CREATE TABLE IF NOT EXISTS `notification` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID 0:全体通知',
  `type_id` int UNSIGNED NOT NULL COMMENT '类型id',
  `content` TEXT NOT NULL COMMENT '通知内容',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '通知表';

-- 为`notification`表添加索引，以提高按用户ID进行的查询效率
ALTER TABLE `notification` ADD INDEX `idx_notification_user_id` (`user_id`);
ALTER TABLE `notification` ADD INDEX `idx_notification_type_id` (`type_id`);
ALTER TABLE `notification` ADD INDEX `idx_notification_status` (`status`);

DROP TABLE IF EXISTS `notification_user_status`;
CREATE TABLE IF NOT EXISTS `notification_user_status` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `notification_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '通知ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:未读 2:已读',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '通知用户状态表';

-- 为`notification_user_status`表添加索引，以提高按用户ID进行的查询效率
ALTER TABLE `notification_user_status` ADD INDEX `idx_notification_user_status_notification_id` (`notification_id`);
ALTER TABLE `notification_user_status` ADD INDEX `idx_notification_user_status_user_id` (`user_id`);
ALTER TABLE `notification_user_status` ADD INDEX `idx_notification_user_status_status` (`status`);


DROP TABLE IF EXISTS `notification_type`;
CREATE TABLE IF NOT EXISTS `notification_type` (
  `id` INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` VARCHAR(20) NOT NULL COMMENT '类型名称',
  `icon` VARCHAR(255) COMMENT '类型图标',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '通知类型表';

-- 为`notification_type`表添加唯一索引，以提高类型名称的唯一性查询效率
ALTER TABLE `notification_type` ADD UNIQUE INDEX `idx_notification_type_name` (`name`);
ALTER TABLE `notification_type` ADD INDEX `idx_notification_type_status` (`status`);

BEGIN;
INSERT INTO `notification_type` (`id`, `name`, `icon`, `status`, `created_at`, `updated_at`) VALUES (1, '系统通知', '', 1, '2024-02-04 23:11:32', '2024-02-04 23:11:32');
COMMIT;
