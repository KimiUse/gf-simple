// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package user

import (
	"context"
	"gf-simple/internal/model"
)

type (
	IUser interface {
		Create(ctx context.Context, in *model.UserCreateInput) (err error)
		GetList(ctx context.Context, in *model.UserListInput) (out []*model.UserListOutput, err error)
		Update(ctx context.Context, in *model.UserUpdateInput) (err error)
		GetPage(ctx context.Context, in *model.UserPageInput) (out []*model.UserListOutput, total int, err error)
		TxCreateUser(ctx context.Context, in *model.TxCreateUserInput) (err error)
		GetUserList(ctx context.Context, in *model.UserModelInput) (out []*model.UserModelOutput, err error)
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
