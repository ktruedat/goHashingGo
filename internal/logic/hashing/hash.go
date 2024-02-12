package hashing

import "crypto/md5"

func MD5Func(bytes []byte) []byte {
	md := md5.New()
	hash := md.Sum(bytes)
	return hash
}
