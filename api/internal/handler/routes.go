// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "github.com/suyuan32/simple-admin-file/api/internal/handler/base"
	file "github.com/suyuan32/simple-admin-file/api/internal/handler/file"
	"github.com/suyuan32/simple-admin-file/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: file.UploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/list",
					Handler: file.FileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file",
					Handler: file.UpdateFileHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/file",
					Handler: file.DeleteFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/status",
					Handler: file.ChangePublicStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/file/download/:id",
					Handler: file.DownloadFileHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
