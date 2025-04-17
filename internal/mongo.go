package internal

import (
	"time"
)

// User Spotify user profile
type User struct {
	ID          string    `bson:"_id"`
	DisplayName string    `bson:"display_name"`
	Email       string    `bson:"email"`
	Country     string    `bson:"country"`
	FetchedAt   time.Time `bson:"fetched_at"`
}

// Playlist Simplified playlist
type Playlist struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	OwnerID   string    `bson:"owner_id"`
	Tracks    int       `bson:"tracks"`
	FetchedAt time.Time `bson:"fetched_at"`
}

// Artist Simplified artist
type Artist struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Genres    []string  `bson:"genres"`
	FetchedAt time.Time `bson:"fetched_at"`
}
