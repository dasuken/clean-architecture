package domain

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Users []User