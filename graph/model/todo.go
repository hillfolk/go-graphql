package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done,omitempty"`
	UserID string `json:"userId"`
	User   *User  `json:"user"`
}

func (t Todo) Owner() *User {
	return t.User
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Ownable interface {
	Owner() *User
}
