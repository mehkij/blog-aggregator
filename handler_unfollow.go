package main

import (
	"context"
	"fmt"

	"github.com/mehkij/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("could not unfollow: %w", err)
	}

	return nil
}
