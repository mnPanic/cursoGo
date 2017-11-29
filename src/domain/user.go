package domain

//User of tweeter
type User struct {
	Name, Password string
}

//NewUser Creates a new user
func NewUser(name, password string) User {
	return User{Name: name, Password: password}
}

//Equals returns if two users are the same
func (user1 User) Equals(user User) bool {
	return (user1.Name == user.Name &&
		user1.Password == user.Password)
}
