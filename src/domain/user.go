package domain

//User of tweeter
type User struct {
	Name string
}

//NewUser Creates a new user
func NewUser(name string) User {
	return User{Name: name}
}
