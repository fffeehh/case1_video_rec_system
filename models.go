package main

import "time"

type Video struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
	Views       int       `json:"views"`
	Likes       int       `json:"likes"`
	CreatedAt   time.Time `json:"created_at"` // Go сам поймет этот формат даты
}

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ViewedVideoIDs []int  `json:"viewed_video_ids"`
}

type Recommender interface {
	Build(user User, allUsers []User, allVideos []Video) []Video
}
