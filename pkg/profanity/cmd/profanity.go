package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wcharczuk/advent/pkg/profanity"
)

func main() {
	var cfg profanity.Config
	flag.StringVar(&cfg.RulesFile, "rules", profanity.DefaultRulesFile, "The rules file to check for in each directory")
	flag.Parse()

	engine := profanity.New(profanity.OptConfig(cfg))
	engine.Stdout = os.Stdout
	engine.Stderr = os.Stderr

	if err := engine.Process(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}
}
