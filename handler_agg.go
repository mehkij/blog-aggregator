package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mehkij/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <1s|1m|1h>", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could not set time between reqs: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		fmt.Println("Sending req...")
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:     time.Now().UTC(),
		FeedID:        next.FeedID,
	})
	if err != nil {
		return fmt.Errorf("unable to mark feed as fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), next.FeedUrl)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
	}

	for _, item := range feed.Channel.Item {
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullString{String: item.PubDate, Valid: true},
			FeedID:      next.FeedID,
		})
		if err != nil {
			return fmt.Errorf("unable to create post: %w", err)
		}
	}

	return nil
}
