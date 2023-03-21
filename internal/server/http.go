package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	v1 "kratos-app/api/helloworld/v1"
	blog_v1 "kratos-app/api/blog/v1"
	account_v1  "kratos-app/api/account/v1"
	"kratos-app/internal/biz"
	"kratos-app/internal/conf"
	"kratos-app/internal/service"
	jwt2 "github.com/golang-jwt/jwt/v4"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, b *conf.Bootstrap, greeter *service.GreeterService, logger log.Logger, blog *service.BlogService, account *service.AccountService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(b.GetAuth().GetJwtSecret()), nil
				}, jwt.WithClaims(func() jwt2.Claims {
					return &biz.MyJwtClaims{}
				})),
			).Match(whiteList(b)).Build(),
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

var whiteList = func(c *conf.Bootstrap) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		for _, v := range c.GetAuth().GetWhiteList() {
			if v == operation {
				return false
			}
		}
		return true
	}
}
