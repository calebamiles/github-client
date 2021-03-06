// This file was generated by counterfeiter
package milestonefakes

import (
	"sync"
	"time"

	"github.com/calebamiles/github-client/milestone"
)

type FakeMilestone struct {
	OpenStub        func() bool
	openMutex       sync.RWMutex
	openArgsForCall []struct{}
	openReturns     struct {
		result1 bool
	}
	DescriptionStub        func() string
	descriptionMutex       sync.RWMutex
	descriptionArgsForCall []struct{}
	descriptionReturns     struct {
		result1 string
	}
	IssuesOpenStub        func() int
	issuesOpenMutex       sync.RWMutex
	issuesOpenArgsForCall []struct{}
	issuesOpenReturns     struct {
		result1 int
	}
	IssuesClosedStub        func() int
	issuesClosedMutex       sync.RWMutex
	issuesClosedArgsForCall []struct{}
	issuesClosedReturns     struct {
		result1 int
	}
	DeadlineStub        func() time.Time
	deadlineMutex       sync.RWMutex
	deadlineArgsForCall []struct{}
	deadlineReturns     struct {
		result1 time.Time
	}
	TitleStub        func() string
	titleMutex       sync.RWMutex
	titleArgsForCall []struct{}
	titleReturns     struct {
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

func (fake *FakeMilestone) Open() bool {
	fake.openMutex.Lock()
	fake.openArgsForCall = append(fake.openArgsForCall, struct{}{})
	fake.recordInvocation("Open", []interface{}{})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub()
	} else {
		return fake.openReturns.result1
	}
}

func (fake *FakeMilestone) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeMilestone) OpenReturns(result1 bool) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMilestone) Description() string {
	fake.descriptionMutex.Lock()
	fake.descriptionArgsForCall = append(fake.descriptionArgsForCall, struct{}{})
	fake.recordInvocation("Description", []interface{}{})
	fake.descriptionMutex.Unlock()
	if fake.DescriptionStub != nil {
		return fake.DescriptionStub()
	} else {
		return fake.descriptionReturns.result1
	}
}

func (fake *FakeMilestone) DescriptionCallCount() int {
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	return len(fake.descriptionArgsForCall)
}

func (fake *FakeMilestone) DescriptionReturns(result1 string) {
	fake.DescriptionStub = nil
	fake.descriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMilestone) IssuesOpen() int {
	fake.issuesOpenMutex.Lock()
	fake.issuesOpenArgsForCall = append(fake.issuesOpenArgsForCall, struct{}{})
	fake.recordInvocation("IssuesOpen", []interface{}{})
	fake.issuesOpenMutex.Unlock()
	if fake.IssuesOpenStub != nil {
		return fake.IssuesOpenStub()
	} else {
		return fake.issuesOpenReturns.result1
	}
}

func (fake *FakeMilestone) IssuesOpenCallCount() int {
	fake.issuesOpenMutex.RLock()
	defer fake.issuesOpenMutex.RUnlock()
	return len(fake.issuesOpenArgsForCall)
}

func (fake *FakeMilestone) IssuesOpenReturns(result1 int) {
	fake.IssuesOpenStub = nil
	fake.issuesOpenReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeMilestone) IssuesClosed() int {
	fake.issuesClosedMutex.Lock()
	fake.issuesClosedArgsForCall = append(fake.issuesClosedArgsForCall, struct{}{})
	fake.recordInvocation("IssuesClosed", []interface{}{})
	fake.issuesClosedMutex.Unlock()
	if fake.IssuesClosedStub != nil {
		return fake.IssuesClosedStub()
	} else {
		return fake.issuesClosedReturns.result1
	}
}

func (fake *FakeMilestone) IssuesClosedCallCount() int {
	fake.issuesClosedMutex.RLock()
	defer fake.issuesClosedMutex.RUnlock()
	return len(fake.issuesClosedArgsForCall)
}

func (fake *FakeMilestone) IssuesClosedReturns(result1 int) {
	fake.IssuesClosedStub = nil
	fake.issuesClosedReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeMilestone) Deadline() time.Time {
	fake.deadlineMutex.Lock()
	fake.deadlineArgsForCall = append(fake.deadlineArgsForCall, struct{}{})
	fake.recordInvocation("Deadline", []interface{}{})
	fake.deadlineMutex.Unlock()
	if fake.DeadlineStub != nil {
		return fake.DeadlineStub()
	} else {
		return fake.deadlineReturns.result1
	}
}

func (fake *FakeMilestone) DeadlineCallCount() int {
	fake.deadlineMutex.RLock()
	defer fake.deadlineMutex.RUnlock()
	return len(fake.deadlineArgsForCall)
}

func (fake *FakeMilestone) DeadlineReturns(result1 time.Time) {
	fake.DeadlineStub = nil
	fake.deadlineReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeMilestone) Title() string {
	fake.titleMutex.Lock()
	fake.titleArgsForCall = append(fake.titleArgsForCall, struct{}{})
	fake.recordInvocation("Title", []interface{}{})
	fake.titleMutex.Unlock()
	if fake.TitleStub != nil {
		return fake.TitleStub()
	} else {
		return fake.titleReturns.result1
	}
}

func (fake *FakeMilestone) TitleCallCount() int {
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	return len(fake.titleArgsForCall)
}

func (fake *FakeMilestone) TitleReturns(result1 string) {
	fake.TitleStub = nil
	fake.titleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMilestone) String() string {
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

func (fake *FakeMilestone) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeMilestone) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMilestone) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	fake.issuesOpenMutex.RLock()
	defer fake.issuesOpenMutex.RUnlock()
	fake.issuesClosedMutex.RLock()
	defer fake.issuesClosedMutex.RUnlock()
	fake.deadlineMutex.RLock()
	defer fake.deadlineMutex.RUnlock()
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMilestone) recordInvocation(key string, args []interface{}) {
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

var _ milestone.Milestone = new(FakeMilestone)
