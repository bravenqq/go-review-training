/*使用sume type 处理发布和订阅事件
* eventChan 处理所有事件的总和类型
* interface{} 接口和切换，使用类型开关实现对值模式匹配
 */
package main

import (
	"fmt"
)

func main() {
	//启动总线
	bus := pubsubBus{make([]chan<- string, 0), make(chan event)}
	go bus.Run()

	//订阅消息
	messageChan := make(chan string, 0)
	bus.eventChan <- subscribeEvent{messageChan}
	bus.eventChan <- publishEvent{"welcome to here!"}

	//推送消息
	fmt.Println("received:", <-messageChan)

}

type subscribeEvent struct {
	messageChan chan<- string
}

func (sub subscribeEvent) isEvent() {}

type publishEvent struct {
	message string
}

func (pub publishEvent) isEvent() {}

type event interface {
	// isEvent()
	visit(p *pubsubBus)
}

type pubsubBus struct {
	subs []chan<- string
	//interface{}能接受任意类型数据，导致运行时可能由于类型问题panic
	// eventChan chan interface{}
	//使用event约束事件类型,必须是实现了event接口
	eventChan chan event
}

func (p *pubsubBus) Run() {
	for event := range p.eventChan {
		// switch e := event.(type) {
		// case subscribeEvent:
		// p.handleSubscribeEvent(e)
		// case publishEvent:
		// 	p.handlePublishEvent(e)
		// default:
		// 	panic("no such event")
		// }
		e.visit(p)
	}
}

func (p *pubsubBus) handlePublishEvent(e publishEvent) {
	for _, sub := range p.subs {
		sub <- e.message
	}
}

func (p *pubsubBus) handleSubscribeEvent(e subscribeEvent) {
	p.subs = append(p.subs, e.messageChan)
}
