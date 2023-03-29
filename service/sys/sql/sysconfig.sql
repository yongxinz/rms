-- rms.sys_config definition

CREATE TABLE `sys_config` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(8) NOT NULL DEFAULT '' COMMENT 'ConfigName',
  `config_key` varchar(32) NOT NULL DEFAULT '' COMMENT 'ConfigKey',
  `config_value` varchar(128) NOT NULL DEFAULT '' COMMENT 'ConfigValue',
  `config_type` tinyint NOT NULL DEFAULT 1 COMMENT 'ConfigType',
  `is_frontend` tinyint NOT NULL DEFAULT 1 COMMENT '是否前台',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_sys_config_create_by` (`create_by`),
  KEY `idx_sys_config_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;