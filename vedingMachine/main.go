// Package main provides ...
package main

type state interface {
	selectItem(count int) error
	giveMoney(money float32) error
	giveItem() error
	addItem(count int) error
}

type VendingMachine struct {
	hasItem      state
	ruquestItem  state
	hasMoney     state
	hasnoItem    state
	itemCount    int
	itemPrice    float32
	currentState state
}
