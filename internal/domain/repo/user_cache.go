package repo

import (
	"context"

	"github.com/daheige/athena/internal/domain/entity"
)

// UserCache 用户缓存接口定义
type UserCache interface {
	GetUser(ctx context.Context, id int64) (*entity.UserEntity, error)
}
