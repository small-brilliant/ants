package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedAuthServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAuthServer(s.Server, &Controller{})
}

func (*Controller) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
