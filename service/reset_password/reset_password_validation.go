package reset_password

import (
	"context"
	"fmt"
	"library/pkg/mail"
	"library/pkg/repository/dto"
	"library/pkg/repository/entadp"
	"library/pkg/utils"
	"time"
)

var _ ServiceResetPass = (*ResetPass)(nil)

type ResetPass struct {
	repo       entadp.RepositoryInterface
	mailSender mail.Mail
}

func New(r entadp.RepositoryInterface, m mail.Mail) *ResetPass {
	return &ResetPass{
		repo:       r,
		mailSender: m,
	}
}

func (r *ResetPass) SendResetPasswordEmail(ctx context.Context, email string) error {
	if _, err := r.repo.User().GetUserByEmail(ctx, email); err != nil {
		return fmt.Errorf("send reset password email :%w", err)
	}

	code := utils.GenerateRandomOTPCode()

	err := r.repo.ResetPass().Create(ctx, &dto.ResetPasswordValidation{
		Email:      email,
		ExpireDate: time.Now().Add(time.Minute * 5),
		Code:       code,
	})

	if err != nil {
		return fmt.Errorf("send reset password email :%w", err)
	}

	err = r.mailSender.SendResetPasswordEmail(ctx, email, code)
	if err != nil {
		return fmt.Errorf("send reset password email :%w", err)
	}
	return nil
}
