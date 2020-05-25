// Package main provides ...
package main

import "errors"

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
	requestItem  state
	hasMoney     state
	hasnoItem    state
	itemCount    int
	itemPrice    float32
	count        int //顾客请求商品数量
	currentState state
}

func (vm *VendingMachine) selectItem(count int) error {
	if count <= 0 {
		return errors.New("request count must >0")
	}
	if vm.count < count {
		return errors.New("not enough")
	}
	vm.count = count
	return vm.currentState.selectItem(count)
}

func (vm *VendingMachine) giveMoney(money float32) error {
	if money < vm.itemPrice*float32(vm.count) {
		return errors.New("not enough money")
	}
	return vm.currentState.giveMoney(money)
}

func (vm *VendingMachine) giveItem() error {
	vm.itemCount -= vm.count
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

type hasItem struct {
	vm *VendingMachine
}

func (s hasItem) selectItem(count int) error {
	s.vm.set(s.vm.requestItem)
	return nil
}

func (s hasItem) giveMoney(m float32) error {
	return errors.New("please select item first")
}

func (s hasItem) giveItem() error {
	return errors.New("please select item first")
}

func (s hasItem) addItem(count int) error {
	if count <= 0 {
		return errors.New("request count must >0")
	}
	s.vm.incrmentItem(count)
	return nil
}

type requestItem struct {
	vm *VendingMachine
}

func (s requestItem) selectItem(count int) error {
	return errors.New("item has select")
}

func (s requestItem) giveMoney(m float32) error {
	s.vm.set(s.vm.hasMoney)
	return nil
}

func (s requestItem) giveItem() error {
	return errors.New("must give money")
}

func (s requestItem) addItem() error {
	return errors.New("must finish request")
}

type hasMoney struct {
	vm *VendingMachine
}

func (s hasMoney) selectItem(count int) error {
	return errors.New("must finish selling")
}

func (s hasMoney) giveMoney(m float32) error {
	return errors.New("had give money")
}

func (s hasMoney) giveItem() error {
	if s.vm.itemCount == 0 {
		s.vm.set(s.vm.hasnoItem)
	} else {
		s.vm.set(s.vm.hasItem)
	}
	return nil
}

func (s hasMoney) addItem() error {
	return errors.New("must finish selling")
}

type hasnoItem struct {
	vm *VendingMachine
}

func (s hasnoItem) selectItem(count int) error {
	return errors.New("has no item to sell")
}

func (s hasnoItem) giveMoney(m float32) error {
	return errors.New("item out of stock")
}

func (s hasnoItem) giveItem() error {
	return errors.New("item out of stock")
}

func (s hasnoItem) addItem(count int) error {
	s.vm.incrmentItem(count)
	s.vm.set(s.vm.hasItem)
	return nil
}
