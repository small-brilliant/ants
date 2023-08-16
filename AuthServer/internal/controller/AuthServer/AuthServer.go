package AuthServer

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	api.UnimplementedAuthServer
}

func Register(s *grpcx.GrpcServer) {
	api.RegisterAuthServer(s.Server, &Controller{})
}

func (*Controller) SignUp(ctx context.Context, req *api.SignUpReq) (res *api.SignUpResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SignIn(ctx context.Context, req *api.SignInReq) (res *api.SignInResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RefreshToken(ctx context.Context, req *api.RefreshTokenReq) (res *api.RefreshTokenResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SignOut(ctx context.Context, req *api.SignOutReq) (res *api.SignOutResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
