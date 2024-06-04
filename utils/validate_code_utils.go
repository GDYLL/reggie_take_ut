package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

// GenerateValidateCode 生成4位或6位数字验证码
func GenerateValidateCode(length int) (int, error) {
	if length != 4 && length != 6 {
		return 0, fmt.Errorf("只能生成4位或6位数字验证码")
	}

	var max int64
	if length == 4 {
		max = 9999
	} else {
		max = 999999
	}

	code, err := rand.Int(rand.Reader, big.NewInt(max+1))
	if err != nil {
		return 0, err
	}

	if length == 4 && code.Int64() < 1000 {
		return int(code.Int64()) + 1000, nil
	} else if length == 6 && code.Int64() < 100000 {
		return int(code.Int64()) + 100000, nil
	}

	return int(code.Int64()), nil
}

// GenerateValidateCode4String 生成指定长度的字符串验证码
func GenerateValidateCode4String(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("验证码长度必须大于0")
	}

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(bytes)[:length]), nil
}
