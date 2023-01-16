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
	ConfigPwReq          = sys.ConfigPwReq
	ConfigPwResp         = sys.ConfigPwResp
	DeptListData         = sys.DeptListData
	DeptTreeData         = sys.DeptTreeData
	DeptTreeReq          = sys.DeptTreeReq
	DeptTreeResp         = sys.DeptTreeResp
	DictDataOp           = sys.DictDataOp
	DictDataOpReq        = sys.DictDataOpReq
	DictDataOpResp       = sys.DictDataOpResp
	LoginRequest         = sys.LoginRequest
	LoginResponse        = sys.LoginResponse
	MenuAddReq           = sys.MenuAddReq
	MenuAddResp          = sys.MenuAddResp
	MenuDeleteReq        = sys.MenuDeleteReq
	MenuDeleteResp       = sys.MenuDeleteResp
	MenuListData         = sys.MenuListData
	MenuListReq          = sys.MenuListReq
	MenuListResp         = sys.MenuListResp
	MenuRoleReq          = sys.MenuRoleReq
	MenuRoleResp         = sys.MenuRoleResp
	MenuTree             = sys.MenuTree
	MenuUpdateReq        = sys.MenuUpdateReq
	MenuUpdateResp       = sys.MenuUpdateResp
	PostAddReq           = sys.PostAddReq
	PostAddResp          = sys.PostAddResp
	PostDeleteReq        = sys.PostDeleteReq
	PostDeleteResp       = sys.PostDeleteResp
	PostListData         = sys.PostListData
	PostListReq          = sys.PostListReq
	PostListResp         = sys.PostListResp
	PostRetrieveReq      = sys.PostRetrieveReq
	PostRetrieveResp     = sys.PostRetrieveResp
	PostUpdateReq        = sys.PostUpdateReq
	PostUpdateResp       = sys.PostUpdateResp
	RoleListData         = sys.RoleListData
	RoleListReq          = sys.RoleListReq
	RoleListResp         = sys.RoleListResp
	SysConfigReq         = sys.SysConfigReq
	SysConfigResp        = sys.SysConfigResp
	UserAddReq           = sys.UserAddReq
	UserAddResp          = sys.UserAddResp
	UserDeleteReq        = sys.UserDeleteReq
	UserDeleteResp       = sys.UserDeleteResp
	UserInfoReq          = sys.UserInfoReq
	UserInfoResp         = sys.UserInfoResp
	UserListData         = sys.UserListData
	UserListReq          = sys.UserListReq
	UserListResp         = sys.UserListResp
	UserRetrieveResp     = sys.UserRetrieveResp
	UserUpdatePwdReq     = sys.UserUpdatePwdReq
	UserUpdatePwdResp    = sys.UserUpdatePwdResp
	UserUpdateReq        = sys.UserUpdateReq
	UserUpdateResp       = sys.UserUpdateResp
	UserUpdateStatusReq  = sys.UserUpdateStatusReq
	UserUpdateStatusResp = sys.UserUpdateStatusResp

	Sys interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		SysConfig(ctx context.Context, in *SysConfigReq, opts ...grpc.CallOption) (*SysConfigResp, error)
		ConfigPw(ctx context.Context, in *ConfigPwReq, opts ...grpc.CallOption) (*ConfigPwResp, error)
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		UserRetrieve(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserRetrieveResp, error)
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error)
		UserUpdateStatus(ctx context.Context, in *UserUpdateStatusReq, opts ...grpc.CallOption) (*UserUpdateStatusResp, error)
		UserUpdatePwd(ctx context.Context, in *UserUpdatePwdReq, opts ...grpc.CallOption) (*UserUpdatePwdResp, error)
		UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error)
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		MenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleResp, error)
		DeptTree(ctx context.Context, in *DeptTreeReq, opts ...grpc.CallOption) (*DeptTreeResp, error)
		DictDataOp(ctx context.Context, in *DictDataOpReq, opts ...grpc.CallOption) (*DictDataOpResp, error)
		PostList(ctx context.Context, in *PostListReq, opts ...grpc.CallOption) (*PostListResp, error)
		PostRetrieve(ctx context.Context, in *PostRetrieveReq, opts ...grpc.CallOption) (*PostRetrieveResp, error)
		PostAdd(ctx context.Context, in *PostAddReq, opts ...grpc.CallOption) (*PostAddResp, error)
		PostUpdate(ctx context.Context, in *PostUpdateReq, opts ...grpc.CallOption) (*PostUpdateResp, error)
		PostDelete(ctx context.Context, in *PostDeleteReq, opts ...grpc.CallOption) (*PostDeleteResp, error)
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

func (m *defaultSys) SysConfig(ctx context.Context, in *SysConfigReq, opts ...grpc.CallOption) (*SysConfigResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysConfig(ctx, in, opts...)
}

func (m *defaultSys) ConfigPw(ctx context.Context, in *ConfigPwReq, opts ...grpc.CallOption) (*ConfigPwResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.ConfigPw(ctx, in, opts...)
}

func (m *defaultSys) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultSys) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultSys) UserRetrieve(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserRetrieveResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserRetrieve(ctx, in, opts...)
}

func (m *defaultSys) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultSys) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultSys) UserUpdateStatus(ctx context.Context, in *UserUpdateStatusReq, opts ...grpc.CallOption) (*UserUpdateStatusResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserUpdateStatus(ctx, in, opts...)
}

func (m *defaultSys) UserUpdatePwd(ctx context.Context, in *UserUpdatePwdReq, opts ...grpc.CallOption) (*UserUpdatePwdResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserUpdatePwd(ctx, in, opts...)
}

func (m *defaultSys) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserDelete(ctx, in, opts...)
}

func (m *defaultSys) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
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

func (m *defaultSys) DeptTree(ctx context.Context, in *DeptTreeReq, opts ...grpc.CallOption) (*DeptTreeResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptTree(ctx, in, opts...)
}

func (m *defaultSys) DictDataOp(ctx context.Context, in *DictDataOpReq, opts ...grpc.CallOption) (*DictDataOpResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictDataOp(ctx, in, opts...)
}

func (m *defaultSys) PostList(ctx context.Context, in *PostListReq, opts ...grpc.CallOption) (*PostListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PostList(ctx, in, opts...)
}

func (m *defaultSys) PostRetrieve(ctx context.Context, in *PostRetrieveReq, opts ...grpc.CallOption) (*PostRetrieveResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PostRetrieve(ctx, in, opts...)
}

func (m *defaultSys) PostAdd(ctx context.Context, in *PostAddReq, opts ...grpc.CallOption) (*PostAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PostAdd(ctx, in, opts...)
}

func (m *defaultSys) PostUpdate(ctx context.Context, in *PostUpdateReq, opts ...grpc.CallOption) (*PostUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PostUpdate(ctx, in, opts...)
}

func (m *defaultSys) PostDelete(ctx context.Context, in *PostDeleteReq, opts ...grpc.CallOption) (*PostDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PostDelete(ctx, in, opts...)
}
