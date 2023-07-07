// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	pb_auth "AuthServer/api/user/v1"
	"context"
)

type (
	IUser interface {
		SignUp(ctx context.Context, req *pb_auth.SignUpReq) (res *pb_auth.SignUpResp, err error)
		SignIn(ctx context.Context, req *pb_auth.SignInReq) (res *pb_auth.SignInResp, err error)
		RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (res *pb_auth.RefreshTokenResp, err error)
		SignOut(ctx context.Context, req *pb_auth.SignOutReq) (res *pb_auth.SignOutResp, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
