package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mehkij/blog-aggregator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	programState := &state{
		config: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	// 1st arg is program name, which we ignore.
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
