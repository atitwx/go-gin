package models

import "time"

type SysUser struct {
	UserId      int8      `gorm:"column:user_id"`
	UserName    string    `gorm:"column:user_name"`
	NickName    string    `gorm:"column:nick_name"`
	Email       string    `gorm:"column:email"`
	PhoneNumber string    `gorm:"column:phonenumber"`
	Sex         string    `gorm:"column:sex"`
	Avatar      string    `gorm:"column:avatar"`
	Password    string    `gorm:"column:_"`
	Status      string    `gorm:"column:status"`
	LoginIp     string    `gorm:"column:login_ip"`
	LoginDate   time.Time `gorm:"column:login_date"`
	CreateBy    string    `gorm:"column:create_by"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateBy    string    `gorm:"column:update_by"`
	UpdateTime  time.Time `gorm:"column:update_time"`
	Remark      string    `gorm:"column:remark"`
}

// TableName 指定数据库表名为 sys_user
func (SysUser) TableName() string {
	return "sys_user"
}
