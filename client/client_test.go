package client_test

import (
	"os"

	"github.com/calebamiles/github-client/client"
	internalClient "github.com/calebamiles/github-client/client/internal/client"
	internalFetcher "github.com/calebamiles/github-client/client/internal/fetcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("creating a client", func() {
	Describe("NewNonCachingClient", func() {
		It("returns a ready to use Client", func() {
			fakeRepoName := "fakeRepo"
			fakeOwnerName := "fakeOwner"
			fakeAccessToken := "fakeToken"

			c := client.NewNonCachingClient(fakeOwnerName, fakeRepoName, fakeAccessToken)
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

	Describe("New", func() {
		It("returns a Client which supports caching", func() {
			fakeRepoName := "fakeRepo"
			fakeOwnerName := "fakeOwner"
			fakeAccessToken := "fakeToken"
			fakeCachePath := "/some/fake/path"

			c, err := client.New(fakeOwnerName, fakeRepoName, fakeAccessToken, fakeCachePath)
			Expect(err).ToNot(HaveOccurred())

			defaultClient, ok := c.(*internalClient.DefaultClient)
			Expect(ok).To(BeTrue())
			Expect(defaultClient.Fetcher).ToNot(BeNil())
			Expect(defaultClient.RepoName).To(Equal(fakeRepoName))
			Expect(defaultClient.RepoOwner).To(Equal(fakeOwnerName))

			_, ok = defaultClient.Fetcher.(*internalFetcher.DefaultCachingFetcher)
			Expect(ok).To(BeTrue())
		})

		It("sets a temporary cache path if none was provied", func() {
			fakeRepoName := "fakeRepo"
			fakeOwnerName := "fakeOwner"
			fakeAccessToken := "fakeToken"

			c, err := client.New(fakeOwnerName, fakeRepoName, fakeAccessToken, "")
			Expect(err).ToNot(HaveOccurred())

			defaultClient, ok := c.(*internalClient.DefaultClient)
			Expect(ok).To(BeTrue())

			tempCachePath := defaultClient.TempCacheFilePath
			Expect(tempCachePath).ToNot(BeEmpty())
			_, err = os.Stat(tempCachePath)
			Expect(err).ToNot(HaveOccurred())

			err = defaultClient.Done()
			Expect(err).ToNot(HaveOccurred())

			_, err = os.Stat(tempCachePath)
			Expect(os.IsNotExist(err)).To(BeTrue())
		})
	})
})
