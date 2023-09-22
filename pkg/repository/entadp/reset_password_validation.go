package entadp

import (
	"context"
	"fmt"
	"library/ent"
	"library/pkg/repository/dto"
)

var _ ResetPasswordInterface = (*ResetPasswordRepository)(nil)

type ResetPasswordRepository struct {
	DBClient *ent.Client
}

func NewResetPassword(db *ent.Client) *ResetPasswordRepository {
	return &ResetPasswordRepository{
		DBClient: db,
	}
}

func (r *ResetPasswordRepository) Create(ctx context.Context, v *dto.ResetPasswordValidation) error {
	_, err := r.DBClient.ResetPasswordValidation.Create().SetEmail(v.Email).SetCode(v.Code).SetExpireDate(v.ExpireDate).Save(ctx)
	if err != nil {
		return fmt.Errorf("ent / reset password create :%w", err)
	}

	return nil
}
