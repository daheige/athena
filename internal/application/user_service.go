package application

import (
	"context"

	"github.com/daheige/athena/internal/domain/entity"
	"github.com/daheige/athena/internal/domain/repo"
)

// UserService 用户相关的业务处理
type UserService struct {
	userCache repo.UserCache
	userRepo  repo.UserRepository
}

// NewUserService 创建一个服务实例
func NewUserService(userRepo repo.UserRepository, userCache repo.UserCache) *UserService {
	return &UserService{
		userCache: userCache,
		userRepo:  userRepo,
	}
}

// GetUser 获取用户信息
func (u *UserService) GetUser(ctx context.Context, id int64) (*entity.UserEntity, error) {
	user, err := u.userCache.GetUser(ctx, id)
	// 省略其他代码....

	return user, err
}

// BatchUsers 根据id批量获取用户
func (u *UserService) BatchUsers(ctx context.Context, ids []int64) ([]entity.UserEntity, error) {
	users, err := u.userRepo.BatchUsers(ctx, ids)
	return users, err
}
