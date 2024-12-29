package config

import (
	"log"
	"path/filepath"

	"github.com/daheige/athena/internal/infras/setting"
)

// Load 加载配置文件内容
func Load(path string) (setting.Config, error) {
	log.Println("config filename:", path, " dir:", filepath.Dir(path))
	c := setting.New(setting.WithConfigFile(path))
	err := c.Load()
	if err != nil {
		return nil, err
	}

	return c, nil
}
