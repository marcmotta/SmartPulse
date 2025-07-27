// cmd/smartpulse/main.go
package main

import (
	"flag"
	"log"
	"os"

	"smartpulse/internal/smartpulse"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := smartpulse.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
