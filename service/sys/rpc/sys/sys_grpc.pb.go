// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: sys.proto

package sys

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SysClient is the client API for Sys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysClient interface {
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
	DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error)
	DictDataOp(ctx context.Context, in *DictDataOpReq, opts ...grpc.CallOption) (*DictDataOpResp, error)
	PostList(ctx context.Context, in *PostListReq, opts ...grpc.CallOption) (*PostListResp, error)
	PostRetrieve(ctx context.Context, in *PostRetrieveReq, opts ...grpc.CallOption) (*PostRetrieveResp, error)
	PostAdd(ctx context.Context, in *PostAddReq, opts ...grpc.CallOption) (*PostAddResp, error)
	PostUpdate(ctx context.Context, in *PostUpdateReq, opts ...grpc.CallOption) (*PostUpdateResp, error)
	PostDelete(ctx context.Context, in *PostDeleteReq, opts ...grpc.CallOption) (*PostDeleteResp, error)
}

type sysClient struct {
	cc grpc.ClientConnInterface
}

func NewSysClient(cc grpc.ClientConnInterface) SysClient {
	return &sysClient{cc}
}

func (c *sysClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) SysConfig(ctx context.Context, in *SysConfigReq, opts ...grpc.CallOption) (*SysConfigResp, error) {
	out := new(SysConfigResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/SysConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) ConfigPw(ctx context.Context, in *ConfigPwReq, opts ...grpc.CallOption) (*ConfigPwResp, error) {
	out := new(ConfigPwResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/ConfigPw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserRetrieve(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserRetrieveResp, error) {
	out := new(UserRetrieveResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	out := new(UserAddResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	out := new(UserUpdateResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserUpdateStatus(ctx context.Context, in *UserUpdateStatusReq, opts ...grpc.CallOption) (*UserUpdateStatusResp, error) {
	out := new(UserUpdateStatusResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserUpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserUpdatePwd(ctx context.Context, in *UserUpdatePwdReq, opts ...grpc.CallOption) (*UserUpdatePwdResp, error) {
	out := new(UserUpdatePwdResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserUpdatePwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	out := new(UserDeleteResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	out := new(RoleListResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/RoleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	out := new(MenuAddResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/MenuAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	out := new(MenuListResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/MenuList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	out := new(MenuUpdateResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/MenuUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuRole(ctx context.Context, in *MenuRoleReq, opts ...grpc.CallOption) (*MenuRoleResp, error) {
	out := new(MenuRoleResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/MenuRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) DeptTree(ctx context.Context, in *DeptTreeReq, opts ...grpc.CallOption) (*DeptTreeResp, error) {
	out := new(DeptTreeResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/DeptTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error) {
	out := new(DeptListResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/DeptList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) DictDataOp(ctx context.Context, in *DictDataOpReq, opts ...grpc.CallOption) (*DictDataOpResp, error) {
	out := new(DictDataOpResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/DictDataOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PostList(ctx context.Context, in *PostListReq, opts ...grpc.CallOption) (*PostListResp, error) {
	out := new(PostListResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/PostList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PostRetrieve(ctx context.Context, in *PostRetrieveReq, opts ...grpc.CallOption) (*PostRetrieveResp, error) {
	out := new(PostRetrieveResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/PostRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PostAdd(ctx context.Context, in *PostAddReq, opts ...grpc.CallOption) (*PostAddResp, error) {
	out := new(PostAddResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/PostAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PostUpdate(ctx context.Context, in *PostUpdateReq, opts ...grpc.CallOption) (*PostUpdateResp, error) {
	out := new(PostUpdateResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/PostUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PostDelete(ctx context.Context, in *PostDeleteReq, opts ...grpc.CallOption) (*PostDeleteResp, error) {
	out := new(PostDeleteResp)
	err := c.cc.Invoke(ctx, "/sysclient.Sys/PostDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysServer is the server API for Sys service.
// All implementations must embed UnimplementedSysServer
// for forward compatibility
type SysServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	SysConfig(context.Context, *SysConfigReq) (*SysConfigResp, error)
	ConfigPw(context.Context, *ConfigPwReq) (*ConfigPwResp, error)
	UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
	UserList(context.Context, *UserListReq) (*UserListResp, error)
	UserRetrieve(context.Context, *UserInfoReq) (*UserRetrieveResp, error)
	UserAdd(context.Context, *UserAddReq) (*UserAddResp, error)
	UserUpdate(context.Context, *UserUpdateReq) (*UserUpdateResp, error)
	UserUpdateStatus(context.Context, *UserUpdateStatusReq) (*UserUpdateStatusResp, error)
	UserUpdatePwd(context.Context, *UserUpdatePwdReq) (*UserUpdatePwdResp, error)
	UserDelete(context.Context, *UserDeleteReq) (*UserDeleteResp, error)
	RoleList(context.Context, *RoleListReq) (*RoleListResp, error)
	MenuAdd(context.Context, *MenuAddReq) (*MenuAddResp, error)
	MenuList(context.Context, *MenuListReq) (*MenuListResp, error)
	MenuUpdate(context.Context, *MenuUpdateReq) (*MenuUpdateResp, error)
	MenuRole(context.Context, *MenuRoleReq) (*MenuRoleResp, error)
	DeptTree(context.Context, *DeptTreeReq) (*DeptTreeResp, error)
	DeptList(context.Context, *DeptListReq) (*DeptListResp, error)
	DictDataOp(context.Context, *DictDataOpReq) (*DictDataOpResp, error)
	PostList(context.Context, *PostListReq) (*PostListResp, error)
	PostRetrieve(context.Context, *PostRetrieveReq) (*PostRetrieveResp, error)
	PostAdd(context.Context, *PostAddReq) (*PostAddResp, error)
	PostUpdate(context.Context, *PostUpdateReq) (*PostUpdateResp, error)
	PostDelete(context.Context, *PostDeleteReq) (*PostDeleteResp, error)
	mustEmbedUnimplementedSysServer()
}

// UnimplementedSysServer must be embedded to have forward compatible implementations.
type UnimplementedSysServer struct {
}

func (UnimplementedSysServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSysServer) SysConfig(context.Context, *SysConfigReq) (*SysConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysConfig not implemented")
}
func (UnimplementedSysServer) ConfigPw(context.Context, *ConfigPwReq) (*ConfigPwResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigPw not implemented")
}
func (UnimplementedSysServer) UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedSysServer) UserList(context.Context, *UserListReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedSysServer) UserRetrieve(context.Context, *UserInfoReq) (*UserRetrieveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRetrieve not implemented")
}
func (UnimplementedSysServer) UserAdd(context.Context, *UserAddReq) (*UserAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (UnimplementedSysServer) UserUpdate(context.Context, *UserUpdateReq) (*UserUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedSysServer) UserUpdateStatus(context.Context, *UserUpdateStatusReq) (*UserUpdateStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdateStatus not implemented")
}
func (UnimplementedSysServer) UserUpdatePwd(context.Context, *UserUpdatePwdReq) (*UserUpdatePwdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdatePwd not implemented")
}
func (UnimplementedSysServer) UserDelete(context.Context, *UserDeleteReq) (*UserDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedSysServer) RoleList(context.Context, *RoleListReq) (*RoleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleList not implemented")
}
func (UnimplementedSysServer) MenuAdd(context.Context, *MenuAddReq) (*MenuAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuAdd not implemented")
}
func (UnimplementedSysServer) MenuList(context.Context, *MenuListReq) (*MenuListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuList not implemented")
}
func (UnimplementedSysServer) MenuUpdate(context.Context, *MenuUpdateReq) (*MenuUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuUpdate not implemented")
}
func (UnimplementedSysServer) MenuRole(context.Context, *MenuRoleReq) (*MenuRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuRole not implemented")
}
func (UnimplementedSysServer) DeptTree(context.Context, *DeptTreeReq) (*DeptTreeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeptTree not implemented")
}
func (UnimplementedSysServer) DeptList(context.Context, *DeptListReq) (*DeptListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeptList not implemented")
}
func (UnimplementedSysServer) DictDataOp(context.Context, *DictDataOpReq) (*DictDataOpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DictDataOp not implemented")
}
func (UnimplementedSysServer) PostList(context.Context, *PostListReq) (*PostListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostList not implemented")
}
func (UnimplementedSysServer) PostRetrieve(context.Context, *PostRetrieveReq) (*PostRetrieveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostRetrieve not implemented")
}
func (UnimplementedSysServer) PostAdd(context.Context, *PostAddReq) (*PostAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAdd not implemented")
}
func (UnimplementedSysServer) PostUpdate(context.Context, *PostUpdateReq) (*PostUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUpdate not implemented")
}
func (UnimplementedSysServer) PostDelete(context.Context, *PostDeleteReq) (*PostDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostDelete not implemented")
}
func (UnimplementedSysServer) mustEmbedUnimplementedSysServer() {}

// UnsafeSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysServer will
// result in compilation errors.
type UnsafeSysServer interface {
	mustEmbedUnimplementedSysServer()
}

func RegisterSysServer(s grpc.ServiceRegistrar, srv SysServer) {
	s.RegisterService(&Sys_ServiceDesc, srv)
}

func _Sys_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_SysConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).SysConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/SysConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).SysConfig(ctx, req.(*SysConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_ConfigPw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigPwReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).ConfigPw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/ConfigPw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).ConfigPw(ctx, req.(*ConfigPwReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserList(ctx, req.(*UserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserRetrieve(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserAdd(ctx, req.(*UserAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserUpdate(ctx, req.(*UserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserUpdateStatus(ctx, req.(*UserUpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserUpdatePwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdatePwdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserUpdatePwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserUpdatePwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserUpdatePwd(ctx, req.(*UserUpdatePwdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserDelete(ctx, req.(*UserDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/RoleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RoleList(ctx, req.(*RoleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/MenuAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuAdd(ctx, req.(*MenuAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/MenuList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuList(ctx, req.(*MenuListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/MenuUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuUpdate(ctx, req.(*MenuUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/MenuRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuRole(ctx, req.(*MenuRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_DeptTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptTreeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).DeptTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/DeptTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).DeptTree(ctx, req.(*DeptTreeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_DeptList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).DeptList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/DeptList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).DeptList(ctx, req.(*DeptListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_DictDataOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictDataOpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).DictDataOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/DictDataOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).DictDataOp(ctx, req.(*DictDataOpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PostList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PostList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/PostList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PostList(ctx, req.(*PostListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PostRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRetrieveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PostRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/PostRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PostRetrieve(ctx, req.(*PostRetrieveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PostAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PostAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/PostAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PostAdd(ctx, req.(*PostAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PostUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PostUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/PostUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PostUpdate(ctx, req.(*PostUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PostDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PostDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysclient.Sys/PostDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PostDelete(ctx, req.(*PostDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sys_ServiceDesc is the grpc.ServiceDesc for Sys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sysclient.Sys",
	HandlerType: (*SysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Sys_Login_Handler,
		},
		{
			MethodName: "SysConfig",
			Handler:    _Sys_SysConfig_Handler,
		},
		{
			MethodName: "ConfigPw",
			Handler:    _Sys_ConfigPw_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Sys_UserInfo_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Sys_UserList_Handler,
		},
		{
			MethodName: "UserRetrieve",
			Handler:    _Sys_UserRetrieve_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _Sys_UserAdd_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _Sys_UserUpdate_Handler,
		},
		{
			MethodName: "UserUpdateStatus",
			Handler:    _Sys_UserUpdateStatus_Handler,
		},
		{
			MethodName: "UserUpdatePwd",
			Handler:    _Sys_UserUpdatePwd_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _Sys_UserDelete_Handler,
		},
		{
			MethodName: "RoleList",
			Handler:    _Sys_RoleList_Handler,
		},
		{
			MethodName: "MenuAdd",
			Handler:    _Sys_MenuAdd_Handler,
		},
		{
			MethodName: "MenuList",
			Handler:    _Sys_MenuList_Handler,
		},
		{
			MethodName: "MenuUpdate",
			Handler:    _Sys_MenuUpdate_Handler,
		},
		{
			MethodName: "MenuRole",
			Handler:    _Sys_MenuRole_Handler,
		},
		{
			MethodName: "DeptTree",
			Handler:    _Sys_DeptTree_Handler,
		},
		{
			MethodName: "DeptList",
			Handler:    _Sys_DeptList_Handler,
		},
		{
			MethodName: "DictDataOp",
			Handler:    _Sys_DictDataOp_Handler,
		},
		{
			MethodName: "PostList",
			Handler:    _Sys_PostList_Handler,
		},
		{
			MethodName: "PostRetrieve",
			Handler:    _Sys_PostRetrieve_Handler,
		},
		{
			MethodName: "PostAdd",
			Handler:    _Sys_PostAdd_Handler,
		},
		{
			MethodName: "PostUpdate",
			Handler:    _Sys_PostUpdate_Handler,
		},
		{
			MethodName: "PostDelete",
			Handler:    _Sys_PostDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys.proto",
}
