package postgres

import "time"

type UserModel struct {
	ID				string `gorm:"primaryKey;type:uuid"`
	Login			string `gorm:"unique;not null"`
	Username		string `gorm:"unique;not null"`
	PasswordHash	string `gorm:"not null"`
	TwoFaSecret 	string
	TwoFaEnabled 	bool
	CreatedAt 		time.Time
	UpdatedAt		time.Time
}