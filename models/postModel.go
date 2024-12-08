package models

type PostRedis struct {
	Id string
	Title string
	Body string
	CreationTime string
	UpdatingTime string
}


type User struct {
	Email string
	Password string
	CreationTime string
	LastLoginTime string
	IsLoggedIn bool
}