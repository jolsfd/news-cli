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
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start local web server",
	Long: `Start local web server where articles will be displayed.
Number of displayed items can be changed in the config file or via flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd, args)
	},
}

// init inits serve command.
func init() {
	rootCmd.AddCommand(serveCmd)
}

// serve is the function run when serveCmd is called.
func serve(cmd *cobra.Command, args []string) {
	getNewsFeeds(args)

	// Output.
	fmt.Printf("Web server started on port %s. End server with key combination [Ctrl + C].\n", cfg.Port)

	// Fileserver for css.
	fileServer := http.FileServer(http.Dir(cssDir))
	http.Handle("/", fileServer)

	// Handle news.
	http.HandleFunc("/news", newsHandler)

	// Start web server.
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}

// newsHandler handles the news site.
func newsHandler(writer http.ResponseWriter, request *http.Request) {
	var news []NewsList

	// Check feeds from config file.
	for _, newsFeed := range feeds {
		// Get news feed.
		newsList, err := newsFeed.GetNews()
		checkError(err)
		news = append(news, newsList)
	}

	// Parse html.
	html, err := template.ParseFiles(htmlPath)
	checkError(err)
	err = html.Execute(writer, news)
	checkError(err)
}
