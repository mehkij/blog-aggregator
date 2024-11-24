package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mehkij/blog-aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	err := s.config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set the current user: %w", err)
	}

	fmt.Println("User successfully logged in!")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username, err := s.db.GetUser(context.Background())
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if username == cmd.Args[0] {
		fmt.Println("User already exists!")
		os.Exit(1)
		return nil
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.Args[0]})
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = s.config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't set the current user: %w", err)
	}

	fmt.Println("User successfully registered!")
	printUser(user)

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
