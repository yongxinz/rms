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
	DeptAddReq           = sys.DeptAddReq
	DeptAddResp          = sys.DeptAddResp
	DeptDeleteReq        = sys.DeptDeleteReq
	DeptDeleteResp       = sys.DeptDeleteResp
	DeptListData         = sys.DeptListData
	DeptListReq          = sys.DeptListReq
	DeptListResp         = sys.DeptListResp
	DeptRetrieveReq      = sys.DeptRetrieveReq
	DeptRetrieveResp     = sys.DeptRetrieveResp
	DeptTreeData         = sys.DeptTreeData
	DeptTreeReq          = sys.DeptTreeReq
	DeptTreeResp         = sys.DeptTreeResp
	DeptUpdateReq        = sys.DeptUpdateReq
	DeptUpdateResp       = sys.DeptUpdateResp
	DictDataOp           = sys.DictDataOp
	DictDataOpReq        = sys.DictDataOpReq
	DictDataOpResp       = sys.DictDataOpResp
	DictTypeListData     = sys.DictTypeListData
	DictTypeListReq      = sys.DictTypeListReq
	DictTypeListResp     = sys.DictTypeListResp
	DictTypeRetrieveReq  = sys.DictTypeRetrieveReq
	DictTypeRetrieveResp = sys.DictTypeRetrieveResp
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
	RoleAddReq           = sys.RoleAddReq
	RoleAddResp          = sys.RoleAddResp
	RoleDeleteReq        = sys.RoleDeleteReq
	RoleDeleteResp       = sys.RoleDeleteResp
	RoleListData         = sys.RoleListData
	RoleListReq          = sys.RoleListReq
	RoleListResp         = sys.RoleListResp
	RoleMenuTreeData     = sys.RoleMenuTreeData
	RoleMenuTreeReq      = sys.RoleMenuTreeReq
	RoleMenuTreeResp     = sys.RoleMenuTreeResp
	RoleRetrieveReq      = sys.RoleRetrieveReq
	RoleRetrieveResp     = sys.RoleRetrieveResp
	RoleUpdateReq        = sys.RoleUpdateReq
	RoleUpdateResp       = sys.RoleUpdateResp
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
		RoleMenuTree(ctx context.Context, in *RoleMenuTreeReq, opts ...grpc.CallOption) (*RoleMenuTreeResp, error)
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error)
		RoleRetrieve(ctx context.Context, in *RoleRetrieveReq, opts ...grpc.CallOption) (*RoleRetrieveResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error)
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		MenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleResp, error)
		DeptTree(ctx context.Context, in *DeptTreeReq, opts ...grpc.CallOption) (*DeptTreeResp, error)
		DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error)
		DeptRetrieve(ctx context.Context, in *DeptRetrieveReq, opts ...grpc.CallOption) (*DeptRetrieveResp, error)
		DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error)
		DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error)
		DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error)
		DictDataOp(ctx context.Context, in *DictDataOpReq, opts ...grpc.CallOption) (*DictDataOpResp, error)
		DictTypeList(ctx context.Context, in *DictTypeListReq, opts ...grpc.CallOption) (*DictTypeListResp, error)
		DictTypeRetrieve(ctx context.Context, in *DictTypeRetrieveReq, opts ...grpc.CallOption) (*DictTypeRetrieveResp, error)
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

func (m *defaultSys) RoleMenuTree(ctx context.Context, in *RoleMenuTreeReq, opts ...grpc.CallOption) (*RoleMenuTreeResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleMenuTree(ctx, in, opts...)
}

func (m *defaultSys) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultSys) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

func (m *defaultSys) RoleRetrieve(ctx context.Context, in *RoleRetrieveReq, opts ...grpc.CallOption) (*RoleRetrieveResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleRetrieve(ctx, in, opts...)
}

func (m *defaultSys) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (m *defaultSys) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
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

func (m *defaultSys) DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptList(ctx, in, opts...)
}

func (m *defaultSys) DeptRetrieve(ctx context.Context, in *DeptRetrieveReq, opts ...grpc.CallOption) (*DeptRetrieveResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptRetrieve(ctx, in, opts...)
}

func (m *defaultSys) DeptAdd(ctx context.Context, in *DeptAddReq, opts ...grpc.CallOption) (*DeptAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptAdd(ctx, in, opts...)
}

func (m *defaultSys) DeptUpdate(ctx context.Context, in *DeptUpdateReq, opts ...grpc.CallOption) (*DeptUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptUpdate(ctx, in, opts...)
}

func (m *defaultSys) DeptDelete(ctx context.Context, in *DeptDeleteReq, opts ...grpc.CallOption) (*DeptDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DeptDelete(ctx, in, opts...)
}

func (m *defaultSys) DictDataOp(ctx context.Context, in *DictDataOpReq, opts ...grpc.CallOption) (*DictDataOpResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictDataOp(ctx, in, opts...)
}

func (m *defaultSys) DictTypeList(ctx context.Context, in *DictTypeListReq, opts ...grpc.CallOption) (*DictTypeListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictTypeList(ctx, in, opts...)
}

func (m *defaultSys) DictTypeRetrieve(ctx context.Context, in *DictTypeRetrieveReq, opts ...grpc.CallOption) (*DictTypeRetrieveResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.DictTypeRetrieve(ctx, in, opts...)
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
