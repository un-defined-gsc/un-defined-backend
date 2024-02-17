package autoload

import "github.com/un-defined-gsc/un-defined-backend/internal/config"

func init() {
	if err := config.InitializeConfig(); err != nil {
		panic(err)
	}
}
