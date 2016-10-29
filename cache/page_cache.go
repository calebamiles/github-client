package cache

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

/*
	TODO move this cache into an internal package since it's only ever used
	by internal code and probably isn't meaningful to consumers of the library
*/

type Page interface {
	KeyForPage(pageURL string) (cacheKey string, err error)
	FetchPageByKey(cacheKey string) (page []byte, err error)
	AddPage(pageURL string, cacheKey string, page []byte) error
	Open() error
	Close() error
}

// New returns a new Page
func New(datastorePath string) Page {
	return &pageCache{
		dbPath: datastorePath,
	}
}

const (
	pagesBucket = "pages"
	keysBucket  = "keys"
)

type pageCache struct {
	dbPath    string
	db        *bolt.DB
	openOnce  sync.Once
	closeOnce sync.Once
}

// AddPage adds a page to the cache
func (c *pageCache) AddPage(pageURL string, cacheKey string, page []byte) error {
	return c.db.Batch(func(tx *bolt.Tx) error {
		var commitErr error

		keys := tx.Bucket([]byte(keysBucket))
		pages := tx.Bucket([]byte(pagesBucket))

		hash := md5.New()
		_, commitErr = io.WriteString(hash, pageURL)
		if commitErr != nil {
			log.Printf("failed to hash page URL: %s: %s\n", pageURL, commitErr)
			return commitErr
		}

		pageKey := hash.Sum(nil)

		previousCacheKey := keys.Get(pageKey)
		commitErr = keys.Put(pageKey, []byte(cacheKey))
		if commitErr != nil {
			log.Printf("failed to put cacheKey %s for URL: %s: %s\n", cacheKey, pageURL, commitErr)
			return commitErr
		}

		commitErr = pages.Put([]byte(cacheKey), page)
		if commitErr != nil {
			log.Printf("failed to put page body for URL: %s: %s\n", pageURL, commitErr)
			return commitErr
		}

		if previousCacheKey != nil {
			commitErr = pages.Delete(previousCacheKey)
			if commitErr != nil {
				log.Printf("failed to purge previously cached page for URL: %s: %s", pageURL, commitErr)
				return commitErr
			}
		}

		return nil
	})
}

// KeyForPage returns the cache key (etag) associated with the provied URL
func (c *pageCache) KeyForPage(pageURL string) (string, error) {
	var cacheKey string

	err := c.db.View(func(tx *bolt.Tx) error {
		var viewErr error

		keys := tx.Bucket([]byte(keysBucket))

		hash := md5.New()
		_, viewErr = io.WriteString(hash, pageURL)
		if viewErr != nil {
			log.Printf("failed to hash page URL: %s\n", viewErr)
			return viewErr
		}

		pageKey := hash.Sum(nil)

		cacheKey = string(keys.Get(pageKey))
		return nil
	})

	if err != nil {
		log.Printf("failed to fetch cache key for URL %s: %s\n", pageURL, err)
		return "", err
	}

	return cacheKey, nil
}

// FetchPageByKey returns the cached page for a given cache key (etag)
// FetchPageByKey returns an error if there is no cached page for the provided key
func (c *pageCache) FetchPageByKey(cacheKey string) ([]byte, error) {
	var fetchedPage []byte

	err := c.db.View(func(tx *bolt.Tx) error {
		pages := tx.Bucket([]byte(pagesBucket))
		pageFromCache := pages.Get([]byte(cacheKey))
		if pageFromCache == nil {
			log.Printf("nil page returned from cache for key %s...hopefully this is due to cache eviction\n", cacheKey)
			return errors.New("page not found in cache")
		}

		//copy our the fetched data out of the transaction
		fetchedPage = append(fetchedPage, pageFromCache...)
		return nil
	})

	if err != nil {
		log.Printf("failed to fetch cached page for cache key %s: %s\n", cacheKey, err)
		return nil, err
	}

	return fetchedPage, nil
}

// Open opens the underlying datastore at most once
func (c *pageCache) Open() error {
	var openErr error
	var db *bolt.DB

	c.openOnce.Do(func() {
		db, openErr = bolt.Open(c.dbPath, 0600, &bolt.Options{Timeout: 10 * time.Second})
		if openErr != nil {
			return
		}

		openErr = db.Update(func(tx *bolt.Tx) error {
			var updateErr error
			_, updateErr = tx.CreateBucketIfNotExists([]byte(pagesBucket))
			if updateErr != nil {
				return fmt.Errorf("create pages bucket: %s", updateErr)
			}

			_, updateErr = tx.CreateBucketIfNotExists([]byte(keysBucket))
			if updateErr != nil {
				return fmt.Errorf("create keys bucket: %s", updateErr)
			}

			return nil
		})

		if openErr != nil {
			log.Printf("failed to open database: %s\n", openErr)
			return
		}

		c.db = db
	})

	return openErr
}

// Close closes the underlying datastore at most once
func (c *pageCache) Close() error {
	var closeErr error

	// ensure the database was opened at least once to prevent panic
	closeErr = c.Open()
	if closeErr != nil {
		return closeErr
	}

	c.closeOnce.Do(func() {
		closeErr = c.db.Close()
	})

	return closeErr
}
