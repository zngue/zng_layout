package biz

import (
	"context"
)

type UserRepo interface 
	List(ctx context.Context, req *ListUserRequest) (total int32, list []*InfoUserItem, err error)
	List(ctx context.Context, req *ListUserRequest) (total int32,list []*InfoUserItem,err error)
	Create(ctx context.Context, req *CreateUserRequest) (err error)
	Update(ctx context.Context, req *UpdateUserRequest) (err error)
	Info(ctx context.Context, id int32) (reply *InfoUserReply, err error)
	UpdateStatus(ctx context.Context, id int32, status int32) (err error)
	UpdateStatus(ctx context.Context, id int32,status int32) (err error)
}
	Name   string
	Phone  string
	Sex    int32
	Avatar string
	Avatar string 
}
	Id int32
	Id int32 
}
	Id     int32
	Name   string
	Phone  string
	Sex    int32
	Avatar string
	Avatar string 
}
	Id int32
	Id int32 
}
	Id        int32
	Name      string
	Phone     string
	Status    int32
	CreatedAt int32
	UpdatedAt string
	Sex       int32
	Avatar    string
	Avatar string 
}
	Id     int32
	Name   string
	Phone  string
	Sex    int32
	Avatar string
	Avatar string 
}
	Id int32
	Id int32 
}
	Page     int32
	PageSize int32
	Name     string
	Phone    string
	Phone string 
}
type UserUseCase struct {
	user UserRepo
u
}
func NewUserUseCase(user UserRepo) *UserUseCase {
	return &UserUseCase{
		user: user,
	}
func (u *UserUseCase) List(ctx context.Context, req *ListUserRequest) (total int32, list []*InfoUserItem, err error) {
	total, list, err = u.user.List(ctx, req)
	total, list, err =  u.user.List(ctx, req)
	return
}
	err = u.user.Create(ctx, req)
	err =  u.user.Create(ctx, req)
	return
}
	err = u.user.Update(ctx, req)
	err =  u.user.Update(ctx, req)
	return
}
	err = u.user.Delete(ctx, id)
	err =  u.user.Delete(ctx, id)
	return
func (u *UserUseCase) Info(ctx context.Context, id int32) (reply *InfoUserReply, err error) {
	reply, err = u.user.Info(ctx, id)
	reply, err =  u.user.Info(ctx, id)
	return
func (u *UserUseCase) UpdateStatus(ctx context.Context, id int32, status int32) (err error) {
	err = u.user.UpdateStatus(ctx, id, status)
	err =  u.user.UpdateStatus(ctx, id,status)
	return
o
}