package main

import (
	"fmt"
	"github.com/banderveloper/go-forms/internal/config"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)
}
