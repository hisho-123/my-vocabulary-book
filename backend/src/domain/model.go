package domain

type User struct {
	Id       *int   `json:"userId"`
	Name     string `json:"userName"`
	Password string `json:"password"`
}

type Book struct {
	Id          *int    `json:"bookId"`
	UserId      int     `json:"userId"`
	Name        string  `json:"bookName"`
	FirstReview *string `json:"firstReview"`
}

type Word struct {
	Id         *int   `json:"wordId"`
	BookId     int    `json:"bookId"`
	Word       string `json:"word"`
	Translated string `json:"translated"`
}
