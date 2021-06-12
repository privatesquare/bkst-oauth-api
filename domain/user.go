package domain

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}
