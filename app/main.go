package main

import (
	"fmt"

	"github.com/Striker87/notes/internal/config"
)

func main() {
	fmt.Println("config init")
	cfg := config.NewConfig()

	fmt.Println("logger init")
}
