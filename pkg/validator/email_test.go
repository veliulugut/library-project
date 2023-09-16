package validator

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsEmailValid_Validate(t *testing.T) {
	tests := []struct {
		email       string
		expectedErr error
	}{
		//valid
		{
			email:       "abc",
			expectedErr: ErrInvalidEmail,
		},

		//invalid
		{
			email:       "veli@test.com",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("ValidEmail_%s", tt.email), func(t *testing.T) {
			vld := IsEmailValid{Email: tt.email}

			err := vld.Validate()

			if !errors.Is(err, tt.expectedErr) {
				t.Error(err)
			}
		})
	}

}
