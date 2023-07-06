package models

type User struct {
	ID 			uint
	Password 	string
	FirstName 	string
	LastName 	string
	Email 		string
}

func (u *User) String() string {
	return u.FirstName + " " + u.LastName
}