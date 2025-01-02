/*
 Navicat Premium Data Transfer

 Source Server         : ddy-docker
 Source Server Type    : MySQL
 Source Server Version : 80402 (8.4.2)
 Source Host           : 127.0.0.1:3306
 Source Schema         : xtpl

 Target Server Type    : MySQL
 Target Server Version : 80402 (8.4.2)
 File Encoding         : 65001

 Date: 02/01/2025 14:06:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for operator
-- ----------------------------
DROP TABLE IF EXISTS `operator`;
CREATE TABLE `operator` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `status` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0' COMMENT '1 启用, 0 禁用',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1 已删除, 0 未删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of operator
-- ----------------------------
BEGIN;
INSERT INTO `operator` (`id`, `username`, `password`, `status`, `is_deleted`, `created_at`, `updated_at`) VALUES (1, 'test', '$2a$10$XH3JxoXyq9M59R7u3IK6ju2si32PvPNKRIymyJiRfy8E9VWj3YSB6', 1, 0, '2025-01-02 13:34:52', '2025-01-02 13:52:52');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
