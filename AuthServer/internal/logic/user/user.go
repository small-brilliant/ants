package user

import (
	pb_auth "AuthServer/api/user/v1"
	"AuthServer/internal/service"
	"context"
)

type sUser struct {
}

func New() *sUser {
	return &sUser{}
}
func int() {
	service.RegisterUser(New())
}

func (*sUser) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (res *pb_auth.SignUpResp, err error) {
	return &pb_auth.SignUpResp{Msg: "ni hao a"}, nil
}

func (*sUser) SignIn(ctx context.Context, req *pb_auth.SignInReq) (res *pb_auth.SignInResp, err error) {
	return nil, nil
}

func (*sUser) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (res *pb_auth.RefreshTokenResp, err error) {
	return nil, nil
}

func (*sUser) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (res *pb_auth.SignOutResp, err error) {
	return nil, nil
}
