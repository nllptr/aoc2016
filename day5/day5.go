package day5

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func hash(door string, index int) string {
	input := door + strconv.Itoa(index)
	sum := md5.Sum([]byte(input))
	if hex := hex.EncodeToString(sum[:]); hex[:5] == "00000" {
		return string(hex[5])
	}
	return ""
}

// CrackPassword cracks the password
func CrackPassword(input string) string {
	trimmed := strings.TrimSpace(input)
	index := 0
	crackedDigits := 0
	password := ""
	for crackedDigits <= 8 {
		if hsh := hash(trimmed, index); hsh != "" {
			password += hsh
			crackedDigits++
		}
		index++
	}
	return password
}
