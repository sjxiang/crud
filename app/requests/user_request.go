package requests


type UserSetupRequest struct {

}


// 应该是用户名、手机号、邮箱（ 3 选 1 ）
type UserLoginRequest struct {
	Email 		string  `json:"email"`
	Password    string  `json:"password"`
}