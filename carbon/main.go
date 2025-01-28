package main

import (
	"fmt"
	"github.com/dromara/carbon/v2"
)

func main() {
	fmt.Printf("%s\n", carbon.Now())
	fmt.Println(carbon.Now().ToString())
	// 今天日期
	fmt.Println(carbon.Now().ToDateString())
	// 今天时间
	fmt.Println(carbon.Now().ToTimeString())
	// 指定时区的今天此刻
	fmt.Println(carbon.Now(carbon.NewYork).ToDateTimeString())
	// 今天的秒级时间戳
	fmt.Println(carbon.Now().Timestamp())
	// 今天的毫秒级时间戳
	fmt.Println(carbon.Now().TimestampMilli())
	// 今天的微秒级时间戳
	fmt.Println(carbon.Now().TimestampMicro())
	// 今天的纳秒级时间戳
	fmt.Println(carbon.Now().TimestampNano())

	fmt.Printf("%s\n", carbon.Yesterday())
	fmt.Println(carbon.Yesterday().ToString())
	// 昨天日期
	fmt.Println(carbon.Yesterday().ToDateString())
	// 昨天时间
	fmt.Println(carbon.Yesterday().ToTimeString())
	// 指定时区的昨天此刻
	fmt.Println(carbon.Yesterday(carbon.NewYork).ToDateTimeString())
	// 昨天的秒级时间戳
	fmt.Println(carbon.Yesterday().Timestamp())
	// 昨天的毫秒级时间戳
	fmt.Println(carbon.Yesterday().TimestampMilli())
	// 昨天的微秒级时间戳
	fmt.Println(carbon.Yesterday().TimestampMicro())
	// 昨天的纳秒级时间戳
	fmt.Println(carbon.Yesterday().TimestampNano())

	fmt.Printf("%s\n", carbon.Tomorrow())
	fmt.Println(carbon.Tomorrow().ToString())
	// 明天日期
	fmt.Println(carbon.Tomorrow().ToDateString())
	// 明天时间
	fmt.Println(carbon.Tomorrow().ToTimeString())
	// 指定时区的明天此刻
	fmt.Println(carbon.Tomorrow(carbon.NewYork).ToDateTimeString())
	// 明天的秒级时间戳
	fmt.Println(carbon.Tomorrow().Timestamp())
	// 明天的毫秒级时间戳
	fmt.Println(carbon.Tomorrow().TimestampMilli())
	// 明天的微秒级时间戳
	fmt.Println(carbon.Tomorrow().TimestampMicro())
	// 明天的纳秒级时间戳
	fmt.Println(carbon.Tomorrow().TimestampNano())
}
