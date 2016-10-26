// This file was generated by counterfeiter
package clientfakes

import (
	"sync"
	"time"

	"github.com/calebamiles/github-client/client"
	"github.com/calebamiles/github-client/commits"
	"github.com/calebamiles/github-client/issues"
	"github.com/calebamiles/github-client/prs"
)

type FakeClient struct {
	FetchCommitsToPathSinceStub        func(string, time.Time) ([]commits.Commit, error)
	fetchCommitsToPathSinceMutex       sync.RWMutex
	fetchCommitsToPathSinceArgsForCall []struct {
		arg1 string
		arg2 time.Time
	}
	fetchCommitsToPathSinceReturns struct {
		result1 []commits.Commit
		result2 error
	}
	FetchCommitsWithCommentsToPathSinceStub        func(string, time.Time) ([]commits.Commit, error)
	fetchCommitsWithCommentsToPathSinceMutex       sync.RWMutex
	fetchCommitsWithCommentsToPathSinceArgsForCall []struct {
		arg1 string
		arg2 time.Time
	}
	fetchCommitsWithCommentsToPathSinceReturns struct {
		result1 []commits.Commit
		result2 error
	}
	FetchIssuesSinceStub        func(time.Time) ([]issues.Issue, error)
	fetchIssuesSinceMutex       sync.RWMutex
	fetchIssuesSinceArgsForCall []struct {
		arg1 time.Time
	}
	fetchIssuesSinceReturns struct {
		result1 []issues.Issue
		result2 error
	}
	FetchPullRequestsSinceStub        func(time.Time) ([]prs.PullRequest, error)
	fetchPullRequestsSinceMutex       sync.RWMutex
	fetchPullRequestsSinceArgsForCall []struct {
		arg1 time.Time
	}
	fetchPullRequestsSinceReturns struct {
		result1 []prs.PullRequest
		result2 error
	}
	FetchPageStub        func(string) ([]byte, error)
	fetchPageMutex       sync.RWMutex
	fetchPageArgsForCall []struct {
		arg1 string
	}
	fetchPageReturns struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) FetchCommitsToPathSince(arg1 string, arg2 time.Time) ([]commits.Commit, error) {
	fake.fetchCommitsToPathSinceMutex.Lock()
	fake.fetchCommitsToPathSinceArgsForCall = append(fake.fetchCommitsToPathSinceArgsForCall, struct {
		arg1 string
		arg2 time.Time
	}{arg1, arg2})
	fake.recordInvocation("FetchCommitsToPathSince", []interface{}{arg1, arg2})
	fake.fetchCommitsToPathSinceMutex.Unlock()
	if fake.FetchCommitsToPathSinceStub != nil {
		return fake.FetchCommitsToPathSinceStub(arg1, arg2)
	} else {
		return fake.fetchCommitsToPathSinceReturns.result1, fake.fetchCommitsToPathSinceReturns.result2
	}
}

func (fake *FakeClient) FetchCommitsToPathSinceCallCount() int {
	fake.fetchCommitsToPathSinceMutex.RLock()
	defer fake.fetchCommitsToPathSinceMutex.RUnlock()
	return len(fake.fetchCommitsToPathSinceArgsForCall)
}

func (fake *FakeClient) FetchCommitsToPathSinceArgsForCall(i int) (string, time.Time) {
	fake.fetchCommitsToPathSinceMutex.RLock()
	defer fake.fetchCommitsToPathSinceMutex.RUnlock()
	return fake.fetchCommitsToPathSinceArgsForCall[i].arg1, fake.fetchCommitsToPathSinceArgsForCall[i].arg2
}

