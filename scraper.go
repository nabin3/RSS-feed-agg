package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

// This func will be added in main func
func (cfg *apiConfig) StartScraping(concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines....", timeBetweenRequest, concurrency)

	// Creating a ticker
	ticker := time.NewTicker(timeBetweenRequest)

	// Each tick will create a new llisy of feed id to be fetched, but using waitgroup we ensure previos list's all feeds are fetched before creting new list of feeds to be fetched
	for ; ; <-ticker.C {
		feeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("couldn't get next feeds to fetch %v", err)
			continue
		}
		log.Printf("Found %d feeds to be fetched", len(feeds))

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go cfg.scrapFeed(feed, wg)
		}
		wg.Wait()
	}
}

// Will be called by StartScraping
func (cfg *apiConfig) scrapFeed(feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	// Marking a feed as it is fetched
	_, err := cfg.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	// Fetching a particular feed
	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("couldn't collect feed %s error: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		log.Println("Found post", item.Title)
		// Retrieving proper formated date from string
		postPublishingDate, err := time.Parse(time.RFC1123Z, item.Pubdate)
		if err != nil {
			log.Printf("can't retrieve publication date from string: %v", err)
			return
		}
		// Adding each item or post to posts table
		_, err = cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: postPublishingDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Printf("can't create post: %v", err)
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}

// Format of a feed to be collected
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

// Format of each posts or item of each feed
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Pubdate     string `xml:"pubDate"`
}

// This function's job is to fetch a feed with it's URL
func fetchFeed(feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}
