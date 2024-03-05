package models

type AdminInput struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Username string `bson:"username" json:"username"`
}

type SignInAdminInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
