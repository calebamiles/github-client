// This file was generated by counterfeiter
package issuesfakes

import (
	"sync"
	"time"

	"github.com/calebamiles/github-client/issues"
	"github.com/calebamiles/github-client/labels"
	"github.com/calebamiles/github-client/milestone"
)

type FakeIssueWithoutComments struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	OpenedByStub        func() string
	openedByMutex       sync.RWMutex
	openedByArgsForCall []struct{}
	openedByReturns     struct {
		result1 string
	}
	BodyStub        func() string
	bodyMutex       sync.RWMutex
	bodyArgsForCall []struct{}
	bodyReturns     struct {
		result1 string
	}
	CommentsURLStub        func() string
	commentsURLMutex       sync.RWMutex
	commentsURLArgsForCall []struct{}
	commentsURLReturns     struct {
		result1 string
	}
	TitleStub        func() string
	titleMutex       sync.RWMutex
	titleArgsForCall []struct{}
	titleReturns     struct {
		result1 string
	}
	OpenStub        func() bool
	openMutex       sync.RWMutex
	openArgsForCall []struct{}
	openReturns     struct {
		result1 bool
	}
	LabelsStub        func() []labels.Label
	labelsMutex       sync.RWMutex
	labelsArgsForCall []struct{}
	labelsReturns     struct {
		result1 []labels.Label
	}
	MilestoneStub        func() milestone.Milestone
	milestoneMutex       sync.RWMutex
	milestoneArgsForCall []struct{}
	milestoneReturns     struct {
		result1 milestone.Milestone
	}
	AssigneesStub        func() []string
	assigneesMutex       sync.RWMutex
	assigneesArgsForCall []struct{}
	assigneesReturns     struct {
		result1 []string
	}
	CreatedAtStub        func() time.Time
	createdAtMutex       sync.RWMutex
	createdAtArgsForCall []struct{}
	createdAtReturns     struct {
		result1 time.Time
	}
	UpdatedAtStub        func() time.Time
	updatedAtMutex       sync.RWMutex
	updatedAtArgsForCall []struct{}
	updatedAtReturns     struct {
		result1 time.Time
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIssueWithoutComments) ID() string {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeIssueWithoutComments) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIssueWithoutComments) OpenedBy() string {
	fake.openedByMutex.Lock()
	fake.openedByArgsForCall = append(fake.openedByArgsForCall, struct{}{})
	fake.recordInvocation("OpenedBy", []interface{}{})
	fake.openedByMutex.Unlock()
	if fake.OpenedByStub != nil {
		return fake.OpenedByStub()
	} else {
		return fake.openedByReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) OpenedByCallCount() int {
	fake.openedByMutex.RLock()
	defer fake.openedByMutex.RUnlock()
	return len(fake.openedByArgsForCall)
}

func (fake *FakeIssueWithoutComments) OpenedByReturns(result1 string) {
	fake.OpenedByStub = nil
	fake.openedByReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIssueWithoutComments) Body() string {
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

func (fake *FakeIssueWithoutComments) BodyCallCount() int {
	fake.bodyMutex.RLock()
	defer fake.bodyMutex.RUnlock()
	return len(fake.bodyArgsForCall)
}

func (fake *FakeIssueWithoutComments) BodyReturns(result1 string) {
	fake.BodyStub = nil
	fake.bodyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIssueWithoutComments) CommentsURL() string {
	fake.commentsURLMutex.Lock()
	fake.commentsURLArgsForCall = append(fake.commentsURLArgsForCall, struct{}{})
	fake.recordInvocation("CommentsURL", []interface{}{})
	fake.commentsURLMutex.Unlock()
	if fake.CommentsURLStub != nil {
		return fake.CommentsURLStub()
	} else {
		return fake.commentsURLReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) CommentsURLCallCount() int {
	fake.commentsURLMutex.RLock()
	defer fake.commentsURLMutex.RUnlock()
	return len(fake.commentsURLArgsForCall)
}

func (fake *FakeIssueWithoutComments) CommentsURLReturns(result1 string) {
	fake.CommentsURLStub = nil
	fake.commentsURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIssueWithoutComments) Title() string {
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

func (fake *FakeIssueWithoutComments) TitleCallCount() int {
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	return len(fake.titleArgsForCall)
}

func (fake *FakeIssueWithoutComments) TitleReturns(result1 string) {
	fake.TitleStub = nil
	fake.titleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIssueWithoutComments) Open() bool {
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

func (fake *FakeIssueWithoutComments) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeIssueWithoutComments) OpenReturns(result1 bool) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeIssueWithoutComments) Labels() []labels.Label {
	fake.labelsMutex.Lock()
	fake.labelsArgsForCall = append(fake.labelsArgsForCall, struct{}{})
	fake.recordInvocation("Labels", []interface{}{})
	fake.labelsMutex.Unlock()
	if fake.LabelsStub != nil {
		return fake.LabelsStub()
	} else {
		return fake.labelsReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) LabelsCallCount() int {
	fake.labelsMutex.RLock()
	defer fake.labelsMutex.RUnlock()
	return len(fake.labelsArgsForCall)
}

func (fake *FakeIssueWithoutComments) LabelsReturns(result1 []labels.Label) {
	fake.LabelsStub = nil
	fake.labelsReturns = struct {
		result1 []labels.Label
	}{result1}
}

func (fake *FakeIssueWithoutComments) Milestone() milestone.Milestone {
	fake.milestoneMutex.Lock()
	fake.milestoneArgsForCall = append(fake.milestoneArgsForCall, struct{}{})
	fake.recordInvocation("Milestone", []interface{}{})
	fake.milestoneMutex.Unlock()
	if fake.MilestoneStub != nil {
		return fake.MilestoneStub()
	} else {
		return fake.milestoneReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) MilestoneCallCount() int {
	fake.milestoneMutex.RLock()
	defer fake.milestoneMutex.RUnlock()
	return len(fake.milestoneArgsForCall)
}

func (fake *FakeIssueWithoutComments) MilestoneReturns(result1 milestone.Milestone) {
	fake.MilestoneStub = nil
	fake.milestoneReturns = struct {
		result1 milestone.Milestone
	}{result1}
}

func (fake *FakeIssueWithoutComments) Assignees() []string {
	fake.assigneesMutex.Lock()
	fake.assigneesArgsForCall = append(fake.assigneesArgsForCall, struct{}{})
	fake.recordInvocation("Assignees", []interface{}{})
	fake.assigneesMutex.Unlock()
	if fake.AssigneesStub != nil {
		return fake.AssigneesStub()
	} else {
		return fake.assigneesReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) AssigneesCallCount() int {
	fake.assigneesMutex.RLock()
	defer fake.assigneesMutex.RUnlock()
	return len(fake.assigneesArgsForCall)
}

func (fake *FakeIssueWithoutComments) AssigneesReturns(result1 []string) {
	fake.AssigneesStub = nil
	fake.assigneesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeIssueWithoutComments) CreatedAt() time.Time {
	fake.createdAtMutex.Lock()
	fake.createdAtArgsForCall = append(fake.createdAtArgsForCall, struct{}{})
	fake.recordInvocation("CreatedAt", []interface{}{})
	fake.createdAtMutex.Unlock()
	if fake.CreatedAtStub != nil {
		return fake.CreatedAtStub()
	} else {
		return fake.createdAtReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) CreatedAtCallCount() int {
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	return len(fake.createdAtArgsForCall)
}

func (fake *FakeIssueWithoutComments) CreatedAtReturns(result1 time.Time) {
	fake.CreatedAtStub = nil
	fake.createdAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeIssueWithoutComments) UpdatedAt() time.Time {
	fake.updatedAtMutex.Lock()
	fake.updatedAtArgsForCall = append(fake.updatedAtArgsForCall, struct{}{})
	fake.recordInvocation("UpdatedAt", []interface{}{})
	fake.updatedAtMutex.Unlock()
	if fake.UpdatedAtStub != nil {
		return fake.UpdatedAtStub()
	} else {
		return fake.updatedAtReturns.result1
	}
}

func (fake *FakeIssueWithoutComments) UpdatedAtCallCount() int {
	fake.updatedAtMutex.RLock()
	defer fake.updatedAtMutex.RUnlock()
	return len(fake.updatedAtArgsForCall)
}

func (fake *FakeIssueWithoutComments) UpdatedAtReturns(result1 time.Time) {
	fake.UpdatedAtStub = nil
	fake.updatedAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeIssueWithoutComments) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.openedByMutex.RLock()
	defer fake.openedByMutex.RUnlock()
	fake.bodyMutex.RLock()
	defer fake.bodyMutex.RUnlock()
	fake.commentsURLMutex.RLock()
	defer fake.commentsURLMutex.RUnlock()
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.labelsMutex.RLock()
	defer fake.labelsMutex.RUnlock()
	fake.milestoneMutex.RLock()
	defer fake.milestoneMutex.RUnlock()
	fake.assigneesMutex.RLock()
	defer fake.assigneesMutex.RUnlock()
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	fake.updatedAtMutex.RLock()
	defer fake.updatedAtMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeIssueWithoutComments) recordInvocation(key string, args []interface{}) {
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

var _ issues.IssueWithoutComments = new(FakeIssueWithoutComments)
