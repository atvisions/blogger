
-- ----------------------------
-- Table structure for announcement
-- ----------------------------
DROP TABLE IF EXISTS `announcement`;
CREATE TABLE `announcement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `imgUrl` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL,
  `detailUrl` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL,
  `createDate` varchar(10) COLLATE utf8_unicode_ci DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
 
-- ----------------------------
-- Records of announcement
-- ----------------------------
INSERT INTO `announcement` VALUES ('1', '/visitshop/img/ann/ann1.jpg', null, '2016-07-20', '0');
INSERT INTO `announcement` VALUES ('2', '/visitshop//img/ann/ann1.jpg', null, '2016-07-20', '0');
INSERT INTO `announcement` VALUES ('3', '/visitshop//img/ann/ann1.jpg', null, '2016-07-20', '0');
INSERT INTO `announcement` VALUES ('4', '/visitshop//img/ann/ann1.jpg', null, '2016-07-20', '0');
