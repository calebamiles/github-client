package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FetchPullRequestsSince", func() {
	It("returns a slice of PullRequests", func() {
		Fail("test not updated")
		// now := time.Now()
		// fetcher := &fetcherfakes.FakeFetcher{}
		// fetcher.FetchStub = func(urlString string) ([]byte, string, error) {
		// 	if strings.HasSuffix(urlString, "comments") {
		// 		return commentsPagesStub, "", nil
		// 	}
		//
		// 	return pullRequestsPagesStub, "", nil
		// }
		//
		// c := &client.DefaultClient{
		// 	Fetcher: fetcher,
		// }
		//
		// pullRequests, err := c.FetchPullRequestsSince(now)
		// Expect(err).ToNot(HaveOccurred())
		// Expect(pullRequests).To(HaveLen(1))
		//
		// pullRequest := pullRequests[0]
		// Expect(pullRequest.Author()).To(Equal("dgoodwin"))
	})

	It("passes the correct URL to the fetcher", func() {
		Fail("test not updated")

		// 	repoName := "test-repo-name"
		// 	repoOwner := "test-repo-owner"
		// 	now := time.Now()
		//
		// 	expectedPath := fmt.Sprintf("/repos/%s/%s/pullRequests", repoOwner, repoName)
		//
		// 	fetcher := &fetcherfakes.FakeFetcher{}
		//
		// 	c := &client.DefaultClient{
		// 		Fetcher:   fetcher,
		// 		RepoName:  repoName,
		// 		RepoOwner: repoOwner,
		// 	}
		//
		// 	c.FetchPullRequestsSince(now)
		// 	urlString := fetcher.FetchArgsForCall(0)
		// 	u, err := url.Parse(urlString)
		// 	Expect(err).ToNot(HaveOccurred())
		//
		// 	query := u.Query()
		// 	Expect(query.Get("per_page")).To(Equal(client.NumberOfPagesToRequest))
		// 	Expect(query.Get("since")).To(Equal(now.Format(client.DateFormat)))
		// 	Expect(u.Host).To(Equal("api.github.com"))
		// 	Expect(u.Path).To(Equal(expectedPath))
		// })
		//
		// It("only adds since if non zero", func() {
		// 	emptyTime := time.Time{}
		// 	fetcher := &fetcherfakes.FakeFetcher{}
		//
		// 	c := &client.DefaultClient{
		// 		Fetcher: fetcher,
		// 	}
		//
		// 	c.FetchPullRequestsSince(emptyTime)
		// 	urlString := fetcher.FetchArgsForCall(0)
		// 	u, err := url.Parse(urlString)
		// 	Expect(err).ToNot(HaveOccurred())
		//
		// 	query := u.Query()
		// 	queryMap := map[string][]string(query)
		//
		// 	_, ok := queryMap["since"]
		// 	Expect(ok).To(BeFalse())
	})
})