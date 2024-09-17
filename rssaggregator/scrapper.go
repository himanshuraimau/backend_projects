//scrapper.go
package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"
	"strings"
	"github.com/google/uuid"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
)

func startScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,

) {
	log.Println("Starting on %v goroutines every %s seconds", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("Cannot get feeds to fetch: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}


func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	// Mark the feed as fetched in the database
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Cannot mark feed as fetched: %v", err)
		return
	}

	// Fetch the RSS feed from the URL
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Cannot fetch feed: %v", err)
		return
	}

	// Iterate over items in the RSS feed
	for _, item := range rssFeed.Channel.Items {
		// Handle empty description
		description := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}

		// Parse the publication date from the RSS item
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			// Use the current time if parsing fails
			log.Printf("Cannot parse pubDate: %v, using current time instead", err)
			pubDate = time.Now().UTC()
		}

		// Generate a new UUID for the post
		postID, err := uuid.NewUUID()
		if err != nil {
			log.Printf("Cannot generate UUID for post: %v", err)
			continue
		}

		// Create the post in the database
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          postID,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			// Ignore duplicate key errors
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Cannot create post: %v", err)
		}
	}

	// Log the completion of feed fetching
	log.Printf("Feed %s fetched, %v posts found", feed.Name, len(rssFeed.Channel.Items))
}
