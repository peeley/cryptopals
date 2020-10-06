package set2

import (
	"cryptopals/set1"
	"fmt"
	"strings"
)

var aesKey = RandomBytes(16)

func Challenge13() {
	fmt.Println()
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

func EncryptCookieFor(email string) []byte {
	cookie := []byte(CookieFor(email))
	return PadAndEncryptECB(cookie, aesKey)
}

func DecryptCookie(encrypted []byte) []byte {
	return set1.DecryptECB(encrypted, aesKey)
}

func CookieFor(email string) string {
	email = strings.ReplaceAll(email, "&", "")
	email = strings.ReplaceAll(email, "=", "")
	email = "email=" + email
	cookie := string(append([]byte(email), []byte("&uid=10&role=user")...))
	return cookie
}

func CutAndPaste() []byte {

	upToRoleCipher := EncryptCookieFor("xxxxxxxx@a.co")
	upToRoleEncrypted := upToRoleCipher[:32] // text ends after role=

	paddedAdmin := PadPKCS([]byte("admin"), 16)
	adminEmail := "xxxxxxxxxx" + string(paddedAdmin) + "@a.co"
	adminCipher := EncryptCookieFor(adminEmail)
	adminRoleEncrypted := adminCipher[16:32] // encrypted text for "admin"
	// for idx := 6; idx < 16; idx++ {
	// 	adminRoleEncrypted[idx] = 0
	// }

	fullEncryptedCookie := append(upToRoleEncrypted, adminRoleEncrypted...)

	return DecryptCookie(fullEncryptedCookie)
}
