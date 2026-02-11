package response

// 自定义一个类型，增加类型安全
type ResCode int

const (
	// 1. 成功状态
	CodeSuccess ResCode = 200

	// 2. 客户端错误 (4000xx)
	CodeInvalidParam ResCode = 400001 // 参数错误
	CodeUserExist    ResCode = 400002 // 用户已存在
	CodeUserNotExist ResCode = 400003 // 用户不存在
	CodeInvalidPass  ResCode = 400004 // 密码错误

	// 3. 服务端错误 (5000xx)
	CodeInternalError ResCode = 500001 // 服务器内部错误
	CodeDBError       ResCode = 500002 // 数据库操作失败
)

// Msg 映射：如果不想在 Controller 里写字符串，可以根据 Code 自动获取 Msg
var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "操作成功",
	CodeInvalidParam:  "请求参数有误",
	CodeUserExist:     "该用户名已被注册",
	CodeUserNotExist:  "用户不存在",
	CodeInvalidPass:   "密码错误",
	CodeInternalError: "服务器开小差了",
	CodeDBError:       "数据库操作失败",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return "未知错误"
	}
	return msg
}
