package main

import (
	"context"
	"fmt"
	stdLog "log"

	"github.com/nightlord189/jsonbeautifier/internal/config"
	"github.com/nightlord189/jsonbeautifier/internal/delivery/http"
	"github.com/nightlord189/jsonbeautifier/pkg/log"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("start #1")

	cfg, err := config.LoadConfig("configs/config.yml")
	if err != nil {
		stdLog.Fatalf("error on load config: %v", err)
	}

	if err := log.InitLogger(cfg.LogLevel, "jsonbeautifier"); err != nil {
		stdLog.Fatalf("error on init logger: %v", err)
	}

	ctx := context.Background()

	zerolog.Ctx(ctx).Debug().Msg("start #2")

	handler := http.New(cfg.HTTP)

	if err := handler.Run(); err != nil {
		stdLog.Fatalf("run handler error: %v", err)
	}
}
