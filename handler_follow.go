package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mehkij/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could not get feed: %w", err)
	}

	_, err = s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("* %s followed: %s\n", s.config.CurrentUserName, feed.Name)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get followed feeds: %w", err)
	}

	for _, follow := range following {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
