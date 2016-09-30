package prs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pull Requests Suite")
}
