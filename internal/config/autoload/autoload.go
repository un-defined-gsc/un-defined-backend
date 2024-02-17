package autoload

import "github.com/ProjectMonWeb/API-Service/internal/config"

func init() {
	if err := config.InitializeConfig(); err != nil {
		panic(err)
	}
}
