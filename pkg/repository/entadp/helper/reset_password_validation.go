package helper

import (
	"library/ent"
	"library/pkg/repository/dto"
)

func DBResetPasswordDTO(b *ent.ResetPasswordValidation) *dto.ResetPasswordValidation {
	return &dto.ResetPasswordValidation{
		Email:      b.Email,
		ExpireDate: b.ExpireDate,
		Code:       b.Code,
	}
}
