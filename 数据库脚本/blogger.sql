/*
SQLyog Professional v12.09 (64 bit)
MySQL - 5.6.40-log : Database - blogger
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`blogger` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `blogger`;

/*Table structure for table `article` */

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `category_id` bigint(20) unsigned NOT NULL COMMENT '文章分类ID',
  `title` varchar(1024) NOT NULL COMMENT '文章标题',
  `content` longtext NOT NULL COMMENT '文章内容',
  `view_count` int(255) unsigned NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) unsigned NOT NULL COMMENT '评论次数',
  `username` varchar(128) NOT NULL COMMENT '作者',
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_view_count` (`view_count`) USING BTREE COMMENT '阅读次数索引',
  KEY `idx_comment_count` (`comment_count`) USING BTREE COMMENT '评论次数索引',
  KEY `idx_category_id` (`category_id`) USING BTREE COMMENT '分类ID索引'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

/*Data for the table `article` */

insert  into `article`(`id`,`category_id`,`title`,`content`,`view_count`,`comment_count`,`username`,`status`,`summary`,`create_time`,`update_time`) values (1,1,'这里是文章的标题','文章内容',1,1,'caicai',1,'文章摘要','2018-12-13 16:41:38','2018-12-13 16:41:38');

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `category_name` varchar(255) NOT NULL COMMENT '分类名称',
  `category_no` int(10) unsigned NOT NULL COMMENT '分类排序',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

/*Data for the table `category` */

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL,
  `content` text NOT NULL COMMENT '留言内容',
  `username` varchar(255) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `article_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

/*Data for the table `comment` */

/*Table structure for table `leave` */

DROP TABLE IF EXISTS `leave`;

CREATE TABLE `leave` (
  `id` bigint(20) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `content` text NOT NULL COMMENT '留言内容',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

/*Data for the table `leave` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
