package entities

import "github.com/go-xorm/xorm"

type userInfoDao struct {
	*xorm.Engine
}

// save new UserInfo into database
func (dao *userInfoDao) Save(userInfo *UserInfo) error {
	_, err := dao.Insert(userInfo)
	return err
}

func (dao *userInfoDao) FindAll() ([]UserInfo, error) {
	userInfolist := make([]UserInfo, 0, 0)
	err := dao.Find(&userInfolist)
	return userInfolist, err
}

func (dao *userInfoDao) FindByID(id int) (*UserInfo, error) {
	var userInfo = &UserInfo{UID: id}
	_, err := dao.Get(userInfo)
	return userInfo, err
}

func (dao *userInfoDao) CountAll() (int64, error) {
	return dao.Count(new(UserInfo))
}

