/*
Navicat MySQL Data Transfer

Source Server         : link
Source Server Version : 80015
Source Host           : localhost:3306
Source Database       : beegoblog

Target Server Type    : MYSQL
Target Server Version : 80015
File Encoding         : 65001

Date: 2019-11-15 17:47:30
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `category`
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `views` bigint(20) NOT NULL DEFAULT '0',
  `topic_time` datetime NOT NULL,
  `topic_count` bigint(20) NOT NULL DEFAULT '0',
  `topic_last_user_id` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`,`title`),
  UNIQUE KEY `title` (`title`),
  UNIQUE KEY `title_2` (`title`),
  KEY `category_created` (`created`),
  KEY `category_views` (`views`),
  KEY `category_topic_time` (`topic_time`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES ('12', 'Java', '2019-05-08 08:54:53', '0', '2019-05-08 08:54:53', '0', '0');
INSERT INTO `category` VALUES ('13', 'Python', '2019-05-08 08:54:55', '0', '2019-05-08 08:54:55', '0', '0');
INSERT INTO `category` VALUES ('14', 'Beego', '2019-05-08 08:54:57', '0', '2019-05-08 08:54:57', '0', '0');
INSERT INTO `category` VALUES ('19', 'Go', '2019-05-08 09:37:41', '0', '2019-05-08 09:37:41', '0', '0');
INSERT INTO `category` VALUES ('22', 'PHP', '2019-05-08 16:43:20', '0', '2019-05-08 16:43:20', '0', '0');
INSERT INTO `category` VALUES ('26', 'C', '2019-05-10 19:00:43', '0', '2019-05-10 19:00:43', '0', '0');
INSERT INTO `category` VALUES ('27', 'Javascript', '2019-05-10 20:20:32', '0', '2019-05-10 20:20:32', '0', '0');

-- ----------------------------
-- Table structure for `comment`
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `website` varchar(255) NOT NULL DEFAULT '',
  `content` varchar(1000) NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `topic_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `comment_created_time` (`created_time`) USING BTREE,
  KEY `topic_id` (`topic_id`),
  CONSTRAINT `topic_id` FOREIGN KEY (`topic_id`) REFERENCES `topic` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('12', '哈哈', '1765770246@qq.com', 'no', '好', '2019-05-08 14:46:41', '9');
INSERT INTO `comment` VALUES ('13', '1', '2', '3', '4', '2019-05-08 14:55:44', '9');
INSERT INTO `comment` VALUES ('14', '', '', '', '', '2019-05-08 15:03:33', '9');
INSERT INTO `comment` VALUES ('15', '1', '2', '3', '', '2019-05-08 17:24:31', '10');
INSERT INTO `comment` VALUES ('16', '嗨', '222', '222', '嗨', '2019-05-09 23:14:14', '9');
INSERT INTO `comment` VALUES ('17', '1', '2', '3', '', '2019-05-10 19:12:27', '9');
INSERT INTO `comment` VALUES ('18', '1', '1765770246@qq.com', '1', '1', '2019-05-10 19:18:44', '9');

-- ----------------------------
-- Table structure for `topic`
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL DEFAULT '',
  `content` varchar(5000) NOT NULL DEFAULT '',
  `attachment` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `views` bigint(20) NOT NULL DEFAULT '0',
  `author` varchar(255) NOT NULL DEFAULT '',
  `reply_time` datetime NOT NULL,
  `reply_count` bigint(20) NOT NULL DEFAULT '0',
  `reply_last_user_id` bigint(20) NOT NULL DEFAULT '0',
  `category_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `topic_created` (`created`),
  KEY `topic_updated` (`updated`),
  KEY `topic_views` (`views`),
  KEY `topic_reply_time` (`reply_time`),
  KEY `fk1` (`category_id`),
  CONSTRAINT `fk1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of topic
-- ----------------------------
INSERT INTO `topic` VALUES ('9', '0', 'Java', '<p>Java2</p><p><br></p>', '', '2019-04-01 06:46:10', '2019-05-07 22:48:39', '2', '', '2019-05-07 06:46:10', '5', '0', '12');
INSERT INTO `topic` VALUES ('10', '0', 'Python', '<p>Py1</p><p><br></p>', '', '2019-04-01 07:33:53', '2019-05-10 19:06:24', '1', '', '2019-05-08 07:33:53', '1', '0', '13');
INSERT INTO `topic` VALUES ('11', '0', 'Spring', '<p>Spring</p>', '', '2019-04-01 15:34:03', '2019-05-08 15:34:03', '0', '', '2019-05-08 15:34:03', '0', '0', '12');
INSERT INTO `topic` VALUES ('12', '0', 'Mybatis', '<p>Mybatis1</p><p><br></p>', '', '2019-05-08 15:34:20', '2019-05-09 22:27:58', '1', '', '2019-05-08 15:34:20', '0', '0', '12');
INSERT INTO `topic` VALUES ('13', '0', 'Hibernate', '<p>Hibernate</p>', '', '2019-04-01 16:25:17', '2019-05-08 16:25:17', '0', '', '2019-05-08 16:25:17', '0', '0', '12');
INSERT INTO `topic` VALUES ('14', '0', 'Beego', '<p>Beego</p>', '', '2019-03-01 16:27:45', '2019-05-08 16:27:45', '0', '', '2019-05-08 16:27:45', '0', '0', '14');
INSERT INTO `topic` VALUES ('15', '0', 'SpringBoot', '<p>SpringBoot<br></p>', '', '2019-05-08 16:40:55', '2019-05-08 16:40:55', '0', '', '2019-05-08 16:40:55', '0', '0', '12');
INSERT INTO `topic` VALUES ('16', '0', 'SpringCloud', '<p>SpringCloud</p>', '', '2019-05-08 16:41:04', '2019-05-08 16:41:04', '0', '', '2019-05-08 16:41:04', '0', '0', '12');
INSERT INTO `topic` VALUES ('17', '0', 'SpringMVC', '<p>SpringMVC</p>', '', '2019-05-08 16:41:31', '2019-05-08 16:41:31', '0', '', '2019-05-08 16:41:31', '0', '0', '12');
INSERT INTO `topic` VALUES ('20', '0', 'Python', '<h1>Python</h1>', '', '2019-05-09 09:39:16', '2019-05-09 09:39:16', '0', '', '2019-05-09 09:39:16', '0', '0', '13');
INSERT INTO `topic` VALUES ('23', '0', 'Go', '<p>1<br></p>', 'spring-1.jpeg', '2019-05-10 09:10:19', '2019-05-10 09:10:19', '0', '', '2019-05-10 09:10:19', '0', '0', '19');
