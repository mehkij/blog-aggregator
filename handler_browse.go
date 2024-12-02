package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mehkij/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32

	if len(cmd.Args) != 1 {
		limit = 2
	} else {
		i, err := strconv.ParseInt(cmd.Args[0], 10, 0)
		if err != nil {
			return fmt.Errorf("unable to parse int: %w", err)
		}

		limit = int32(i)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("unable to get posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("* Title: %s\n Description: %s\n", post.Title.String, post.Description.String)
	}

	return nil
}
