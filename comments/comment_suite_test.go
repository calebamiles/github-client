package comments_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestComment(t *testing.T) {
	log.SetOutput(GinkgoWriter)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Comments Suite")
}
