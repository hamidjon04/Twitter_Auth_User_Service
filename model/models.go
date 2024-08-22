package model

type RegisterReq struct {
	Email    string `json:"emil"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogenResp struct {
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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassResp struct {
	Message string `json:"message"`
}

type ChangePassReq struct {
	NowPassword string `json:"now_password"`
	NewPassword string `json:"new_password"`
}

type ChangePassResp struct {
	Message string `json:"message"`
}
type SuccessResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
