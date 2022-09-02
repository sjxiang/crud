package requests


type UserSetupRequest struct {
	NickName string `json:"nickname" valid:"nickname"`
	Password string `json:"password" valid:"password"`  // TODO 敏感信息，json 转义忽略
	Email    string `json:"email"    valid:"email"`
	Phone    string `json:"phone"    valid:"phone"`
}


// 应该是用户名、手机号、邮箱（ 3 选 1 ）
type UserLoginRequest struct {
	Email 		string  `json:"email"    valid:"email"`
	Password    string  `json:"password" valid:"password"` 
}
