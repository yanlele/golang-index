## 密码加密和解密

### 工具方法
```go
func Md5(str string) string {
	data := []byte(str)
	rest := fmt.Sprintf("%x", md5.Sum(data))

	return rest
}

// 加密用户密码
func CryptUserPassword(password string, salt string) string {
	return Md5(password + salt)
}

// 获取4位密码的盐值
func Salt() string {
	rand.Seed(time.Now().UnixNano()) // 伪随机种子
	baseStr := "abcdefghigklmnopqistuvwxyzABCDEFGHIGKLMNOPQISTUVWXYZ0123456789"
	saltLen := 4
	salt := make([]byte, saltLen)
	for n := 0; n < saltLen; n++ {
		salt[n] = baseStr[rand.Int31n(int32(len(baseStr)))]
	}

	return string(salt)
}

// 验证用户的密码是否正确
func VerifyUserPassword(user *modules.User, oldPsd string) bool {
	password := CryptUserPassword(oldPsd, user.Salt)
	if password == user.Password {
		return true
	}
	return false
}
```

使用加密方法加密密码
```go
salt := utils.Salt()
user = modules.User {
    Name:       regData.Name,
    Password:   utils.CryptUserPassword(regData.Password, salt),
    Salt:       salt,
}
```


### 解密方法
自己去研究吧
