package commits_test

import (
	"github.com/calebamiles/github-client/commits"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("building commits from JSON", func() {
	Describe("New", func() {
		It("returns a slice of basic commits from JSON", func() {
			rawJSON := []byte(commitsStub)
			cs, err := commits.New(rawJSON)
			Expect(err).ToNot(HaveOccurred())

			Expect(cs).To(HaveLen(1))

			Expect(cs[0].SHA()).To(Equal("20e764ab5d4b64e39f16a6bb441b52563757e156"))
			Expect(cs[0].ParentSHAs()).To(ConsistOf("b840a837c5dc01007d6771a3b349975562124c48", "edcf97db1dfeafe973f7e2ad203957b07f69b198"))
			Expect(cs[0].CommentsURL()).To(Equal("https://api.github.com/repos/kubernetes/kubernetes/commits/20e764ab5d4b64e39f16a6bb441b52563757e156/comments"))
			Expect(cs[0].Author().GitHubID()).To(Equal("k8s-merge-robot"))
			Expect(cs[0].Author().Email()).To(Equal("k8s-merge-robot@users.noreply.github.com"))
			Expect(cs[0].Author().Name()).To(Equal("Kubernetes Submit Queue"))
		})
	})
})

const commitsStub = `
[
  {
    "sha": "20e764ab5d4b64e39f16a6bb441b52563757e156",
    "commit": {
      "author": {
        "name": "Kubernetes Submit Queue",
        "email": "k8s-merge-robot@users.noreply.github.com",
        "date": "2016-09-30T21:54:13Z"
      },
      "committer": {
        "name": "GitHub",
        "email": "noreply@github.com",
        "date": "2016-09-30T21:54:13Z"
      },
      "message": "Merge pull request #33848 from mtaufen/fix-configure-helper\n\nAutomatic merge from submit-queue\n\nCorrect env var name in configure-helper",
      "tree": {
        "sha": "3d482c44910474f26ad184d426c052dad3d5d8ec",
        "url": "https://api.github.com/repos/kubernetes/kubernetes/git/trees/3d482c44910474f26ad184d426c052dad3d5d8ec"
      },
      "url": "https://api.github.com/repos/kubernetes/kubernetes/git/commits/20e764ab5d4b64e39f16a6bb441b52563757e156",
      "comment_count": 0
    },
    "url": "https://api.github.com/repos/kubernetes/kubernetes/commits/20e764ab5d4b64e39f16a6bb441b52563757e156",
    "html_url": "https://github.com/kubernetes/kubernetes/commit/20e764ab5d4b64e39f16a6bb441b52563757e156",
    "comments_url": "https://api.github.com/repos/kubernetes/kubernetes/commits/20e764ab5d4b64e39f16a6bb441b52563757e156/comments",
    "author": {
      "login": "k8s-merge-robot",
      "id": 13653959,
      "avatar_url": "https://avatars.githubusercontent.com/u/13653959?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/k8s-merge-robot",
      "html_url": "https://github.com/k8s-merge-robot",
      "followers_url": "https://api.github.com/users/k8s-merge-robot/followers",
      "following_url": "https://api.github.com/users/k8s-merge-robot/following{/other_user}",
      "gists_url": "https://api.github.com/users/k8s-merge-robot/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/k8s-merge-robot/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/k8s-merge-robot/subscriptions",
      "organizations_url": "https://api.github.com/users/k8s-merge-robot/orgs",
      "repos_url": "https://api.github.com/users/k8s-merge-robot/repos",
      "events_url": "https://api.github.com/users/k8s-merge-robot/events{/privacy}",
      "received_events_url": "https://api.github.com/users/k8s-merge-robot/received_events",
      "type": "User",
      "site_admin": false
    },
    "committer": {
      "login": "web-flow",
      "id": 19864447,
      "avatar_url": "https://avatars.githubusercontent.com/u/19864447?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/web-flow",
      "html_url": "https://github.com/web-flow",
      "followers_url": "https://api.github.com/users/web-flow/followers",
      "following_url": "https://api.github.com/users/web-flow/following{/other_user}",
      "gists_url": "https://api.github.com/users/web-flow/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/web-flow/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/web-flow/subscriptions",
      "organizations_url": "https://api.github.com/users/web-flow/orgs",
      "repos_url": "https://api.github.com/users/web-flow/repos",
      "events_url": "https://api.github.com/users/web-flow/events{/privacy}",
      "received_events_url": "https://api.github.com/users/web-flow/received_events",
      "type": "User",
      "site_admin": false
    },
    "parents": [
      {
        "sha": "b840a837c5dc01007d6771a3b349975562124c48",
        "url": "https://api.github.com/repos/kubernetes/kubernetes/commits/b840a837c5dc01007d6771a3b349975562124c48",
        "html_url": "https://github.com/kubernetes/kubernetes/commit/b840a837c5dc01007d6771a3b349975562124c48"
      },
      {
        "sha": "edcf97db1dfeafe973f7e2ad203957b07f69b198",
        "url": "https://api.github.com/repos/kubernetes/kubernetes/commits/edcf97db1dfeafe973f7e2ad203957b07f69b198",
        "html_url": "https://github.com/kubernetes/kubernetes/commit/edcf97db1dfeafe973f7e2ad203957b07f69b198"
      }
    ]
  }
]
`
