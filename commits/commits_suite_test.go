package commits_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCommit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commits Suite")
}
