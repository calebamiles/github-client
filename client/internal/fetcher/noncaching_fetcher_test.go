package fetcher_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"

	"github.com/calebamiles/github-client/client/internal/fetcher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	numberOfPagesToServe = 2
)

var _ = Describe("DefaultFetcher", func() {
	It("uses a pagination function to crawl through reachable pages from a base URL", func() {
		handler := newPageSender()

		s := httptest.NewServer(handler)
		defer s.Close()

		f := fetcher.DefaultFetcher{
			Paginate: fakePaginator,
		}

		fetchedPages, err := f.Fetch(s.URL)
		Expect(err).ToNot(HaveOccurred())

		serverHits := atomic.LoadUint32(&handler.requests)
		Expect(serverHits).To(BeEquivalentTo(numberOfPagesToServe))

		testPages := []testPage{}
		err = json.Unmarshal(fetchedPages, &testPages)
		Expect(err).ToNot(HaveOccurred())

		Expect(testPages).To(HaveLen(numberOfPagesToServe))
		Expect(testPages[0].Content).To(Equal("there have been: 1 requests"))
		Expect(testPages[1].Content).To(Equal("there have been: 2 requests"))
	})

	It("returns an error for non 200 status", func() {
		handler := httpStatus(http.StatusTooManyRequests)

		s := httptest.NewServer(handler)
		defer s.Close()

		f := fetcher.DefaultFetcher{
			Paginate: fakePaginator,
		}

		serverURL, err := url.Parse(s.URL)
		Expect(err).ToNot(HaveOccurred())

		_, err = f.Fetch(serverURL.String())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("got unexpected status code: 429"))
	})
})

const (
	nextPageHeaderKey = "nextPage"
)

type testPage struct {
	Content string `json:"content"`
}

type multiplePageSender struct {
	requests       uint32
	pathsRequested map[string]bool
}

func newPageSender() *multiplePageSender {
	return &multiplePageSender{
		0,
		make(map[string]bool),
	}
}

func (handler *multiplePageSender) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	atomic.AddUint32(&handler.requests, 1)

	// the client is supposed` to be following a new link so return immediately if that isn't true
	if handler.pathsRequested[req.RequestURI] {
		return
	}

	handler.pathsRequested[req.RequestURI] = true

	if atomic.LoadUint32(&handler.requests) < numberOfPagesToServe {
		w.Header().Set(nextPageHeaderKey, fmt.Sprintf("http://%s/%d", req.Host, atomic.LoadUint32(&handler.requests)))
	}

	responseString := fmt.Sprintf("there have been: %d requests", atomic.LoadUint32(&handler.requests))
	pageContent := []testPage{testPage{Content: responseString}}

	pageBytes, err := json.Marshal(pageContent)
	Expect(err).ToNot(HaveOccurred())

	w.Write(pageBytes)
}

func fakePaginator(r *http.Response) ([]byte, string, error) {
	defer r.Body.Close()
	nextPageURL := r.Header.Get(nextPageHeaderKey)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, "", err
	}

	return body, nextPageURL, nil
}

type httpStatus int

func (s httpStatus) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(int(s))
}
