package entities

type UserInfoAtomicService struct{}

var UserInfoService = UserInfoAtomicService{}

func (*UserInfoAtomicService) Save(userInfo *UserInfo) error {
	dao := userInfoDao{xormEngine}
	return dao.Save(userInfo)
}

func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{xormEngine}
	userInfolist, err := dao.FindAll()
	checkErr(err)
	return userInfolist
}

func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	dao := userInfoDao{xormEngine}
	userInfo, err := dao.FindByID(id)
	checkErr(err)
	return userInfo
}

func (*UserInfoAtomicService) Count() int64 {
	dao := userInfoDao{xormEngine}
	count, err := dao.CountAll()
	checkErr(err)
	return count
}
