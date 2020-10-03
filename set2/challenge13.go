package set2

import (
	"cryptopals/set1"
	"fmt"
	"math/rand"
	"strings"
)

func Challenge13() {
	fmt.Println("SOLUTION 13:", string(CutAndPaste()))
}

func ParseCookie(input string) map[string]string {
	parsed := make(map[string]string)
	lines := strings.Split(input, "&")
	for _, line := range lines {
		kvs := strings.Split(line, "=")
		parsed[kvs[0]] = kvs[1]
	}
	return parsed
}

func EncryptCookie(cookie, key []byte) []byte {
	input := PadPKCS(cookie, GetPaddedLength(cookie))
	return set1.EncryptECB(input, key)
}

func DecryptCookie(encrypted, key []byte) []byte {
	return set1.DecryptECB(encrypted, key)
}

func CookieFor(email string) string {
	email = strings.ReplaceAll(email, "&", "")
	email = strings.ReplaceAll(email, "=", "")
	email = "email=" + email
	cookie := string(append([]byte(email), []byte("&uid=10&role=user")...))
	return cookie
}

func CutAndPaste() []byte {
	rand.Seed(69)
	aesKey := RandomBytes(16)
	blockCookie := []byte(CookieFor("aaaaaaaa@a.co")) // blocks end right at `role=`
	blockCipher := CookieOracle(blockCookie, aesKey)
	upToRoleEncrypted := blockCipher[:len(blockCipher)-4]
	adminCookie := []byte(CookieFor("admin@a.co")) // will give us ciphered `admin`
	adminCipher := CookieOracle(adminCookie, aesKey)
	adminRoleEncrypted := adminCipher[:6]
	fullEncryptedCookie := append(upToRoleEncrypted, adminRoleEncrypted...)
	fullEncryptedCookie = PadPKCS(fullEncryptedCookie, 48)
	return DecryptCookie(fullEncryptedCookie, aesKey)
}

func CookieOracle(inputCookie, aesKey []byte) []byte {
	encryptedCookie := PadAndEncryptECB(inputCookie, aesKey)
	return encryptedCookie
}
