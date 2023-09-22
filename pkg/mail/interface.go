package mail

import "context"

type Mail interface {
	SendResetPasswordEmail(ctx context.Context, email, otp string) error
}
