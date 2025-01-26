package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"time"
)

func main() {
	// 监控单个目录变化 监听多个目录变化的话只需要多加几个watcher.Add
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	dir := "."
	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	// 事件处理
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("Event: ", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("创建文件：", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("修改文件：", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("删除文件：", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Println("重命名文件：", event.Name)
				}
			case err = <-watcher.Errors:
				fmt.Println("Error: ", err)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Creating file test.txt...")
	_, err = createFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Deleting file test.txt...")
	err = deleteFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	select {}
}

func createFile(filename string) (*os.File, error) {
	return os.Create(filename)
}

func deleteFile(filename string) error {
	return os.Remove(filename)
}
