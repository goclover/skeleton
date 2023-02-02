package render

import (
	"context"

	"github.com/goclover/clover/render"
)

type jsonData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type rError struct {
	code int
	msg  string
}

func (re rError) Error() string {
	return re.msg
}

var JSON = func(ctx context.Context, data interface{}, err error) (r render.Render) {
	r = render.JSON(jsonData{Code: 0, Msg: "success", Data: data})
	if err != nil {
		if e, ok := err.(rError); ok {
			r = render.JSON(jsonData{Code: e.code, Msg: e.msg, Data: data})
			return
		}
		r = render.JSON(jsonData{Code: 5000, Msg: err.Error(), Data: data})
	}
	return
}

var RError = func(msg string, codes ...int) rError {
	var code = 1
	if len(codes) > 0 {
		code = codes[0]
	}
	return rError{code: code, msg: msg}
}

var ParamsError = RError("请求异常，请重试", 4000)
var SystemError = RError("系统繁忙，请稍后再试", 4001)
