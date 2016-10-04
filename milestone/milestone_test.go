package milestone_test

import (
	"github.com/calebamiles/github-client/milestone"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const milestoneStub = `
{
"url": "https://api.github.com/repos/octocat/Hello-World/milestones/1",
		 "html_url": "https://github.com/octocat/Hello-World/milestones/v1.0",
		 "labels_url": "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels",
		 "id": 1002604,
		 "number": 1,
		 "state": "open",
		 "title": "v1.0",
		 "description": "Tracking milestone for version 1.0",
		 "open_issues": 4,
		 "closed_issues": 8,
		 "created_at": "2011-04-10T20:09:31Z",
		 "updated_at": "2014-03-03T18:58:10Z",
		 "closed_at": "2013-02-12T13:22:01Z",
		 "due_on": "2012-10-09T23:39:01Z"
}
`

var _ = Describe("building a milestone from JSON", func() {
	Describe("New", func() {
		It("builds a Milestone from raw JSON", func() {
			m, err := milestone.New([]byte(milestoneStub))
			Expect(err).ToNot(HaveOccurred())

			Expect(m.Open()).To(BeTrue())
			Expect(m.IssuesOpen()).To(Equal(4))
			Expect(m.IssuesClosed()).To(Equal(8))
			Expect(m.Title()).To(Equal("v1.0"))
			Expect(m.Description()).To(Equal("Tracking milestone for version 1.0"))
		})

		It("returns an empty milestone with passed an empty slice of bytes", func() {
			m, err := milestone.New([]byte{})
			Expect(err).ToNot(HaveOccurred())
			Expect(m.Title()).To(BeEmpty())
		})
	})

})
