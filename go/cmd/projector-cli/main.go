package main

import (
	"encoding/json"
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

	proj := projector.NewProjector(config)
	if config.Operation == projector.Print {
		if len(config.Args) == 0 {
			data := proj.GetValueAll()
			jsonStr, err := json.Marshal(data)

			if err != nil {
				log.Fatalf("unable to get config %v", err)
			}

			fmt.Printf("%v", string(jsonStr))
		} else if value, ok := proj.GetValue(config.Args[0]); ok {
			fmt.Printf("%v", value)
		}
	}

	if config.Operation == projector.Add {
		proj.SetValue(config.Args[0], config.Args[1])
		proj.Save()
	}

	if config.Operation == projector.Remove {
		proj.RemoveValue(config.Args[0])
		proj.Save()
	}
}
