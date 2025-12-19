package main

import (
	"fmt"

	"github.com/aDolgosheev/restApiService/cmd/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
