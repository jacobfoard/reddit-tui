package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/jacobfoard/reddit-tui/pkg/reddit"
	"github.com/jesseduffield/gocui"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()

	g, err := gocui.NewGui(gocui.Output256, false, log.NewEntry(logger))
	if err != nil {
		logger.Fatal(err)
	}
	defer g.Close()

	g.SetManager(gocui.ManagerFunc(layout))
	// if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
	// 	log.fmt.Printlnln(err)
	// }
	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		fmt.Println(err)

	}
	// subreddit := flag.String("subreddit", "all", "Subreddit to list")
	// flag.Parse()
	// creds := mira.Credentials{
	// 	ClientId:     os.Getenv("REDDIT_CLIENT"),
	// 	ClientSecret: os.Getenv("REDDIT_SECRET"),
	// 	Username:     os.Getenv("REDDIT_USERNAME"),
	// 	Password:     os.Getenv("REDDIT_PASSWORD"),
	// 	UserAgent:    "reddit-tui-golang/v1",
	// }

	// r, err := mira.Init(creds)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// posts, err := r.Subreddit(*subreddit).Submissions("", "", 25)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, post := range posts {
	// 	fmt.Printf("%s\n", post.GetTitle())
	// 	fmt.Printf("Author: %s   Points: %.0f\n", post.GetAuthor(), post.GetKarma())
	// 	fmt.Println("")
	// }

}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("hello", 0, 0, maxX-1, maxY-1, 0); err != nil {

		if errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		body, err := ioutil.ReadFile("popular.json")
		if err != nil {
			fmt.Println(err)
		}

		subreddits := &reddit.SubredditListing{}
		err = json.Unmarshal(body, subreddits)
		if err != nil {
			fmt.Println(err)
		}
		for idx, subreddit := range subreddits.ListingData.Children {
			if idx == 0 {
				fmt.Println("")
				fmt.Print(" > ")
			} else {
				fmt.Print("  ")
			}
			fmt.Println(subreddit.T5Data.DisplayNamePrefixed)
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
