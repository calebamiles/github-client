// This file was generated by counterfeiter
package commentsfakes

import (
	"sync"

	"github.com/calebamiles/github-client/comments"
)

type FakeComment struct {
	AuthorStub        func() string
	authorMutex       sync.RWMutex
	authorArgsForCall []struct{}
	authorReturns     struct {
		result1 string
	}
	BodyStub        func() string
	bodyMutex       sync.RWMutex
	bodyArgsForCall []struct{}
	bodyReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeComment) Author() string {
	fake.authorMutex.Lock()
	fake.authorArgsForCall = append(fake.authorArgsForCall, struct{}{})
	fake.recordInvocation("Author", []interface{}{})
	fake.authorMutex.Unlock()
	if fake.AuthorStub != nil {
		return fake.AuthorStub()
	} else {
		return fake.authorReturns.result1
	}
}

func (fake *FakeComment) AuthorCallCount() int {
	fake.authorMutex.RLock()
	defer fake.authorMutex.RUnlock()
	return len(fake.authorArgsForCall)
}

func (fake *FakeComment) AuthorReturns(result1 string) {
	fake.AuthorStub = nil
	fake.authorReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComment) Body() string {
	fake.bodyMutex.Lock()
	fake.bodyArgsForCall = append(fake.bodyArgsForCall, struct{}{})
	fake.recordInvocation("Body", []interface{}{})
	fake.bodyMutex.Unlock()
	if fake.BodyStub != nil {
		return fake.BodyStub()
	} else {
		return fake.bodyReturns.result1
	}
}

func (fake *FakeComment) BodyCallCount() int {
	fake.bodyMutex.RLock()
	defer fake.bodyMutex.RUnlock()
	return len(fake.bodyArgsForCall)
}

func (fake *FakeComment) BodyReturns(result1 string) {
	fake.BodyStub = nil
	fake.bodyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeComment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authorMutex.RLock()
	defer fake.authorMutex.RUnlock()
	fake.bodyMutex.RLock()
	defer fake.bodyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeComment) recordInvocation(key string, args []interface{}) {
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

var _ comments.Comment = new(FakeComment)