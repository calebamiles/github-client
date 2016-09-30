package comments_test

import (
	"encoding/json"

	"github.com/calebamiles/github-client/comments"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const commentsStub = `
[
  {
    "url": "https://api.github.com/repos/kubernetes/kubernetes/issues/comments/250834951",
    "html_url": "https://github.com/kubernetes/kubernetes/pull/33854#issuecomment-250834951",
    "issue_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33854",
    "id": 250834951,
    "user": {
      "login": "k8s-ci-robot",
      "id": 20407524,
      "avatar_url": "https://avatars.githubusercontent.com/u/20407524?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/k8s-ci-robot",
      "html_url": "https://github.com/k8s-ci-robot",
      "followers_url": "https://api.github.com/users/k8s-ci-robot/followers",
      "following_url": "https://api.github.com/users/k8s-ci-robot/following{/other_user}",
      "gists_url": "https://api.github.com/users/k8s-ci-robot/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/k8s-ci-robot/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/k8s-ci-robot/subscriptions",
      "organizations_url": "https://api.github.com/users/k8s-ci-robot/orgs",
      "repos_url": "https://api.github.com/users/k8s-ci-robot/repos",
      "events_url": "https://api.github.com/users/k8s-ci-robot/events{/privacy}",
      "received_events_url": "https://api.github.com/users/k8s-ci-robot/received_events",
      "type": "User",
      "site_admin": false
    },
    "created_at": "2016-09-30T19:43:21Z",
    "updated_at": "2016-09-30T19:43:21Z",
    "body": "Jenkins Kubemark GCE e2e [**failed**](https://k8s-gubernator.appspot.com/build/kubernetes-jenkins/pr-logs/pull/33854/kubernetes-pull-build-test-kubemark-e2e-gce/1953/) for commit 5cba1d98929ff5928a2bf2ea29a6edfdea2e9152. [Full PR test history](http://pr-test.k8s.io/33854). Please help us cut down flakes by linking to an [open flake issue](https://github.com/kubernetes/kubernetes/issues?q=is:issue+label:kind/flake+is:open) when you hit one in your PR."
  }
]
`

var _ = Describe("building comments from JSON", func() {
	Describe("New", func() {
		It("returns a slice of comments from JSON", func() {
			rawJSON := json.RawMessage(commentsStub)
			cs, err := comments.New(rawJSON)
			Expect(err).ToNot(HaveOccurred())
			Expect(cs).To(HaveLen(1))
			Expect(cs[0].Author()).To(Equal("k8s-ci-robot"))
			Expect(cs[0].Body()).To(ContainSubstring("Jenkins Kubemark GCE e2e"))
		})
	})
})
