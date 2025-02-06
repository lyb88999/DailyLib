package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// 不同的值转为布尔值
	fmt.Println(cast.ToBool("true"))
	fmt.Println(cast.ToBool("false"))
	fmt.Println(cast.ToBool(1))
	fmt.Println(cast.ToBool(0))
	fmt.Println(cast.ToBool("notabool"))

	// 不同的值转为整数
	fmt.Println(cast.ToInt("123"))
	fmt.Println(cast.ToInt("abc"))
	fmt.Println(cast.ToInt(10.5))
	fmt.Println(cast.ToInt(false))

	// 不同的值转为int64
	fmt.Println(cast.ToInt64("1234567890"))
	fmt.Println(cast.ToInt64(10.99))

	// 不同的值转为float64
	fmt.Println(cast.ToFloat64("123.45"))
	fmt.Println(cast.ToFloat64(10))
	fmt.Println(cast.ToFloat64("notafloat"))

	// 不同的值转为string
	fmt.Println(cast.ToString(123))
	fmt.Println(cast.ToString(true))
	fmt.Println(cast.ToString(10.99))
	fmt.Println(cast.ToString(nil))

	// 不同的值转为切片
	fmt.Println(cast.ToIntSlice([]int{1, 2}))
	fmt.Println(cast.ToStringSlice("hello"))

	// 不同的值转换为映射
	fmt.Println(cast.ToStringMapIntE(map[string]int{"a": 1, "b": 2}))

	// 不同的值转换为时间
	fmt.Println(cast.ToDuration("10s"))
	fmt.Println(cast.ToDuration(20))
}
