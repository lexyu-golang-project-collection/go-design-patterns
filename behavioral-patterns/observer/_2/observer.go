package main

import (
	"fmt"
	"sync"
	"time"
)

// Event 代表可以被觀察的事件
type Event struct {
	Name string
	Data interface{}
}

// Observer 定義了觀察者的接口
type Observer interface {
	Update(Event)
}

// EventManager 管理事件和觀察者
type EventManager struct {
	observers map[string][]Observer
	mu        sync.RWMutex
}

// NewEventManager 創建一個新的 EventManager
func NewEventManager() *EventManager {
	return &EventManager{
		observers: make(map[string][]Observer),
	}
}

func (em *EventManager) getObserversCount(eventName string) int {
	em.mu.RLock()
	defer em.mu.RUnlock()
	return len(em.observers[eventName])
}

func (em *EventManager) getObservers(eventName string) []Observer {
	em.mu.RLock()
	defer em.mu.RUnlock()
	return em.observers[eventName]
}

// Subscribe 訂閱特定事件
func (em *EventManager) Subscribe(eventName string, observer Observer) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.observers[eventName] = append(em.observers[eventName], observer)
}

// Unsubscribe 取消訂閱特定事件
func (em *EventManager) Unsubscribe(eventName string, observer Observer) {
	em.mu.Lock()
	defer em.mu.Unlock()
	if observers, found := em.observers[eventName]; found {
		for i, obs := range observers {
			if obs == observer {
				em.observers[eventName] = append(observers[:i], observers[i+1:]...)
				break
			}
		}
	}
}

// Notify 通知特定事件的所有觀察者
func (em *EventManager) Notify(event Event) {
	em.mu.RLock()
	defer em.mu.RUnlock()
	if observers, found := em.observers[event.Name]; found {
		for _, observer := range observers {
			observer.Update(event)
		}
	}
}

// ConcreteObserver 是 Observer 接口的具體實現
type ConcreteObserver struct {
	ID string
}

func (co *ConcreteObserver) Update(event Event) {
	fmt.Printf("Observer %s received event: %s with data: %v\n", co.ID, event.Name, event.Data)
}

func main() {
	concurrentDemo()
}

func demo() {
	// 創建 EventManager
	manager := NewEventManager()

	// 創建觀察者
	observer1 := &ConcreteObserver{ID: "1"}
	observer2 := &ConcreteObserver{ID: "2"}

	// 訂閱事件
	manager.Subscribe("stockUpdate", observer1)
	manager.Subscribe("stockUpdate", observer2)
	manager.Subscribe("priceChange", observer1)

	// 觸發事件
	manager.Notify(Event{Name: "stockUpdate", Data: 100})
	manager.Notify(Event{Name: "priceChange", Data: 9.99})

	// 取消訂閱
	manager.Unsubscribe("stockUpdate", observer2)

	// 再次觸發事件
	manager.Notify(Event{Name: "stockUpdate", Data: 150})
}

func concurrentDemo() {
	// 創建 EventManager
	manager := NewEventManager()

	// 創建觀察者
	observer1 := &ConcreteObserver{ID: "1"}
	observer2 := &ConcreteObserver{ID: "2"}

	// 訂閱事件
	manager.Subscribe("stockUpdate", observer1)
	manager.Subscribe("stockUpdate", observer2)

	// 使用 WaitGroup 來等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 模擬多個線程同時觸發事件
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模擬一些處理時間
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			manager.Notify(Event{Name: "stockUpdate", Data: 100 + id})
		}(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	fmt.Println("All notifications sent.")

	// 展示取消訂閱功能
	manager.Unsubscribe("stockUpdate", observer2)

	fmt.Println("\nAfter unsubscribing observer2:")

	// 再次觸發事件
	manager.Notify(Event{Name: "stockUpdate", Data: 200})
}
