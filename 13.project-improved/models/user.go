package models

type User struct {
	ID   int
	Name string
}

var (
	users  []*User // slice of pointers of users
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
