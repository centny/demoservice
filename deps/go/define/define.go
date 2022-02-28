package define

import "fmt"

type Codable interface {
	Code() int
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Debug   string `json:"debug"`
}

/***** metadata:ReturnCode *****/

const (
	//Success 正常返回
	Success = 0
	//ArgsInvalid 接口参数错误，可能是少了必须参数或参数值非法
	ArgsInvalid = 1100
	//CodeInvalid 验证码错误
	CodeInvalid = 1200
	//UserInvalid 用户异常
	UserInvalid = 1300
	//SignInvalid 签名异常
	SignInvalid = 1400
	//Duplicate 数据重复
	Duplicate = 1500
	//Frequently 调用过于频繁
	Frequently = 1600
	//ServerError 服务器异常
	ServerError = 2000
	//Redirect 接口需要登录但未登录
	Redirect = 301 //redirect code
	//NotAccess 接口没有权限访问
	NotAccess = 401 //not access
	//NotFound 未找到所需数据
	NotFound = 404
)

type Error struct {
	Code_  int    `json:"code"`
	Error_ string `json:"error"`
	Inner  error  `json:"inner"`
}

func NewError(code int, message string, inner error) (err *Error) {
	err = &Error{Code_: code, Error_: message, Inner: inner}
	return
}

func (e *Error) Error() string {
	if e.Inner == nil {
		return fmt.Sprintf("%v", e.Error_)
	}
	return fmt.Sprintf("%v:%v", e.Error_, e.Inner)
}

func (e *Error) Code() int { return e.Code_ }

func (e *Error) String() string { return e.Error() }

var ErrArgsInvalid = &Error{Code_: ArgsInvalid, Error_: "ArgsInvalid"}
var ErrCodeInvalid = &Error{Code_: CodeInvalid, Error_: "CodeInvalid"}
var ErrUserInvalid = &Error{Code_: UserInvalid, Error_: "UserInvalid"}
var ErrSignInvalid = &Error{Code_: SignInvalid, Error_: "SignInvalid"}
var ErrDuplicate = &Error{Code_: Duplicate, Error_: "Duplicate"}
var ErrFrequently = &Error{Code_: Frequently, Error_: "Frequently"}
var ErrServerError = &Error{Code_: ServerError, Error_: "ServerError"}
var ErrNotAccess = &Error{Code_: NotAccess, Error_: "NotAccess"}
var ErrNotFound = &Error{Code_: NotFound, Error_: "NotFound"}

func IsCodeError(err error, code int) bool {
	codable, ok := err.(Codable)
	return ok && codable.Code() == code
}
