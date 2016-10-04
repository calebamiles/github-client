// This file was generated by counterfeiter
package prsfakes

import (
	"sync"
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/labels"
	"github.com/calebamiles/github-client/milestone"
	"github.com/calebamiles/github-client/prs"
)

type FakePullRequest struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
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
	TitleStub        func() string
	titleMutex       sync.RWMutex
	titleArgsForCall []struct{}
	titleReturns     struct {
		result1 string
	}
	CommentsURLStub        func() string
	commentsURLMutex       sync.RWMutex
	commentsURLArgsForCall []struct{}
	commentsURLReturns     struct {
		result1 string
	}
	OpenStub        func() bool
	openMutex       sync.RWMutex
	openArgsForCall []struct{}
	openReturns     struct {
		result1 bool
	}
	MergedStub        func() bool
	mergedMutex       sync.RWMutex
	mergedArgsForCall []struct{}
	mergedReturns     struct {
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
	ReviewersStub        func() []string
	reviewersMutex       sync.RWMutex
	reviewersArgsForCall []struct{}
	reviewersReturns     struct {
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
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct{}
	stringReturns     struct {
		result1 string
	}
	CommentsStub        func() []comments.Comment
	commentsMutex       sync.RWMutex
	commentsArgsForCall []struct{}
	commentsReturns     struct {
		result1 []comments.Comment
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePullRequest) ID() string {
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

func (fake *FakePullRequest) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakePullRequest) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) Author() string {
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

func (fake *FakePullRequest) AuthorCallCount() int {
	fake.authorMutex.RLock()
	defer fake.authorMutex.RUnlock()
	return len(fake.authorArgsForCall)
}

func (fake *FakePullRequest) AuthorReturns(result1 string) {
	fake.AuthorStub = nil
	fake.authorReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) Body() string {
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

func (fake *FakePullRequest) BodyCallCount() int {
	fake.bodyMutex.RLock()
	defer fake.bodyMutex.RUnlock()
	return len(fake.bodyArgsForCall)
}

func (fake *FakePullRequest) BodyReturns(result1 string) {
	fake.BodyStub = nil
	fake.bodyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) Title() string {
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

func (fake *FakePullRequest) TitleCallCount() int {
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	return len(fake.titleArgsForCall)
}

func (fake *FakePullRequest) TitleReturns(result1 string) {
	fake.TitleStub = nil
	fake.titleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) CommentsURL() string {
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

func (fake *FakePullRequest) CommentsURLCallCount() int {
	fake.commentsURLMutex.RLock()
	defer fake.commentsURLMutex.RUnlock()
	return len(fake.commentsURLArgsForCall)
}

func (fake *FakePullRequest) CommentsURLReturns(result1 string) {
	fake.CommentsURLStub = nil
	fake.commentsURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) Open() bool {
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

func (fake *FakePullRequest) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakePullRequest) OpenReturns(result1 bool) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePullRequest) Merged() bool {
	fake.mergedMutex.Lock()
	fake.mergedArgsForCall = append(fake.mergedArgsForCall, struct{}{})
	fake.recordInvocation("Merged", []interface{}{})
	fake.mergedMutex.Unlock()
	if fake.MergedStub != nil {
		return fake.MergedStub()
	} else {
		return fake.mergedReturns.result1
	}
}

func (fake *FakePullRequest) MergedCallCount() int {
	fake.mergedMutex.RLock()
	defer fake.mergedMutex.RUnlock()
	return len(fake.mergedArgsForCall)
}

func (fake *FakePullRequest) MergedReturns(result1 bool) {
	fake.MergedStub = nil
	fake.mergedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePullRequest) Labels() []labels.Label {
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

func (fake *FakePullRequest) LabelsCallCount() int {
	fake.labelsMutex.RLock()
	defer fake.labelsMutex.RUnlock()
	return len(fake.labelsArgsForCall)
}

func (fake *FakePullRequest) LabelsReturns(result1 []labels.Label) {
	fake.LabelsStub = nil
	fake.labelsReturns = struct {
		result1 []labels.Label
	}{result1}
}

func (fake *FakePullRequest) Milestone() milestone.Milestone {
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

func (fake *FakePullRequest) MilestoneCallCount() int {
	fake.milestoneMutex.RLock()
	defer fake.milestoneMutex.RUnlock()
	return len(fake.milestoneArgsForCall)
}

func (fake *FakePullRequest) MilestoneReturns(result1 milestone.Milestone) {
	fake.MilestoneStub = nil
	fake.milestoneReturns = struct {
		result1 milestone.Milestone
	}{result1}
}

func (fake *FakePullRequest) Reviewers() []string {
	fake.reviewersMutex.Lock()
	fake.reviewersArgsForCall = append(fake.reviewersArgsForCall, struct{}{})
	fake.recordInvocation("Reviewers", []interface{}{})
	fake.reviewersMutex.Unlock()
	if fake.ReviewersStub != nil {
		return fake.ReviewersStub()
	} else {
		return fake.reviewersReturns.result1
	}
}

func (fake *FakePullRequest) ReviewersCallCount() int {
	fake.reviewersMutex.RLock()
	defer fake.reviewersMutex.RUnlock()
	return len(fake.reviewersArgsForCall)
}

func (fake *FakePullRequest) ReviewersReturns(result1 []string) {
	fake.ReviewersStub = nil
	fake.reviewersReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakePullRequest) CreatedAt() time.Time {
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

func (fake *FakePullRequest) CreatedAtCallCount() int {
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	return len(fake.createdAtArgsForCall)
}

func (fake *FakePullRequest) CreatedAtReturns(result1 time.Time) {
	fake.CreatedAtStub = nil
	fake.createdAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakePullRequest) UpdatedAt() time.Time {
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

func (fake *FakePullRequest) UpdatedAtCallCount() int {
	fake.updatedAtMutex.RLock()
	defer fake.updatedAtMutex.RUnlock()
	return len(fake.updatedAtArgsForCall)
}

func (fake *FakePullRequest) UpdatedAtReturns(result1 time.Time) {
	fake.UpdatedAtStub = nil
	fake.updatedAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakePullRequest) String() string {
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

func (fake *FakePullRequest) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakePullRequest) StringReturns(result1 string) {
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) Comments() []comments.Comment {
	fake.commentsMutex.Lock()
	fake.commentsArgsForCall = append(fake.commentsArgsForCall, struct{}{})
	fake.recordInvocation("Comments", []interface{}{})
	fake.commentsMutex.Unlock()
	if fake.CommentsStub != nil {
		return fake.CommentsStub()
	} else {
		return fake.commentsReturns.result1
	}
}

func (fake *FakePullRequest) CommentsCallCount() int {
	fake.commentsMutex.RLock()
	defer fake.commentsMutex.RUnlock()
	return len(fake.commentsArgsForCall)
}

func (fake *FakePullRequest) CommentsReturns(result1 []comments.Comment) {
	fake.CommentsStub = nil
	fake.commentsReturns = struct {
		result1 []comments.Comment
	}{result1}
}

func (fake *FakePullRequest) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.authorMutex.RLock()
	defer fake.authorMutex.RUnlock()
	fake.bodyMutex.RLock()
	defer fake.bodyMutex.RUnlock()
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	fake.commentsURLMutex.RLock()
	defer fake.commentsURLMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.mergedMutex.RLock()
	defer fake.mergedMutex.RUnlock()
	fake.labelsMutex.RLock()
	defer fake.labelsMutex.RUnlock()
	fake.milestoneMutex.RLock()
	defer fake.milestoneMutex.RUnlock()
	fake.reviewersMutex.RLock()
	defer fake.reviewersMutex.RUnlock()
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	fake.updatedAtMutex.RLock()
	defer fake.updatedAtMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	fake.commentsMutex.RLock()
	defer fake.commentsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePullRequest) recordInvocation(key string, args []interface{}) {
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

var _ prs.PullRequest = new(FakePullRequest)