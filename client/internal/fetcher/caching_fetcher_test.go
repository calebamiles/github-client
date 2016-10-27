package fetcher_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/calebamiles/github-client/cache/cachefakes"
	"github.com/calebamiles/github-client/client/fetcher/fetcherfakes"
	"github.com/calebamiles/github-client/client/internal/fetcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CachingFetcher", func() {
	Describe("Fetch", func() {
		Context("when the page exists in the cache", func() {
			It("returns the cached page", func() {
				handler := etag(fakeEtag)

				s := httptest.NewServer(handler)
				defer s.Close()

				fakeCache := &cachefakes.FakePage{}
				fakeCache.KeyForPageReturns(fakeEtag, nil)

				fakeFetcher := &fetcherfakes.FakeFetcher{}

				f := fetcher.DefaultCachingFetcher{
					Fetcher: fakeFetcher,
					Cache:   fakeCache,
				}

				_, err := f.Fetch(s.URL)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeFetcher.FetchCallCount()).To(BeZero())
			})
		})

		Context("when the page does not exist in the cache", func() {
			It("uses a fetcher to fetch the page and caches it", func() {
				handler := etag(fakeEtag)

				s := httptest.NewServer(handler)
				defer s.Close()

				fakeCache := &cachefakes.FakePage{}
				fakeCache.KeyForPageReturns("some-other-etag", nil)

				fakeFetcher := &fetcherfakes.FakeFetcher{}

				f := fetcher.DefaultCachingFetcher{
					Fetcher: fakeFetcher,
					Cache:   fakeCache,
				}

				_, err := f.Fetch(s.URL)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeFetcher.FetchCallCount()).ToNot(BeZero())
			})
		})
	})
})

const fakeEtag = "fake-etag"

type etag string

func (handler etag) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	suppliedEtag := req.Header.Get(fetcher.CacheRequestHeader)

	if suppliedEtag == string(handler) {
		w.WriteHeader(http.StatusNotModified)
		w.Header().Set(fetcher.CacheResponseHeader, fakeEtag)
		return
	}

	pageBytes, err := json.Marshal([]struct{}{})
	Expect(err).ToNot(HaveOccurred())

	w.WriteHeader(http.StatusOK)
	w.Header().Set(fetcher.CacheResponseHeader, string(handler))
	w.Write(pageBytes)
}
