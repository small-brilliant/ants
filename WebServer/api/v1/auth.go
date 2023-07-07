package v1

import "github.com/gogf/gf/v2/frame/g"

type SingInReq struct {
	g.Meta `path:"/sing_in" method:"post"`
}
type SingInRes struct {
	Msg string `json:"msg"`
}
