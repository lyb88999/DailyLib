package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"net/http"
)

func morning(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "morning!")
}

func evening(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "evening!")
}

type Morning struct{}

func (m Morning) Start() {
	http.HandleFunc("/morning", morning)
	http.ListenAndServe(":8080", nil)
}

func (m Morning) Stop() {
	fmt.Println("Morning service stopped")
}

type Evening struct{}

func (e Evening) Start() {
	http.HandleFunc("/evening", evening)
	http.ListenAndServe(":8081", nil)
}
func (e Evening) Stop() {
	fmt.Println("Evening service stopped")
}

//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		fmt.Println("Starting morning service...")
//		m := Morning{}
//		m.Start()
//		defer m.Stop()
//	}()
//	go func() {
//		defer wg.Done()
//		fmt.Println("Starting evening service...")
//		e := Evening{}
//		e.Start()
//		defer e.Stop()
//	}()
//	wg.Wait()
//}

func main() {
	// ServiceGroup 不光只是管理了每个服务的 Start/Stop，同时也提供了 graceful shutdown，当收到 SIGTERM 信号的时候会主动调用每个服务的 Stop 方法
	group := service.NewServiceGroup()
	defer group.Stop()
	group.Add(Morning{})
	group.Add(Evening{})
	group.Start()
}
