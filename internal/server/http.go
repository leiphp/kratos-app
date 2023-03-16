package server

import (
	v1 "kratos-app/api/helloworld/v1"
	blog_v1 "kratos-app/api/blog/v1"
	account_v1  "kratos-app/api/account/v1"
	"kratos-app/internal/conf"
	"kratos-app/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger, blog *service.BlogService, account *service.AccountService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	blog_v1.RegisterBlogServiceHTTPServer(srv, blog)
	account_v1.RegisterAccountHTTPServer(srv, account)
	return srv
}
