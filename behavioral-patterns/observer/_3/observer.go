package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Event struct {
	Name string
	Data interface{}
}

type Observer interface {
	Update(Event)
	GetID() string
}

type ConcreteObserver struct {
	ID string
}

func (co *ConcreteObserver) Update(event Event) {
	fmt.Printf("Observer %s received event: %s with data: %v\n", co.ID, event.Name, event.Data)
}

func (co *ConcreteObserver) GetID() string {
	return co.ID
}

type EventManager struct {
	observers map[string][]Observer
	mu        sync.RWMutex
	logMu     sync.Mutex
}

func NewEventManager() *EventManager {
	return &EventManager{
		observers: make(map[string][]Observer),
	}
}

func (em *EventManager) Subscribe(eventName string, observer Observer) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.observers[eventName] = append(em.observers[eventName], observer)
	em.log(fmt.Sprintf("Subscribed: %s to %s", observer.GetID(), eventName))
}

func (em *EventManager) Unsubscribe(eventName string, observer Observer) {
	em.mu.Lock()
	defer em.mu.Unlock()
	if observers, found := em.observers[eventName]; found {
		for i, obs := range observers {
			if obs.GetID() == observer.GetID() {
				em.observers[eventName] = append(observers[:i], observers[i+1:]...)
				em.log(fmt.Sprintf("Unsubscribed: %s from %s", observer.GetID(), eventName))
				break
			}
		}
	}
}

func (em *EventManager) Notify(event Event) {
	em.mu.RLock()
	defer em.mu.RUnlock()
	if observers, found := em.observers[event.Name]; found {
		for _, observer := range observers {
			observer.Update(event)
		}
	}
	em.log(fmt.Sprintf("Notified event: %s", event.Name))
}

func (em *EventManager) NotifyAll(event Event) {
	em.mu.RLock()
	defer em.mu.RUnlock()

	for eventName, observers := range em.observers {
		for _, observer := range observers {
			observer.Update(Event{Name: eventName, Data: event.Data})
		}
	}
	em.log(fmt.Sprintf("NotifyAll: event data %v sent to all observers", event.Data))
}

func (em *EventManager) getObserversCount(eventName string) int {
	em.mu.RLock()
	defer em.mu.RUnlock()
	count := len(em.observers[eventName])
	em.log(fmt.Sprintf("Observer count for %s: %d", eventName, count))
	return count
}

func (em *EventManager) log(message string) {
	em.logMu.Lock()
	defer em.logMu.Unlock()
	fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05.000"), message)
}

func main() {
	manager := NewEventManager()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 初始化觀察者
	for i := 1; i <= 3; i++ {
		observer := &ConcreteObserver{ID: fmt.Sprintf("StockObserver-%d", i)}
		manager.Subscribe("stockUpdate", observer)
	}
	for i := 1; i <= 2; i++ {
		observer := &ConcreteObserver{ID: fmt.Sprintf("PriceObserver-%d", i)}
		manager.Subscribe("priceChange", observer)
	}

	var wg sync.WaitGroup

	// 模擬並發操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			switch rand.Intn(5) {
			case 0: // 讀取操作
				manager.getObserversCount("stockUpdate")
			case 1: // 訂閱操作
				eventType := []string{"stockUpdate", "priceChange"}[rand.Intn(2)]
				newObserver := &ConcreteObserver{ID: fmt.Sprintf("New-%s-Observer-%d", eventType, id)}
				manager.Subscribe(eventType, newObserver)
			case 2: // 取消訂閱操作
				eventType := []string{"stockUpdate", "priceChange"}[rand.Intn(2)]
				observerID := fmt.Sprintf("%s-Observer-%d", eventType, rand.Intn(3)+1)
				manager.Unsubscribe(eventType, &ConcreteObserver{ID: observerID})
			case 3: // 特定事件通知
				eventType := []string{"stockUpdate", "priceChange"}[rand.Intn(2)]
				manager.Notify(Event{Name: eventType, Data: fmt.Sprintf("%s-Update-%d", eventType, id)})
			case 4: // NotifyAll 操作
				manager.NotifyAll(Event{Data: fmt.Sprintf("GlobalUpdate-%d", id)})
			}
		}(i)
	}
	wg.Wait()

	fmt.Printf("\nFinal observer counts:\n")
	fmt.Printf("stockUpdate: %d\n", manager.getObserversCount("stockUpdate"))
	fmt.Printf("priceChange: %d\n", manager.getObserversCount("priceChange"))
}
