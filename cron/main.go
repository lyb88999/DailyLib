package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	// 基础定时任务(开启秒级调度)
	// BasicScheduledTask()
	// 多个定时任务(开启秒级调度)
	// MultipleScheduledTask()
	// 自定义任务(开启秒级调度)
	RegisterTaskUsingAddJob()
}

func BasicScheduledTask() {
	c := cron.New(cron.WithSeconds())
	defer c.Stop()
	// cron表达式写在README.md里面
	// 注册一个定时任务 每5秒执行一次
	_, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("Task executed at: ", time.Now())
	})
	if err != nil {
		log.Fatalf("Failed to add task: %v", err)
	}
	// 启动cron调度器
	c.Start()
	select {}
}

func MultipleScheduledTask() {
	c := cron.New(cron.WithSeconds())
	defer c.Stop()
	_, err := c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Task1 executed at: ", time.Now())
	})
	if err != nil {
		log.Fatalf("Failed to add task1: %v", err)
	}
	_, err = c.AddFunc("*/15 * * * * *", func() {
		fmt.Println("Task2 executed at: ", time.Now())
	})
	if err != nil {
		log.Fatalf("Failed to add task2: %v", err)
	}
	c.Start()
	select {}
}

type MyJob struct {
}

func (m *MyJob) Run() {
	fmt.Println("Custom job executed at: ", time.Now())
}

func RegisterTaskUsingAddJob() {
	c := cron.New(cron.WithSeconds())
	defer c.Stop()
	_, err := c.AddJob("*/10 * * * * *", &MyJob{})
	if err != nil {
		log.Fatalf("Failed to add custom job: %v", err)
	}
	c.Start()
	select {}
}
