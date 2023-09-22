package dto

import "time"

type ResetPasswordValidation struct {
	Email      string    `json:"email"`
	ExpireDate time.Time `json:"expire_date"`
	Code       string    `json:"code"`
}
