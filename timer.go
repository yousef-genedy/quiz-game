package main

import (
	"fmt"
	"sync"
	"time"
)

type Score struct {
	mu             sync.Mutex
	correct, total int
}

func (s *Score) Add(isCorrect bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if isCorrect {
		s.correct++
	}
}

func (s *Score) Snapshot() (int, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.correct, s.total
}

func Timer(limit int, problems [][]string, isInterrupted bool) {
	timeUp := time.After(time.Duration(limit) * time.Second)

	resultCh := make(chan struct{})

	score := Score{
		mu:      sync.Mutex{},
		correct: 0,
		total:   len(problems),
	}

	go func() {
		if !isInterrupted {
			Process(problems, &score)
		}
		close(resultCh)
	}()

	select {
	case <-timeUp:
		fmt.Println("\nTime's up!")
	case <-resultCh:
	}

	correct, total := score.Snapshot()

	fmt.Printf("You scored %d out of %d.\n", correct, total)
}
