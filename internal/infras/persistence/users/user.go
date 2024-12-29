package users

import (
	"context"

	"gorm.io/gorm"

	"github.com/daheige/athena/internal/domain/entity"
	"github.com/daheige/athena/internal/domain/repo"
)

var _ repo.UserRepository = (*UserRepoImpl)(nil)

// UserRepoImpl 用户接口实现
type UserRepoImpl struct {
	db *gorm.DB
}

// NewUserRepo 创建用户repo接口
func NewUserRepo(db *gorm.DB) repo.UserRepository {
	return &UserRepoImpl{
		db: db,
	}
}

// GetUser 根据id获取用户信息
func (u *UserRepoImpl) GetUser(ctx context.Context, id int64) (*entity.UserEntity, error) {
	user := &entity.UserEntity{}
	// 由于gorm存在db session将其嵌套到结构体字段会发生交叉污染的问题
	// 因此这里通过WithContext创建一个新的session会话
	err := u.db.WithContext(ctx).Table(entity.UserTable).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// BatchUsers 批量获取用户信息
func (u *UserRepoImpl) BatchUsers(ctx context.Context, ids []int64) ([]entity.UserEntity, error) {
	users := make([]entity.UserEntity, 0, len(ids))
	err := u.db.WithContext(ctx).Table(entity.UserTable).Where("id in (?)", ids).Find(&users).Error
	return users, err
}
