package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"-"`
	Email      string             `bson:"email" json:"email"`
	TelegramID string             `bson:"telegram_id" json:"telegram_id"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

func (u *User) SetPassword(password string) error {
	// Generate a salt and hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password
	u.Password = string(hashedPassword)

	return nil
}
