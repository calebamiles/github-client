package client_test

import (
	"github.com/calebamiles/github-client/client/fetcher/fetcherfakes"
	"github.com/calebamiles/github-client/client/internal/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FetchPages", func() {
	It("delgates to a Fetcher", func() {
		testURL := "https://www.example.com/"
		fetcher := &fetcherfakes.FakeFetcher{}

		c := &client.DefaultClient{
			Fetcher: fetcher,
		}

		_, err := c.FetchPages(testURL)
		Expect(err).ToNot(HaveOccurred())

		calledURL := fetcher.FetchArgsForCall(0)
		Expect(calledURL).To(Equal(testURL))
	})
})
