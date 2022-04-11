package model

import "time"

type Link struct {
	Id            uint64    `json:"id"`
	OriginLink    string    `json:"origin_link"`
	ShortenedLink string    `json:"shortened_link"`
	CreatedAt     time.Time `json:"created_at"`
}
