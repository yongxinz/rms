// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	dept "rms/service/sys/api/internal/handler/dept"
	dictdata "rms/service/sys/api/internal/handler/dictdata"
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
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
