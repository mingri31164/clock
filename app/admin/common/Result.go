package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReJson 定义返回结构体
type ReJson struct {
	Code int `json:"code"` //返回代码
	Msg  any `json:"msg"`  //返回提示信息
	//Count int `json:"count"` //返回总数
	Data any `json:"data"` //返回数据
}

// ResOK 请求成功的返回体，传入请求成功的数据和总数
func ResOK(c *gin.Context, msg any, data any) {
	//将参数赋值给你结构体
	Json := ReJson{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, Json)
}

func Success(c *gin.Context) {
	//将参数赋值给你结构体
	Json := ReJson{
		Code: 200,
		Msg:  "请求成功！",
		Data: nil,
	}
	c.JSON(http.StatusOK, Json)
}

// ResErr 请求成功但是有错误的返回体，把错误提示信息传入就行
func ResErr(c *gin.Context, msg any) {
	//将参数赋值给你结构体
	Json := ReJson{
		Code: 400,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, Json)
}

func Error(c *gin.Context) {
	//将参数赋值给你结构体
	Json := ReJson{
		Code: 400,
		Msg:  "请求失败！",
		Data: nil,
	}
	c.JSON(http.StatusOK, Json)
}

// ResFail 请求失败的返回体（网络不通），只需要传入请求失败的信息回来就行了
func ResFail(c *gin.Context, msg any, data any) {
	//将参数赋值给你结构体
	Json := ReJson{
		Code: 400,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusNotFound, Json)
}

//type Result[T any] struct {
//	Code int    `json:"code"`
//	Msg  string `json:"msg"`
//	Data T      `json:"data"`
//}
//
//func Success[T any](data T) *Result[T] {
//	return &Result[T]{
//		Code: 1,
//		Data: data,
//	}
//}
//
//func Error(msg string) *Result[any] {
//	return &Result[any]{
//		Code: 0,
//		Msg:  msg,
//	}
//}
//
//func (r *Result[any]) String() interface{} {
//	data, err := json.Marshal(r)
//	if err != nil {
//		return ""
//	}
//	return string(data)
//}
