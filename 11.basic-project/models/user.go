package models

type User struct {
	ID   int
	Name string
}

var (
	users  []*User // slice of pointers of users
	nextID = 1
)
