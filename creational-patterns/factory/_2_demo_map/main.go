package main

import "fmt"

var (
	PaymentFactory = map[string]func() PaymentMethod{
		"paypal":     func() PaymentMethod { return &Paypal{} },
		"creditcard": func() PaymentMethod { return &CreditCard{} },
	}
)

// Interface Factory
type PaymentMethod interface {
	Pay(amount float64) string
}

// Concrete
type Paypal struct{}

func (p *Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f via Paypal", amount)
}

// Concrete
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f using Credit Card", amount)
}

// Concrete Factory
func GetPaymentMethod(method string) PaymentMethod {
	if factory, ok := PaymentFactory[method]; ok {
		return factory()
	}
	return nil
}

func main() {
	payment := GetPaymentMethod("paypal")
	fmt.Println(payment.Pay(100.5))
}
