package main

import (
	"os"

	"github.com/h4shu/shiritori-go/cmd"
)

func main() {
	port := os.Getenv("PORT")
	cmd.Server(port)
}
