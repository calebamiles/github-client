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
	Expected workflow:
		1.  call KeyForPage() -> get cache key (etag)
			1a. if cacheKey != "" make conditional API request
			1b. if 302 (unchanged) call FetchPageByKey() -> response body
			1c. if resource page changed make API request
			1d. call AddPage -> page body cached
		2.  if cacheKey == "" make API request
		  2a. call AddPage -> page body cached
*/

type PageCache interface {
	KeyForPage(pageURL string) (cacheKey string, exists bool, err error)
	FetchPageByKey(cacheKey string) (page []byte, err error)
	AddPage(pageURL string, cacheKey string, page []byte) error
	Open() error
	Close() error
}

// New returns a new PageCache or an error relating to opening the datastore
func New(datastorePath string) PageCache {
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

func (c *pageCache) Close() error {
	var closeErr error

	c.closeOnce.Do(func() {
		closeErr = c.db.Close()
	})

	return closeErr
}

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

		return nil
	})
}

func (c *pageCache) KeyForPage(pageURL string) (string, bool, error) {
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
		return "", false, err
	}

	if cacheKey == "" {
		return "", false, nil
	}

	return cacheKey, true, nil
}

func (c *pageCache) FetchPageByKey(cacheKey string) ([]byte, error) {
	var fetchedPage []byte

	err := c.db.View(func(tx *bolt.Tx) error {
		pages := tx.Bucket([]byte(pagesBucket))
		pageFromCache := pages.Get([]byte(cacheKey))
		if pageFromCache == nil {
			log.Printf("nil page returned from cache for key %s...this should not have happened\n", cacheKey)
			return errors.New("page not found in cache")
		}

		copy(fetchedPage, pageFromCache)
		return nil
	})

	if err != nil {
		log.Printf("failed to fetch cached page for cache key %s: %s\n", cacheKey, err)
		return nil, err
	}

	return fetchedPage, nil
}
