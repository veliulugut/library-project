package reset_password

type SendResetPasswordReq struct {
	Email string `json:"email"`
}

type ValidateReq struct {
	Email              string `json:"email"`
	Code               string `json:"code"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}
