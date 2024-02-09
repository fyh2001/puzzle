package models

import (
	"puzzle/utils"
	"time"
)

// User 用户模型
type User struct {
	Id        int64     `json:"id" gorm:"primaryKey"`            // 主键ID
	Username  string    `json:"username"`                        // 用户名
	Password  string    `json:"password"`                        // 密码
	Nickname  string    `json:"nickname"`                        // 昵称
	Email     string    `json:"email" gorm:"default NULL"`       // 邮箱
	Phone     string    `json:"phone" gorm:"default NULL"`       // 手机号
	Avatar    string    `json:"avatar" gorm:"default NULL"`      // 头像
	Status    int       `json:"status" gorm:"default 0"`         // 状态 0:启用 1:冻结 2:删除
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

// UserReq 用户请求模型
type UserReq struct {
	Id         int64            `json:"id"`       // 主键ID
	Username   string           `json:"username"` // 用户名
	Password   string           `json:"password"` // 密码
	Nickname   string           `json:"nickname"` // 昵称
	Email      string           `json:"email"`    // 邮箱
	Phone      string           `json:"phone"`    // 手机号
	Status     int              `json:"status"`   // 状态 0:启用 1:冻结 2:删除
	Pagination utils.Pagination `gorm:"embedded"` // 分页
	Ids        []int64          `json:"ids"`      // 主键ID集合
}

// UserResp 用户响应模型
type UserResp struct {
	Id        string    `json:"id"`        // 主键ID
	Username  string    `json:"username"`  // 用户名
	Nickname  string    `json:"nickname"`  // 昵称
	Email     string    `json:"email"`     // 邮箱
	Phone     string    `json:"phone"`     // 手机号
	Avatar    string    `json:"avatar"`    // 头像
	Status    int       `json:"status"`    // 状态 0:启用 1:冻结 2:删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}

// UserListResp 用户列表响应模型
type UserListResp struct {
	Total   int64      `json:"total"`   // 总数
	Records []UserResp `json:"records"` // 用户列表
}

// UserRegisterReq 用户注册请求模型
type UserRegisterReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
}

// UserLoginReq 用户登录请求模型
type UserLoginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// UserLoginResp 用户登录响应模型
type UserLoginResp struct {
	Token string   `json:"token"` // token
	User  UserResp `json:"user"`  // 用户信息
}
