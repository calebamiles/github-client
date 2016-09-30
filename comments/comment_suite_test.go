package comments_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestComment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Comments Suite")
}
