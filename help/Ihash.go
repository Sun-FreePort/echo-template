package help

import "golang.org/x/crypto/bcrypt"

// IHashPassword 生成安全的哈希化密码
func IHashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// IHashCheckPassword 检查哈希密码与提交密码的一致性
func IHashCheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
