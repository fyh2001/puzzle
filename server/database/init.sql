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
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_user_username` (`username`),
  UNIQUE INDEX `idx_user_nickname` (`nickname`),
  INDEX `idx_user_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';

DROP TABLE IF EXISTS `accolade`;
CREATE TABLE IF NOT EXISTS `accolade` (
  `id` INT(3) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` VARCHAR(20) NOT NULL COMMENT '头衔名称',
  `icon` VARCHAR(255) NOT NULL COMMENT '头衔图标',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_accolade_name` (`name`),
  INDEX `idx_accolade_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '头衔表';

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
  PRIMARY KEY (`id`),
  INDEX `idx_record_user_id` (`user_id`),
  INDEX `idx_record_dimension` (`dimension`),
  INDEX `idx_record_type` (`type`),
  INDEX `idx_record_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '记录表';

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
  PRIMARY KEY (`id`),
  INDEX `idx_record_best_single_user_id` (`user_id`),
  INDEX `idx_record_best_single_dimension` (`dimension`),
  INDEX `idx_record_best_single_record_duration` (`record_duration`),
  INDEX `idx_record_best_single_ranked` (`ranked`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳单次记录表';

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
  PRIMARY KEY (`id`),
  INDEX `idx_record_best_average_user_id` (`user_id`),
  INDEX `idx_record_best_average_dimension` (`dimension`),
  INDEX `idx_record_best_average_type` (`type`),
  INDEX `idx_record_best_average_record_average_duration` (`record_average_duration`),
  INDEX `idx_record_best_average_ranked` (`ranked`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳平均记录表';

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
  PRIMARY KEY (`id`),
  INDEX `idx_record_best_step_user_id` (`user_id`),
  INDEX `idx_record_best_step_dimension` (`dimension`),
  INDEX `idx_record_best_step_record_step` (`record_step`),
  INDEX `idx_record_best_step_ranked` (`ranked`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳步数记录表';

DROP TABLE IF EXISTS `scramble`;
CREATE TABLE IF NOT EXISTS `scramble` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `idx` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱随机数',
  `scramble` VARCHAR(255) NOT NULL COMMENT '打乱公式',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `idx_scramble_dimension` (`dimension`),
  INDEX `idx_scramble_idx` (`idx`),
  INDEX `idx_scramble_idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '打乱表';

DROP TABLE IF EXISTS `scrambled_user_status`;
CREATE TABLE IF NOT EXISTS `scrambled_user_status`(
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `scramble_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱ID',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:未完成 2:已完成',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `idx_scrambled_user_status_user_id` (`user_id`),
  INDEX `idx_scrambled_user_status_dimension` (`dimension`),
  INDEX `idx_scrambled_user_status_scramble_id` (`scramble_id`),
  INDEX `idx_scrambled_user_status_user_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户打乱状态表';

DROP TABLE IF EXISTS `notification`;
CREATE TABLE IF NOT EXISTS `notification` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID 0:全体通知',
  `type_id` int UNSIGNED NOT NULL COMMENT '类型id',
  `content` TEXT NOT NULL COMMENT '通知内容',
  `read_status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '已读状态 1:未读 2:已读',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `idx_notification_user_id` (`user_id`),
  INDEX `idx_notification_type_id` (`type_id`),
  INDEX `idx_notification_status` (`status`),
  INDEX `idx_notification_read` (`read_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '通知表';

DROP TABLE IF EXISTS `notification_type`;
CREATE TABLE IF NOT EXISTS `notification_type` (
  `id` INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` VARCHAR(20) NOT NULL COMMENT '类型名称',
  `icon` VARCHAR(255) COMMENT '类型图标',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_notification_type_name` (`name`),
  INDEX `idx_notification_type_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '通知类型表';

DROP TABLE IF EXISTS `update_log`;
CREATE TABLE IF NOT EXISTS `update_log` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `version` VARCHAR(20) NOT NULL COMMENT '版本号',
  `type` TINYINT(1) NOT NULL COMMENT '类型 1:日常更新 2:重大更新',
  `content` TEXT NOT NULL COMMENT '更新内容',
  `description` TEXT COMMENT '更新描述',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_update_log_version` (`version`),
  INDEX `idx_update_log_type` (`type`),
  INDEX `idx_update_log_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '更新日志表';