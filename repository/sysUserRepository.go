package repository

import (
	"go-gin/models"
)

// userRepository 实现 UserRepository 接口
type sysUserRepository struct {
}

var SysUserRepository = new(sysUserRepository)

// GetAllUsers 获取所有用户
func (s *sysUserRepository) GetAllUsers() ([]models.SysUser, error) {
	var users []models.SysUser
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
