/*使用sume type 处理发布和订阅事件
* eventChan 处理所有事件的总和类型
* interface{} 接口和切换，使用类型开关实现对值模式匹配
 */
package main

import "fmt"

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

// func (sub subscribeEvent) isEvent() {}
// func (sub subscribeEvent) visit(p *pubsubBus) {
// 	p.handleSubscribeEvent(sub)
// }
func (sub subscribeEvent) visit(visitor eventVisitor) {
	visitor.visitSubscribe(sub)
}

type publishEvent struct {
	message string
}

// func (pub publishEvent) isEvent() {}
// func (pub publishEvent) visit(p *pubsubBus) {
// 	p.handlePublishEvent(pub)
// }
func (pub publishEvent) visit(visitor eventVisitor) {
	visitor.visitPublish(pub)
}

type event interface {
	// isEvent()
	// visit(p *pubsubBus)
	visit(visitor eventVisitor)
}

type eventVisitor struct {
	visitSubscribe func(sub subscribeEvent)
	visitPublish   func(pub publishEvent)
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
		// event.visit(p)
		event.visit(eventVisitor{
			visitSubscribe: p.handleSubscribeEvent,
			visitPublish:   p.handlePublishEvent,
		})
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
