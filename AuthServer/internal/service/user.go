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
