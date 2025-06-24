package postgres

import "time"

type UserModel struct {
	ID				string `gorm:"primaryKey;type:uuid"`
	Login			string `gorm:"unique;not null"`
	Username		string `gorm:"unique;not null"`
	PasswordHash	string `gorm:"not null"`
	TwoFaSecret 	string	`gorm:"two_fa_secret"`
	TwoFaEnabled 	bool	`gorm:"two_fa_enabled"`
	CreatedAt 		time.Time
	UpdatedAt		time.Time
}