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

func (ce rError) Error() string {
	return ce.msg
}

var JSON = func(ctx context.Context, data interface{}, err error) render.Render {
	if e, ok := err.(rError); ok {
		return render.JSON(jsonData{
			Code: e.code,
			Msg:  e.msg,
			Data: data,
		})
	}
	return render.JSON(jsonData{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

var RError = func(msg string, codes ...int) rError {
	var code = 0
	if len(codes) > 0 {
		code = codes[0]
	}
	return rError{
		code: code,
		msg:  msg,
	}
}
