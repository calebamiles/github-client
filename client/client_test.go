package client_test

import (
	"github.com/calebamiles/github-client/client"
	"github.com/calebamiles/github-client/client/internal"

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
			defaultClient, ok := c.(*internal.DefaultClient)
			Expect(ok).To(BeTrue())
			Expect(defaultClient.Fetcher).ToNot(BeNil())
			Expect(defaultClient.RepoName).To(Equal(fakeRepoName))
			Expect(defaultClient.RepoOwner).To(Equal(fakeOwnerName))

			defaultFetcher, ok := defaultClient.Fetcher.(*internal.DefaultFetcher)
			Expect(ok).To(BeTrue())
			Expect(defaultFetcher.AccessToken).To(Equal(fakeAccessToken))
			Expect(defaultFetcher.Paginate).ToNot(BeNil())
		})
	})
})
