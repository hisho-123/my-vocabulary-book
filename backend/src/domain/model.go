package domain

type User struct {
	Id       *int   `json:"userId"`
	Name     string `json:"userName"`
	Password string `json:"password"`
}

type Book struct {
	Id          int    `json:"bookId"`
	UserId      int     `json:"userId"`
	Name        string  `json:"bookName"`
	FirstReview string `json:"firstReview"`
}

type Word struct {
	Id         *int   `json:"wordId"`
	BookId     int    `json:"bookId"`
	Word       string `json:"word"`
	Translated string `json:"translated"`
}

type AuthInput struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AuthOutput struct {
	UserId int `json:"userId"`
	Token  string `json:"token"`
}

type CreateWordInput struct {
	Word       string `json:"word"`
	Translated string `json:"translated"`
}

type CreateBookInput struct {
	BookName string `json:"bookName"`
	Words    []Word `json:"words"`
}

type GetBookListOutput struct {
	BookId 		int			`json:"bookId"`
	BookName	string 	`json:"bookName"`
}
