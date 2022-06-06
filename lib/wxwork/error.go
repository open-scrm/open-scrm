package wxwork

import "fmt"

type response interface {
	ErrCode() int
	ErrMsg() string
}

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (r *Response) Error() string {
	return fmt.Sprintf("code=%v msg=%v", r.Errcode, r.Errmsg)
}

func (r Response) ErrCode() int {
	return r.Errcode
}

func (r Response) ErrMsg() string {
	return r.Errmsg
}
