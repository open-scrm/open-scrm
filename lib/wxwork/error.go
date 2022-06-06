package wxwork

import "fmt"

type Response interface {
	ErrCode() int
	ErrMsg() string
}

type response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (r *response) Error() string {
	return fmt.Sprintf("code=%v msg=%v", r.Errcode, r.Errmsg)
}

func (r response) ErrCode() int {
	return r.Errcode
}

func (r response) ErrMsg() string {
	return r.Errmsg
}
