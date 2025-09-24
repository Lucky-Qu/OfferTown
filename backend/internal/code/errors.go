// Package code errors.go
//
// 功能:
// - 封装错误
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-09-24

package code

// Code 定义代码类型，便于编写方法
type Code int

const (
	Success Code = 1001

	InvalidUsername      Code = 2001
	UsernameAlreadyExist Code = 2002
	InvalidPassword      Code = 2003

	DatabaseError Code = 3001

	BindFailed Code = 4001
)

// Msg 返回代码对应的信息
func (code Code) Msg() string {
	switch code {
	case Success:
		return "请求成功"
	case InvalidUsername:
		return "用户名不符合规范，请修改后重试！"
	case UsernameAlreadyExist:
		return "用户名已存在，请修改后重试"
	case InvalidPassword:
		return "密码不符合规范，请修改后重试！"
	case DatabaseError:
		return "数据库错误，请稍后重试"
	case BindFailed:
		return "绑定数据失败"
	}
	return "未知错误"
}
