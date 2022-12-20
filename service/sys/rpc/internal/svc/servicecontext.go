package svc

import (
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.SysUserModel
	SysConfigModel model.SysConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewSysUserModel(conn, c.CacheRedis),
		SysConfigModel: model.NewSysConfigModel(conn, c.CacheRedis),
	}
}
