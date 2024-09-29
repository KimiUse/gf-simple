package model

import (
	"gf-simple/internal/model/system/entity"
	"gf-simple/utility/pagehelper"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type UserCreateInput struct {
	Id       int64
	Passport string
	Password string
	Nickname string
	Brief    string
}

type UserListInput struct {
}

type UserListOutput struct {
	Id       int64
	Passport string
	Nickname string
	Status   string
	Brief    string
	CreateAt *gtime.Time
	UpdateAt *gtime.Time
}

type UserUpdateInput struct {
	Id       int64
	Passport string
	Password string
	Nickname string
	Brief    string
}

type UserPageInput struct {
	Passport string
	Nickname string
	Page     pagehelper.PageReq
}

type TxCreateUserInput struct {
	Id       int64
	Passport string
	Password string
	Nickname string
	Brief    string
	Address  string
}

type UserModelInput struct {
}

type UserModelOutput struct {
	User       *entity.User
	UserDetail *entity.UserDetail
	UserScores []*entity.UserScores
}

type UserWithOutput struct {
	gmeta.Meta `orm:"table:user"`
	Id         int64
	Passport   string
	Password   string
	Nickname   string
	Brief      string
	UserDetail *entity.UserDetail   `orm:"with:uid=id"`
	UserScores []*entity.UserScores `orm:"with:uid=id"`
}
