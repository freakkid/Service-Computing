package entities

import "time"

// UserInfo
type UserInfo struct {
	UID        int `xorm:"pk autoincr"` // primary key and auto increase
	UserName   string
	DepartName string     `xorm:"default ''"`
	CreateAt   *time.Time `xorm:"created"` // auto create time
}

// User is an entity to save user info with username is primary key.
type User struct {
	UserName string `xorm:"pk varchar(255) notnull unique"`
	Password string `xorm:"varchar(255) notnull"`
	Email    string `xorm:"varchar(255) notnull"`
	Phone    string `xorm:"varchar(255) notnull"`
}
type Meeting struct {
	Title       string     `xorm:"pk varchar(255) notnull unique"`
	Sponsor     string     `xorm:"varchar(255) notnull"`
	Particulars string     `xorm:"text notnull"`
	StartTime   *time.Time `xorm:"DateTime"`
	EndTime     *time.Time `xorm:"DateTime"`
}
