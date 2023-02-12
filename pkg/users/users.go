package users

import (
	"SnowLynxSoftware/lynx-identity-server/pkg/database"
	"fmt"
	"time"
)

type UserRegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEntity struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PassHash     string    `json:"pass_hash"`
	Verified     bool      `json:"verified"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
	ArchivedAt   time.Time `json:"archived_at"`
	BannedAt     time.Time `json:"banned_at"`
}

func GetUserByEmail(email string) (*UserEntity, error) {
	var userEntity UserEntity

	sql := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)

	err := database.DBConnection.QueryRow(sql).Scan(&userEntity.ID, &userEntity.Email, &userEntity.BannedAt, &userEntity.PassHash, &userEntity.CreatedAt, &userEntity.LastModified, &userEntity.Verified, &userEntity.ArchivedAt)
	if err != nil {
		return &UserEntity{}, err
	} else {
		return &userEntity, nil
	}
}

func GetUserByID(id int) *UserEntity {
	return nil
}
