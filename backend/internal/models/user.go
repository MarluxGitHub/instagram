package models

type User struct {
	ID 			uint 	`gorm:"primary_key, index, AUTO_INCREMENT"`
	Password 	string
	Username 	string
	Email 		string
}

func (u *User) String() string {
	return "User: " + u.Username + " " + u.Email
}