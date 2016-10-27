// This file was generated by counterfeiter
package fetcherfakes

import (
	"sync"

	"github.com/calebamiles/github-client/client/fetcher"
)

type FakeFetcher struct {
	FetchStub        func(url string) (pageContent []byte, err error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		url string
	}
	fetchReturns struct {
		result1 []byte
		result2 error
	}
	DoneStub        func() error
	doneMutex       sync.RWMutex
	doneArgsForCall []struct{}
	doneReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetcher) Fetch(url string) (pageContent []byte, err error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		url string
	}{url})
	fake.recordInvocation("Fetch", []interface{}{url})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(url)
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2
	}
}

func (fake *FakeFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeFetcher) FetchArgsForCall(i int) string {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].url
}

func (fake *FakeFetcher) FetchReturns(result1 []byte, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) Done() error {
	fake.doneMutex.Lock()
	fake.doneArgsForCall = append(fake.doneArgsForCall, struct{}{})
	fake.recordInvocation("Done", []interface{}{})
	fake.doneMutex.Unlock()
	if fake.DoneStub != nil {
		return fake.DoneStub()
	} else {
		return fake.doneReturns.result1
	}
}

func (fake *FakeFetcher) DoneCallCount() int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return len(fake.doneArgsForCall)
}

func (fake *FakeFetcher) DoneReturns(result1 error) {
	fake.DoneStub = nil
	fake.doneReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFetcher) recordInvocation(key string, args []interface{}) {
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

var _ fetcher.Fetcher = new(FakeFetcher)
