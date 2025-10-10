// Package code errors.go
//
// 功能:
// - 封装错误
//
// 作者: LuckyQu
// 创建日期: 2025-09-24
// 修改日期: 2025-10-09

package code

// Code 定义代码类型，便于编写方法
type Code int

const (
	Success Code = 1001

	InvalidUsername           Code = 2001
	UsernameAlreadyExists     Code = 2002
	InvalidPassword           Code = 2003
	UserNotExists             Code = 2004
	PasswordWrong             Code = 2005
	UnLoginUser               Code = 2006
	OverlongSignature         Code = 2007
	InvalidCategoryName       Code = 2008
	CategoryNameAlreadyExists Code = 2009
	CategoryNotExists         Code = 2010
	InvalidQuestion           Code = 2011
	QuestionNotExists         Code = 2012
	QuestionAlreadyExists     Code = 2013
	UnknownTarget             Code = 2014

	DatabaseError Code = 3001

	BindFailed Code = 4001

	EncryptError Code = 5001
	VerifyError  Code = 5002

	JWTSignFail  Code = 6001
	InvalidToken Code = 6002

	CacheError Code = 7001

	ServerError Code = 8001

	PermissionDenied Code = 9001
)

// Msg 返回代码对应的信息
func (code Code) Msg() string {
	switch code {
	case Success:
		return "请求成功"
	case InvalidUsername:
		return "用户名不符合规范，请修改后重试！"
	case UsernameAlreadyExists:
		return "用户名已存在，请修改后重试"
	case InvalidPassword:
		return "密码不符合规范，请修改后重试！"
	case DatabaseError:
		return "数据库错误，请稍后重试"
	case BindFailed:
		return "绑定数据失败"
	case EncryptError:
		return "加密错误"
	case VerifyError:
		return "验证错误"
	case JWTSignFail:
		return "Token签名出错"
	case InvalidToken:
		return "非法Token"
	case UserNotExists:
		return "用户不存在"
	case PasswordWrong:
		return "密码错误"
	case CacheError:
		return "缓存出错"
	case UnLoginUser:
		return "用户尚未登录"
	case ServerError:
		return "服务器出错，请稍后再试"
	case OverlongSignature:
		return "个性签名过长，请修改后重试"
	case PermissionDenied:
		return "权限不足"
	case InvalidCategoryName:
		return "分类名不符合规范，请修改后重试"
	case CategoryNameAlreadyExists:
		return "分类名已存在，请修改后重试"
	case CategoryNotExists:
		return "分类名不存在，请检查后再试"
	case InvalidQuestion:
		return "问题信息不合法，请修改后上传"
	case QuestionNotExists:
		return "题目不存在"
	case QuestionAlreadyExists:
		return "题目名已存在"
	case UnknownTarget:
		return "未知Target"
	}
	return "未知错误"
}
