package main

import (
	"fmt"
	"log"

	"github.com/modern-dev-dude/polyglot-programming/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("Unable to get the options %v", err)
	}

	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("Unable to get the config %v", err)
	}

	fmt.Printf("opts: %+v", opts)
	fmt.Printf("config: %+v", config)
}
