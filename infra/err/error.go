package err

import "fmt"

type GimError struct {
	Code   int32
	Msg    string
	Detail string //错误详情
}

func (m *GimError) Error() string {
	return fmt.Sprintf("err code:%d msg:%s detail:%s", m.Code, m.Msg, m.Detail)
}

func (m *GimError) SetDetail(detail string) *GimError {
	return &GimError{Code: m.Code, Msg: m.Msg, Detail: detail}
}

var (
	UnknownError = &GimError{Code: 1001, Msg: "未知错误"}
	//错误码定义规范 1000-1999 系统性错误 比如数据库连接异常 网络错误等 2000-2999 业务性错误，比如用户已登录 用户未登录等
	UserNotLoginError = &GimError{Code: 2001, Msg: "用户未登录"}
)
