package auth

type User struct {
	Id       string `bson:"_id"`
	UserName string `bson:"userName"`
	PassWord string `bson:"password"`
}
