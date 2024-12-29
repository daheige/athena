package commands

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/daheige/athena/internal/application"
	"github.com/daheige/athena/internal/infras/config"
	"github.com/daheige/athena/internal/providers"
)

// newUserCommand 创建 user command
func newUserCommand() *cobra.Command {
	var id int64
	var userCmd = &cobra.Command{
		Use:   "output-user",
		Short: "Print the user by id",
		Run: func(cmd *cobra.Command, args []string) {
			if id == 0 {
				log.Fatalln("id invalid")
			}

			// // 初始化db和redis
			gormClient := config.InitDB("db_conf")
			redisClient := config.InitRedisClient("redis_conf")
			// init repos and user service
			repos := providers.NewRepositories(gormClient, redisClient)
			userService := application.NewUserService(repos.UserRepo, repos.UserCache)

			users, err := userService.GetUser(context.Background(), id)
			if err != nil {
				log.Println("failed to get user err: ", err)
				return
			}

			log.Println("users: ", users)
		},
	}

	// 指定当前命令行的参数id
	userCmd.Flags().Int64VarP(&id, "id", "i", 0, "user id")

	return userCmd
}
