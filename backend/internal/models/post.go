package models


type Post struct {
	ID 			uint 	`gorm:"primary_key, index, AUTO_INCREMENT"`
	Image_url 	string
	Image_url_type string
	Caption 	string
	Timestamp 	string

	// Foreign key for user
	UserID 		uint
}