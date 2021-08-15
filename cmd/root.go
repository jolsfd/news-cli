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
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var cfg Config

var count int
var ignore bool

// feeds contains the news feeds structs.
var feeds []NewsFeed

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "news-cli",
	Short: "A CLI news reader for atom feeds",
	Long:  `Read atom feeds (e.g. heise online newsfeed and show details on the command line`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init inits the rootCmd.
func init() {
	cobra.OnInitialize(initConfig)

	// Global flags.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.news-cli.yaml)")
	rootCmd.PersistentFlags().IntVarP(&count, "count", "c", 5, "number of news")
	rootCmd.PersistentFlags().BoolVarP(&ignore, "ignore", "i", false, "ignore value from config file")

	// Version.
	rootCmd.Version = verison
}

// initConfig reads in config file.
func initConfig() {
	// Set config constants
	SetConfig()

	// Get home dir.
	home, err := os.UserHomeDir()
	checkError(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory.
		viper.AddConfigPath(home)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// If no config file is found, create it.
			err = viper.WriteConfigAs(filepath.Join(home, ConfigPath))
			checkError(err)
		}
	}
	err = viper.Unmarshal(&cfg)
	checkError(err)
}

// checkError checks if an error is not nil.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// getNewsFeeds searches for NewsFeeds struct from the config file
// or builds NewsFeeds struct from args.
func getNewsFeeds(args []string) {
	if len(args) > 0 {
		// News feeds from args.
		feeds = BuildNewsFeeds(args, count)
	} else {
		// News feeds from config file.
		for _, newsFeed := range cfg.Feeds {
			// Use count value instead of config.
			if ignore {
				newsFeed.Count = count
			}

			// Append struct to feeds.
			feeds = append(feeds, newsFeed)
		}
	}
}
