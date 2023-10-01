package reset_password

import (
	"context"
	"fmt"
	"library/pkg/mail"
	"library/pkg/passwd"
	"library/pkg/repository/dto"
	"library/pkg/repository/entadp"
	"library/pkg/utils"
	"time"
)

var _ ServiceResetPass = (*ResetPass)(nil)

func New(r entadp.RepositoryInterface, m mail.Mail, e passwd.Interface) *ResetPass {
	return &ResetPass{
		repo:       r,
		mailSender: m,
		enc:        e,
	}
}

type ResetPass struct {
	repo       entadp.RepositoryInterface
	mailSender mail.Mail
	enc        passwd.Interface
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

func (r *ResetPass) Validate(ctx context.Context, email, code, newPassword, confirmNewPassword string) error {
	rpv, err := r.repo.ResetPass().Get(ctx, email, code)
	if err != nil {
		return fmt.Errorf("validate :%w", err)
	}

	if rpv.ExpireDate.Before(time.Now()) {
		return fmt.Errorf("validation :%w", ErrExpiredValidationCode)
	}

	if newPassword != confirmNewPassword {
		return fmt.Errorf("validation :%w", ErrPasswordNotMatch)
	}

	hashed, err := r.enc.Generate(newPassword)
	if err != nil {
		return fmt.Errorf("validate :%w", err)
	}

	err = r.repo.User().UpdatePassword(ctx, email, hashed)
	if err != nil {
		return fmt.Errorf("validate :%w", err)
	}

	err = r.repo.ResetPass().Delete(ctx, email)
	if err != nil {
		return fmt.Errorf("validate :%w", err)
	}

	return nil
}
