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
	"backend/internal/auth"
	"backend/internal/cache"
	"backend/internal/code"
	"backend/internal/crypto"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/validator"
)

// RegisterUser 注册用户功能
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

// UserLogin 用户登录功能
func UserLogin(userDTO *dto.UserLoginDTO) (string, code.Code) {
	// 判断用户名是否存在
	if ok, err := repository.CheckUsername(userDTO.Username); !ok {
		if err != nil {
			return "", code.DatabaseError
		}
		return "", code.UserNotExists
	}

	user, err := repository.GetUserByUsername(userDTO.Username)
	if err != nil {
		return "", code.DatabaseError
	}

	//比对密码
	ok, eCode := crypto.Verify(user.EncryptedPassword, userDTO.Password)
	if eCode != code.Success {
		return "", eCode
	}
	if !ok {
		return "", code.PasswordWrong
	}
	// 检查是否已在登陆状态，避免重复生成Token
	check, err := cache.CheckJWTIsExists(user.Username)
	if err != nil {
		return "", code.CacheError
	}
	if check {
		token, err := cache.GetJWTByUsername(user.Username)
		if err != nil {
			return "", code.CacheError
		}
		return token, code.Success
	}

	token, eCode := auth.GetToken(int(user.ID), user.Username)
	if eCode != code.Success {
		return "", eCode
	}
	//存入缓存
	cache.SetJWT(token)
	return token, code.Success
}

// GetUserInfoByUsername 通过用户名获得用户DTO对象
func GetUserInfoByUsername(username string) (*dto.UserInfoDTO, code.Code) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return nil, code.DatabaseError
	}
	userDTO := dto.UserInfoDTO{
		Username:   user.Username,
		Signature:  user.Signature,
		AvatarPath: user.AvatarPath,
	}
	return &userDTO, code.Success
}
