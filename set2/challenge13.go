package set2

import (
	"cryptopals/set1"
	"fmt"
	"strings"
)

func Challenge13() {
	profileString := (ProfileFor("poop@gmail.com"))
	fmt.Println(string(profileString), len(profileString))
	roleIndex := strings.Index(profileString, "role=user")
	fmt.Println(roleIndex, string(profileString[roleIndex:]))
	profile := []byte(profileString)
	key := []byte("YELLOW SUBMARINE")
	encryptedCookie := set1.EncryptECB(PadPKCS(profile, GetPaddedLength(profile)), key)
	adminCookie := MangleToAdmin(encryptedCookie, key, roleIndex)
	fmt.Println("SOLUTION 13:", string(adminCookie), len(adminCookie))
	fmt.Println(string(set1.DecryptECB(adminCookie, key)))
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

func MangleToAdmin(encrypted, key []byte, roleIndex int) []byte {
	roleIndex += 5 // go from index of 'r' in role to index of 'u' in user
	targetSubstr := []byte("admin")
	candidate := encrypted
	fmt.Println("starting with,", string(set1.DecryptECB(candidate, key)))
	for _, targetChar := range targetSubstr {
		for byte := byte(0); byte < 255; byte++ {
			candidate[roleIndex] = byte
			decryptedCandidate := set1.DecryptECB(candidate, key)
			if decryptedCandidate[roleIndex] == targetChar {
				fmt.Printf("matched char %v at idx %v: %v\n", string(byte), roleIndex, string(set1.DecryptECB(candidate, key)))
				roleIndex += 1
				break
			}
		}
	}
	return candidate
}
