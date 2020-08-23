package user

type user struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=10"`
	Password string `form:"password" binding:"required,min=8,max=20"`
}
