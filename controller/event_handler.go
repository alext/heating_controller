package controller

import (
	"sync"
	"time"

	"github.com/alext/heating-controller/scheduler"
	"github.com/alext/heating-controller/units"
)

//go:generate counterfeiter . EventHandler
type EventHandler interface {
	AddEvent(Event) error
	RemoveEvent(units.TimeOfDay) error
	ReadEvents() []Event
	NextEvent() *Event

	Boost(time.Duration)
	CancelBoost()
	Boosted() bool
}

type eventHandler struct {
	lock    sync.RWMutex
	events  []Event
	demand  func(Event)
	sched   scheduler.Scheduler
	boosted bool
}

func NewEventHandler(s scheduler.Scheduler, demand func(Event)) EventHandler {
	return &eventHandler{
		sched:  s,
		demand: demand,
		events: make([]Event, 0),
	}
}

func (eh *eventHandler) trigger(e Event) {
	eh.lock.Lock()
	defer eh.lock.Unlock()
	eh.boosted = false
	eh.demand(e)
}

// nextEvent returns the next regular event, ignoring any overrides. This is
// different from NextEvent, which queries the scheduler.
func (eh *eventHandler) nextEvent() *Event {
	if len(eh.events) < 1 {
		return nil
	}
	currentToD := units.NewTimeOfDay(timeNow().Local().Clock())
	for _, e := range eh.events {
		if e.Time > currentToD {
			return &e
		}
	}
	return &eh.events[0]
}

func (eh *eventHandler) previousEvent() *Event {
	if len(eh.events) < 1 {
		return nil
	}
	currentToD := units.NewTimeOfDay(timeNow().Local().Clock())
	for i, e := range eh.events {
		if e.Time > currentToD {
			if i > 0 {
				return &eh.events[i-1]
			}
			break
		}
	}
	return &eh.events[len(eh.events)-1]
}

func (eh *eventHandler) AddEvent(e Event) error {
	if !e.Valid() {
		return ErrInvalidEvent
	}
	eh.lock.Lock()
	defer eh.lock.Unlock()

	eh.events = append(eh.events, e)
	sortEvents(eh.events)

	return eh.sched.AddJob(e.buildSchedulerJob(eh.trigger))
}

func (eh *eventHandler) RemoveEvent(t units.TimeOfDay) error {
	eh.lock.Lock()
	defer eh.lock.Unlock()

	newEvents := make([]Event, 0)
	for _, ee := range eh.events {
		if ee.Time != t {
			newEvents = append(newEvents, ee)
		}
	}
	eh.events = newEvents
	return eh.sched.SetJobs(buildSchedulerJobs(eh.events, eh.demand))
}

func (eh *eventHandler) NextEvent() *Event {
	j := eh.sched.NextJob()
	if j == nil {
		return nil
	}
	eh.lock.RLock()
	defer eh.lock.RUnlock()
	for _, e := range eh.events {
		if j.Time == e.Time {
			return &e
		}
	}

	// scheduler is boosted, construct event representing end.
	return &Event{
		Time: j.Time,
	}
}

func (eh *eventHandler) ReadEvents() []Event {
	eh.lock.RLock()
	defer eh.lock.RUnlock()

	events := make([]Event, 0, len(eh.events))
	for _, e := range eh.events {
		events = append(events, e)
	}
	return events
}

func (eh *eventHandler) Boosted() bool {
	eh.lock.RLock()
	defer eh.lock.RUnlock()
	return eh.boosted
}

func (eh *eventHandler) Boost(d time.Duration) {
	eh.lock.Lock()
	defer eh.lock.Unlock()
	eh.boosted = true
	eh.demand(Event{Action: On})

	if d == 0 {
		return
	}

	endTime := timeNow().Local().Add(d)
	endEvent := Event{
		Time:   units.NewTimeOfDay(endTime.Clock()),
		Action: Off,
	}

	nextEvent := eh.nextEvent()

	if nextEvent == nil || endEvent.NextOccurance().Before(nextEvent.NextOccurance()) {
		eh.sched.Override(endEvent.buildSchedulerJob(eh.trigger))
	}
}

func (eh *eventHandler) CancelBoost() {
	eh.lock.Lock()
	defer eh.lock.Unlock()
	if !eh.boosted {
		return
	}
	eh.boosted = false
	eh.sched.CancelOverride()

	previous := eh.previousEvent()
	if previous != nil {
		eh.demand(*previous)
	} else {
		eh.demand(Event{Action: Off})
	}
}
