package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var openID int

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open on browser by ID",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(fmt.Sprintf("https://dev.to/api/articles/%d", openID))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		var post Post
		err = json.NewDecoder(resp.Body).Decode(&post)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		exec.Command("open", "-a", "firefox", post.URL).Run()
	},
}

func init() {
	openCmd.Flags().IntVarP(&openID, "id", "", 0, "ID of article")
}
