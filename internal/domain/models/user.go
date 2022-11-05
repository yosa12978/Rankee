package models

import "time"

type User struct {
	UUID     string    `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Icon     string    `json:"icon"`
	Ranks    []Rank    `json:"ranks"`
	Subs     []User    `json:"subs"`
	RegDate  time.Time `json:"regDate"`
	Deleted  bool      `json:"deleted"`
	Active   bool      `json:"active"`
}
