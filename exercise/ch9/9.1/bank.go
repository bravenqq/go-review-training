// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var widthdraw = make(chan drawinfo)

type drawinfo struct {
	amount int
	res    chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Widthdraw(amount int) bool {
	result := make(chan bool)
	go func() {
		widthdraw <- drawinfo{amount: amount, res: result}
	}()

	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-widthdraw:
			if w.amount > balance {
				w.res <- false
			} else {
				balance -= w.amount
				w.res <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
