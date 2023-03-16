// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v4.22.1
// source: api/account/v1/account.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAccountInfo = "/api.account.v1.Account/Info"
const OperationAccountLogin = "/api.account.v1.Account/Login"
const OperationAccountRegister = "/api.account.v1.Account/Register"

type AccountHTTPServer interface {
	// Info 获取用户信息
	Info(context.Context, *InfoRequest) (*InfoReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.POST("/account/login", _Account_Login0_HTTP_Handler(srv))
	r.POST("/account/register", _Account_Register0_HTTP_Handler(srv))
	r.GET("/account/info", _Account_Info0_HTTP_Handler(srv))
}

func _Account_Login0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Account_Register0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Account_Info0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Info(ctx, req.(*InfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InfoReply)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	Info(ctx context.Context, req *InfoRequest, opts ...http.CallOption) (rsp *InfoReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) Info(ctx context.Context, in *InfoRequest, opts ...http.CallOption) (*InfoReply, error) {
	var out InfoReply
	pattern := "/account/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAccountInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/account/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/account/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
