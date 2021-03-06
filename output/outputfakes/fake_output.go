// Code generated by counterfeiter. DO NOT EDIT.
package outputfakes

import (
	sync "sync"

	output "github.com/alext/heating-controller/output"
)

type FakeOutput struct {
	ActivateStub        func() error
	activateMutex       sync.RWMutex
	activateArgsForCall []struct {
	}
	activateReturns struct {
		result1 error
	}
	activateReturnsOnCall map[int]struct {
		result1 error
	}
	ActiveStub        func() (bool, error)
	activeMutex       sync.RWMutex
	activeArgsForCall []struct {
	}
	activeReturns struct {
		result1 bool
		result2 error
	}
	activeReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DeactivateStub        func() error
	deactivateMutex       sync.RWMutex
	deactivateArgsForCall []struct {
	}
	deactivateReturns struct {
		result1 error
	}
	deactivateReturnsOnCall map[int]struct {
		result1 error
	}
	IdStub        func() string
	idMutex       sync.RWMutex
	idArgsForCall []struct {
	}
	idReturns struct {
		result1 string
	}
	idReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOutput) Activate() error {
	fake.activateMutex.Lock()
	ret, specificReturn := fake.activateReturnsOnCall[len(fake.activateArgsForCall)]
	fake.activateArgsForCall = append(fake.activateArgsForCall, struct {
	}{})
	fake.recordInvocation("Activate", []interface{}{})
	fake.activateMutex.Unlock()
	if fake.ActivateStub != nil {
		return fake.ActivateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.activateReturns
	return fakeReturns.result1
}

func (fake *FakeOutput) ActivateCallCount() int {
	fake.activateMutex.RLock()
	defer fake.activateMutex.RUnlock()
	return len(fake.activateArgsForCall)
}

func (fake *FakeOutput) ActivateReturns(result1 error) {
	fake.ActivateStub = nil
	fake.activateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutput) ActivateReturnsOnCall(i int, result1 error) {
	fake.ActivateStub = nil
	if fake.activateReturnsOnCall == nil {
		fake.activateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.activateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutput) Active() (bool, error) {
	fake.activeMutex.Lock()
	ret, specificReturn := fake.activeReturnsOnCall[len(fake.activeArgsForCall)]
	fake.activeArgsForCall = append(fake.activeArgsForCall, struct {
	}{})
	fake.recordInvocation("Active", []interface{}{})
	fake.activeMutex.Unlock()
	if fake.ActiveStub != nil {
		return fake.ActiveStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.activeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOutput) ActiveCallCount() int {
	fake.activeMutex.RLock()
	defer fake.activeMutex.RUnlock()
	return len(fake.activeArgsForCall)
}

func (fake *FakeOutput) ActiveReturns(result1 bool, result2 error) {
	fake.ActiveStub = nil
	fake.activeReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeOutput) ActiveReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ActiveStub = nil
	if fake.activeReturnsOnCall == nil {
		fake.activeReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.activeReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeOutput) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeOutput) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeOutput) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutput) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutput) Deactivate() error {
	fake.deactivateMutex.Lock()
	ret, specificReturn := fake.deactivateReturnsOnCall[len(fake.deactivateArgsForCall)]
	fake.deactivateArgsForCall = append(fake.deactivateArgsForCall, struct {
	}{})
	fake.recordInvocation("Deactivate", []interface{}{})
	fake.deactivateMutex.Unlock()
	if fake.DeactivateStub != nil {
		return fake.DeactivateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deactivateReturns
	return fakeReturns.result1
}

func (fake *FakeOutput) DeactivateCallCount() int {
	fake.deactivateMutex.RLock()
	defer fake.deactivateMutex.RUnlock()
	return len(fake.deactivateArgsForCall)
}

func (fake *FakeOutput) DeactivateReturns(result1 error) {
	fake.DeactivateStub = nil
	fake.deactivateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutput) DeactivateReturnsOnCall(i int, result1 error) {
	fake.DeactivateStub = nil
	if fake.deactivateReturnsOnCall == nil {
		fake.deactivateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deactivateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutput) Id() string {
	fake.idMutex.Lock()
	ret, specificReturn := fake.idReturnsOnCall[len(fake.idArgsForCall)]
	fake.idArgsForCall = append(fake.idArgsForCall, struct {
	}{})
	fake.recordInvocation("Id", []interface{}{})
	fake.idMutex.Unlock()
	if fake.IdStub != nil {
		return fake.IdStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.idReturns
	return fakeReturns.result1
}

func (fake *FakeOutput) IdCallCount() int {
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	return len(fake.idArgsForCall)
}

func (fake *FakeOutput) IdReturns(result1 string) {
	fake.IdStub = nil
	fake.idReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeOutput) IdReturnsOnCall(i int, result1 string) {
	fake.IdStub = nil
	if fake.idReturnsOnCall == nil {
		fake.idReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.idReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeOutput) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.activateMutex.RLock()
	defer fake.activateMutex.RUnlock()
	fake.activeMutex.RLock()
	defer fake.activeMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.deactivateMutex.RLock()
	defer fake.deactivateMutex.RUnlock()
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOutput) recordInvocation(key string, args []interface{}) {
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

var _ output.Output = new(FakeOutput)
