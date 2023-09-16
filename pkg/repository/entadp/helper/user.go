package helper

import (
	"library/ent"
	"library/pkg/repository/dto"
)

func DBUserToDTO(d *ent.User) *dto.User {
	return &dto.User{
		ID:        uint8(d.ID),
		Email:     d.Email,
		Password:  d.Password,
		Username:  d.Username,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
