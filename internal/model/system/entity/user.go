// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int64       `json:"id"       description:"用户ID"`
	Passport string      `json:"passport" description:"账号"`
	Password string      `json:"password" description:"密码"`
	Nickname string      `json:"nickname" description:"昵称"`
	Status   string      `json:"status"   description:"状态"`
	Brief    string      `json:"brief"    description:"备注信息"`
	CreateAt *gtime.Time `json:"createAt" description:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" description:"修改时间"`
}
