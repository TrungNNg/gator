package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/TrungNNg/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("feed name and url is required")
	}

	feed, err := s.queries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	_, err = s.queries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	return nil
}

func handlerGetAllFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return errors.New("extra args")
	}
	feeds, err := s.queries.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		user, err := s.queries.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
		fmt.Println(feed.LastFetchedAt)
	}
	return nil
}
