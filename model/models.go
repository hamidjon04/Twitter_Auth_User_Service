package model

type RegisterReq struct {
	Email    string `json:"emil"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RegisterResp struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type LogoutReq struct {
	Id string `json:"id"`
}

type LogoutResp struct {
	Message string `json:"message"`
}

type ResetPassReq struct {
	Id string `json:"id"`
	Password string `json:"password"`
}

type ResetPassResp struct {
	Message string `json:"message"`
}

type ChangePassReq struct {
	UserId      string `json:"user_id"`
	NowPassword string `json:"now_password"`
	NewPassword string `json:"new_password"`
}

type UserInfo struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SaveToken struct {
	UserId       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type ChangePassResp struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Error struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type ForgotPassReq struct{
	Email string `json:"email"`
}

type ForgotPassResp struct{
	Message string `json:"message"`
}
