package set2

import (
	//	"cryptopals/set1"
	"fmt"
	"strings"
)

func Challenge13() {
	profile := []byte(ProfileFor("master@me.com"))
	key := []byte("YELLOW SUBMARINE")
	fmt.Println("SOLUTION 13:")
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

func ProfileFor(email string) string {
	email = strings.ReplaceAll(email, "&", "")
	email = strings.ReplaceAll(email, "=", "")
	cookie := string(append([]byte(email), []byte("&uid=10&role=user")...))
	return cookie
}

func CutAndPaste(cookie1, cookie2 []byte) []byte {
	return []byte{}
}
