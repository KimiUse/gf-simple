// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"gf-simple/api/system/v1"
)

type ISystemV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	GetPage(ctx context.Context, req *v1.GetPageReq) (res *v1.GetPageRes, err error)
	TxCreateUser(ctx context.Context, req *v1.TxCreateUserReq) (res *v1.TxCreateUserRes, err error)
	GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error)
}
