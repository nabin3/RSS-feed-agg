package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

// Format of responsing a newly created user's data
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

// This function maps databse user data to response-usaer data
func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIKey:    user.ApiKey,
	}
}

// Format for responsing newly created feed
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

/*// This function maps databse feed data to response-feed data
func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}*/

// Format of responding feed and feed_follow of a newly created feed
type FeedAndFeedFollow struct {
	NewFeed       Feed        `json:"feed"`
	NewFeedFollow Feed_Follow `json:"feed_follow"`
}

// This function take feed and feed_follow and stich them together
func databseFeedAndfeedFollowToRespFeedAndFeedFollow(feed database.Feed, feedFollow database.FeedFollow) FeedAndFeedFollow {
	return FeedAndFeedFollow{
		NewFeed: Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
			UserID:    feed.UserID,
		},

		NewFeedFollow: Feed_Follow{
			ID:        feedFollow.ID,
			CreatedAt: feedFollow.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			UserID:    feedFollow.UserID,
			FeedID:    feedFollow.FeedID,
		},
	}
}

// This function maps databse all-feed data to array of all response-feed data
func allDatabaseFeedToAllFeed(feeds []database.Feed) []Feed {
	// Cretaing an empty slice to store all new mapped feeds
	all_feeds := make([]Feed, 0)

	// Iterating each databse feed, map them to response Feed and adding them to our all_feeds response payload one by one
	for _, each_retrieved_feed := range feeds {
		new_feed := Feed{
			ID:        each_retrieved_feed.ID,
			CreatedAt: each_retrieved_feed.CreatedAt,
			UpdatedAt: each_retrieved_feed.UpdatedAt,
			Name:      each_retrieved_feed.Name,
			Url:       each_retrieved_feed.Url,
			UserID:    each_retrieved_feed.UserID,
		}
		all_feeds = append(all_feeds, new_feed)
	}

	return all_feeds
}

// Format of feed-follow object
type Feed_Follow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

// Function to convert database feed-follow to response feed-follow
func databaseFeedFollowToRespFeedFollow(feedFollow database.FeedFollow) Feed_Follow {
	return Feed_Follow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

// Function to represent all feed-follows of an user
func allDatabaseFeedFollowsToAllFeedFollowsOfUser(FeedFollows []database.FeedFollow) []Feed_Follow {
	// Cretaing an empty slice to store all new mapped feed_follows
	feed_follows := make([]Feed_Follow, 0)

	// Iterating each database FeedFollow, map them to response feed_follow and adding them to our feed_follows response payload one by one
	for _, eachRetrievedFeedFollow := range FeedFollows {
		newFeedFollow := Feed_Follow{
			ID:        eachRetrievedFeedFollow.ID,
			CreatedAt: eachRetrievedFeedFollow.CreatedAt,
			UpdatedAt: eachRetrievedFeedFollow.UpdatedAt,
			UserID:    eachRetrievedFeedFollow.UserID,
			FeedID:    eachRetrievedFeedFollow.FeedID,
		}
		feed_follows = append(feed_follows, newFeedFollow)
	}

	return feed_follows
}

// Format of response of a post
type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
}

// Function to respond all posts
func allDatabasePostsToPosts(posts []database.GetPostsByUserRow) []Post {
	// Cretaing an empty slice to store all new mapped posts
	all_posts := make([]Post, 0)

	// Iterating each database post, map them to response post and adding them to our all_posts response payload one by one
	for _, eachRetrievedPosts := range posts {
		newPost := Post{
			ID:          eachRetrievedPosts.ID,
			Title:       eachRetrievedPosts.Title,
			Url:         eachRetrievedPosts.Url,
			Description: eachRetrievedPosts.Description,
			PublishedAt: eachRetrievedPosts.PublishedAt,
		}
		all_posts = append(all_posts, newPost)
	}

	return all_posts
}
