package models

import "fmt"

// AutoId 用户ID自增涨
func AutoId() int {
	id := 0
	for _, user := range UsersDb {
		if user.Id > id {
			id = user.Id
		}
	}
	return id + 1
}

// GetUserId 获取用户Id
func (u Userinfo) GetUserId(id int) (Userinfo, error) {
	if err := NewJsonByUser.DecoderUsersDb(); err != nil {
		return Userinfo{}, err
	}
	for _, user := range UsersDb {
		if user.Id == id {
			return user, nil
		}
	}
	return Userinfo{}, fmt.Errorf("用户ID不存在")
}

// GetUserName 获取用户名
func (u Userinfo) GetUserName(name string) (Userinfo, error) {
	if err := NewJsonByUser.DecoderUsersDb(); err != nil {
		return Userinfo{}, err
	}

	for _, user := range UsersDb {
		if user.Name == name {
			return user, nil
		}
	}
	return Userinfo{}, fmt.Errorf("用户名不存在")
}

// AddNewUser 添加用户
func (u Userinfo) AddNewUser(user Userinfo) {
	UsersDb = append(UsersDb, user)
	if err := NewJsonByUser.EncoderUsersDb(); err != nil {
		fmt.Println(err)
		return
	}
}

// DeleteUserById 删除用户信息
func (u Userinfo) DeleteUserById(user Userinfo) {
	tmpUsersDb := make([]Userinfo, 0)
	for _, tuser := range UsersDb {
		if tuser.Id != user.Id {
			tmpUsersDb = append(tmpUsersDb, tuser)
		}
	}
	UsersDb = tmpUsersDb
	if err := NewJsonByUser.EncoderUsersDb(); err != nil {
		fmt.Println(err)
		return
	}
}

// ModifyUserById 修改用户信息
func (u Userinfo) ModifyUserById(user Userinfo) {
	for i, tuser := range UsersDb {
		if tuser.Id == user.Id {
			UsersDb[i] = user
		}
	}
	if err := NewJsonByUser.EncoderUsersDb(); err != nil {
		fmt.Println(err)
		return
	}
}

func (u Userinfo) QueryUserAll() ([]Userinfo, error) {
	if err := NewJsonByUser.DecoderUsersDb(); err != nil {
		return []Userinfo{}, err
	}
	return UsersDb, nil
}
