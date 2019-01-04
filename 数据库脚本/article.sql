
-- ----------------------------
-- Table structure for announcement
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `category_id` bigint(20) UNSIGNED NOT NULL COMMENT '文章分类ID',
  `title` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `view_count` int(255) UNSIGNED NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) UNSIGNED NOT NULL COMMENT '评论次数',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  `status` int(10) UNSIGNED NOT NULL COMMENT '状态',
  `summary` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_view_count`(`view_count`) USING BTREE COMMENT '阅读次数索引',
  INDEX `idx_comment_count`(`comment_count`) USING BTREE COMMENT '评论次数索引',
  INDEX `idx_category_id`(`category_id`) USING BTREE COMMENT '分类ID索引'
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE utf8mb4_general_ci ROW_FORMAT = Dynamic;
