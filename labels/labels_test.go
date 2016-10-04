package labels_test

import (
	"github.com/calebamiles/github-client/labels"
	"github.com/calebamiles/github-client/labels/labelsfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// ensure our fakes can be used by consumers if desired
var _ labels.Label = &labelsfakes.FakeLabel{}

var _ = Describe("building labels from JSON", func() {
	Describe("New", func() {
		It("builds a slice of Label from raw JSON", func() {
			lbls, err := labels.New([]byte(labelsStub))
			Expect(err).ToNot(HaveOccurred())
			Expect(lbls).To(HaveLen(4))
		})

		It("returns an empty slice of Label when passed an empty slice of bytes", func() {
			lbls, err := labels.New([]byte{})
			Expect(err).ToNot(HaveOccurred())
			Expect(lbls).To(BeEmpty())
		})
	})
})

const labelsStub = `
[
 {
	 "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/area/cluster-federation",
	 "name": "area/cluster-federation",
	 "color": "fad8c7"
 },
 {
	 "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/component/controller-manager",
	 "name": "component/controller-manager",
	 "color": "0052cc"
 },
 {
	 "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/priority/P1",
	 "name": "priority/P1",
	 "color": "eb6420"
 },
 {
	 "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/team/ux",
	 "name": "team/ux",
	 "color": "d2b48c"
 }
]
`
