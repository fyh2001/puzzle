CREATE DATABASE IF NOT EXISTS puzzle DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

USE puzzle;

DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `username` VARCHAR(100) NOT NULL COMMENT '用户名',
  `password` VARCHAR(255) NOT NULL COMMENT '密码',
  `nickname` VARCHAR(15) NOT NULL COMMENT '昵称',
  `email` VARCHAR(50) COMMENT '邮箱',
  `phone` VARCHAR(11) COMMENT '手机号',
  `avatar` VARCHAR(255) COMMENT '头像',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态 0:启用 1:冻结 2:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';

-- 为`user`表的`username`列添加唯一索引，以提高用户名的唯一性查询效率
ALTER TABLE `user` ADD UNIQUE INDEX `idx_username` (`username`);

-- ----------------------------
-- 新增admin账号
-- ----------------------------
-- BEGIN;
-- INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `phone`, `avatar`, `status`, `created_at`, `updated_at`) VALUES (271060023935762432, 'admin1', 'eb7f6a2e87831026c14971228cb27b37', '黄叶', '', '', '', 0, '2024-02-04 23:11:32', '2024-02-04 23:11:32');
-- COMMIT;

SET FOREIGN_KEY_CHECKS = 1;


DROP TABLE IF EXISTS `record`;
CREATE TABLE IF NOT EXISTS `record` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `type` TINYINT(1) NOT NULL COMMENT '类型 0:个人 1:排行榜 2:对战',
  `duration` INT NOT NULL COMMENT '耗时',
  `step` INT NOT NULL COMMENT '步数',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态 0:启用 1:冻结 2:删除',
  `scramble` VARCHAR(255) NOT NULL COMMENT '打乱公式',
  `solution` TEXT NOT NULL COMMENT '还原公式',
  `idx` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱随机数',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '记录表';

-- 为`record`表添加索引，以提高按用户ID进行的查询效率
ALTER TABLE `record` ADD INDEX `idx_user_id` (`user_id`);
ALTER TABLE `record` ADD INDEX `dimension` (`dimension`);
ALTER TABLE `record` ADD INDEX `idx_type` (`type`);
ALTER TABLE `record` ADD INDEX `idx_status` (`status`);

DROP TABLE IF EXISTS `record_best_single`;
CREATE TABLE IF NOT EXISTS `record_best_single` (
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `record_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '记录ID',
  `record_duration` INT NOT NULL COMMENT '记录耗时',
  `record_step` INT NOT NULL COMMENT '记录步数',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳单次记录表';


-- 为`record_best_single`表添加索引，以提高按事件类型进行的查询效率
ALTER TABLE `record_best_single` ADD INDEX `idx_user_id` (`user_id`);
ALTER TABLE `record_best_single` ADD INDEX `dimension` (`dimension`);


DROP TABLE IF EXISTS `record_best_average`;
CREATE TABLE IF NOT EXISTS `record_best_average` (
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `type` TINYINT(5) NOT NULL COMMENT '类型 5:5次平均 12:12次平均',
  `record_ids` TEXT NOT NULL COMMENT '记录ID',
  `record_average_duration` INT NOT NULL COMMENT '记录平均用时',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳平均记录表';

-- 为`record_best_average`表添加索引，以提高按事件类型进行的查询效率
ALTER TABLE `record_best_average` ADD INDEX `idx_user_id` (`user_id`);
ALTER TABLE `record_best_average` ADD INDEX `dimension` (`dimension`);
ALTER TABLE `record_best_average` ADD INDEX `idx_type` (`type`);

DROP TABLE IF EXISTS `record_best_step`;
CREATE TABLE IF NOT EXISTS `record_best_step` (
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `record_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '记录ID',
  `record_step` INT NOT NULL COMMENT '记录步数',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '最佳步数记录表';


-- 为`record_best_step`表添加索引，以提高按事件类型进行的查询效率
ALTER TABLE `record_best_step` ADD INDEX `idx_user_id` (`user_id`);
ALTER TABLE `record_best_step` ADD INDEX `dimension` (`dimension`);

DROP TABLE IF EXISTS `scramble`;
CREATE TABLE IF NOT EXISTS `scramble` (
  `id` BIGINT(20) UNSIGNED NOT NULL COMMENT '主键ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `idx` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱随机数',
  `scramble` VARCHAR(255) NOT NULL COMMENT '打乱公式',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态 0:启用 1:冻结 2:删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '打乱表';

-- 为`scramble`表添加索引，以提高按阶数进行的查询效率
ALTER TABLE `scramble` ADD INDEX `dimension` (`dimension`);
ALTER TABLE `scramble` ADD INDEX `idx` (`idx`);
ALTER TABLE `scramble` ADD INDEX `idx_status` (`status`);

DROP TABLE IF EXISTS `scrambled_user_status`;
CREATE TABLE IF NOT EXISTS `scrambled_user_status`(
  `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `dimension` TINYINT(1) NOT NULL COMMENT '阶数 3 | 4 | 5 | 6',
  `scramble_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '打乱ID',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态 0:未完成 1:已完成',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户打乱状态表';

-- 为`scrambled_user_status`表添加索引，以提高按用户ID进行的查询效率
ALTER TABLE `scrambled_user_status` ADD INDEX `idx_user_id` (`user_id`);
ALTER TABLE `scrambled_user_status` ADD INDEX `dimension` (`dimension`);
ALTER TABLE `scrambled_user_status` ADD INDEX `scramble_id` (`scramble_id`);
ALTER TABLE `scrambled_user_status` ADD INDEX `user_status` (`status`);





