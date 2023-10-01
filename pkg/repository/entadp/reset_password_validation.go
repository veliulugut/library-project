package entadp

import (
	"context"
	"fmt"
	"library/ent"
	"library/ent/resetpasswordvalidation"
	"library/pkg/repository/dto"
	"library/pkg/repository/entadp/helper"
)

var _ ResetPasswordInterface = (*ResetPasswordRepository)(nil)

func NewResetPassword(db *ent.Client) *ResetPasswordRepository {
	return &ResetPasswordRepository{
		DBClient: db,
	}
}

type ResetPasswordRepository struct {
	DBClient *ent.Client
}

func (r *ResetPasswordRepository) Create(ctx context.Context, v *dto.ResetPasswordValidation) error {
	_, err := r.DBClient.ResetPasswordValidation.
		Create().
		SetEmail(v.Email).
		SetCode(v.Code).
		SetExpireDate(v.ExpireDate).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("ent / reset password create :%w", err)
	}

	return nil
}

func (r *ResetPasswordRepository) Get(ctx context.Context, email string, code string) (*dto.ResetPasswordValidation, error) {
	rpv, err := r.DBClient.ResetPasswordValidation.Query().
		Where(
			resetpasswordvalidation.Email(email),
			resetpasswordvalidation.Code(code),
		).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("ent / get  :%w", ErrNotFound)
		}
		return nil, fmt.Errorf("ent / get :%w", err)
	}

	return helper.DBResetPasswordDTO(rpv), nil
}

func (r *ResetPasswordRepository) Delete(ctx context.Context, email string) error {
	_, err := r.DBClient.ResetPasswordValidation.Delete().Where(resetpasswordvalidation.Email(email)).Exec(ctx)
	if err != nil {
		return fmt.Errorf("ent / delete :%w", err)
	}

	return nil
}
