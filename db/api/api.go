package api

import "github.com/gin-gonic/gin"

var (
	SuccessCode    = 200
	ErrorCode      = 100
	ErrorParameter = 422
	SuccessMsg     = "success"
	ErrorMsg       = "error"
	ParameterMsg   = "参数错误"
)

type Response struct {
	Code int    `json:"code" `
	Msg  string `json:"msg" `
	Data any    `json:"data" `
}
type Fn func(response *Response)

func Code(code int) Fn {
	return func(response *Response) {
		response.Code = code
	}
}
func Err(err error) Fn {
	return func(response *Response) {
		if err != nil {
			response.Msg = err.Error()
		}
	}
}
func Data(data any) Fn {
	return func(response *Response) {
		response.Data = data
	}
}

// Msg /*
func Msg(msg string) Fn {
	return func(response *Response) {
		response.Msg = msg
	}
}

// Success /*
func DataSuccess(ctx *gin.Context, fns ...Fn) {
	var success = &Response{
		Code: SuccessCode,
		Msg:  SuccessMsg,
		Data: nil,
	}
	if len(fns) > 0 {
		for _, fn := range fns {
			fn(success)
		}
	}
	ctx.JSON(200, success)
}
func DataWithErr(ctx *gin.Context, err error, data any, fns ...Fn) {
	if err != nil {
		DataError(ctx, err, fns...)
	} else {
		var fnArr []Fn
		if data != nil {
			fnArr = append(fnArr, Data(data))
		}
		fnArr = append(fnArr, fns...)
		DataSuccess(ctx, fnArr...)
	}
}

// Error /*
func DataError(ctx *gin.Context, err error, fns ...Fn) {
	var data = &Response{
		Code: ErrorCode,
		Msg:  ErrorMsg,
		Data: nil,
	}
	if err != nil {
		fns = append(fns, Err(err))
	}
	if len(fns) > 0 {
		for _, fn := range fns {
			fn(data)
		}
	}
	ctx.JSON(200, data)
}

// WeChatPayError /*
func WeChatPayError(ctx *gin.Context) {
	ctx.JSON(500, gin.H{
		"code":    "FAILED",
		"message": "支付失败",
	})
}
func WeChatPaySuccess(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    "SUCCESS",
		"message": "成功",
	})
}
