package api

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/zngue/zng_layout/api/user/v1"
	"github.com/zngue/zng_layout/internal/biz"
)

import (
	"context"
)

type UserService struct {
	user *biz.UserUseCase
}

func NewUserService(user *biz.UserUseCase) v1.UserGinHttpService {
	return &UserService{
		user: user,
	}
}
func (s *UserService) List(ctx context.Context, req *v1.ListUserRequest) (rs *v1.ListUserReply, err error) {
	var reqData = &biz.ListUserRequest{}
	total, list, err := s.user.List(ctx, reqData)
	if err != nil {
		return
	}
	fmt.Println("UserService->List", total, list, err)
	return
}
func (s *UserService) Create(ctx context.Context, req *v1.CreateUserRequest) (rs *empty.Empty, err error) {
	var reqData = &biz.CreateUserRequest{}
	err = s.user.Create(ctx, reqData)
	if err != nil {
		return
	}
	fmt.Println("UserService->Create", err)
	return
}
func (s *UserService) Update(ctx context.Context, req *v1.UpdateUserRequest) (rs *empty.Empty, err error) {
	var reqData = &biz.UpdateUserRequest{}
	err = s.user.Update(ctx, reqData)
	if err != nil {
		return
	}
	fmt.Println("UserService->Update", err)
	return
}
func (s *UserService) Delete(ctx context.Context, req *v1.DeleteUserRequest) (rs *empty.Empty, err error) {
	err = s.user.Delete(ctx, req.Id)
	if err != nil {
		return
	}
	fmt.Println("UserService->Delete", err)
	return
}
func (s *UserService) Info(ctx context.Context, req *v1.InfoUserRequest) (rs *v1.InfoUserReply, err error) {
	reply, err := s.user.Info(ctx, req.Id)
	if err != nil {
		return
	}
	fmt.Println("UserService->Info", reply, err)
	return
}
func (s *UserService) UpdateStatus(ctx context.Context, req *v1.UpdateStatusUserRequest) (rs *empty.Empty, err error) {
	err = s.user.UpdateStatus(ctx, req.Id, req.Status)
	if err != nil {
		return
	}
	fmt.Println("UserService->UpdateStatus", err)
	return
}
