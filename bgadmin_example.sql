/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50562
Source Host           : localhost:3306
Source Database       : beego_admin

Target Server Type    : MYSQL
Target Server Version : 50562
File Encoding         : 65001

Date: 2019-12-06 17:02:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `description` varchar(255) DEFAULT NULL,
  `content` longtext NOT NULL,
  `category_id` int(11) NOT NULL,
  `read` int(11) NOT NULL DEFAULT '0',
  `like` int(11) NOT NULL DEFAULT '0',
  `publish` int(11) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `publishtime` datetime DEFAULT NULL,
  `createtime` datetime NOT NULL,
  `member_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('3', 'Laravel框架', 'The PHP Framework for Web Artisans。', '<h1>Web Artisans的PHP框架</h1><p>Laravel是一个具有表达力，优雅语法的Web应用程序框架。</p><p>我们已经奠定了基础-使您可以自由创作，而不会费力小事。</p><p><br></p>', '1', '0', '0', '1', '1', null, '2019-11-29 17:37:27', '1');
INSERT INTO `article` VALUES ('4', 'Vue.js简介', 'Vue  是一套用于构建用户界面的渐进式框架。', '<p>Vue (读音 /vjuː/，类似于&nbsp;<strong>view</strong>) 是一套用于构建用户界面的<strong>渐进式框架</strong>。</p><p>与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用。</p><p>Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合。</p><p>另一方面，当与<a href=\"https://cn.vuejs.org/v2/guide/single-file-components.html\" style=\"background-color: rgb(255, 255, 255);\">现代化的工具链</a>以及各种<a href=\"https://github.com/vuejs/awesome-vue#libraries--plugins\" target=\"_blank\" rel=\"noopener\" style=\"background-color: rgb(255, 255, 255);\">支持类库</a>结合使用时，Vue 也完全能够为复杂的单页应用提供驱动。</p><p><br></p>', '4', '0', '0', '0', '1', null, '2019-11-29 17:38:39', '1');
INSERT INTO `article` VALUES ('5', 'Android Studio简介', 'Android Studio是Google推出基于IntelliJ IDEA的Android应用开发集成开发环境（IDE），', '<p>Android Studio是Google推出基于IntelliJ IDEA的Android应用开发集成开发环境（IDE），而且提供了更多提高Android应用的构建效率的功能；<br><br>&nbsp; &nbsp; &nbsp;1) 基于Gradle的灵活构建系统<br><br>&nbsp; &nbsp; &nbsp;2）Instant Run可以将变更推送到正在运行的应用中，无需重新构建Apk；<br><br>&nbsp; &nbsp; &nbsp;3）快速和功能丰富的模拟器；<br><br>&nbsp; &nbsp; &nbsp;4）丰富的测试工具、性能工具（CPU Profile和Memory Profile)和网络监控工具（Network Profiler）；<br><br>&nbsp; &nbsp; &nbsp;5）C++和NDK支持，以及LLDB可以调试原生代码；<br><br>&nbsp; &nbsp; &nbsp;6）使用Room将数据持久化数据库(SQLite)<br><br>&nbsp; &nbsp; &nbsp;7）使用apkanalyzer对预构建APK进行分析和调试<br><br>&nbsp; &nbsp; &nbsp;8）强大的布局编辑器<br><br>&nbsp; &nbsp; &nbsp;9）支持Koltin编码和Lua编码(通过NDK开发)<br><br></p><p><br></p>', '3', '0', '0', '1', '1', null, '2019-11-30 09:49:41', '3');
INSERT INTO `article` VALUES ('6', '', '', '', '0', '0', '0', '0', '0', null, '2019-12-03 15:05:07', '0');
INSERT INTO `article` VALUES ('7', '', '', '', '0', '0', '0', '0', '0', null, '2019-12-03 15:05:13', '0');

-- ----------------------------
-- Table structure for article_tags
-- ----------------------------
DROP TABLE IF EXISTS `article_tags`;
CREATE TABLE `article_tags` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `article_id` bigint(20) NOT NULL,
  `tag_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article_tags
-- ----------------------------
INSERT INTO `article_tags` VALUES ('37', '4', '28');
INSERT INTO `article_tags` VALUES ('40', '3', '6');
INSERT INTO `article_tags` VALUES ('41', '3', '19');
INSERT INTO `article_tags` VALUES ('42', '5', '4');

-- ----------------------------
-- Table structure for auth_group
-- ----------------------------
DROP TABLE IF EXISTS `auth_group`;
CREATE TABLE `auth_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `describe` varchar(255) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `rules` varchar(255) DEFAULT NULL,
  `member_id` int(11) DEFAULT NULL,
  `update_at` datetime NOT NULL,
  `create_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of auth_group
