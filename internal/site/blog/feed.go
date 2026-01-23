package blog

import (
	"encoding/xml"
	"net/http"
	"time"

	"circuitcamel.com/internal/cache"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language,omitempty"`
	PubDate     string `xml:"pubDate,omitempty"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description,omitempty"`
	PubDate     string `xml:"pubDate,omitempty"`
	GUID        string `xml:"guid,omitempty"`
}

func Feed(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://circuitcamel.com"

	items := make([]Item, 0, len(cache.BlogPosts))
	for _, post := range cache.BlogPosts {
		t, err := time.Parse("02-01-2006", post.Date)
		pubDate := ""
		if err == nil {
			pubDate = t.Format(time.RFC1123Z)
		}

		items = append(items, Item{
			Title:       post.Title,
			Link:        baseURL + "/blog/" + post.Slug,
			Description: post.Body,
			PubDate:     pubDate,
			GUID:        baseURL + "/blog/" + post.Slug,
		})
	}

	channelPubDate := ""
	if len(cache.BlogPosts) > 0 {
		t, err := time.Parse("02-01-2006", cache.BlogPosts[0].Date)
		if err == nil {
			channelPubDate = t.Format(time.RFC1123Z)
		}
	}

	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       "CircuitCamel",
			Link:        baseURL + "/blog",
			Description: "Blog posts from CircuitCamel",
			Language:    "en-us",
			PubDate:     channelPubDate,
			Items:       items,
		},
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(xml.Header))
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")
	encoder.Encode(rss)
}
