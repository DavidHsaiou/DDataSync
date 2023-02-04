package main

import (
	"dsync/logger"
	"dsync/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(server.NewHttpServer),
		fx.Invoke(func(server server.IServer) {
			server.Run()
		}),
	).Run()
}
