package store

import (
	"bytes"
	"crypto/md5"
	"log"

	"github.com/boltdb/bolt"
	"github.com/calebamiles/github-client/client/fetcher"
	"github.com/calebamiles/github-client/prs/pr"
)

type PullRequest interface {
	FetchPullRequest(basicPR pr.BasicPullRequest) (pullRequest pr.PullRequest, err error)
}

type pullRequestStore struct {
	db      *bolt.DB
	fetcher fetcher.Fetcher
}

const (
	cacheKeyBucket = "keys"
	prBucket       = "prs"
)

func (s *pullRequestStore) FetchPullRequest(basicPR pr.BasicPullRequest) (pullRequest pr.PullRequest, err error) {
	var fetchedPR pr.PullRequest

	err := s.db.Batch(func(tx *bolt.Tx) error {
		var commitErr error

		keys := tx.Bucket([]byte(cacheKeyBucket))
		pulls := tx.Bucket([]byte(prBucket))

		pageHasher := md5.New()
		prPage, _, commitErr := s.fetcher.Fetch(basicPR.URL())
		if commitErr != nil {
			log.Printf("failed to fetch PR URL: %s: %s", basicPR.URL(), commitErr)
			return commitErr
		}

		_, commitErr = pageHasher.Write(prPage)
		if commitErr != nil {
			return commitErr
		}

		currentPageHash := pageHasher.Sum(nil)
		previousPageHash := keys.Get([]byte(basicPR.ID))

		if bytes.Equal(previousPageHash, currentPageHash) {
			previousPRBytes := pulls.Get([]byte(basicPR.ID))
		}

		return nil
	})

}
