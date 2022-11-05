package models

import "time"

type Rank struct {
	UUID    string    `json:"uuid"`
	Message string    `json:"message"`
	PubDate time.Time `json:"pubDate"`
	Edited  bool      `json:"edited"`
	Deleted bool      `json:"deleted"`
	Author  User      `json:"author"`
}
