package golang_jwt

import "testing"

func TestJWT_Generate(t *testing.T) {
	j := New("verysecretkey", 24)

	if _, err := j.Generate(1); err != nil {
		t.Error(err)
	}

}
