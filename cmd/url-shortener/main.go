package main

import (
	"fmt"

	"github.com/Prrromanssss/URLShortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
