package v1

import (
	"gf-simple/utility/pagehelper"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateReq struct {
	g.Meta   `path:"/create" tags:"create" method:"post" summary:"create user"`
	Passport string `json:"passport" description:"账号"`
	Password string `json:"password" description:"密码"`
	Nickname string `json:"nickname" description:"昵称"`
	Brief    string `json:"brief"    description:"备注信息"`
}

type CreateRes struct {
	Id int64 `json:"id"       description:"用户ID"`
}

type GetListReq struct {
	g.Meta `path:"/list" tags:"list" method:"post" summary:"list user"`
}

type GetListRes struct {
	ItemList []*GetListItem `json:"list"`
}

type GetListItem struct {
	Id       int64       `json:"id"       description:"用户ID"`
	Passport string      `json:"passport" description:"账号"`
	Nickname string      `json:"nickname" description:"昵称"`
	Status   string      `json:"status"   description:"状态"`
	Brief    string      `json:"brief"    description:"备注信息"`
	CreateAt *gtime.Time `json:"createAt" description:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" description:"修改时间"`
}

type UpdateReq struct {
	g.Meta   `path:"/update" tags:"update" method:"post" summary:"create user"`
	Id       int64  `json:"id"       description:"用户ID"`
	Passport string `json:"passport" description:"账号"`
	Password string `json:"password" description:"密码"`
	Nickname string `json:"nickname" description:"昵称"`
	Brief    string `json:"brief"    description:"备注信息"`
}

type UpdateRes struct {
}

type GetPageReq struct {
	g.Meta   `path:"/page" tags:"page" method:"post" summary:"page user"`
	Passport string             `json:"passport" description:"账号"`
	Nickname string             `json:"nickname" description:"昵称"`
	Page     pagehelper.PageReq `json:"page" description:"分页参数"`
}

type GetPageRes struct {
	GetListRes
	pagehelper.PageRes
}

// TxCreateUserReq 带事务的保存
type TxCreateUserReq struct {
	g.Meta   `path:"/tx-create" tags:"tx create" method:"post" summary:"tx create user"`
	Passport string `json:"passport" description:"账号"`
	Password string `json:"password" description:"密码"`
	Nickname string `json:"nickname" description:"昵称"`
	Brief    string `json:"brief"    description:"备注信息"`
	Address  string `json:"address" ds:"地址"`
}
type TxCreateUserRes struct {
	Id int64 `json:"id"       description:"用户ID"`
}
type GetUserListReq struct {
	g.Meta `path:"/list-associate" tags:"list associate" method:"post" summary:"list associate user"`
}

type GetUserListRes struct {
	ItemList []*UserEntity `json:"list"`
}

type UserEntity struct {
	Id       int64  `json:"id"       description:"用户ID"`
	Passport string `json:"passport" description:"账号"`
	Password string `json:"password" description:"密码"`
	Nickname string `json:"nickname" description:"昵称"`
	Brief    string `json:"brief"    description:"备注信息"`
	Address  string `json:"address" ds:"地址"`
	Score    []uint `json:"score"  description:"分数"`
}
