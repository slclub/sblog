
-- 会员表
DROP TABLE IF EXISTS `s_member`;
CREATE TABLE `s_member` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NOT NULL DEFAULT '',
    `password` VARCHAR(64) NOT NULL DEFAULT '',
    `name` VARCHAR(64) NULL DEFAULT '',
    `email` VARCHAR(64) NULL DEFAULT '',
    `lg_ip` VARCHAR(20) NULL DEFAULT '',
    `created_time` INT NOT NULL DEFAULT '0',
    `timestampe` int  not null,
    PRIMARY KEY (`uid`)
);

INSERT INTO `s_member`(uid, username, password, name, email,created_time,timestampe)values
(1,'sadmin','e10adc3949ba59abbe56e057f20f883e','admin','','0','0');

-- 管理员表
DROP TABLE IF EXISTS `s_admin`;
CREATE TABLE `s_admin` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NOT NULL DEFAULT '',
    `password` VARCHAR(64) NOT NULL DEFAULT '',
    `name` VARCHAR(64) NULL DEFAULT '',
    `email` VARCHAR(64) NULL DEFAULT '',
    `created_time` DATE NULL DEFAULT '0',
    `timestampe` int  not null,
    PRIMARY KEY (`uid`)
);
INSERT INTO `s_admin`(uid, username, password, name, email,created_time,timestampe)values
(1,'sadmin','e10adc3949ba59abbe56e057f20f883e','admin','','0','0');

-- 类型
DROP TABLE IF EXISTS `s_category`;
CREATE TABLE `s_category` (
    `c_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '帖子ID', 
    `parent_id` INT(10) NOT NULL DEFAULT '0', 
    `name` VARCHAR(255) NULL DEFAULT '',
    `timestampe` int  not null,
    PRIMARY KEY (`c_id`)
);

-- 标签
DROP TABLE IF EXISTS `s_tag`;
CREATE TABLE `s_tag` (
    `tag_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '标签ID',
    `name` VARCHAR(255) NULL DEFAULT '' COMMENT '标签名',
    `timestampe` int  not null,
    PRIMARY KEY (`tag_id`)
);

-- 帖子
DROP TABLE IF EXISTS `s_post`;
CREATE TABLE `s_post` (
    `p_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
    `c_id` INT(10) NOT NULL DEFAULT '0',
    `uid` INT(10) NULL DEFAULT '0',
    `author` VARCHAR(255) NULL DEFAULT '' COMMENT '', 
    `tags` VARCHAR(255) NULL DEFAULT '',
    `title` VARCHAR(255) NULL DEFAULT '',
    `sort` INT NOT NULL,
    `content` text,
    `created_time` INT NOT NULL,
    `modified_time` INT NOT NULL,
    PRIMARY KEY (`p_id`)
);

DROP TABLE IF EXISTS `s_comment`;
CREATE TABLE `s_comment` (
    `c_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
    `p_id` INT(10) NOT NULL DEFAULT '0',
    `uid` INT(10) NULL DEFAULT '0',
    `author` VARCHAR(255) NULL DEFAULT '' COMMENT '', `tags` VARCHAR(255) NULL DEFAULT '',
    `content` VARCHAR(6000) NULL DEFAULT '',
    `created_time` INT NOT NULL,
    PRIMARY KEY (`c_id`)
);

DROP TABLE IF EXISTS `s_setting`;
CREATE TABLE `s_setting` (
    `s_id` INT(10) NOT NULL AUTO_INCREMENT COMMENT 'settingID',
    `key` VARCHAR(255) NULL DEFAULT '' COMMENT '设置键名',
    `value` VARCHAR(255) NULL DEFAULT '' COMMENT '键值',
    `created_time` INT NOT NULL,
    PRIMARY KEY (`s_id`)
);

