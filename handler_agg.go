package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("could not fetch RSSFeed: %w", err)
	}

	fmt.Printf("%v\n", feed)

	return nil
}
