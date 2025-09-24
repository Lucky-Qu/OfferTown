// Package validator user.go
//
// 功能:
// - 检查用户名合法性
// - 检查密码合法性
// - 检查用户名是否重名
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-24

package validator

import (
	"backend/internal/repository"
)

func UsernameCheck(username string) bool {
	//用户名长度不小于2，不大于20
	if len(username) < 2 || len(username) > 20 {
		return false
	}
	//用户名只允许包含0~9,a~z,A~Z,-,_,.
	for _, v := range username {
		if !((v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || v == '-' || v == '_' || v == '.') {
			return false
		}
	}
	return true
}

func PasswordCheck(password string) bool {
	//密码必须大于等于八位,小于等于二十位
	if !(len(password) >= 8 && len(password) <= 20) {
		return false
	}
	return true
}

func UsernameExistCheck(username string) bool {
	//用户名不能重名
	exist, err := repository.CheckUsername(username)
	if err != nil {
		//TODO 处理数据库错误
		return false
	}
	if exist {
		return true
	}
	return false
}
