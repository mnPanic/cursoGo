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
func (user User) Equals(other User) bool {
	return (other.Name == user.Name &&
		other.Password == user.Password)
}

//String returns the user as a printable string (just the name)
func (user User) String() string {
	return user.Name
}
