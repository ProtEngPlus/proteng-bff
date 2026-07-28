package models

type UserInput struct {
	Email    string   `bson:"email" json:"email"`
	Password string   `bson:"password" json:"password"`
	Name     string   `bson:"name" json:"name"`
	Surname  string   `bson:"surname" json:"surname"`
	Role     []string `bson:"role" json:"role"`
	UserRole string   `bson:"user_role" json:"user_role"`
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

type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}
type SendVerificationInput struct {
	Email string `json:"email" binding:"required"`
}
