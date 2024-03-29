// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	dept "rms/service/sys/api/internal/handler/dept"
	dictdata "rms/service/sys/api/internal/handler/dictdata"
	dicttype "rms/service/sys/api/internal/handler/dicttype"
	menu "rms/service/sys/api/internal/handler/menu"
	post "rms/service/sys/api/internal/handler/post"
	role "rms/service/sys/api/internal/handler/role"
	sysconfig "rms/service/sys/api/internal/handler/sysconfig"
	user "rms/service/sys/api/internal/handler/user"
	"rms/service/sys/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: user.CaptchaHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getinfo",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/profile",
				Handler: user.ProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sys-user",
				Handler: user.UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sys-user/:userId",
				Handler: user.UserRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sys-user",
				Handler: user.UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/sys-user",
				Handler: user.UserUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/status",
				Handler: user.UserUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/pwd/reset",
				Handler: user.UserUpdatePwdHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/sys-user",
				Handler: user.UserDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/role",
				Handler: role.RoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/:roleId",
				Handler: role.RoleRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role",
				Handler: role.RoleAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/role/:roleId",
				Handler: role.RoleUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/role",
				Handler: role.RoleDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/roleMenuTreeselect/:roleId",
				Handler: role.RoleMenuTreeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/deptTree",
				Handler: dept.DeptTreeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/dept",
				Handler: dept.DeptListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/dept/:deptId",
				Handler: dept.DeptRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/dept",
				Handler: dept.DeptAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/dept",
				Handler: dept.DeptUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/dept",
				Handler: dept.DeptDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/app-config",
				Handler: sysconfig.SysConfigHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/configKey/:configKey",
				Handler: sysconfig.ConfigPwHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/config",
				Handler: sysconfig.ConfigListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/config/:id",
				Handler: sysconfig.ConfigRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/config",
				Handler: sysconfig.ConfigAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/config/:id",
				Handler: sysconfig.ConfigUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/config",
				Handler: sysconfig.ConfigDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/menurole",
				Handler: menu.MenuRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu",
				Handler: menu.MenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/:menuId",
				Handler: menu.MenuRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu",
				Handler: menu.MenuAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/menu/:menuId",
				Handler: menu.MenuUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/menu",
				Handler: menu.MenuDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/dict-data/option-select",
				Handler: dictdata.DictDataOpHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/dict/data",
				Handler: dictdata.DictDataListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/dict/data/:dictId",
				Handler: dictdata.DictDataRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/dict/data",
				Handler: dictdata.DictDataAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/dict/data/:dictId",
				Handler: dictdata.DictDataUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/dict/data",
				Handler: dictdata.DictDataDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/dict/type",
				Handler: dicttype.DictTypeListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/dict/type/:dictId",
				Handler: dicttype.DictTypeRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/dict/type",
				Handler: dicttype.DictTypeAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/dict/type/:dictId",
				Handler: dicttype.DictTypeUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/dict/type",
				Handler: dicttype.DictTypeDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/post",
				Handler: post.PostListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/post/:postId",
				Handler: post.PostRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/post",
				Handler: post.PostAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/post",
				Handler: post.PostUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/post",
				Handler: post.PostDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
