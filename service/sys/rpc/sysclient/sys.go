// Code generated by goctl. DO NOT EDIT!
// Source: sys.proto

package sysclient

import (
	"context"

	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeepTreeData   = sys.DeepTreeData
	DeepTreeReq    = sys.DeepTreeReq
	DeepTreeResp   = sys.DeepTreeResp
	DeptListData   = sys.DeptListData
	LoginRequest   = sys.LoginRequest
	LoginResponse  = sys.LoginResponse
	MenuAddReq     = sys.MenuAddReq
	MenuAddResp    = sys.MenuAddResp
	MenuDeleteReq  = sys.MenuDeleteReq
	MenuDeleteResp = sys.MenuDeleteResp
	MenuListData   = sys.MenuListData
	MenuListReq    = sys.MenuListReq
	MenuListResp   = sys.MenuListResp
	MenuRoleReq    = sys.MenuRoleReq
	MenuRoleResp   = sys.MenuRoleResp
	MenuTree       = sys.MenuTree
	MenuUpdateReq  = sys.MenuUpdateReq
	MenuUpdateResp = sys.MenuUpdateResp
	SysConfigReq   = sys.SysConfigReq
	SysConfigResp  = sys.SysConfigResp
	UserInfoReq    = sys.UserInfoReq
	UserInfoResp   = sys.UserInfoResp
	UserListData   = sys.UserListData
	UserListReq    = sys.UserListReq
	UserListResp   = sys.UserListResp

	Sys interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		SysConfig(ctx context.Context, in *SysConfigReq, opts ...grpc.CallOption) (*SysConfigResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		MenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleResp, error)
		DeepTree(ctx context.Context, in *DeepTreeReq, opts ...grpc.CallOption) (*DeepTreeResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

func (m *defaultSys) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultSys) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultSys) SysConfig(ctx context.Context, in *SysConfigReq, opts ...grpc.CallOption) (*SysConfigResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysConfig(ctx, in, opts...)
}

func (m *defaultSys) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultSys) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuAdd(ctx, in, opts...)
}

func (m *defaultSys) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

func (m *defaultSys) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

func (m *defaultSys) MenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuRole(ctx, in, opts...)
}

func (m *defaultSys) DeepTree(ctx context.Context, in *DeepTreeReq, opts ...grpc.CallOption) (*DeepTreeResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeepTree(ctx, in, opts...)
}
