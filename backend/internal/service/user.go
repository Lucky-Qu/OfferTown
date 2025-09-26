// Package service user.go
//
// 功能:
// - 用户注册
// - 用户登录
// - 用户更新信息
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-26

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
	"strconv"
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
	check, err := cache.CheckJWTIsExists(strconv.Itoa(int(user.ID)))
	if err != nil {
		return "", code.CacheError
	}
	if check {
		token, err := cache.GetJWTByUserid(strconv.Itoa(int(user.ID)))
		if err != nil {
			return "", code.CacheError
		}
		return token, code.Success
	}

	token, eCode := auth.GetToken(int(user.ID))
	if eCode != code.Success {
		return "", eCode
	}
	//存入缓存
	cache.SetJWT(token)
	return token, code.Success
}

// GetUserInfoByUserid 通过用户id获得用户DTO对象
func GetUserInfoByUserid(userid string) (*dto.UserInfoDTO, code.Code) {
	user, err := repository.GetUserByUserid(userid)
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

// UserUpdate 更新用户信息
func UserUpdate(userDTO *dto.UserUpdateDTO, userIdStr string) code.Code {
	//处理传入的userId
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return code.ServerError
	}
	userId := uint(userIdInt)
	//取出要修改的用户信息
	user, err := repository.GetUserByUserid(userIdStr)
	if err != nil {
		return code.DatabaseError
	}
	//构建map进行更新
	updates := make(map[string]interface{})
	//筛选出要更新的内容并加入到map
	//-用户名
	if userDTO.Username != "" && userDTO.Username != user.Username {
		if !validator.UsernameCheck(userDTO.Username) {
			return code.InvalidUsername
		}
		if validator.UsernameExistCheck(userDTO.Username) {
			return code.UsernameAlreadyExist
		}
		updates["username"] = userDTO.Username
	}
	//-密码
	if userDTO.Password != "" {
		if !validator.PasswordCheck(userDTO.Password) {
			return code.PasswordWrong
		}

		result, eCode := crypto.Verify(user.EncryptedPassword, userDTO.Password)
		if eCode != code.Success {
			return eCode
		}
		if !result {
			encryptedPassword, eCode := crypto.Encrypt(userDTO.Password)
			if eCode != code.Success {
				return eCode
			}
			updates["encrypted_password"] = encryptedPassword
			//如果修改密码，需要重新登录
			err := cache.DeleteTokenByUserid(strconv.Itoa(int(user.ID)))
			if err != nil {
				return code.CacheError
			}
		}
	}
	//-个性签名
	if userDTO.Signature != user.Signature {
		if !validator.SignatureCheck(userDTO.Signature) {
			return code.OverlongSignature
		}
		updates["signature"] = userDTO.Signature
	}

	if len(updates) == 0 {
		return code.Success
	}
	//传入数据层操作
	err = repository.UpdateUser(updates, uint(userId))
	if err != nil {
		return code.DatabaseError
	}
	return code.Success
}
