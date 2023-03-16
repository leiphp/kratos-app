/**
 * @package data
 * @file      : user.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2023/3/16 16:03
 **/
package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-app/internal/biz"
	"gorm.io/gorm"
)

type userRepo struct {
	data  *Data
	log   *log.Helper
	table *gorm.DB
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:  data,
		log:   log.NewHelper(logger),
		table: data.db2.Table("users"),
	}
}

func (u *userRepo) FetchByUsername(ctx context.Context, username string) (user *biz.User, err error) {
	user = &biz.User{}
	u.table.WithContext(ctx).First(user, "username = ?", username)
	if user.ID == 0 {
		return nil, biz.ErrUserNotExist
	}
	return user, nil
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (id uint32, err error) {
	result := u.table.WithContext(ctx).Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (u *userRepo) FetchByUid(ctx context.Context, uid uint32) (user *biz.User, err error) {
	user = &biz.User{}
	u.table.WithContext(ctx).First(user, "ID = ?", uid)
	if user.ID == 0 {
		return nil, biz.ErrUserNotExist
	}
	return user, nil
}

func (u *userRepo) FetchByUidList(ctx context.Context, uidList []uint32) (user map[uint32]*biz.User, err error) {
	result := make(map[uint32]*biz.User, len(uidList))
	uidInfoList := make([]*biz.User, 0, len(uidList))
	if len(uidList) == 0 {
		return result, nil
	}
	err = u.table.WithContext(ctx).Where("ID in (?)", uidList).Find(&uidInfoList).Error
	if err != nil {
		return nil, err
	}
	for _, user := range uidInfoList {
		result[user.ID] = user
	}
	return result, nil
}