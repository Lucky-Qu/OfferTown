// Package service user.go
//
// 功能:
// - 用户注册
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-25

package service

import (
	"backend/internal/code"
	"backend/internal/crypto"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/validator"
)

func RegisterUser(userCreateDTO *dto.UserCreateDTO) code.Code {
	//对传入的dto进行合法性检查
	if ok := validator.UsernameCheck(userCreateDTO.Username); !ok {
		return code.InvalidUsername
	}
	if ok := validator.UsernameExistCheck(userCreateDTO.Username); ok {
		return code.UsernameAlreadyExist
	}
	if ok := validator.PasswordCheck(userCreateDTO.Password); !ok {
		return code.InvalidPassword
	}
	//对用户密码进行加密，避免明文存储
	encryptedPassword, eCode := crypto.Encrypt(userCreateDTO.Password)
	if eCode != code.Success {
		return eCode
	}
	//构建用户对象
	user := model.User{
		Username:          userCreateDTO.Username,
		EncryptedPassword: encryptedPassword,
		Signature:         "",
		AvatarPath:        "./static/images/defaultAvatar.jpg",
	}
	err := repository.RegisterUser(&user)
	if err != nil {
		return code.DatabaseError
	}
	return code.Success
}
