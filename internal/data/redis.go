/**
 * @package data
 * @file      : redis.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2023/3/10 19:01
 **/
package data

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func likeKey(id int64) string {
	return fmt.Sprintf("like:%d", id)
}

func (ar *articleRepo) GetArticleLike(ctx context.Context, id int64) (rv int64, err error) {
	get := ar.data.rdb.Get(ctx, likeKey(id))
	rv, err = get.Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return
}

func (ar *articleRepo) IncArticleLike(ctx context.Context, id int64) error {
	_, err := ar.data.rdb.Incr(ctx, likeKey(id)).Result()
	return err
}