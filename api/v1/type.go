package api

type userInfo struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=10"`
	Password string `form:"password" binding:"required,min=8,max=20"`
}

type userCreateResp struct {
	Username string `json:"username"`
}

type loginResp struct {
	JWT string
}
