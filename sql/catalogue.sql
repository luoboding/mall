CREATE TABLE IF NOT EXISTS CATALOGUE (
    ID INT(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    TITLE VARCHAR(12) NOT NULL COMMENT '分类标题',
    THUMBNAIL VARCHAR(64) NOT NULL COMMENT '分类图片',
    SORT INT(16) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
    STATUS TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0 正常 1 禁用',
    CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW() COMMENT '创建时间'
)