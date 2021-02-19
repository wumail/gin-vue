package dto

import (
	"main/model"
)

//UserDto 返回的信息
type UserDto struct {
	Name  string `json:"name"`
	Group string `json:"group"`
	// Telephone string `json:"telephone"`
}

//ToUserDto 转化model.User中的信息，只返回name 和 权限group
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Group: user.Group,
		// Telephone: user.Telephone,
	}
}
