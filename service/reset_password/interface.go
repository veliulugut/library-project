//go:generate mockgen -destination=service_mocks.go -package=reset_password library/service/reset_password ServiceResetPass
package reset_password

import (
	"context"
)

type ServiceResetPass interface {
	SendResetPasswordEmail(ctx context.Context, email string) error
}
