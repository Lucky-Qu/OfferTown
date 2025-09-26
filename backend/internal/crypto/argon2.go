// Package crypto argon2.go
//
// 功能:
// - 对明文密码进行argon2id哈希
// - 比较明文密码和哈希
//
// 作者: LuckyQu
// 创建日期: 2025-09-25
// 修改日期: 2025-09-26

package crypto

import (
	"backend/configs"
	"backend/internal/code"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// Encrypt 对传入的字符串类型的密码进行加密
func Encrypt(password string) (string, code.Code) {
	salt, err := randSalt16()
	if err != nil {
		return "", code.EncryptError
	}
	version := argon2.Version
	hashPassword := argon2.IDKey(
		[]byte(password),
		salt,
		configs.Config.Argon2Id.Times,
		configs.Config.Argon2Id.Memory,
		configs.Config.Argon2Id.Threads,
		configs.Config.Argon2Id.KeyLen,
	)
	encodedHashPassword := base64.RawStdEncoding.EncodeToString(hashPassword)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encryptedPassword := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		version,
		configs.Config.Argon2Id.Memory,
		configs.Config.Argon2Id.Times,
		configs.Config.Argon2Id.Threads,
		encodedSalt,
		encodedHashPassword,
	)
	return encryptedPassword, code.Success
}

// Verify 验证传来的密码是否正确
func Verify(encryptedPassword string, password string) (bool, code.Code) {
	// 截断取得各部分
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 6 {
		return false, code.VerifyError
	}
	if parts[1] != "argon2id" {
		return false, code.VerifyError
	}
	var (
		memory  uint32
		times   uint32
		threads uint8
	)
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &times, &threads)
	if err != nil {
		return false, code.VerifyError
	}
	encodedSalt := parts[4]
	encodedHash := parts[5]
	decodedSalt, err := base64.RawStdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false, code.VerifyError
	}
	decodedHash, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, code.VerifyError
	}
	result := subtle.ConstantTimeCompare(decodedHash, argon2.IDKey([]byte(password), decodedSalt, times, memory, threads, uint32(len(decodedHash))))
	if result != 1 {
		return false, code.Success
	}
	return true, code.Success
}

// randSalt16 取得安全的盐值
func randSalt16() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	return salt, err
}
