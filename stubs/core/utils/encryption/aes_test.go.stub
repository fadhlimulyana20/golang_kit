package encryption

import "testing"

func TestEncrypt(t *testing.T) {
	aesEncrypt := NewAESEncrypt("ABC5dasar1231412")
	c, err := aesEncrypt.Encrypt("halo dunia")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("Encrypted data: %s", c)
}

func TestDecrypt(t *testing.T) {
	aesEncrypt := NewAESEncrypt("ABC5dasar1231412")
	c, err := aesEncrypt.Encrypt("halo dunia")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("Encrypted data: %s", c)

	cd, err := aesEncrypt.Decrypt(c)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("Decrypted data: %s", cd)
}
