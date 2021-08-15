/*
Copyright © 2021 jolsfd

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list articles",
	Long: `Show headline, description, link and time from articles.
Number of displayed items can be changed in the config file or via flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		list(cmd, args)
	},
}

// init inits listCmd.
func init() {
	rootCmd.AddCommand(listCmd)
}

// list is the function run when listCmd is called.
func list(cmd *cobra.Command, args []string) {
	getNewsFeeds(args)

	for _, newsFeed := range feeds {
		// Get feed news.
		newsList, err := newsFeed.GetNews()
		checkError(err)

		// Show news.
		ListNews(newsList)
	}
}
