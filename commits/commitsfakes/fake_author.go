// This file was generated by counterfeiter
package commitsfakes

import (
	"sync"

	"github.com/calebamiles/github-client/commits"
)

type FakeAuthor struct {
	GitHubIDStub        func() string
	gitHubIDMutex       sync.RWMutex
	gitHubIDArgsForCall []struct{}
	gitHubIDReturns     struct {
		result1 string
	}
	EmailStub        func() string
	emailMutex       sync.RWMutex
	emailArgsForCall []struct{}
	emailReturns     struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct{}
	stringReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthor) GitHubID() string {
	fake.gitHubIDMutex.Lock()
	fake.gitHubIDArgsForCall = append(fake.gitHubIDArgsForCall, struct{}{})
	fake.recordInvocation("GitHubID", []interface{}{})
	fake.gitHubIDMutex.Unlock()
	if fake.GitHubIDStub != nil {
		return fake.GitHubIDStub()
	} else {
		return fake.gitHubIDReturns.result1
	}
}

func (fake *FakeAuthor) GitHubIDCallCount() int {
	fake.gitHubIDMutex.RLock()
	defer fake.gitHubIDMutex.RUnlock()
	return len(fake.gitHubIDArgsForCall)
}

func (fake *FakeAuthor) GitHubIDReturns(result1 string) {
	fake.GitHubIDStub = nil
	fake.gitHubIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) Email() string {
	fake.emailMutex.Lock()
	fake.emailArgsForCall = append(fake.emailArgsForCall, struct{}{})
	fake.recordInvocation("Email", []interface{}{})
	fake.emailMutex.Unlock()
	if fake.EmailStub != nil {
		return fake.EmailStub()
	} else {
		return fake.emailReturns.result1
	}
}

func (fake *FakeAuthor) EmailCallCount() int {
	fake.emailMutex.RLock()
	defer fake.emailMutex.RUnlock()
	return len(fake.emailArgsForCall)
}

func (fake *FakeAuthor) EmailReturns(result1 string) {
	fake.EmailStub = nil
	fake.emailReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeAuthor) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeAuthor) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) String() string {
	fake.stringMutex.Lock()
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct{}{})
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if fake.StringStub != nil {
		return fake.StringStub()
	} else {
		return fake.stringReturns.result1
	}
}

func (fake *FakeAuthor) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeAuthor) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.gitHubIDMutex.RLock()
	defer fake.gitHubIDMutex.RUnlock()
	fake.emailMutex.RLock()
	defer fake.emailMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAuthor) recordInvocation(key string, args []interface{}) {
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

var _ commits.Author = new(FakeAuthor)
