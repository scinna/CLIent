package main

import (
	"fmt"
	"github.com/scinna/CLIent/config"
	"os"
)

func main() {
	fmt.Println("Scinna CLIent")

	cfg, err := config.ReadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if !cfg.IsConfigValid() {
		config.Setup(cfg)
	}

	if !cfg.AreDefaultValid() {
		fmt.Println("The defaults are not correctly set. Please check the config file")
		return
	}

}
