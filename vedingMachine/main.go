// Package main provides ...
package main

//state 自动售卖机的行为
type state interface {
	selectItem(count int) error
	giveMoney(money float32) error
	giveItem() error
	addItem(count int) error
}

//VendingMachine 自动售卖机，售卖商品
type VendingMachine struct {
	hasItem      state
	ruquestItem  state
	hasMoney     state
	hasnoItem    state
	itemCount    int
	itemPrice    float32
	count        int //顾客请求商品数量
	currentState state
}

func (vm *VendingMachine) selectItem(count int) error {
	return vm.currentState.selectItem(count)
}

func (vm *VendingMachine) giveMoney(money float32) error {
	return vm.currentState.giveMoney(money)
}

func (vm *VendingMachine) giveItem() error {
	return vm.currentState.giveItem()
}

func (vm *VendingMachine) addItem(count int) error {
	return vm.addItem(count)
}

func (vm *VendingMachine) set(s state) {
	vm.currentState = s
}

func (vm *VendingMachine) incrmentItem(count int) {
	vm.itemCount += count
}
