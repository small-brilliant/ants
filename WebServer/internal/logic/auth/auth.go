package auth

import (
	pb_auth "AuthServer/api/user/v1"
	v1 "WebServer/api/v1"
	"WebServer/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
)

type sAuth struct {
}

func New() *sAuth {
	return &sAuth{}
}

func init() {
	service.RegisterAuth(New())
}

func (c *sAuth) SingIn(ctx context.Context, req *v1.SingInReq) (res *v1.SingInRes, err error) {
	sir, err := GetConn(ctx).SignIn(ctx, &pb_auth.SignInReq{})
	fmt.Print("replay", sir)
	return &v1.SingInRes{Msg: sir.Msg}, nil
}

func (c *sAuth) SingUp(ctx context.Context, req *v1.SingUpReq) (res *v1.SingUpRes, err error) {
	sir, err := GetConn(ctx).SignIn(ctx, &pb_auth.SignInReq{})
	fmt.Print("replay", sir)
	return &v1.SingUpRes{Msg: sir.Msg}, nil
}

func GetConn(ctx context.Context) pb_auth.AuthClient {
	name := g.Cfg().MustGet(ctx, "grpc.name").String()
	return pb_auth.NewAuthClient(grpcx.Client.MustNewGrpcClientConn(name))
}
