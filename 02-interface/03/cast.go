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

type Card struct {
	Balance    int
	Number     string
	Cardholder string
	ValidUntil string
	CVV        string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	c.Balance -= amount
	return nil
}

//------------------------------------------------

type ApplePay struct {
	Money   int
	AppleId string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	a.Money -= amount
	return nil
}

//------------------------------------------------

func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличными?")
	case *Card:
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("Не удалось преобразовать к типу *Card")
		}
		fmt.Println("Вставьте катру,", plasticCard.Cardholder)
	default:
		fmt.Println("Что-то новое!")
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

	var MyMoney Payer

	MyMoney = &Card{Balance: 500, Cardholder: "Denn"}
	Buy(MyMoney)

	MyMoney = &ApplePay{Money: 9}
	Buy(MyMoney)
}
