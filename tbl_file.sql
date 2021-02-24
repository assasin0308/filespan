/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 127.0.0.1:3306
 Source Schema         : filespan

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 24/02/2021 22:48:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tbl_file
-- ----------------------------
DROP TABLE IF EXISTS `tbl_file`;
CREATE TABLE `tbl_file`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `file_sha1` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件hash',
  `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件名',
  `file_size` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `file_addr` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件存储位置',
  `create_at` datetime(0) NOT NULL COMMENT '创建日期',
  `update_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新日期',
  `status` tinyint(255) UNSIGNED NULL DEFAULT 0 COMMENT '状态(可用/禁用/已删除等状态)',
  `ext1` int(11) NOT NULL DEFAULT 0 COMMENT '备用字段1',
  `ext2` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '备用字段2',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_file_hash`(`file_sha1`) USING BTREE,
  INDEX `idx_status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
