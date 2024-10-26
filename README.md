---
# Gator
---
Gator is a command-line tool that fetches and manages RSS feeds. It allows users to add, follow, and view RSS feeds and posts conveniently through a terminal interface.

## Prerequisites
Before you can run Gator, you'll need to have both PostgreSQL and Go installed on your system:
* PostgreSQL: Make sure you have a running instance of PostgreSQL and that you know the connection details (e.g., host, port, user, password, database).
* Go: You need Go installed to compile the program. You can download it from https://golang.org/dl/.

## Installation
To install the Gator CLI tool, run the following command:
``` bash
go install github.com/TrungNNg/gator@latest
```
This will download and compile the Gator binary, making it available in your Go binary path (usually `$GOPATH/bin`).

## Configuration
Gator requires a configuration file to be set up before running. This file should be located at `~/.gatorconfig.json` and contain the following structure:

```json
{
    "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
    "current_user_name": "your_username"
}
```
**db_url**: Replace username, password, localhost, and 5432 with your PostgreSQL connection details.
**current_user_name**: The username that you want to use for Gator's commands.

To create the configuration file, you can use the following command:
```bash
echo '{"db_url":"postgres://username:password@localhost:5432/gator?sslmode=disable","current_user_name":"your_username"}' > ~/.gatorconfig.json
```

## Running the Program
To start using Gator, run the following command in your terminal:

```bash
Copy code
gator <command> [arguments]
```

Some of the commands you can use:
* register: Register a new user.
* login: Log in an existing user.
* addfeed: Add a new RSS feed.
* follow: Follow a feed.
* following: List the feeds that the current user is following.

For example, to register a new user:
```bash
Copy code
gator register
```

To add a new feed:
```bash
Copy code
gator addfeed <feed-name> <feed-url>
```