package prs_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPr(t *testing.T) {
	log.SetOutput(GinkgoWriter)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Pull Requests Suite")
}
