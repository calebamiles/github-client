package cache_test

import (
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("PageCache", func() {
	Describe("Open", func() {
		It("attempts to open or create a database", func() {
			Fail("untested")

			By("timing out")

			By("being idempotent")
		})
	})

	Describe("Close", func() {
		It("closes the database", func() {
			Fail("untested")
		})
	})

	Describe("AddPage", func() {
		It("adds a new page to the cache by URL and cache key", func() {
			Fail("untested")

			By("purging previously cached pages for the URL")
		})
	})

	Describe("KeyForPage", func() {
		It("returns the stored cache key if it exists for a given URL", func() {
			Fail("untested")
		})
	})

	Describe("FetchPageByKey", func() {
		It("returns a cached page body for a given cache key", func() {
			Fail("untested")
		})
	})
})
