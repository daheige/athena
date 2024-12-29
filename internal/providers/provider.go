package providers

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/daheige/athena/internal/domain/repo"
	"github.com/daheige/athena/internal/infras/persistence/cache"
	"github.com/daheige/athena/internal/infras/persistence/users"
)

// Repositories 这个providers层可以根据实际情况看是否要添加
type Repositories struct {
	UserRepo  repo.UserRepository
	UserCache repo.UserCache
}

// NewRepositories 创建 Repositories
func NewRepositories(gormClient *gorm.DB, redisClient redis.UniversalClient) *Repositories {
	userRepo := users.NewUserRepo(gormClient)
	userCache := cache.NewUserCache(userRepo, redisClient)
	repos := &Repositories{
		UserRepo:  userRepo,
		UserCache: userCache,
	}

	return repos
}
