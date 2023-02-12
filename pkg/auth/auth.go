package auth

import (
	"SnowLynxSoftware/lynx-identity-server/pkg/database"
	"SnowLynxSoftware/lynx-identity-server/pkg/users"
	"SnowLynxSoftware/lynx-identity-server/pkg/util"
	"fmt"
)

func RegisterNewUserAccount(registerDTO *users.UserRegisterDTO) *users.UserEntity {
	password, err := util.HashPassword(registerDTO.Password)
	if err != nil {
		return nil
	}

	userEntity := users.UserEntity{Email: registerDTO.Email, PassHash: password}

	sql := fmt.Sprintf("INSERT INTO users (email, passHash) VALUES ('%s', '%s')", userEntity.Email, userEntity.PassHash)

	insert, err := database.DBConnection.Query(sql)

	err = insert.Close()
	if err != nil {
		return nil
	}

	createdUser, findErr := users.GetUserByEmail(userEntity.Email)
	if findErr != nil {
		return nil
	} else {
		return createdUser
	}
}
