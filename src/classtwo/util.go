package classtwo


import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * @Author: gaoz
 * @Date: 2020/9/28
 */

const SUCCESS  = 1000
const FAIL  = 999

func MD5Password (password string) string {
	md := md5.New()
	md.Write([]byte(password))

	return hex.EncodeToString(md.Sum(nil))
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Msg string
	Code int
	Data interface{}
}

func SuccessResponse(data interface{}, msg string) *Response {
	return &Response{
		Msg:  msg,
		Code: SUCCESS,
		Data: data,
	}
}

func FailResponse(msg string) *Response {
	return &Response{
		Msg:  msg,
		Code: FAIL,
	}
}