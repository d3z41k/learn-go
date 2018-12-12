package main

import "fmt"

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

//------------------------------------------------

func Buy(in interface{}) {
	var p Payer
	var ok bool

	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не является платежным средством \n\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWalet := &Wallet{Cash: 100}
	Buy(myWalet)
	fmt.Printf("Оставшиеся средства: %d\n\n", myWalet.Cash)

	Buy([]int{1, 2, 3})
	Buy(3.1)
}