-- ----------------------------
INSERT INTO `auth_group` VALUES ('1', 'Developer', '开发者', '-2', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45', '0', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `auth_group` VALUES ('2', 'SuperAdmin', '超级管理员', '-2', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25', '1', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `auth_group` VALUES ('3', 'Admin', '管理员', '-2', '1,2,3,4,5,6,7,8,9,10,11,12,13,15,16,17,18,20,21,22,23,24', '2', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `auth_group` VALUES ('4', 'User', '普通用户', '-2', '1,2,3,4,5,6', '2', '2019-12-02 12:11:03', '2019-12-02 12:11:03');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `sort` int(11) NOT NULL DEFAULT '0',
  `types` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES ('1', '0', '后端', '1', '2', '0');
INSERT INTO `category` VALUES ('3', '0', '移动端', '1', '3', '0');
INSERT INTO `category` VALUES ('4', '0', '前端', '1', '1', '0');
INSERT INTO `category` VALUES ('5', '1', '测试类2', '1', '1', '0');
INSERT INTO `category` VALUES ('6', '5', '测试222', '1', '1', '0');
INSERT INTO `category` VALUES ('7', '6', '测试22222', '1', '1', '0');
INSERT INTO `category` VALUES ('8', '0', '操作系统', '1', '1', '0');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `writer` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL,
  `context` varchar(255) NOT NULL,
  `article_id` int(11) NOT NULL,
  `createtime` datetime NOT NULL,
  `floor` int(11) NOT NULL DEFAULT '1',
  `reply` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('39', 'hello', 'hello@163.com', 'hello world', '5', '2019-12-06 17:00:50', '1', '0');
INSERT INTO `comment` VALUES ('40', 'Nice', 'nice@qq.com', 'Nice to meet you.', '5', '2019-12-06 17:01:25', '2', '1');

-- ----------------------------
-- Table structure for front_menu
-- ----------------------------
DROP TABLE IF EXISTS `front_menu`;
CREATE TABLE `front_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `type` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '1',
  `sort` int(11) NOT NULL DEFAULT '0',
  `level` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `url` (`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of front_menu
-- ----------------------------

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(255) NOT NULL DEFAULT '',
  `member_name` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `mobile` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '1',
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `is_share` int(11) NOT NULL DEFAULT '1',
  `is_inside` int(11) NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL,
  `create_at` datetime NOT NULL,
  `auth_group_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nick_name` (`nick_name`),
  UNIQUE KEY `member_name` (`member_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of member
-- ----------------------------
INSERT INTO `member` VALUES ('1', 'developer', 'developer', '4badaee57fed5610012a296273158f5f', 'admin@admin.com', '13111111111', '-2', '0', '0', '0', '2019-12-02 12:11:03', '2019-12-02 12:11:03', '1');
INSERT INTO `member` VALUES ('2', 'super_administrator', 'super', 'e10adc3949ba59abbe56e057f20f883e', 'admin@admin.com', '13111111111', '-2', '1', '0', '0', '2019-12-02 12:11:03', '2019-12-02 12:11:03', '2');
INSERT INTO `member` VALUES ('3', 'administrator', 'admin', '96e79218965eb72c92a549dd5a330112', 'admin@admin.com', '13111111111', '-2', '2', '0', '0', '2019-12-02 12:11:03', '2019-12-02 12:11:03', '3');
INSERT INTO `member` VALUES ('4', 'TestUser', 'testuser', '96e79218965eb72c92a549dd5a330112', 'admin@admin.com', '13111111112', '-2', '2', '0', '0', '2019-12-02 12:11:03', '2019-12-02 12:11:03', '4');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `pid` int(11) NOT NULL DEFAULT '0',
  `sort` int(11) NOT NULL DEFAULT '1',
  `module` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `is_hide` int(11) NOT NULL DEFAULT '1',
  `is_shortcut` int(11) NOT NULL DEFAULT '1',
  `status` int(11) NOT NULL DEFAULT '1',
  `icon` varchar(255) DEFAULT NULL,
  `update_at` datetime NOT NULL,
  `create_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('1', '系统首页', '0', '1', 'admin', '/index', '1', '1', '-2', 'fa fa-tachometer', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('2', '基本操作', '0', '1', 'admin', '/operate', '2', '1', '-2', 'fa fa-tachometer', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('3', '权限警告', '2', '1', 'admin', '/operate/jump', '2', '1', '-2', 'fa fa-warning', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('4', '跳转提示', '2', '1', 'admin', '/operate/tips', '2', '1', '-2', 'fa fa-info', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('5', '个人信息', '2', '1', 'admin', '/operate/person', '2', '1', '-2', 'fa fa-user-secret', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('6', '修改信息', '2', '1', 'admin', '/operate/savemsg', '2', '1', '-2', 'fa fa-send', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('7', '会员管理', '0', '1', 'admin', '/member/index', '1', '1', '-2', 'fa fa-users', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('8', '权限管理', '0', '1', 'admin', '/auth/index', '1', '1', '-2', 'fa fa-key', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('9', '菜单管理', '0', '1', 'admin', '/menu/index', '1', '1', '-2', 'fa fa-th-large', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('10', '会员列表', '7', '1', 'admin', '/member/list', '1', '1', '-2', 'fa fa-user', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('11', '会员新建', '7', '1', 'admin', '/member/add', '1', '1', '-2', 'fa fa-user-plus', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('12', '会员编辑', '7', '1', 'admin', '/member/edit', '2', '1', '-2', 'fa fa-edit', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('13', '会员保存', '7', '1', 'admin', '/member/doedit', '2', '1', '-2', 'fa fa-edit', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('14', '会员删除', '7', '1', 'admin', '/member/del', '2', '1', '-2', 'fa fa-user-times', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('15', '权限组列表', '8', '1', 'admin', '/auth/list', '1', '1', '-2', 'fa fa-unlock-alt', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('16', '权限组添加', '8', '1', 'admin', '/auth/add', '1', '1', '-2', 'fa fa-plus-square', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('17', '权限组编辑', '8', '1', 'admin', '/auth/edit', '2', '1', '-2', 'fa fa-pencil-square', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('18', '权限组保存', '8', '1', 'admin', '/auth/doedit', '2', '1', '-2', 'fa fa-pencil-square', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('19', '权限组删除', '8', '1', 'admin', '/auth/del', '2', '1', '-2', 'fa fa-minus-square', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('20', '权限授权', '8', '1', 'admin', '/auth/authorize', '2', '1', '-2', 'fa fa-child', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('21', '菜单列表', '9', '1', 'admin', '/menu/list', '1', '1', '-2', 'fa fa-list-alt', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('22', '菜单新建', '9', '1', 'admin', '/menu/add', '1', '1', '-2', 'fa fa-plus-circle', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('23', '菜单编辑', '9', '1', 'admin', '/menu/edit', '2', '1', '-2', 'fa fa-pencil', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('24', '菜单保存', '9', '1', 'admin', '/menu/doedit', '2', '1', '-2', 'fa fa-pencil', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('25', '菜单删除', '9', '1', 'admin', '/menu/del', '2', '1', '-2', 'fa fa-minus-circle', '2019-12-02 12:11:03', '2019-12-02 12:11:03');
INSERT INTO `menu` VALUES ('27', 'menu2', '26', '1', 'admin', '/category/index', '1', '1', '1', 'fa fa-users', '2019-12-02 14:02:16', '2019-12-02 14:02:16');
INSERT INTO `menu` VALUES ('28', '文章管理', '0', '1', 'admin', '/article/index', '1', '1', '1', 'fa fa-book', '2019-12-03 14:59:55', '2019-12-03 14:59:55');
INSERT INTO `menu` VALUES ('29', '文章列表', '28', '1', 'admin', '/article/list', '1', '1', '1', 'fa fa-book', '2019-12-03 15:01:05', '2019-12-03 15:01:05');
INSERT INTO `menu` VALUES ('30', '分类列表', '28', '1', 'admin', '/articlecategory/list', '1', '1', '1', 'fa fa-list', '2019-12-03 15:01:44', '2019-12-03 15:01:44');
INSERT INTO `menu` VALUES ('31', '标签列表', '28', '1', 'admin', '/articletag/list', '1', '1', '1', 'fa fa-tags', '2019-12-03 15:02:31', '2019-12-03 15:02:31');
INSERT INTO `menu` VALUES ('32', '文章添加', '28', '1', 'admin', '/article/add', '2', '1', '1', 'fa fa-plus', '2019-12-03 15:03:34', '2019-12-03 15:03:34');
INSERT INTO `menu` VALUES ('33', '文章修改', '28', '1', 'admin', '/article/edit', '2', '1', '1', 'fa fa-book', '2019-12-03 15:04:10', '2019-12-03 15:04:10');
INSERT INTO `menu` VALUES ('34', '文章保存', '28', '1', 'admin', '/article/save', '2', '1', '1', 'fa fa-book', '2019-12-03 15:05:02', '2019-12-03 15:05:02');
INSERT INTO `menu` VALUES ('35', '分类添加', '28', '1', 'admin', '/articlecategory/add', '2', '1', '1', 'fa fa-list', '2019-12-03 15:06:21', '2019-12-03 15:06:21');
INSERT INTO `menu` VALUES ('36', '分类编辑', '28', '1', 'admin', '/articlecategory/edit', '2', '1', '1', 'fa fa-list', '2019-12-03 15:07:02', '2019-12-03 15:07:02');
INSERT INTO `menu` VALUES ('37', '分类保存', '28', '1', 'admin', '/articlecategory/save', '2', '1', '1', 'fa fa-list', '2019-12-03 15:07:30', '2019-12-03 15:07:30');
INSERT INTO `menu` VALUES ('38', '分类删除', '28', '1', 'admin', '/articlecategory/del', '2', '1', '1', 'fa fa-list', '2019-12-03 15:08:40', '2019-12-03 15:08:40');
INSERT INTO `menu` VALUES ('39', '文章删除', '28', '1', 'admin', '/article/del', '2', '1', '1', 'fa fa-book', '2019-12-03 15:09:19', '2019-12-03 15:09:19');
INSERT INTO `menu` VALUES ('40', '标签添加', '28', '1', 'admin', '/articletag/add', '2', '1', '1', 'fa fa-plus', '2019-12-03 15:10:21', '2019-12-03 15:10:21');
INSERT INTO `menu` VALUES ('41', '标签编辑', '28', '1', 'admin', '/articletag/edit', '2', '1', '1', 'fa fa-tag', '2019-12-03 15:10:51', '2019-12-03 15:10:51');
INSERT INTO `menu` VALUES ('42', '标签保存', '28', '1', 'admin', '/articletag/save', '2', '1', '1', 'fa fa-tag', '2019-12-03 15:11:28', '2019-12-03 15:11:28');
INSERT INTO `menu` VALUES ('43', '标签删除', '28', '1', 'admin', '/articletag/del', '2', '1', '1', 'fa fa-tag', '2019-12-03 15:12:06', '2019-12-03 15:12:06');
INSERT INTO `menu` VALUES ('44', '评论列表', '28', '1', 'admin', '/articlecomment/list', '2', '1', '1', 'fa fa-list-o', '2019-12-06 11:22:33', '2019-12-06 11:22:33');
INSERT INTO `menu` VALUES ('45', '评论删除', '28', '1', 'admin', '/articlecomment/del', '2', '1', '1', 'fa fa-list', '2019-12-06 11:23:15', '2019-12-06 11:23:15');

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '1',
  `createtime` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES ('3', 'Python', '1', '2019-11-28 15:56:42');
INSERT INTO `tag` VALUES ('4', 'Java', '1', '2019-11-28 15:56:49');
INSERT INTO `tag` VALUES ('5', 'Golang', '1', '2019-11-29 16:32:30');
INSERT INTO `tag` VALUES ('6', 'PHP', '1', '2019-11-30 16:27:07');
INSERT INTO `tag` VALUES ('12', 'Delphi', '1', '2019-12-06 16:36:34');
INSERT INTO `tag` VALUES ('13', 'C/C++', '1', '2019-12-06 16:36:50');
INSERT INTO `tag` VALUES ('14', 'C#', '1', '2019-12-06 16:36:56');
INSERT INTO `tag` VALUES ('15', '.net', '1', '2019-12-06 16:37:07');
INSERT INTO `tag` VALUES ('16', 'Beego', '1', '2019-12-06 16:37:19');
INSERT INTO `tag` VALUES ('17', 'gin', '1', '2019-12-06 16:37:28');
INSERT INTO `tag` VALUES ('18', 'thinkphp', '1', '2019-12-06 16:37:51');
INSERT INTO `tag` VALUES ('19', 'Laravel', '1', '2019-12-06 16:38:03');
INSERT INTO `tag` VALUES ('20', 'Django', '1', '2019-12-06 16:38:15');
INSERT INTO `tag` VALUES ('21', 'Linux', '1', '2019-12-06 16:38:30');
INSERT INTO `tag` VALUES ('22', 'Windows', '1', '2019-12-06 16:38:45');
INSERT INTO `tag` VALUES ('23', 'gRPC', '1', '2019-12-06 16:38:51');
INSERT INTO `tag` VALUES ('24', 'Docker', '1', '2019-12-06 16:39:35');
INSERT INTO `tag` VALUES ('25', 'MySql', '1', '2019-12-06 16:39:54');
INSERT INTO `tag` VALUES ('26', 'Redis', '1', '2019-12-06 16:40:04');
INSERT INTO `tag` VALUES ('27', 'Mongo', '1', '2019-12-06 16:41:23');
INSERT INTO `tag` VALUES ('28', 'JavaScript', '1', '2019-12-06 16:41:33');


