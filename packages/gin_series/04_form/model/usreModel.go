package model

type UserModel struct {
	Email string `form:"email" binding:"email"`
	Password string `form:"password"`
	PasswordAgain string `form:"password" binding:"eqfield=Password"`
}
