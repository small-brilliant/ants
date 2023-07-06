package user

import (
	pb_auth "AuthServer/api/user/v1"
	"context"
)

type sUser struct {
}

func (*sUser) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (res *pb_auth.SignUpResp, err error) {
	return nil, nil
}
