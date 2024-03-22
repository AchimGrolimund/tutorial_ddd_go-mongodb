package domain

type User struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
}

type UserRepository interface {
	GetUser(id string) (User, error)
	CreateUser(user User) (string, error)
}