func (fake *FakeClient) FetchCommitsToPathSinceReturns(result1 []commits.Commit, result2 error) {
	fake.FetchCommitsToPathSinceStub = nil
	fake.fetchCommitsToPathSinceReturns = struct {
		result1 []commits.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FetchCommitsWithCommentsToPathSince(arg1 string, arg2 time.Time) ([]commits.Commit, error) {
	fake.fetchCommitsWithCommentsToPathSinceMutex.Lock()
	fake.fetchCommitsWithCommentsToPathSinceArgsForCall = append(fake.fetchCommitsWithCommentsToPathSinceArgsForCall, struct {
		arg1 string
		arg2 time.Time
	}{arg1, arg2})
	fake.recordInvocation("FetchCommitsWithCommentsToPathSince", []interface{}{arg1, arg2})
	fake.fetchCommitsWithCommentsToPathSinceMutex.Unlock()
	if fake.FetchCommitsWithCommentsToPathSinceStub != nil {
		return fake.FetchCommitsWithCommentsToPathSinceStub(arg1, arg2)
	} else {
		return fake.fetchCommitsWithCommentsToPathSinceReturns.result1, fake.fetchCommitsWithCommentsToPathSinceReturns.result2
	}
}

func (fake *FakeClient) FetchCommitsWithCommentsToPathSinceCallCount() int {
	fake.fetchCommitsWithCommentsToPathSinceMutex.RLock()
	defer fake.fetchCommitsWithCommentsToPathSinceMutex.RUnlock()
	return len(fake.fetchCommitsWithCommentsToPathSinceArgsForCall)
}

func (fake *FakeClient) FetchCommitsWithCommentsToPathSinceArgsForCall(i int) (string, time.Time) {
	fake.fetchCommitsWithCommentsToPathSinceMutex.RLock()
	defer fake.fetchCommitsWithCommentsToPathSinceMutex.RUnlock()
	return fake.fetchCommitsWithCommentsToPathSinceArgsForCall[i].arg1, fake.fetchCommitsWithCommentsToPathSinceArgsForCall[i].arg2
}

func (fake *FakeClient) FetchCommitsWithCommentsToPathSinceReturns(result1 []commits.Commit, result2 error) {
	fake.FetchCommitsWithCommentsToPathSinceStub = nil
	fake.fetchCommitsWithCommentsToPathSinceReturns = struct {
		result1 []commits.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FetchIssuesSince(arg1 time.Time) ([]issues.Issue, error) {
	fake.fetchIssuesSinceMutex.Lock()
	fake.fetchIssuesSinceArgsForCall = append(fake.fetchIssuesSinceArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	fake.recordInvocation("FetchIssuesSince", []interface{}{arg1})
	fake.fetchIssuesSinceMutex.Unlock()
	if fake.FetchIssuesSinceStub != nil {
		return fake.FetchIssuesSinceStub(arg1)
	} else {
		return fake.fetchIssuesSinceReturns.result1, fake.fetchIssuesSinceReturns.result2
	}
}

func (fake *FakeClient) FetchIssuesSinceCallCount() int {
	fake.fetchIssuesSinceMutex.RLock()
	defer fake.fetchIssuesSinceMutex.RUnlock()
	return len(fake.fetchIssuesSinceArgsForCall)
}

func (fake *FakeClient) FetchIssuesSinceArgsForCall(i int) time.Time {
	fake.fetchIssuesSinceMutex.RLock()
	defer fake.fetchIssuesSinceMutex.RUnlock()
	return fake.fetchIssuesSinceArgsForCall[i].arg1
}

func (fake *FakeClient) FetchIssuesSinceReturns(result1 []issues.Issue, result2 error) {
	fake.FetchIssuesSinceStub = nil
	fake.fetchIssuesSinceReturns = struct {
		result1 []issues.Issue
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FetchPullRequestsSince(arg1 time.Time) ([]prs.PullRequest, error) {
	fake.fetchPullRequestsSinceMutex.Lock()
	fake.fetchPullRequestsSinceArgsForCall = append(fake.fetchPullRequestsSinceArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	fake.recordInvocation("FetchPullRequestsSince", []interface{}{arg1})
	fake.fetchPullRequestsSinceMutex.Unlock()
	if fake.FetchPullRequestsSinceStub != nil {
		return fake.FetchPullRequestsSinceStub(arg1)
	} else {
		return fake.fetchPullRequestsSinceReturns.result1, fake.fetchPullRequestsSinceReturns.result2
	}
}

func (fake *FakeClient) FetchPullRequestsSinceCallCount() int {
	fake.fetchPullRequestsSinceMutex.RLock()
	defer fake.fetchPullRequestsSinceMutex.RUnlock()
	return len(fake.fetchPullRequestsSinceArgsForCall)
}

func (fake *FakeClient) FetchPullRequestsSinceArgsForCall(i int) time.Time {
	fake.fetchPullRequestsSinceMutex.RLock()
	defer fake.fetchPullRequestsSinceMutex.RUnlock()
	return fake.fetchPullRequestsSinceArgsForCall[i].arg1
}

func (fake *FakeClient) FetchPullRequestsSinceReturns(result1 []prs.PullRequest, result2 error) {
	fake.FetchPullRequestsSinceStub = nil
	fake.fetchPullRequestsSinceReturns = struct {
		result1 []prs.PullRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FetchPage(arg1 string) ([]byte, error) {
	fake.fetchPageMutex.Lock()
	fake.fetchPageArgsForCall = append(fake.fetchPageArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FetchPage", []interface{}{arg1})
	fake.fetchPageMutex.Unlock()
	if fake.FetchPageStub != nil {
		return fake.FetchPageStub(arg1)
	} else {
		return fake.fetchPageReturns.result1, fake.fetchPageReturns.result2
	}
}

func (fake *FakeClient) FetchPageCallCount() int {
	fake.fetchPageMutex.RLock()
	defer fake.fetchPageMutex.RUnlock()
	return len(fake.fetchPageArgsForCall)
}

func (fake *FakeClient) FetchPageArgsForCall(i int) string {
	fake.fetchPageMutex.RLock()
	defer fake.fetchPageMutex.RUnlock()
	return fake.fetchPageArgsForCall[i].arg1
}

func (fake *FakeClient) FetchPageReturns(result1 []byte, result2 error) {
	fake.FetchPageStub = nil
	fake.fetchPageReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchCommitsToPathSinceMutex.RLock()
	defer fake.fetchCommitsToPathSinceMutex.RUnlock()
	fake.fetchCommitsWithCommentsToPathSinceMutex.RLock()
	defer fake.fetchCommitsWithCommentsToPathSinceMutex.RUnlock()
	fake.fetchIssuesSinceMutex.RLock()
	defer fake.fetchIssuesSinceMutex.RUnlock()
	fake.fetchPullRequestsSinceMutex.RLock()
	defer fake.fetchPullRequestsSinceMutex.RUnlock()
	fake.fetchPageMutex.RLock()
	defer fake.fetchPageMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ client.Client = new(FakeClient)
