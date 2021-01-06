package main

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	once             sync.Once
	ErrExceededLimit             = errors.New("Too many requests, exceeded the limit. ")
	_                RateLimiter = &slidingWindowCounter{}
)

type RateLimiter interface {
	Take() error
}

type slidingWindowCounter struct {
	incurRequests    int32
	durationRequests chan int32
	accuracy         time.Duration
	snippet          time.Duration
	currentRequests  int32
	allowRequests    int32
}

func New(accuracy time.Duration, snippet time.Duration, allowRequests int32) *slidingWindowCounter {
	return &slidingWindowCounter{durationRequests: make(chan int32, snippet/accuracy/1000), accuracy: accuracy, snippet: snippet, allowRequests: allowRequests}
}

func (l *slidingWindowCounter) Take() error {
	once.Do(func() {
		go sliding(l)
		go calculate(l)
	})
	curRequest := atomic.LoadInt32(&l.currentRequests)
	if curRequest >= l.allowRequests {
		return ErrExceededLimit
	}
	if !atomic.CompareAndSwapInt32(&l.currentRequests, curRequest, curRequest+1) {
		return ErrExceededLimit
	}
	atomic.AddInt32(&l.incurRequests, 1)
	return nil

}

func sliding(l *slidingWindowCounter) {
	for {
		select {
		case <-time.After(l.accuracy):
			t := atomic.SwapInt32(&l.incurRequests, 0)
			l.durationRequests <- t
		}
	}
}

func calculate(l *slidingWindowCounter) {
	for {
		<-time.After(l.accuracy)
		if len(l.durationRequests) == cap(l.durationRequests) {
			break
		}
	}
	for {
		<-time.After(l.accuracy)
		t := <-l.durationRequests
		if t != 0 {
			atomic.AddInt32(&l.currentRequests, -t)
		}
	}
}
