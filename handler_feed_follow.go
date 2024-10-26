package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/TrungNNg/gator/internal/database"
	"github.com/google/uuid"
)

func handlerDeleteFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("url is required")
	}
	feed, err := s.queries.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	err = s.queries.DeleteFeedByUserAndFeed(context.Background(), database.DeleteFeedByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	return err
}

func handlerGetFeedFollowsForUser(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("extra argument")
	}

	feedsOfUser, err := s.queries.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feed := range feedsOfUser {
		fmt.Println(feed.Name)
	}
	return nil
}

func handlerCreateFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("url is required")
	}
	feed, err := s.queries.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	feedFollowRow, err := s.queries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feedFollowRow.UserName)
	fmt.Println(feedFollowRow.FeedName)
	return nil
}
