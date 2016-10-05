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

var _ = Describe("FetchPullRequestsSince", func() {
	It("returns a slice of PullRequests with comments", func() {
		now := time.Now()
		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchStub = func(urlString string) ([][]byte, error) {
			if strings.HasSuffix(urlString, "comments") {
				return commentsPagesStub, nil
			}

			return pullRequestsPagesStub, nil
		}

		c := &client.DefaultClient{
			Fetcher: fetcher,
		}

		pullRequests, err := c.FetchPullRequestsSince(now)
		Expect(err).ToNot(HaveOccurred())
		Expect(pullRequests).To(HaveLen(1))

		pullRequest := pullRequests[0]
		Expect(pullRequest.Author()).To(Equal("dgoodwin"))

		comments := pullRequest.Comments()
		Expect(comments).To(HaveLen(1))

		comment := comments[0]
		Expect(comment.Author()).To(Equal("k8s-ci-robot"))

	})

	It("passes the correct URL to the fetcher", func() {
		repoName := "test-repo-name"
		repoOwner := "test-repo-owner"
		now := time.Now()

		expectedPath := fmt.Sprintf("/repos/%s/%s/pullRequests", repoOwner, repoName)

		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchReturns(nil, nil)

		c := &client.DefaultClient{
			Fetcher:   fetcher,
			RepoName:  repoName,
			RepoOwner: repoOwner,
		}

		_, err := c.FetchPullRequestsSince(now)
		Expect(err).ToNot(HaveOccurred())

		urlString := fetcher.FetchArgsForCall(0)
		u, err := url.Parse(urlString)
		Expect(err).ToNot(HaveOccurred())

		query := u.Query()
		Expect(query.Get("per_page")).To(Equal(client.NumberOfPagesToRequest))
		Expect(query.Get("since")).To(Equal(now.Format(client.DateFormat)))
		Expect(u.Host).To(Equal("api.github.com"))
		Expect(u.Path).To(Equal(expectedPath))
	})

	It("only adds since if non zero", func() {
		emptyTime := time.Time{}
		fetcher := &fetcherfakes.FakeFetcher{}
		fetcher.FetchReturns(nil, nil)

		c := &client.DefaultClient{
			Fetcher: fetcher,
		}

		_, err := c.FetchPullRequestsSince(emptyTime)
		Expect(err).ToNot(HaveOccurred())

		urlString := fetcher.FetchArgsForCall(0)
		u, err := url.Parse(urlString)
		Expect(err).ToNot(HaveOccurred())

		query := u.Query()
		queryMap := map[string][]string(query)

		_, ok := queryMap["since"]
		Expect(ok).To(BeFalse())
	})
})
