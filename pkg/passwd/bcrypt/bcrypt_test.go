package bcrypt

import "testing"

func TestBcrypt_Generate(t *testing.T) {
	secret := "verysecretkey"
	bc := NewBcrypt(secret, 10)

	if _, err := bc.Generate("test1"); err != nil {
		t.Error(err)
	}
}

func TestBcrypt_Compare(t *testing.T) {
	secret := "verysecretkey"
	bc := NewBcrypt(secret, 10)

	var (
		hashed    string
		plainPass = "abc"
		err       error
	)
	if hashed, err = bc.Generate(plainPass); err != nil {
		t.Error(err)
	}

	t.Run("Compare", func(t *testing.T) {
		if err = bc.Compare(hashed, plainPass); err != nil {
			t.Error(err)
		}
	})

	t.Run("ErrorTesting", func(t *testing.T) {
		t.Run("Wrong_Password", func(t *testing.T) {
			if err = bc.Compare(hashed, "abcd"); err == nil {
				t.Error("expected: hashedPassword is not the hash of the given password")
			}
		})
	})

}
