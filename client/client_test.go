package client_test

import (
	"github.com/calebamiles/github-client/client"
	internalClient "github.com/calebamiles/github-client/client/internal/client"
	internalFetcher "github.com/calebamiles/github-client/client/internal/fetcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("creating a client", func() {
	Describe("New", func() {
		It("returns a ready to use Client", func() {
			fakeRepoName := "fakeRepo"
			fakeOwnerName := "fakeOwner"
			fakeAccessToken := "fakeToken"

			c := client.New(fakeOwnerName, fakeRepoName, fakeAccessToken)
			defaultClient, ok := c.(*internalClient.DefaultClient)
			Expect(ok).To(BeTrue())
			Expect(defaultClient.Fetcher).ToNot(BeNil())
			Expect(defaultClient.RepoName).To(Equal(fakeRepoName))
			Expect(defaultClient.RepoOwner).To(Equal(fakeOwnerName))

			defaultFetcher, ok := defaultClient.Fetcher.(*internalFetcher.DefaultFetcher)
			Expect(ok).To(BeTrue())
			Expect(defaultFetcher.AccessToken).To(Equal(fakeAccessToken))
			Expect(defaultFetcher.Paginate).ToNot(BeNil())
		})
	})
})
