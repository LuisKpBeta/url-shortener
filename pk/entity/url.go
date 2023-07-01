package entity

import "time"

type Url struct {
	Id        int       `json:"id"`
	Original  string    `json:"original"`
	Shortened string    `json:"shortened_url"`
	CreatedAt time.Time `json:"created_at"`
}
