package models

type UserInterface interface {
	GetUserId(int) (Userinfo, error)
	GetUserName(string) (Userinfo, error)
	AddNewUser(Userinfo)
	DeleteUserById(Userinfo)
	ModifyUserById(Userinfo)
	QueryUserAll() ([]Userinfo, error)
}

type JsonByUser interface {
	EncoderUsersDb() error
	DecoderUsersDb() error
}

var NewUserInterface UserInterface = new(Userinfo)
var NewJsonByUser JsonByUser = new(Userinfo)
