package user

import (
	"context"
	dao "gf-simple/internal/dao/system"
	"gf-simple/internal/model"
	"gf-simple/internal/model/system/entity"
	"gf-simple/internal/service/user"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sUser struct {
}

func init() {
	user.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Create(ctx context.Context, in *model.UserCreateInput) (err error) {
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}

	id, err := dao.User.Ctx(ctx).InsertAndGetId(entity.User{
		Passport: in.Passport,
		Password: in.Password,
		Nickname: in.Nickname,
		Brief:    in.Brief,
		CreateAt: gtime.Now(),
		UpdateAt: gtime.Now(),
	})
	in.Id = id
	return
}

func (s *sUser) GetList(ctx context.Context, in *model.UserListInput) (out []*model.UserListOutput, err error) {
	var list []*entity.User
	err = dao.User.Ctx(ctx).Scan(&list)
	out = make([]*model.UserListOutput, len(list))
	for k, v := range list {
		out[k] = &model.UserListOutput{
			Id:       v.Id,
			Nickname: v.Nickname,
			Passport: v.Passport,
			Brief:    v.Brief,
			UpdateAt: v.UpdateAt,
			CreateAt: v.CreateAt,
		}
	}
	return
}

func (s *sUser) Update(ctx context.Context, in *model.UserUpdateInput) (err error) {
	userColumns := dao.User.Columns()
	_, err = dao.User.Ctx(ctx).Where(userColumns.Id, in.Id).Data(g.Map{
		userColumns.Passport: in.Passport,
		userColumns.UpdateAt: gtime.Now(),
		userColumns.Brief:    in.Brief,
		userColumns.Nickname: in.Nickname,
	}).Update()
	return
}

func (s *sUser) GetPage(ctx context.Context, in *model.UserPageInput) (out []*model.UserListOutput, total int, err error) {
	var users []*entity.User
	userDao := dao.User
	err = userDao.Ctx(ctx).
		Where(g.Map{
			userDao.Columns().Nickname: in.Nickname,
			userDao.Columns().Passport: in.Passport,
		}).
		// 空值过滤
		OmitEmpty().
		OrderDesc(userDao.Columns().CreateAt).
		Limit(in.Page.GetPageNum(), in.Page.GetPageSize()).
		ScanAndCount(&users, &total, true)

	out = make([]*model.UserListOutput, len(users))

	for k, v := range users {
		out[k] = &model.UserListOutput{
			Id:       v.Id,
			Nickname: v.Nickname,
			Passport: v.Passport,
			Brief:    v.Brief,
			UpdateAt: v.UpdateAt,
			CreateAt: v.CreateAt,
		}
	}
	return
}

func (s *sUser) TxCreateUser(ctx context.Context, in *model.TxCreateUserInput) (err error) {
	err = dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// save user
		id, err := dao.User.Ctx(ctx).InsertAndGetId(entity.User{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
			Brief:    in.Brief,
			CreateAt: gtime.Now(),
			UpdateAt: gtime.Now(),
		})
		in.Id = id
		if err != nil {
			return err
		}
		// save detail
		_, err = dao.UserDetail.Ctx(ctx).Insert(entity.UserDetail{
			Uid:     id,
			Address: in.Address,
		})
		return err
	})

	return err
}

func (s *sUser) GetUserList(ctx context.Context, in *model.UserModelInput) (out []*model.UserModelOutput, err error) {
	s.scanList(ctx, in)
	s.with(ctx)
	return
}

func (s *sUser) scanList(ctx context.Context, in *model.UserModelInput) (out []*model.UserModelOutput, err error) {
	g.Log().Debug(ctx, "Debug-ScanList-Start")
	var user []model.UserModelOutput
	err = dao.User.Ctx(ctx).ScanList(&user, "User")

	err = dao.UserDetail.Ctx(ctx).Where(dao.UserDetail.Columns().Uid, gdb.ListItemValues(user, "User", "Id")).
		ScanList(&user, "UserDetail", "User", dao.UserDetail.Columns().Uid+":"+"Id")
	g.Log().Debug(ctx, "Debug-ScanList-End")

	return
}

func (s *sUser) with(ctx context.Context) (out []*model.UserWithOutput, err error) {
	g.Log().Debug(ctx, "Debug-With-Start")

	var user []*model.UserWithOutput

	dao.User.Ctx(ctx).WithAll().Scan(&user)
	g.Log().Debug(ctx, "Debug-With-End")
	return
}
