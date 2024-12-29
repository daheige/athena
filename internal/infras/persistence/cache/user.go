package cache

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/daheige/athena/internal/domain/entity"
	"github.com/daheige/athena/internal/domain/repo"
)

var _ repo.UserCache = (*UserCacheImpl)(nil)

// UserCacheImpl 用户缓存接口实现
type UserCacheImpl struct {
	redisClient redis.UniversalClient
	userRepo    repo.UserRepository
}

// NewUserCache 创建一个用户cache repo
func NewUserCache(userRepo repo.UserRepository, redisClient redis.UniversalClient) repo.UserCache {
	u := &UserCacheImpl{
		redisClient: redisClient,
		userRepo:    userRepo,
	}

	return u
}

// GetUser 获取用户信息
func (u *UserCacheImpl) GetUser(ctx context.Context, id int64) (*entity.UserEntity, error) {
	key := "athena:user:" + strconv.FormatInt(id, 10)
	str, err := u.redisClient.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// log.Println("str: ", str, err)
	// 如果缓存存在，就直接返回
	user := &entity.UserEntity{}
	if str != "" {
		err = json.Unmarshal([]byte(str), user)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	// 从数据库中获取，并更新cache
	user, err = u.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(user)
	err = u.redisClient.Set(ctx, key, string(b), 3600*time.Second).Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}
