package main

import (
	"fmt"
	"time"
)

// Basic Publisher
type Subject interface {
	register(Observer Observer)
	deregister(Observer Observer)
	notifyAll()
}

// Customer 訂閱的商品, Concrete Publisher
type Item struct {
	observerList []Observer
	name         string
	inStock      bool
	price        float64
}

func newItem(name string, price float64) *Item {
	return &Item{
		name:  name,
		price: price,
	}
}

// 有存貨則通知訂閱客戶
func (item *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", item.name)
	item.inStock = true
	item.notifyAll()
}

func (item *Item) register(observer Observer) {
	item.observerList = append(item.observerList, observer)
}

func (item *Item) deregister(observer Observer) {
	item.observerList = removeObserverFromSlice(item.observerList, observer)
}

func (item *Item) notifyAll() {
	for _, observer := range item.observerList {
		observer.update(item.name)
	}
}

func removeObserverFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// Subscriber
type Observer interface {
	update(itemName string)
	getID() string
}

// Concrete Subscriber
type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	// use "net/smtp" package send email to customer
	fmt.Printf("Sending email to notify customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

func main() {

	// Create items
	gpuCard := newItem("NVIDIA A100", 10000.00)
	cpu := newItem("AMD Ryzen 9 5950X", 799.99)

	// Create customers
	customer1 := &Customer{id: "john@example.com"}
	customer2 := &Customer{id: "alice@example.com"}
	customer3 := &Customer{id: "bob@example.com"}

	// Register customers for GPU
	gpuCard.register(customer1)
	gpuCard.register(customer2)

	// Register customers for CPU
	cpu.register(customer2)
	cpu.register(customer3)

	// Simulate item becoming available
	time.Sleep(2 * time.Second)
	fmt.Println("\nUpdating GPU availability:")
	gpuCard.updateAvailability()

	time.Sleep(2 * time.Second)
	fmt.Println("\nUpdating CPU availability:")
	cpu.updateAvailability()

	// Deregister a customer
	time.Sleep(2 * time.Second)
	fmt.Println("\nDeregistering customer2 from GPU:")
	gpuCard.deregister(customer2)

	// Update availability again
	time.Sleep(2 * time.Second)
	fmt.Println("\nUpdating GPU availability after deregistration:")
	gpuCard.updateAvailability()
}
