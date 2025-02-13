package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func main() {
	FunctionalVerification()
	StructVerification()
	CustomValidation()
}

func FunctionalVerification() {
	// 验证邮箱
	email := "test@example.com"
	if govalidator.IsEmail(email) {
		fmt.Println(email, "is a valid email address")
	} else {
		fmt.Println(email, "is not a valid email address")
	}

	// 验证url
	url := "https://www.example.com"
	if govalidator.IsURL(url) {
		fmt.Println(url, "is a valid url")
	} else {
		fmt.Println(url, "is not a valid url")
	}

	// 验证IP地址
	ip := "192.168.0.1"
	if govalidator.IsIP(ip) {
		fmt.Println(ip, "is a valid IP address")
	} else {
		fmt.Println(ip, "is not a valid IP address")
	}
}

type User struct {
	Email string `valid:"email"`
	Age   int    `valid:"range(18|60)"`
}

func StructVerification() {
	userValid := User{
		Email: "test@example.com",
		Age:   25,
	}
	valid, err := govalidator.ValidateStruct(userValid)
	if err != nil {
		fmt.Println("invalid struct: ", err)
	}
	if valid {
		fmt.Println("valid struct")
	}
	fmt.Println(userValid)
	userInvalid := User{
		Email: "test@example.com",
		Age:   15,
	}
	valid, err = govalidator.ValidateStruct(userInvalid)
	if err != nil {
		fmt.Println("invalid struct: ", err)
	}
	if valid {
		fmt.Println("valid struct")
	}
	fmt.Println(userInvalid)
}

type UserWithPassword struct {
	Password string `valid:"valid_password"`
}

func CustomValidation() {
	// 自定义验证规则
	govalidator.CustomTypeTagMap.Set("valid_password", func(i interface{}, o interface{}) bool {
		str, ok := i.(string)
		if !ok {
			return false // 如果不是字符串类型直接返回验证失败
		}
		return len(str) >= 6 && containsDigit(str) && containsLetter(str)
	})
	userValid := UserWithPassword{
		Password: "lyb8889999",
	}
	valid, err := govalidator.ValidateStruct(userValid)
	if err != nil {
		fmt.Println("invalid struct: ", err)
	}
	if valid {
		fmt.Println("valid struct")
	}
	fmt.Println(userValid)
	userInvalid := UserWithPassword{
		Password: "123456",
	}
	valid, err = govalidator.ValidateStruct(userInvalid)
	if err != nil {
		fmt.Println("invalid struct: ", err)
	}
	if valid {
		fmt.Println("valid struct")
	}
	fmt.Println(userInvalid)
}

func containsDigit(str string) bool {
	for _, c := range str {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func containsLetter(str string) bool {
	for _, c := range str {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}
