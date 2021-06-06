package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List latest Devdose blog post",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://dev.to/api/articles?username=iporsut")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		var posts []Post
		err = json.NewDecoder(resp.Body).Decode(&posts)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, post := range posts {
			fmt.Printf("ID: %d\n", post.ID)
			fmt.Printf("Title: %s\n", post.Title)
			fmt.Printf("URL: %s\n", post.URL)
			fmt.Println()
		}
	},
}
