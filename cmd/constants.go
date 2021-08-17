package cmd

// Constants for web server.
const (
	htmlPath = "template/view.html"
	cssPath  = "template"
)

// Constants for config file
const (
	portString = "port"
	feedString = "feed"
	port       = ":8080"

	configName = ".news-cli"
	configType = "yaml"
	ConfigPath = ".news-cli.yaml"
)

// defaultFeed is the feed at first start.
var (
	defaultFeed = []NewsFeed{{"Heise Online", "https://www.heise.de/rss/heise-atom.xml", 5}}
)

// Version.
const verison = "0.0.0"
