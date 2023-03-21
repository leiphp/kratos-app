/**
 * @package service
 * @file      : account.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2023/3/16 16:10
 **/
package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	v1 "kratos-app/api/account/v1"
	"kratos-app/internal/biz"
)

type AccountService struct {
	v1.UnimplementedAccountServer
	log *log.Helper
	auc *biz.AccountUseCase
}

func NewAccountService(logger log.Logger, auc *biz.AccountUseCase) *AccountService {
	return &AccountService{
		log: log.NewHelper(logger),
		auc: auc,
	}
}

func (a *AccountService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginReply, error) {
	token, err := a.auc.Login(ctx, request.GetPhone(), request.GetPassword())
	if err != nil {
		return nil, errors.New(500, "登录失败", err.Error())
	}
	return &v1.LoginReply{
		Token: token,
	}, nil
}

func (a *AccountService) Register(ctx context.Context, request *v1.RegisterRequest) (*v1.RegisterReply, error) {
	err := a.auc.Register(ctx, request.GetPhone(), request.GetPassword())
	if err != nil {
		return nil, errors.New(500, "注册失败", err.Error())
	}
	return &v1.RegisterReply{}, nil
}

func (a *AccountService) Info(ctx context.Context, request *v1.InfoRequest) (*v1.InfoReply, error) {
	claims, _ := jwt.FromContext(ctx)
	userInfo, err := a.auc.UserInfo(ctx, claims.(*biz.MyJwtClaims).Uid)
	if err != nil {
		return nil, errors.New(500, "获取用户信息失败", err.Error())
	}
	return &v1.InfoReply{
		Id:       userInfo.ID,
		Username: userInfo.Username,
		Avatar:   userInfo.Avatar,
	}, nil
}