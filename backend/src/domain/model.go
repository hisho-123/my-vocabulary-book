package domain

import (
	// "time"
)

type User struct {
	Id int
	Name string
	Password string
}

type Book struct {
	Id int
	Name string
	// RecentReview time.Time
}

type Word struct {
	Id int
	Word string
	Translated string
}
