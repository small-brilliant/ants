package user

import (
	pb_auth "AuthServer/api/user/v1"
	"AuthServer/internal/service"
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	pb_auth.UnimplementedAuthServer
}

func Register(s *grpcx.GrpcServer) {
	pb_auth.RegisterAuthServer(s.Server, &Controller{})

}

func (*Controller) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (res *pb_auth.SignUpResp, err error) {
	return service.User().SignUp(ctx, req)
}

func (*Controller) SignIn(ctx context.Context, req *pb_auth.SignInReq) (res *pb_auth.SignInResp, err error) {
	return service.User().SignIn(ctx, req)
}

func (*Controller) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (res *pb_auth.RefreshTokenResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (res *pb_auth.SignOutResp, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
