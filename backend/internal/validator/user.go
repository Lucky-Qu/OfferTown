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
	"unicode"
)

func UsernameCheck(username string) bool {
	//用户名长度不小于2，不大于20(中文算作两个字符)
	if len(username) < 2 || len(username) > 20 {
		return false
	}
	//用户名只允许包含中文,字母,数字,-,_,.
	for _, v := range []rune(username) {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) && v != '_' && v != '-' && v != '.' {
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
	//密码只能包含0~9,a~z,A~Z,!~/
	for _, v := range password {
		if !(('0' <= v && v <= '9') || // 数字
			(v >= 'a' && v <= 'z') || // 小写字母
			(v >= 'A' && v <= 'Z') || // 大写字母
			(v >= '!' && v <= '~')) { // 特殊字符
			return false
		}
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
