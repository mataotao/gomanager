/*
 Navicat Premium Data Transfer

 Source Server         : go
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : localhost:3306
 Source Schema         : db_apiserver

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 01/09/2018 23:15:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for manager_permission
-- ----------------------------
DROP TABLE IF EXISTS `manager_permission`;
CREATE TABLE `manager_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `icon` varchar(255) DEFAULT NULL COMMENT 'Icon\n',
  `label` varchar(255) NOT NULL COMMENT '菜单名',
  `is_contain_menu` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是包含类型的菜单1是0否',
  `pid` int(11) DEFAULT '0' COMMENT '父id',
  `url` varchar(255) DEFAULT NULL COMMENT '菜单的url',
  `level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '层级',
  `cond` varchar(2000) DEFAULT NULL COMMENT '条件',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of manager_permission
-- ----------------------------
BEGIN;
INSERT INTO `manager_permission` VALUES (1, 'el-icon-setting', '系统设置', 1, 0, '', 1, ' ', 500, '2018-08-11 14:31:37', '2018-09-01 14:44:39');
INSERT INTO `manager_permission` VALUES (2, NULL, '权限管理', 0, 1, '/manager/menu', 2, 'apiserver/handler/admin/manager/permission.List,apiserver/handler/admin/manager/permission.Get,apiserver/handler/admin/manager/permission.Update,apiserver/handler/admin/manager/permission.Delete,apiserver/handler/admin/manager/permission.Create', 500, '2018-08-11 14:32:10', '2018-08-30 10:38:32');
INSERT INTO `manager_permission` VALUES (3, NULL, '角色管理', 0, 1, '/manager/role', 2, 'apiserver/handler/admin/manager/role.List,apiserver/handler/admin/manager/role.Get,apiserver/handler/admin/manager/role.Update,apiserver/handler/admin/manager/role.Delete,apiserver/handler/admin/manager/role.Create', 500, '2018-08-11 14:34:40', '2018-08-30 10:40:31');
INSERT INTO `manager_permission` VALUES (4, NULL, '账号管理', 0, 1, '/manager/user', 2, 'apiserver/handler/admin/manager/user.List,apiserver/handler/admin/manager/user.Get,apiserver/handler/admin/manager/user.Update,apiserver/handler/admin/manager/user.Freeze,apiserver/handler/admin/manager/user.Create,apiserver/handler/admin/manager/user.Pwd', 500, '2018-08-11 14:36:19', '2018-08-30 10:41:47');
INSERT INTO `manager_permission` VALUES (5, NULL, '日志管理', 0, 1, '', 2, 'apiserver/handler/admin/manager/user.List,apiserver/handler/admin/manager/user.Get,apiserver/handler/admin/manager/user.Update,apiserver/handler/admin/manager/user.Freeze,apiserver/handler/admin/manager/user.Create,apiserver/handler/admin/manager/user.Pwd', 500, '2018-08-11 14:38:07', '2018-08-31 01:15:37');
INSERT INTO `manager_permission` VALUES (6, NULL, '其他', 1, 0, '', 1, '', 500, '2018-09-01 14:19:07', '2018-09-01 14:19:07');
INSERT INTO `manager_permission` VALUES (7, 'el-icon-tickets', '系统首页', 2, 6, '/dashboard', 2, '', 500, '2018-09-01 14:21:44', '2018-09-01 14:45:00');
INSERT INTO `manager_permission` VALUES (8, NULL, '基础表格', 2, 6, '/table', 2, '', 500, '2018-09-01 14:22:04', '2018-09-01 14:23:03');
INSERT INTO `manager_permission` VALUES (9, NULL, 'tab选项卡', 2, 6, '/tabs', 2, '', 500, '2018-09-01 14:22:23', '2018-09-01 14:23:00');
INSERT INTO `manager_permission` VALUES (10, NULL, '基本表单', 2, 6, '/form', 2, '', 500, '2018-09-01 14:22:37', '2018-09-01 14:22:55');
INSERT INTO `manager_permission` VALUES (11, NULL, '富文本编辑器', 2, 6, '/editor', 2, '', 500, '2018-09-01 14:23:23', '2018-09-01 14:23:51');
INSERT INTO `manager_permission` VALUES (12, NULL, 'markdown编辑器', 2, 6, '/markdown', 2, '', 500, '2018-09-01 14:23:47', '2018-09-01 14:23:47');
INSERT INTO `manager_permission` VALUES (13, NULL, '文件上传', 2, 6, '/upload', 2, '', 500, '2018-09-01 14:24:09', '2018-09-01 14:24:15');
INSERT INTO `manager_permission` VALUES (14, NULL, 'schart图表', 2, 6, '/charts', 2, '', 500, '2018-09-01 14:24:48', '2018-09-01 14:24:48');
INSERT INTO `manager_permission` VALUES (15, NULL, '拖拽列表', 2, 6, '/drag', 2, '', 500, '2018-09-01 14:26:55', '2018-09-01 14:26:55');
INSERT INTO `manager_permission` VALUES (18, '', '测试', 2, 0, '', 1, '', 500, '2018-09-01 15:11:34', '2018-09-01 15:11:34');
COMMIT;

-- ----------------------------
-- Table structure for manager_role
-- ----------------------------
DROP TABLE IF EXISTS `manager_role`;
CREATE TABLE `manager_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '名称',
  `description` varchar(255) DEFAULT NULL COMMENT '简介',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of manager_role
-- ----------------------------
BEGIN;
INSERT INTO `manager_role` VALUES (1, '系统管理员', '系统管理.........', '2018-08-12 04:38:19', '2018-08-12 04:38:28');
INSERT INTO `manager_role` VALUES (2, 'ces', 'ces', '2018-08-25 11:00:25', '2018-08-25 11:00:25');
COMMIT;

-- ----------------------------
-- Table structure for manager_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `manager_role_permission`;
CREATE TABLE `manager_role_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `permission_id` int(11) NOT NULL COMMENT '权限id',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `role_id` (`role_id`,`permission_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of manager_role_permission
-- ----------------------------
BEGIN;
INSERT INTO `manager_role_permission` VALUES (10, 1, 1, NULL, NULL);
INSERT INTO `manager_role_permission` VALUES (11, 1, 2, '2018-08-31 09:40:25', '2018-08-31 09:40:25');
INSERT INTO `manager_role_permission` VALUES (12, 1, 3, NULL, NULL);
INSERT INTO `manager_role_permission` VALUES (13, 1, 4, NULL, NULL);
INSERT INTO `manager_role_permission` VALUES (14, 1, 5, '2018-08-30 02:38:23', '2018-08-30 02:38:23');
INSERT INTO `manager_role_permission` VALUES (15, 2, 1, NULL, NULL);
INSERT INTO `manager_role_permission` VALUES (16, 2, 2, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for manager_user
-- ----------------------------
DROP TABLE IF EXISTS `manager_user`;
CREATE TABLE `manager_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `mobile` varchar(20) NOT NULL COMMENT '用户手机号',
  `password` varchar(100) NOT NULL COMMENT '用户密码',
  `head_img` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `last_time` datetime DEFAULT NULL COMMENT '上次登录时间',
  `last_ip` varchar(20) DEFAULT NULL COMMENT '上次登录ip',
  `is_root` tinyint(4) DEFAULT '0' COMMENT '是否为超级管理员1是0否',
  `status` tinyint(4) DEFAULT '1' COMMENT '用户状态1正常0冻结',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='后台用户表';

-- ----------------------------
-- Records of manager_user
-- ----------------------------
BEGIN;
INSERT INTO `manager_user` VALUES (1, '111', 'ces11', '28615944108', '$2a$10$8PScrOt36kzuUfay4UM8qu5lf4cqCK4Z7fJ8qlrMNMEEiqncYy9xq', '3/dsadasd/dsadasda.jpg', '2018-09-01 15:12:02', '172.18.0.1', 2, 1, '2018-08-15 13:13:18', '2018-09-01 15:12:02');
INSERT INTO `manager_user` VALUES (2, '测试111222', '姓名1111', '18615944108', '$2a$10$Eb.8U5H30Tlk/lXlnTNbFuQZf35HsdMhkQXr7ynDdtitjbl7wLnWm', 'dsadasd/dsadasda.jpg', '2018-08-27 07:10:23', '172.19.0.1', 0, 1, '2018-08-16 06:24:29', '2018-08-27 07:10:23');
INSERT INTO `manager_user` VALUES (3, '测试11122211', '姓名1111', '18615944108', '$2a$10$qHrz0XFLROCvz1C32zrwiO0tOrC/dayWV4lhMAqJCuoBK3sGsKQt6', 'dsadasd/dsadasda.jpg', '2018-08-27 07:10:23', '172.19.0.1', 0, 1, '2018-08-21 08:28:01', '2018-08-27 07:10:23');
INSERT INTO `manager_user` VALUES (4, 'dsadas', 'das', '18681944981', '$2a$10$xw7Bs2RI6/UlSsJmoJfY7OeDyVCp8nDL3.zN9.G0khGQ4UwH0Ifba', 'uploads/2018/08/3ef9d291-a903-11e8-b611-0242ac13000a.jpg', '2018-08-27 07:10:23', '172.19.0.1', 0, 1, '2018-08-26 07:40:09', '2018-08-27 07:10:23');
INSERT INTO `manager_user` VALUES (5, '测试111', 'xiuv测试赛1', '27628293814', '$2a$10$pawSJEf6ofjnxPHHcNxw7.6y5Z5IlmWAHHC9RMH8sA6sxbKO15MWq', 'uploads/2018/08/8a5cfb7f-a903-11e8-b611-0242ac13000a.jpg', '2018-08-27 07:10:23', '172.19.0.1', 0, 2, '2018-08-26 07:42:16', '2018-08-27 07:10:23');
COMMIT;

-- ----------------------------
-- Table structure for manager_user_role
-- ----------------------------
DROP TABLE IF EXISTS `manager_user_role`;
CREATE TABLE `manager_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '管理员id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `member_id` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of manager_user_role
-- ----------------------------
BEGIN;
INSERT INTO `manager_user_role` VALUES (22, 1, 1);
INSERT INTO `manager_user_role` VALUES (4, 2, 1);
INSERT INTO `manager_user_role` VALUES (25, 3, 1);
INSERT INTO `manager_user_role` VALUES (27, 4, 1);
INSERT INTO `manager_user_role` VALUES (26, 4, 2);
INSERT INTO `manager_user_role` VALUES (31, 5, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
