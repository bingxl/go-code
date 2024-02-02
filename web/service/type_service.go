package service

// service 返回给 controller 层的数据
type ResT struct {
	Code   int    `json:"code"`           //状态码
	Status int    `json:"status"`         // 状态表示, 成功为1 失败为0
	Msg    string `json:"msg"`            // 提示消息
	Data   any    `json:"data,omitempty"` // 返回数据
}

func success(data any) ResT {
	return ResT{
		Code:   200,
		Status: 1,
		Msg:    "success",
		Data:   data,
	}
}

func fail(msg string) ResT {
	return ResT{
		Code:   401,
		Status: 0,
		Msg:    msg,
	}
}

func handle(data any, err error) ResT {
	if err != nil {
		return fail(err.Error())
	}
	return success(data)
}
