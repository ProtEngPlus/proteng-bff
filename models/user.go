package models

type UserInput struct {
	Email     string   `bson:"email" json:"email"`
	Password  string   `bson:"password" json:"password"`
	Name      string   `bson:"name" json:"name"`
	Surname   string   `bson:"surname" json:"surname"`
	CitizenId string   `bson:"citizen_id" json:"citizen_id"`
	Role      []string `bson:"role" json:"role"`
}

type SignInInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Role     string `bson:"role" json:"role"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password string `json:"password" binding:"required"`
}
