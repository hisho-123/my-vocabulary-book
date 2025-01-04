package domain

// "time"

type User struct {
	Id       *int
	Name     string
	Password string
}

type Book struct {
	Id     *int
	UserId int
	Name   string
	// RecentReview time.Time
}

type Word struct {
	Id         *int
	BookId     int
	Word       string
	Translated string
}
