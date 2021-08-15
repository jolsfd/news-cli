package cmd_test

import (
	"testing"

	"github.com/jolsfd/news-cli/cmd"
	"github.com/spf13/viper"
)

func TestSetConfig(t *testing.T) {
	var got cmd.Config

	cmd.SetConfig()

	err := viper.Unmarshal(&got)
	if err != nil {
		t.Error(err)
	}

	var want cmd.Config

	want.Port = ":8080"
	want.Feeds = []cmd.NewsFeed{{Name: "Heise Online", URL: "https://www.heise.de/rss/heise-atom.xml", Count: 5}}

	if want.Port != got.Port {
		t.Errorf("got: %s; want: %s", got.Port, want.Port)
	}

	for i, newsFeed := range got.Feeds {
		if newsFeed != want.Feeds[i] {
			t.Errorf("got: %v; want: %v", got.Feeds, want.Feeds)
		}
	}
}

func TestGetNews(t *testing.T) {
	newsFeed := cmd.NewsFeed{Name: "Heise Online", URL: "https://www.heise.de/rss/heise-atom.xml", Count: 5}
	newsList, err := newsFeed.GetNews()
	if err != nil {
		t.Error(err)
	}
	cmd.ListNews(newsList)
}

func TestBuildNewsFeeds(t *testing.T) {
	urls := []string{"https://www.heise.de/rss/heise-atom.xml"}
	count := 10

	want := []cmd.NewsFeed{{Name: "heise.de", URL: "https://www.heise.de/rss/heise-atom.xml", Count: 10}}
	got := cmd.BuildNewsFeeds(urls, count)

	for i, newsFeed := range got {
		if newsFeed != want[i] {
			t.Errorf("got: %v; want: %v;", got, want)
		}
	}
}
