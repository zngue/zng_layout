package model

import (
	"context"
	"fmt"
	"github.com/zngue/zng_app/db/data"
	"github.com/zngue/zng_app/db/data/page"
	"github.com/zngue/zng_layout/internal/biz"
	"github.com/zngue/zng_layout/internal/model/db"
	"gorm.io/gorm"
)

func NewUserRepo(conn *gorm.DB) biz.UserRepo {
	return &UserRepo{
		conn: conn,
	}
}

type UserRepo struct {
	conn *gorm.DB
}

// List 请求方法 list
func (u *UserRepo) List(ctx context.Context, req *biz.ListUserRequest) (total int32, list []*biz.InfoUserItem, err error) {
	var conn = u.conn.WithContext(ctx)
	var dbConn = data.NewDB[db.User](conn)
	var (
		rs    []*db.User
		where = make(map[string]any)
	)
	if req.Name != "" {
		where["name like ?"] = "%" + req.Name + "%"
	}
	if req.Phone != "" {
		where["phone like ?"] = "%" + req.Phone + "%"
	}
	rs, err = dbConn.List(&data.ListRequest{
		Page: &page.Page{
			Page:     int(req.Page),
			PageSize: int(req.PageSize),
		},
		Where: where,
		Order: []string{"id desc"},
	})
	fmt.Println(rs)
	fmt.Println(dbConn)
	//TODO implement me
	panic("UserRepo->List implement me")
	return
}

// Create 请求方法 add
func (u *UserRepo) Create(ctx context.Context, req *biz.CreateUserRequest) (err error) {
	var conn = u.conn.WithContext(ctx)
	var dbConn = data.NewDB[db.User](conn)
	err = dbConn.Add(&db.User{
		Name:   req.Name,
		Phone:  req.Phone,
		Sex:    req.Sex,
		Avatar: req.Avatar,
	})
	fmt.Println(dbConn)
	//TODO implement me
	panic("UserRepo->Create implement me")
	return
}

// Update 请求方法 update
func (u *UserRepo) Update(ctx context.Context, req *biz.UpdateUserRequest) (err error) {
	var conn = u.conn.WithContext(ctx)
	var dbConn = data.NewDB[db.User](conn)
	var where = map[string]any{
		"id = ?": req.Id,
	}
	var updateData = map[string]any{
		"name":   req.Name,
		"phone":  req.Phone,
		"sex":    req.Sex,
		"avatar": req.Avatar,
	}
	err = dbConn.Update(where, updateData)
	fmt.Println(dbConn)
	//TODO implement me
	panic("UserRepo->Update implement me")
	return
}

// Delete 请求方法 delete
func (u *UserRepo) Delete(ctx context.Context, id int32) (err error) {
	var conn = u.conn.WithContext(ctx)
	var dbConn = data.NewDB[db.User](conn)
	var where = map[string]any{
		"id = ?": id,
	}
	err = dbConn.Delete(where)
	fmt.Println(dbConn)
	//TODO implement me
	panic("UserRepo->Delete implement me")
	return
}

// Info 请求方法 query
func (u *UserRepo) Info(ctx context.Context, id int32) (reply *biz.InfoUserReply, err error) {
	var conn = u.conn.WithContext(ctx)
	var dbConn = data.NewDB[db.User](conn)
	var where = map[string]any{
		"id = ?": id,
	}
	var rs *db.User
	rs, err = dbConn.Content(&data.ContentRequest{
		Where: where,
	})
	fmt.Println(rs)
	fmt.Println(dbConn)
	//TODO implement me
	panic("UserRepo->Info implement me")
	return
}

// UpdateStatus 请求方法 update
func (u *UserRepo) UpdateStatus(ctx context.Context, id int32, status int32) (err error) {
	var conn = u.conn.WithContext(ctx)
	var dbConn = data.NewDB[db.User](conn)
	var where = map[string]any{
		"id = ?": id,
	}
	var updateData = map[string]any{
		"status": status,
	}
	err = dbConn.Update(where, updateData)
	fmt.Println(dbConn)
	//TODO implement me
	panic("UserRepo->UpdateStatus implement me")
	return
}
