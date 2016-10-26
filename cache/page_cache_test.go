package cache_test

import (
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/calebamiles/github-client/cache"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PageCache", func() {
	Describe("AddPage", func() {
		It("adds a new page to the cache by URL and cache key", func() {
			fakeCacheKey := "fake-etag-for-page"
			fakeURL := "fake-url"
			fakePage := []byte("fake-web-page-response-body")

			dbName := "test-db"
			dbDir, err := ioutil.TempDir("", dbName)
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dbDir)

			dbPath := path.Join(dbDir, dbName)

			pageCache := cache.New(dbPath)
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())

			err = pageCache.AddPage(fakeURL, fakeCacheKey, fakePage)
			Expect(err).ToNot(HaveOccurred())

			fetchedBody, err := pageCache.FetchPageByKey(fakeCacheKey)
			Expect(err).ToNot(HaveOccurred())

			Expect(fetchedBody).To(Equal(fakePage))
		})

		It("purges previously cached pages for the URL", func() {
			firstCacheKey := "first-fake-etag-for-page"
			secondCacheKey := "second-fake-etag-for-page"
			fakeURL := "fake-url"
			fakePage := []byte("fake-web-page-response-body")

			dbName := "test-db"
			dbDir, err := ioutil.TempDir("", dbName)
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dbDir)

			dbPath := path.Join(dbDir, dbName)

			pageCache := cache.New(dbPath)
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())

			err = pageCache.AddPage(fakeURL, firstCacheKey, fakePage)
			Expect(err).ToNot(HaveOccurred())

			err = pageCache.AddPage(fakeURL, secondCacheKey, fakePage)
			Expect(err).ToNot(HaveOccurred())

			_, err = pageCache.FetchPageByKey(firstCacheKey)
			Expect(err).To(MatchError("page not found in cache"))
		})
	})

	Describe("KeyForPage", func() {
		It("returns the stored cache key if it exists for a given URL", func() {
			fakeCacheKey := "fake-etag-for-page"
			fakeURL := "fake-url"
			fakePage := []byte("fake-web-page-response-body")

			dbName := "test-db"
			dbDir, err := ioutil.TempDir("", dbName)
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dbDir)

			dbPath := path.Join(dbDir, dbName)

			pageCache := cache.New(dbPath)
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())

			err = pageCache.AddPage(fakeURL, fakeCacheKey, fakePage)
			Expect(err).ToNot(HaveOccurred())

			fetchedCacheKey, err := pageCache.KeyForPage(fakeURL)
			Expect(err).ToNot(HaveOccurred())
			Expect(fetchedCacheKey).To(Equal(fakeCacheKey))
		})

		It("does not return an error if a cache key was not found", func() {
			fakeURL := "fake-url"

			dbName := "test-db"
			dbDir, err := ioutil.TempDir("", dbName)
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dbDir)

			dbPath := path.Join(dbDir, dbName)

			pageCache := cache.New(dbPath)
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())

			cacheKey, err := pageCache.KeyForPage(fakeURL)
			Expect(err).ToNot(HaveOccurred())
			Expect(cacheKey).To(BeEmpty())
		})
	})

	Describe("FetchPageByKey", func() {
		It("returns a cached page body for a given cache key", func() {
			fakeCacheKey := "fake-etag-for-page"
			fakeURL := "fake-url"
			fakePage := []byte("fake-web-page-response-body")

			dbName := "test-db"
			dbDir, err := ioutil.TempDir("", dbName)
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dbDir)

			dbPath := path.Join(dbDir, dbName)

			pageCache := cache.New(dbPath)
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())

			err = pageCache.AddPage(fakeURL, fakeCacheKey, fakePage)
			Expect(err).ToNot(HaveOccurred())

			fetchedPage, err := pageCache.FetchPageByKey(fakeCacheKey)
			Expect(err).ToNot(HaveOccurred())

			Expect(fetchedPage).To(Equal(fakePage))
		})
	})

	Describe("Open", func() {
		It("attempts to open or create a database", func() {
			dbName := "test-db"
			dbDir, err := ioutil.TempDir("", dbName)
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dbDir)

			dbPath := path.Join(dbDir, dbName)

			pageCache := cache.New(dbPath)
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())

			_, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 10 * time.Millisecond})
			Expect(err).To(MatchError("timeout"))

			By("being idempotent")
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Close", func() {
		It("closes the database", func() {
			dbName := "test-db"
			dbDir, err := ioutil.TempDir("", dbName)
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dbDir)

			dbPath := path.Join(dbDir, dbName)

			pageCache := cache.New(dbPath)
			err = pageCache.Open()
			Expect(err).ToNot(HaveOccurred())

			err = pageCache.Close()
			Expect(err).ToNot(HaveOccurred())

			db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 10 * time.Millisecond})
			Expect(err).ToNot(HaveOccurred())

			err = db.Close()
			Expect(err).ToNot(HaveOccurred())

			By("being idempotent")
			err = db.Close()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
