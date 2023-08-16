package controller

import (
	v1 "WebServer/api/v1"
	"WebServer/internal/service"
	"context"
)

var (
	Auth = cAuth{}
)

type cAuth struct{}

func (c *cAuth) SingIn(ctx context.Context, req *v1.SingInReq) (res *v1.SingInRes, err error) {
	return service.Auth().SingIn(ctx, req)
}

func (c *cAuth) SingUp(ctx context.Context, req *v1.SingUpReq) (res *v1.SingUpRes, err error) {
	return service.Auth().SingUp(ctx, req)
}
