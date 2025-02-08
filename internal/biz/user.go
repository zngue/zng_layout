package biz

import (
	"context"
)

type UserRepo interface {
	List(ctx context.Context, req *ListUserRequest) (total int32, list []*InfoUserItem, err error)
	Create(ctx context.Context, req *CreateUserRequest) (err error)
	Update(ctx context.Context, req *UpdateUserRequest) (err error)
	Delete(ctx context.Context, id int32) (err error)
	Info(ctx context.Context, id int32) (reply *InfoUserReply, err error)
	UpdateStatus(ctx context.Context, id int32, status int32) (err error)
}
type UpdateStatusUserReply struct {
	Id int32
}
type ListUserRequest struct {
	Page     int32
	PageSize int32
	Name     string
	Phone    string
}
type InfoUserItem struct {
	Id        int32
	Name      string
	Phone     string
	Status    int32
	CreatedAt int32
	UpdatedAt string
	Sex       int32
	Avatar    string
}
type CreateUserReply struct {
	Id int32
}
type UpdateUserRequest struct {
	Id     int32
	Name   string
	Phone  string
	Sex    int32
	Avatar string
}
type UpdateUserReply struct {
	Id int32
}
type InfoUserReply struct {
	Id     int32
	Name   string
	Phone  string
	Sex    int32
	Avatar string
}
type CreateUserRequest struct {
	Name   string
	Phone  string
	Sex    int32
	Avatar string
}
type UserUseCase struct {
	user UserRepo
}

func NewUserUseCase(user UserRepo) *UserUseCase {
	return &UserUseCase{
		user: user,
	}
}
func (u *UserUseCase) List(ctx context.Context, req *ListUserRequest) (total int32, list []*InfoUserItem, err error) {
	total, list, err = u.user.List(ctx, req)
	return
}
func (u *UserUseCase) Create(ctx context.Context, req *CreateUserRequest) (err error) {
	err = u.user.Create(ctx, req)
	return
}
func (u *UserUseCase) Update(ctx context.Context, req *UpdateUserRequest) (err error) {
	err = u.user.Update(ctx, req)
	return
}
func (u *UserUseCase) Delete(ctx context.Context, id int32) (err error) {
	err = u.user.Delete(ctx, id)
	return
}
func (u *UserUseCase) Info(ctx context.Context, id int32) (reply *InfoUserReply, err error) {
	reply, err = u.user.Info(ctx, id)
	return
}
func (u *UserUseCase) UpdateStatus(ctx context.Context, id int32, status int32) (err error) {
	err = u.user.UpdateStatus(ctx, id, status)
	return
}
