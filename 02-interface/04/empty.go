package main

import "fmt"
import "strconv"

type Payer interface {
	Pay(int) error
}

//------------------------------------------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

func (w *Wallet) String() string {
	return "Кошелек в котором " + strconv.Itoa(w.Cash) + " денег"
}

func main() {
	myWalet := &Wallet{Cash: 100}
	fmt.Printf("Raw payment : %#v\n", myWalet)
	fmt.Printf("Способ оплаты : %s\n", myWalet)
}
