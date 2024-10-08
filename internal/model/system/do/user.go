// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	Id       interface{} // 用户ID
	Passport interface{} // 账号
	Password interface{} // 密码
	Nickname interface{} // 昵称
	Status   interface{} // 状态
	Brief    interface{} // 备注信息
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 修改时间
}
