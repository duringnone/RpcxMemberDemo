package model

type User struct {
	UID      int64  `json:"id"`
	AddTime  int64  `json:"addtime"`
	UserType int32  `json:"utype"`
	Uface    string `json:"uface"`
	UserName string `json:"uname"`
}