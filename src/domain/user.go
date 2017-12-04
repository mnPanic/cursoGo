package domain

//User of tweeter
type User struct {
	Name, Password string
	Following      []User
}

//NewUser Creates a new user
func NewUser(name, password string) User {
	return User{Name: name, Password: password}
}

//Equals returns if two users are the same
func (u User) Equals(other User) bool {
	return (other.Name == u.Name &&
		other.Password == u.Password)
}

//String returns the user as a printable string (just the name)
func (u User) String() string {
	return u.Name
}

//Follow follows another user
func (u *User) Follow(toFollow User) {
	u.Following = append(u.Following, toFollow)
}

//IsFollowing returns if a user is following another one
func (u User) IsFollowing(userToCheck User) bool {
	for _, followedUsers := range u.Following {
		if userToCheck.Equals(followedUsers) {
			return true
		}
	}
	return false
}
