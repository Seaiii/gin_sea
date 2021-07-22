package model

type Admin struct {
	Id         int    `json:"id" gorm:"Column:id"`
	UserName   string `json:"username" gorm:"Column:username" form:"username"`
	PassWord   string `json:"password" gorm:"Column:password" form:"password"`
	HeadImgUrl string `json:"headimgurl" gorm:"Column:headimgurl"`
	Role       int    `json:"role"`
}

type AdminPassword struct {
	Admin
	NewPassword      string `json:"new_password" form:"new_password"`
	NewAgainPassword string `json:"new_again_password" form:"new_again_password"`
}
