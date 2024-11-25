package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve feeds: %w", err)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get users: %w", err)
	}

	for _, feed := range feeds {
		for _, user := range users {
			if feed.UserID == user.ID {
				fmt.Printf("* Name: %s, URL: %s, Created by: %s\n", feed.Name, feed.Url, user.Name)
			}
		}
	}

	return nil
}
