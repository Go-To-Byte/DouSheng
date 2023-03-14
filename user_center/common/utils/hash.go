// @Author: Ciusyan 2023/1/30
package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash Bcrypt 将敏感数据做强类型慢速哈希
func BcryptHash(data any) string {

	// 先将数据序列化
	dataMarshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// 对数据做 hash
	bytesPassword, err := bcrypt.GenerateFromPassword(dataMarshal, 14)

	if err != nil {
		panic(err)
	}

	return string(bytesPassword)
}

// VerifyBcryptHash 验证 BcryptHash
func VerifyBcryptHash(data any, hash string) bool {

	dataMarshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), dataMarshal)

	return err == nil
}

// Md5Hash 对不敏感数据做快速 Hash
func Md5Hash(data any) string {

	dataMarshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	hash := md5.New()
	_, err = hash.Write(dataMarshal)

	if err != nil {
		// TODO：快哈希 没必要 panic
		panic(err)
	}

	hd := hash.Sum(nil)

	return fmt.Sprintf("%x", hd)
}

// VerifyMd5Hash 验证 Md5Hash
func VerifyMd5Hash(data any, hash string) bool {
	return Md5Hash(data) == hash
}
