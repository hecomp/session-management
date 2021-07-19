// Code generated by counterfeiter. DO NOT EDIT.
package session_managementfakes

import (
	"sync"

	"github.com/hecomp/session-management/internal/models"
	"github.com/hecomp/session-management/pkg/session_management"
)

type FakeSessionMgmntService struct {
	CreateStub        func(*models.SessionRequest) (*session_management.SessionMgmntResponse, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *models.SessionRequest
	}
	createReturns struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	DestroyStub        func(*models.DestroyRequest) (*session_management.SessionMgmntResponse, error)
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 *models.DestroyRequest
	}
	destroyReturns struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	ExtendStub        func(*models.ExtendRequest) (*session_management.SessionMgmntResponse, error)
	extendMutex       sync.RWMutex
	extendArgsForCall []struct {
		arg1 *models.ExtendRequest
	}
	extendReturns struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	extendReturnsOnCall map[int]struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	ListStub        func() (*session_management.SessionMgmntResponse, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSessionMgmntService) Create(arg1 *models.SessionRequest) (*session_management.SessionMgmntResponse, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *models.SessionRequest
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionMgmntService) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSessionMgmntService) CreateCalls(stub func(*models.SessionRequest) (*session_management.SessionMgmntResponse, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeSessionMgmntService) CreateArgsForCall(i int) *models.SessionRequest {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSessionMgmntService) CreateReturns(result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) CreateReturnsOnCall(i int, result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *session_management.SessionMgmntResponse
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) Destroy(arg1 *models.DestroyRequest) (*session_management.SessionMgmntResponse, error) {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 *models.DestroyRequest
	}{arg1})
	stub := fake.DestroyStub
	fakeReturns := fake.destroyReturns
	fake.recordInvocation("Destroy", []interface{}{arg1})
	fake.destroyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionMgmntService) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeSessionMgmntService) DestroyCalls(stub func(*models.DestroyRequest) (*session_management.SessionMgmntResponse, error)) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeSessionMgmntService) DestroyArgsForCall(i int) *models.DestroyRequest {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSessionMgmntService) DestroyReturns(result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) DestroyReturnsOnCall(i int, result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 *session_management.SessionMgmntResponse
			result2 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) Extend(arg1 *models.ExtendRequest) (*session_management.SessionMgmntResponse, error) {
	fake.extendMutex.Lock()
	ret, specificReturn := fake.extendReturnsOnCall[len(fake.extendArgsForCall)]
	fake.extendArgsForCall = append(fake.extendArgsForCall, struct {
		arg1 *models.ExtendRequest
	}{arg1})
	stub := fake.ExtendStub
	fakeReturns := fake.extendReturns
	fake.recordInvocation("Extend", []interface{}{arg1})
	fake.extendMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionMgmntService) ExtendCallCount() int {
	fake.extendMutex.RLock()
	defer fake.extendMutex.RUnlock()
	return len(fake.extendArgsForCall)
}

func (fake *FakeSessionMgmntService) ExtendCalls(stub func(*models.ExtendRequest) (*session_management.SessionMgmntResponse, error)) {
	fake.extendMutex.Lock()
	defer fake.extendMutex.Unlock()
	fake.ExtendStub = stub
}

func (fake *FakeSessionMgmntService) ExtendArgsForCall(i int) *models.ExtendRequest {
	fake.extendMutex.RLock()
	defer fake.extendMutex.RUnlock()
	argsForCall := fake.extendArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSessionMgmntService) ExtendReturns(result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.extendMutex.Lock()
	defer fake.extendMutex.Unlock()
	fake.ExtendStub = nil
	fake.extendReturns = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) ExtendReturnsOnCall(i int, result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.extendMutex.Lock()
	defer fake.extendMutex.Unlock()
	fake.ExtendStub = nil
	if fake.extendReturnsOnCall == nil {
		fake.extendReturnsOnCall = make(map[int]struct {
			result1 *session_management.SessionMgmntResponse
			result2 error
		})
	}
	fake.extendReturnsOnCall[i] = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) List() (*session_management.SessionMgmntResponse, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionMgmntService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeSessionMgmntService) ListCalls(stub func() (*session_management.SessionMgmntResponse, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeSessionMgmntService) ListReturns(result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) ListReturnsOnCall(i int, result1 *session_management.SessionMgmntResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *session_management.SessionMgmntResponse
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *session_management.SessionMgmntResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionMgmntService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.extendMutex.RLock()
	defer fake.extendMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSessionMgmntService) recordInvocation(key string, args []interface{}) {
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

var _ session_management.SessionMgmntService = new(FakeSessionMgmntService)