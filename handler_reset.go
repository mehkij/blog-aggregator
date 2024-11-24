package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users table: %w", err)
	}

	fmt.Println("Successfully reset users table!")
	return nil
}
