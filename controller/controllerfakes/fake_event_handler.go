// Code generated by counterfeiter. DO NOT EDIT.
package controllerfakes

import (
	sync "sync"
	time "time"

	controller "github.com/alext/heating-controller/controller"
)

type FakeEventHandler struct {
	AddEventStub        func(controller.Event) error
	addEventMutex       sync.RWMutex
	addEventArgsForCall []struct {
		arg1 controller.Event
	}
	addEventReturns struct {
		result1 error
	}
	addEventReturnsOnCall map[int]struct {
		result1 error
	}
	BoostStub        func(time.Duration)
	boostMutex       sync.RWMutex
	boostArgsForCall []struct {
		arg1 time.Duration
	}
	BoostedStub        func() bool
	boostedMutex       sync.RWMutex
	boostedArgsForCall []struct {
	}
	boostedReturns struct {
		result1 bool
	}
	boostedReturnsOnCall map[int]struct {
		result1 bool
	}
	CancelBoostStub        func()
	cancelBoostMutex       sync.RWMutex
	cancelBoostArgsForCall []struct {
	}
	NextEventStub        func() *controller.Event
	nextEventMutex       sync.RWMutex
	nextEventArgsForCall []struct {
	}
	nextEventReturns struct {
		result1 *controller.Event
	}
	nextEventReturnsOnCall map[int]struct {
		result1 *controller.Event
	}
	ReadEventsStub        func() []controller.Event
	readEventsMutex       sync.RWMutex
	readEventsArgsForCall []struct {
	}
	readEventsReturns struct {
		result1 []controller.Event
	}
	readEventsReturnsOnCall map[int]struct {
		result1 []controller.Event
	}
	RemoveEventStub        func(controller.Event)
	removeEventMutex       sync.RWMutex
	removeEventArgsForCall []struct {
		arg1 controller.Event
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventHandler) AddEvent(arg1 controller.Event) error {
	fake.addEventMutex.Lock()
	ret, specificReturn := fake.addEventReturnsOnCall[len(fake.addEventArgsForCall)]
	fake.addEventArgsForCall = append(fake.addEventArgsForCall, struct {
		arg1 controller.Event
	}{arg1})
	fake.recordInvocation("AddEvent", []interface{}{arg1})
	fake.addEventMutex.Unlock()
	if fake.AddEventStub != nil {
		return fake.AddEventStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addEventReturns
	return fakeReturns.result1
}

func (fake *FakeEventHandler) AddEventCallCount() int {
	fake.addEventMutex.RLock()
	defer fake.addEventMutex.RUnlock()
	return len(fake.addEventArgsForCall)
}

func (fake *FakeEventHandler) AddEventArgsForCall(i int) controller.Event {
	fake.addEventMutex.RLock()
	defer fake.addEventMutex.RUnlock()
	argsForCall := fake.addEventArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEventHandler) AddEventReturns(result1 error) {
	fake.AddEventStub = nil
	fake.addEventReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventHandler) AddEventReturnsOnCall(i int, result1 error) {
	fake.AddEventStub = nil
	if fake.addEventReturnsOnCall == nil {
		fake.addEventReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addEventReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventHandler) Boost(arg1 time.Duration) {
	fake.boostMutex.Lock()
	fake.boostArgsForCall = append(fake.boostArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("Boost", []interface{}{arg1})
	fake.boostMutex.Unlock()
	if fake.BoostStub != nil {
		fake.BoostStub(arg1)
	}
}

func (fake *FakeEventHandler) BoostCallCount() int {
	fake.boostMutex.RLock()
	defer fake.boostMutex.RUnlock()
	return len(fake.boostArgsForCall)
}

func (fake *FakeEventHandler) BoostArgsForCall(i int) time.Duration {
	fake.boostMutex.RLock()
	defer fake.boostMutex.RUnlock()
	argsForCall := fake.boostArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEventHandler) Boosted() bool {
	fake.boostedMutex.Lock()
	ret, specificReturn := fake.boostedReturnsOnCall[len(fake.boostedArgsForCall)]
	fake.boostedArgsForCall = append(fake.boostedArgsForCall, struct {
	}{})
	fake.recordInvocation("Boosted", []interface{}{})
	fake.boostedMutex.Unlock()
	if fake.BoostedStub != nil {
		return fake.BoostedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.boostedReturns
	return fakeReturns.result1
}

func (fake *FakeEventHandler) BoostedCallCount() int {
	fake.boostedMutex.RLock()
	defer fake.boostedMutex.RUnlock()
	return len(fake.boostedArgsForCall)
}

func (fake *FakeEventHandler) BoostedReturns(result1 bool) {
	fake.BoostedStub = nil
	fake.boostedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEventHandler) BoostedReturnsOnCall(i int, result1 bool) {
	fake.BoostedStub = nil
	if fake.boostedReturnsOnCall == nil {
		fake.boostedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.boostedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEventHandler) CancelBoost() {
	fake.cancelBoostMutex.Lock()
	fake.cancelBoostArgsForCall = append(fake.cancelBoostArgsForCall, struct {
	}{})
	fake.recordInvocation("CancelBoost", []interface{}{})
	fake.cancelBoostMutex.Unlock()
	if fake.CancelBoostStub != nil {
		fake.CancelBoostStub()
	}
}

func (fake *FakeEventHandler) CancelBoostCallCount() int {
	fake.cancelBoostMutex.RLock()
	defer fake.cancelBoostMutex.RUnlock()
	return len(fake.cancelBoostArgsForCall)
}

func (fake *FakeEventHandler) NextEvent() *controller.Event {
	fake.nextEventMutex.Lock()
	ret, specificReturn := fake.nextEventReturnsOnCall[len(fake.nextEventArgsForCall)]
	fake.nextEventArgsForCall = append(fake.nextEventArgsForCall, struct {
	}{})
	fake.recordInvocation("NextEvent", []interface{}{})
	fake.nextEventMutex.Unlock()
	if fake.NextEventStub != nil {
		return fake.NextEventStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nextEventReturns
	return fakeReturns.result1
}

func (fake *FakeEventHandler) NextEventCallCount() int {
	fake.nextEventMutex.RLock()
	defer fake.nextEventMutex.RUnlock()
	return len(fake.nextEventArgsForCall)
}

func (fake *FakeEventHandler) NextEventReturns(result1 *controller.Event) {
	fake.NextEventStub = nil
	fake.nextEventReturns = struct {
		result1 *controller.Event
	}{result1}
}

func (fake *FakeEventHandler) NextEventReturnsOnCall(i int, result1 *controller.Event) {
	fake.NextEventStub = nil
	if fake.nextEventReturnsOnCall == nil {
		fake.nextEventReturnsOnCall = make(map[int]struct {
			result1 *controller.Event
		})
	}
	fake.nextEventReturnsOnCall[i] = struct {
		result1 *controller.Event
	}{result1}
}

func (fake *FakeEventHandler) ReadEvents() []controller.Event {
	fake.readEventsMutex.Lock()
	ret, specificReturn := fake.readEventsReturnsOnCall[len(fake.readEventsArgsForCall)]
	fake.readEventsArgsForCall = append(fake.readEventsArgsForCall, struct {
	}{})
	fake.recordInvocation("ReadEvents", []interface{}{})
	fake.readEventsMutex.Unlock()
	if fake.ReadEventsStub != nil {
		return fake.ReadEventsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.readEventsReturns
	return fakeReturns.result1
}

func (fake *FakeEventHandler) ReadEventsCallCount() int {
	fake.readEventsMutex.RLock()
	defer fake.readEventsMutex.RUnlock()
	return len(fake.readEventsArgsForCall)
}

func (fake *FakeEventHandler) ReadEventsReturns(result1 []controller.Event) {
	fake.ReadEventsStub = nil
	fake.readEventsReturns = struct {
		result1 []controller.Event
	}{result1}
}

func (fake *FakeEventHandler) ReadEventsReturnsOnCall(i int, result1 []controller.Event) {
	fake.ReadEventsStub = nil
	if fake.readEventsReturnsOnCall == nil {
		fake.readEventsReturnsOnCall = make(map[int]struct {
			result1 []controller.Event
		})
	}
	fake.readEventsReturnsOnCall[i] = struct {
		result1 []controller.Event
	}{result1}
}

func (fake *FakeEventHandler) RemoveEvent(arg1 controller.Event) {
	fake.removeEventMutex.Lock()
	fake.removeEventArgsForCall = append(fake.removeEventArgsForCall, struct {
		arg1 controller.Event
	}{arg1})
	fake.recordInvocation("RemoveEvent", []interface{}{arg1})
	fake.removeEventMutex.Unlock()
	if fake.RemoveEventStub != nil {
		fake.RemoveEventStub(arg1)
	}
}

func (fake *FakeEventHandler) RemoveEventCallCount() int {
	fake.removeEventMutex.RLock()
	defer fake.removeEventMutex.RUnlock()
	return len(fake.removeEventArgsForCall)
}

func (fake *FakeEventHandler) RemoveEventArgsForCall(i int) controller.Event {
	fake.removeEventMutex.RLock()
	defer fake.removeEventMutex.RUnlock()
	argsForCall := fake.removeEventArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEventHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addEventMutex.RLock()
	defer fake.addEventMutex.RUnlock()
	fake.boostMutex.RLock()
	defer fake.boostMutex.RUnlock()
	fake.boostedMutex.RLock()
	defer fake.boostedMutex.RUnlock()
	fake.cancelBoostMutex.RLock()
	defer fake.cancelBoostMutex.RUnlock()
	fake.nextEventMutex.RLock()
	defer fake.nextEventMutex.RUnlock()
	fake.readEventsMutex.RLock()
	defer fake.readEventsMutex.RUnlock()
	fake.removeEventMutex.RLock()
	defer fake.removeEventMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ controller.EventHandler = new(FakeEventHandler)
