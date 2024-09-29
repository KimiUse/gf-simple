package system

import (
	"context"
	"gf-simple/api/system/v1"
	"gf-simple/internal/model"
	service "gf-simple/internal/service/user"
	"gf-simple/utility/pagehelper"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	var in = model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
		Brief:    req.Brief,
	}
	err = service.User().Create(ctx, &in)
	res = new(v1.CreateRes)
	res.Id = in.Id
	return
}
func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	out, err := service.User().GetList(ctx, nil)
	res = &v1.GetListRes{
		ItemList: make([]*v1.GetListItem, len(out)),
	}
	for k, v := range out {
		res.ItemList[k] = &v1.GetListItem{
			Id:       v.Id,
			Passport: v.Passport,
			Nickname: v.Nickname,
			Brief:    v.Brief,
			CreateAt: v.CreateAt,
			UpdateAt: v.UpdateAt,
		}
	}
	return
}
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, service.User().Update(ctx, &model.UserUpdateInput{
		Id:       req.Id,
		Passport: req.Passport,
		Nickname: req.Nickname,
		Brief:    req.Brief,
	})
}
func (c *ControllerV1) GetPage(ctx context.Context, req *v1.GetPageReq) (res *v1.GetPageRes, err error) {
	list, total, err := service.User().GetPage(ctx, &model.UserPageInput{
		Passport: req.Passport,
		Nickname: req.Nickname,
		Page:     req.Page,
	})
	res = &v1.GetPageRes{
		GetListRes: v1.GetListRes{
			ItemList: make([]*v1.GetListItem, len(list)),
		},
		PageRes: pagehelper.PageRes{
			Total: total,
		},
	}
	for k, v := range list {
		res.ItemList[k] = &v1.GetListItem{
			Id:       v.Id,
			Passport: v.Passport,
			Nickname: v.Nickname,
			Brief:    v.Brief,
			CreateAt: v.CreateAt,
			UpdateAt: v.UpdateAt,
		}
	}
	return
}
func (c *ControllerV1) TxCreateUser(ctx context.Context, req *v1.TxCreateUserReq) (res *v1.TxCreateUserRes, err error) {
	var in = &model.TxCreateUserInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
		Brief:    req.Brief,
		Address:  req.Address,
	}
	err = service.User().TxCreateUser(ctx, in)
	res = &v1.TxCreateUserRes{
		Id: in.Id,
	}
	return
}
func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {

	service.User().GetUserList(ctx, &model.UserModelInput{})

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
