package cmd

import (
	"regexp"

	"github.com/fatih/color"
	"github.com/mmcdole/gofeed"
)

// NewsList contains a list of News and a name.
type NewsList struct {
	News []News `json:"news,omitempty"`
	Name string `json:"name,omitempty"`
}

// News is a structure which contains informations about a news.
type News struct {
	Date        string `json:"date,omitempty"`
	GUID        string `json:"guid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link,omitempty"`
}

// NewsFeed holds configurations for a news feed.
type NewsFeed struct {
	Name  string `mapstructure:"name"`
	URL   string `mapstructure:"url"`
	Count int    `mapstructure:"count"`
}

// GetNews fetches a certain number of news from a feed and returns a NewsList struct.
func (n *NewsFeed) GetNews() (NewsList, error) {
	list := NewsList{Name: n.Name}
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(n.URL)
	if err != nil {
		return list, err
	}

	for i, item := range feed.Items {
		if i < n.Count {
			news := News{item.Published, item.GUID, item.Title, item.Description, item.Link}
			list.News = append(list.News, news)
		} else {
			break
		}
	}
	return list, nil
}

// ListNews takes a NewsList and prints the elements to the terminal.
func ListNews(feed NewsList) {
	// Define colors.
	feedNameColor := color.New(color.Bold, color.FgBlue)
	headlineColor := color.New(color.Bold, color.FgWhite)
	descriptionColor := color.New()
	linkColor := color.New(color.Italic)
	dateColor := color.New(color.Italic)

	// Output.
	feedNameColor.Printf("%s\n\n", feed.Name)

	for _, news := range feed.News {
		headlineColor.Printf("%s\n\n", news.Title)
		descriptionColor.Printf("%s\n\n", news.Description)
		linkColor.Printf("%s\n\n", news.Link)
		dateColor.Printf("%s\n\n", news.Date)
	}
}

// BuildNewsFeeds takes an slice of urls and a number. It returns an slice of NewsFeeds sructs.
func BuildNewsFeeds(urls []string, count int) (newsFeeds []NewsFeed) {
	// Regular Expression.
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)

	// Check feeds from args.
	for _, url := range urls {
		// Create NewsFeed struct.
		newsFeeds = append(newsFeeds, NewsFeed{Name: re.FindStringSubmatch(url)[1], URL: url, Count: count})
	}

	return newsFeeds
}
