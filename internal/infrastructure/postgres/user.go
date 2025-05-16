package postgres

import "github.com/LavaJover/shvark-user-service/internal/domain"

type UserModel struct {
	ID				string 		`gorm:"primaryKey;type:uuid"`
	Login			string 		`gorm:"unique;not null"`
	Username		string 		`gorm:"unique;not null"`
	PasswordHash	string 		`gorm:"not null"`
	Role			domain.Role	
}