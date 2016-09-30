package milestone_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMilestone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Milestone Suite")
}
