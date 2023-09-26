package main

import (
	"log"
	"sync"
)

type Event struct {
	Message string
}

type Subscriber struct {
	ID       int
	Events   chan Event
	QuitChan chan struct{}
}

func NewSubscriber(id int) *Subscriber {
	return &Subscriber{
		ID:       id,
		Events:   make(chan Event),
		QuitChan: make(chan struct{}),
	}
}

func (s *Subscriber) Start() {
	log.Printf("Subscriber %d started", s.ID)
	for {
		select {
		case event := <-s.Events:
			log.Printf("Subscriber %d received event: %s", s.ID, event.Message)
		case <-s.QuitChan:
			log.Printf("Subscriber %d stopped", s.ID)
			return
		}
	}
}

func (s *Subscriber) Stop() {
	close(s.QuitChan)
}

type Publisher struct {
	Subscribers []*Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{
		Subscribers: make([]*Subscriber, 0),
	}
}

func (p *Publisher) AddSubscriber(subscriber *Subscriber) {
	p.Subscribers = append(p.Subscribers, subscriber)
}

func (p *Publisher) PublishEvent(event Event) {
	for _, subscriber := range p.Subscribers {
		subscriber.Events <- event
	}
}

func (p *Publisher) Stop() {
	for _, subscriber := range p.Subscribers {
		subscriber.Stop()
	}
}

func main() {
	publisher := NewPublisher()

	subscriber1 := NewSubscriber(1)
	subscriber2 := NewSubscriber(2)

	publisher.AddSubscriber(subscriber1)
	publisher.AddSubscriber(subscriber2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		subscriber1.Start()
		wg.Done()
	}()

	go func() {
		subscriber2.Start()
		wg.Done()
	}()

	publisher.PublishEvent(Event{Message: "Hello, World!"})

	// 停止发布者和订阅者
	publisher.Stop()

	wg.Wait()
}
