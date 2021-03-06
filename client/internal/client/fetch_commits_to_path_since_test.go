package client_test

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/calebamiles/github-client/client/fetcher/fetcherfakes"
	"github.com/calebamiles/github-client/client/internal/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FetchCommitsWithCommentsToPathSince", func() {
	It("returns a slice of commits with comments", func() {
		now := time.Now()
		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchStub = func(urlString string) ([]byte, error) {
			if strings.HasSuffix(urlString, "comments") {
				return commentsPagesStub, nil
			}

			return commitsPagesStub, nil
		}

		c := &client.DefaultClient{
			Fetcher: fetcher,
		}

		commits, err := c.FetchCommitsWithCommentsToPathSince("", now)
		Expect(err).ToNot(HaveOccurred())
		Expect(commits).To(HaveLen(1))

		commit := commits[0]
		Expect(commit.Author().GitHubID()).To(Equal("k8s-merge-robot"))

		comments := commit.Comments()
		Expect(comments).To(HaveLen(1))

		comment := comments[0]
		Expect(comment.Author()).To(Equal("k8s-ci-robot"))
	})
})

var _ = Describe("FetchCommitsToPathSince", func() {
	It("returns a slice of commits without comments", func() {
		now := time.Now()
		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchStub = func(urlString string) ([]byte, error) {
			if strings.HasSuffix(urlString, "comments") {
				return commentsPagesStub, nil
			}

			return commitsPagesStub, nil
		}

		c := &client.DefaultClient{
			Fetcher: fetcher,
		}

		commits, err := c.FetchCommitsToPathSince("", now)
		Expect(err).ToNot(HaveOccurred())
		Expect(commits).To(HaveLen(1))

		commit := commits[0]
		Expect(commit.Author().GitHubID()).To(Equal("k8s-merge-robot"))

		comments := commit.Comments()
		Expect(comments).To(BeEmpty())
	})
})

var _ = Describe("shared behavior", func() {
	It("passes the correct URL to the fetcher", func() {
		repoName := "test-repo-name"
		repoOwner := "test-repo-owner"
		testPath := "test-path"
		now := time.Now()

		expectedPath := fmt.Sprintf("/repos/%s/%s/commits", repoOwner, repoName)

		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchReturns(nil, nil)

		c := &client.DefaultClient{
			Fetcher:   fetcher,
			RepoName:  repoName,
			RepoOwner: repoOwner,
		}

		c.FetchCommitsToPathSince(testPath, now)
		urlString := fetcher.FetchArgsForCall(0)
		u, err := url.Parse(urlString)
		Expect(err).ToNot(HaveOccurred())

		query := u.Query()
		Expect(query.Get("per_page")).To(Equal(client.NumberOfPagesToRequest))
		Expect(query.Get("path")).To(Equal(testPath))
		Expect(query.Get("since")).To(Equal(now.Format(client.DateFormat)))
		Expect(u.Host).To(Equal("api.github.com"))
		Expect(u.Path).To(Equal(expectedPath))
	})

	It("only adds a path when non empty", func() {
		now := time.Now()
		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchReturns(nil, nil)

		c := &client.DefaultClient{
			Fetcher: fetcher,
		}

		c.FetchCommitsToPathSince("", now)
		urlString := fetcher.FetchArgsForCall(0)
		u, err := url.Parse(urlString)
		Expect(err).ToNot(HaveOccurred())

		query := u.Query()
		queryMap := map[string][]string(query)

		_, ok := queryMap["path"]
		Expect(ok).To(BeFalse())
	})

	It("only adds since if non zero", func() {
		emptyTime := time.Time{}
		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchReturns(nil, nil)

		c := &client.DefaultClient{
			Fetcher: fetcher,
		}

		c.FetchCommitsToPathSince("", emptyTime)
		urlString := fetcher.FetchArgsForCall(0)
		u, err := url.Parse(urlString)
		Expect(err).ToNot(HaveOccurred())

		query := u.Query()
		queryMap := map[string][]string(query)

		_, ok := queryMap["since"]
		Expect(ok).To(BeFalse())
	})
})
