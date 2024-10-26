package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/TrungNNg/gator/internal/config"
	"github.com/TrungNNg/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg     *config.Config
	queries *database.Queries
}

func main() {
	cfg := config.Read()

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("can not open database")
	}
	queries := database.New(db)
	programState := &state{
		cfg:     &cfg,
		queries: queries,
	}
	cmds := commands{
		cmds: map[string]func(*state, command) error{},
	}

	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerFetchRss)
	cmds.register("feeds", handlerGetAllFeeds)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("follow", middlewareLoggedIn(handlerCreateFeedFollow))
	cmds.register("following", middlewareLoggedIn(handlerGetFeedFollowsForUser))
	cmds.register("unfollow", middlewareLoggedIn(handlerDeleteFeed))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}
	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.queries.GetUser(context.Background(), s.cfg.UserName)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
