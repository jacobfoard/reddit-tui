package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thecsw/mira"
)

func main() {
	subreddit := flag.String("subreddit", "all", "Subreddit to list")
	flag.Parse()
	creds := mira.Credentials{
		ClientId:     os.Getenv("REDDIT_CLIENT"),
		ClientSecret: os.Getenv("REDDIT_SECRET"),
		Username:     os.Getenv("REDDIT_USERNAME"),
		Password:     os.Getenv("REDDIT_PASSWORD"),
		UserAgent:    "reddit-tui-golang/v1",
	}

	r, err := mira.Init(creds)
	if err != nil {
		panic(err)
	}

	posts, err := r.Subreddit(*subreddit).Submissions("top", "all", 25)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		fmt.Printf("%s\n", post.GetTitle())
		fmt.Printf("Author: %s   Points: %.0f\n", post.GetAuthor(), post.GetKarma())
		fmt.Println("")
	}

}
