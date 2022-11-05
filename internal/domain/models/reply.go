package models

type Reply struct {
	Rank
	Parent Rank `json:"parent"`
}
