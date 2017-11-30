package entities

import "time"

// UserInfo
type UserInfo struct {
	UID        int `xorm:"pk autoincr"` // primary key and auto increase
	UserName   string
	DepartName string     `xorm:"default ''"`
	CreateAt   *time.Time `xorm:"created"` // auto create time
}
