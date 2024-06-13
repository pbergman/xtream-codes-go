package xtream_codes_go

import (
	"time"
)

type stopwatch struct {
	events map[string]*stopwatchEvent
}

func (s *stopwatch) getContext() map[string]interface{} {

	var ctx = make(map[string]interface{})

	for name, event := range s.events {
		ctx[name] = event.String()
	}

	return ctx
}

func (s *stopwatch) GetEvent(name string) *stopwatchEvent {
	_, ok := s.events[name]

	if !ok {
		s.events[name] = new(stopwatchEvent)
	}

	return s.events[name]
}

type stopwatchEvent struct {
	start time.Time
	stop  time.Time
}

func (s *stopwatchEvent) Start() {
	s.start = time.Now()
}

func (s *stopwatchEvent) Stop() {
	s.stop = time.Now()
}

func (s *stopwatchEvent) String() string {
	return s.stop.Sub(s.start).String()
}
