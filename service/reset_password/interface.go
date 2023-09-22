package reset_password

import (
	"context"
)

type ServiceResetPass interface {
	SendResetPasswordEmail(ctx context.Context, email string) error
}
