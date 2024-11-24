package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mehkij/blog-aggregator/internal/config"
	"github.com/mehkij/blog-aggregator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{
		db:     dbQueries,
		config: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

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
