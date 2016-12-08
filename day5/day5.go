package day5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func hash(door string, index int) byte {
	input := door + strconv.Itoa(index)
	sum := md5.Sum([]byte(input))
	if hex := hex.EncodeToString(sum[:]); hex[:5] == "00000" {
		return hex[5]
	}
	return ' '
}

func hash2(door string, index int, password []byte) (int, byte) {
	input := door + strconv.Itoa(index)
	sum := md5.Sum([]byte(input))
	if hex := hex.EncodeToString(sum[:]); hex[:5] == "00000" {
		if pwdIndex, err := strconv.Atoi(string(hex[5])); pwdIndex < len(password) {
			if err != nil {
				return -1, '_'
			}
			return pwdIndex, hex[6]
		}
	}
	return -1, '_'
}

// CrackPassword cracks the password
func CrackPassword(input string) string {
	trimmed := strings.TrimSpace(input)
	index := 0
	crackedDigits := 0
	password := []byte("________")
	for crackedDigits < len(password) {
		if hsh := hash(trimmed, index); hsh != ' ' {
			password[crackedDigits] = hsh
			crackedDigits++
			fmt.Printf("\rCracking: %v", string(password))
		}
		index++
	}
	return string(password)
}

// CrackPassword2 cracks the password - unordered!
func CrackPassword2(input string) string {
	trimmed := strings.TrimSpace(input)
	index := 0
	crackedDigits := 0
	password := []byte("________")
	for crackedDigits < len(password) {
		if pos, hsh := hash2(trimmed, index, password); pos >= 0 {
			if password[pos] == '_' {
				crackedDigits++
				password[pos] = hsh
				fmt.Printf("\rCracking: %v", string(password))
			}
		}
		index++
	}
	return string(password)
}
