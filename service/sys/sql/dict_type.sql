-- rms.sys_dict_type definition

CREATE TABLE `sys_dict_type` (
  `dict_id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `dict_name` varchar(32) NOT NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(32) NOT NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`dict_id`),
  KEY `idx_sys_dict_type_create_by` (`create_by`),
  KEY `idx_sys_dict_type_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
