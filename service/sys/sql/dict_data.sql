-- rms.sys_dict_data definition

CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT,
  `dict_sort` bigint(20) DEFAULT NULL,
  `dict_label` varchar(128) DEFAULT NULL,
  `dict_value` varchar(255) DEFAULT NULL,
  `dict_type` varchar(64) DEFAULT NULL,
  `css_class` varchar(128) DEFAULT NULL,
  `list_class` varchar(128) DEFAULT NULL,
  `is_default` varchar(8) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `default` varchar(8) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `update_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`),
  KEY `idx_sys_dict_data_create_by` (`create_by`),
  KEY `idx_sys_dict_data_update_by` (`update_by`),
  KEY `idx_sys_dict_data_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4;