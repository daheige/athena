package repo

import (
	"context"

	"github.com/daheige/athena/internal/domain/entity"
)

// UserRepository 用户接口定义
type UserRepository interface {
	GetUser(ctx context.Context, id int64) (*entity.UserEntity, error)
	BatchUsers(ctx context.Context, ids []int64) ([]entity.UserEntity, error)
}
