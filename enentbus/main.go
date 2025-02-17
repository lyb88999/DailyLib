package main

import (
	"github.com/asaskevich/EventBus"
	"log"
)

type Order struct {
	ID    string
	Price float64
}

func main() {
	// 创建一个新的事件总线实例
	bus := EventBus.New()

	// 定义一个函数 用于处理订单创建事件
	var orderCreatedHandler = func(order Order) {
		log.Printf("订单 %s 创建成功，价格为 %.2f\n", order.ID, order.Price)
	}

	// 订阅订单创建事件
	err := bus.Subscribe("order:created", orderCreatedHandler)
	if err != nil {
		log.Fatalf("failed to subscribe order:created event\n")
	}

	// 创建一个订单实例
	newOrder := Order{
		ID:    "12345",
		Price: 99.99,
	}

	// 发布订单创建事件
	bus.Publish("order:created", newOrder)

	//// 取消订阅订单创建事件
	//err = bus.Unsubscribe("order:created", orderCreatedHandler)
	//if err != nil {
	//	log.Fatalf("failed to unsubscribe order:created event")
	//}
}
