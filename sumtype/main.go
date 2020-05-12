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
	bus := pubsubBus{make([]chan<- string, 0), make(chan interface{})}
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

type publishEvent struct {
	message string
}

type pubsubBus struct {
	subs      []chan<- string
	eventChan chan interface{}
}

func (p *pubsubBus) Run() {
	for event := range p.eventChan {
		switch e := event.(type) {
		case subscribeEvent:
			p.handleSubscribeEvent(e)
		case publishEvent:
			p.handlePublishEvent(e)
		default:
			panic("no such event")
		}
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
