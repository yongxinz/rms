package svc

import (
	"rms/service/sys/api/internal/config"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	SysRpc sysclient.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SysRpc: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
