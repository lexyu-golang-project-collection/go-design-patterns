package main

import (
	"fmt"
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
}

func newItem(name string) *Item {
	return &Item{
		name: name,
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

	gpu_card := newItem("A100")

	subscriber_1 := &Customer{id: "first@gmail.com"}
	subscriber_2 := &Customer{id: "second@gmail.com"}

	gpu_card.register(subscriber_1)
	gpu_card.register(subscriber_2)

	gpu_card.notifyAll()
}
