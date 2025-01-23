package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/zngue/zng_layout/api/user/v1"
)

type UserService struct {
}

func (u *UserService) List(ctx *gin.Context, req *v1.ListUserRequest) (rs *v1.ListUserReply, err error) {
	rs = &v1.ListUserReply{
		Total: 2000,
		List:  nil,
	}

	return
}

func (u *UserService) Create(ctx *gin.Context, req *v1.CreateUserRequest) (rs *v1.CreateUserReply, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Update(ctx *gin.Context, req *v1.UpdateUserRequest) (rs *v1.UpdateUserReply, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Delete(ctx *gin.Context, req *v1.DeleteUserRequest) (rs *empty.Empty, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Info(ctx *gin.Context, req *v1.InfoUserRequest) (rs *v1.InfoUserReply, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateStatus(ctx *gin.Context, req *v1.UpdateStatusUserRequest) (rs *v1.UpdateStatusUserReply, err error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService() v1.UserGinHttpService {
	return &UserService{}
}
