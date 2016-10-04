package issues_test

import (
	"encoding/json"
	"time"

	"github.com/calebamiles/github-client/issues"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("building issues from JSON", func() {
	It("builds a slice of basic issues from raw JSON", func() {
		ts := []times{}
		err := json.Unmarshal([]byte(issuesStub), &ts)
		Expect(err).ToNot(HaveOccurred())
		Expect(ts).To(HaveLen(2))
		t := ts[0]

		is, err := issues.New([]byte(issuesStub))
		Expect(err).ToNot(HaveOccurred())

		By("filtering out pull requests")
		Expect(is).To(HaveLen(1))

		issue := is[0]

		Expect(issue.Author()).To(Equal("madhusudancs"))
		Expect(issue.Body()).To(ContainSubstring("As described here because the provider"))
		Expect(issue.Milestone().Title()).To(Equal("v1.4"))
		Expect(issue.Title()).To(Equal("New federation deployment mechanism exits with an error for non-GCP clusters."))
		Expect(issue.Open()).To(BeTrue())
		Expect(issue.ID()).To(Equal("33886"))
		Expect(issue.CommentsURL()).To(Equal("https://api.github.com/repos/kubernetes/kubernetes/issues/33886/comments"))
		Expect(issue.Assignees()).To(HaveLen(1))
		Expect(issue.Assignees()).To(ConsistOf("madhusudancs"))
		Expect(issue.CreatedAt()).To(Equal(t.CreatedAt))
		Expect(issue.UpdatedAt()).To(Equal(t.UpdatedAt))

		labelsStrings := []string{}
		Expect(issue.Labels()).To(HaveLen(4))
		for _, l := range issue.Labels() {
			labelsStrings = append(labelsStrings, l.Name())
		}

		Expect(labelsStrings).To(ConsistOf("area/cluster-federation", "component/controller-manager", "priority/P1", "team/ux"))

	})
})

type times struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const issuesStub = `
[
  {
    "url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33886",
    "repository_url": "https://api.github.com/repos/kubernetes/kubernetes",
    "labels_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33886/labels{/name}",
    "comments_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33886/comments",
    "events_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33886/events",
    "html_url": "https://github.com/kubernetes/kubernetes/issues/33886",
    "id": 180477743,
    "number": 33886,
    "title": "New federation deployment mechanism exits with an error for non-GCP clusters.",
    "user": {
      "login": "madhusudancs",
      "id": 10183,
      "avatar_url": "https://avatars.githubusercontent.com/u/10183?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/madhusudancs",
      "html_url": "https://github.com/madhusudancs",
      "followers_url": "https://api.github.com/users/madhusudancs/followers",
      "following_url": "https://api.github.com/users/madhusudancs/following{/other_user}",
      "gists_url": "https://api.github.com/users/madhusudancs/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/madhusudancs/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/madhusudancs/subscriptions",
      "organizations_url": "https://api.github.com/users/madhusudancs/orgs",
      "repos_url": "https://api.github.com/users/madhusudancs/repos",
      "events_url": "https://api.github.com/users/madhusudancs/events{/privacy}",
      "received_events_url": "https://api.github.com/users/madhusudancs/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
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
    ],
    "state": "open",
    "locked": false,
    "assignee": {
      "login": "madhusudancs",
      "id": 10183,
      "avatar_url": "https://avatars.githubusercontent.com/u/10183?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/madhusudancs",
      "html_url": "https://github.com/madhusudancs",
      "followers_url": "https://api.github.com/users/madhusudancs/followers",
      "following_url": "https://api.github.com/users/madhusudancs/following{/other_user}",
      "gists_url": "https://api.github.com/users/madhusudancs/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/madhusudancs/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/madhusudancs/subscriptions",
      "organizations_url": "https://api.github.com/users/madhusudancs/orgs",
      "repos_url": "https://api.github.com/users/madhusudancs/repos",
      "events_url": "https://api.github.com/users/madhusudancs/events{/privacy}",
      "received_events_url": "https://api.github.com/users/madhusudancs/received_events",
      "type": "User",
      "site_admin": false
    },
    "assignees": [
      {
        "login": "madhusudancs",
        "id": 10183,
        "avatar_url": "https://avatars.githubusercontent.com/u/10183?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/madhusudancs",
        "html_url": "https://github.com/madhusudancs",
        "followers_url": "https://api.github.com/users/madhusudancs/followers",
        "following_url": "https://api.github.com/users/madhusudancs/following{/other_user}",
        "gists_url": "https://api.github.com/users/madhusudancs/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/madhusudancs/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/madhusudancs/subscriptions",
        "organizations_url": "https://api.github.com/users/madhusudancs/orgs",
        "repos_url": "https://api.github.com/users/madhusudancs/repos",
        "events_url": "https://api.github.com/users/madhusudancs/events{/privacy}",
        "received_events_url": "https://api.github.com/users/madhusudancs/received_events",
        "type": "User",
        "site_admin": false
      }
    ],
    "milestone": {
      "url": "https://api.github.com/repos/kubernetes/kubernetes/milestones/22",
      "html_url": "https://github.com/kubernetes/kubernetes/milestones/v1.4",
      "labels_url": "https://api.github.com/repos/kubernetes/kubernetes/milestones/22/labels",
      "id": 1777384,
      "number": 22,
      "title": "v1.4",
      "description": "https://github.com/kubernetes/features/blob/master/release-1.4/Release-1.4.md",
      "creator": {
        "login": "bgrant0607",
        "id": 7725777,
        "avatar_url": "https://avatars.githubusercontent.com/u/7725777?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/bgrant0607",
        "html_url": "https://github.com/bgrant0607",
        "followers_url": "https://api.github.com/users/bgrant0607/followers",
        "following_url": "https://api.github.com/users/bgrant0607/following{/other_user}",
        "gists_url": "https://api.github.com/users/bgrant0607/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/bgrant0607/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/bgrant0607/subscriptions",
        "organizations_url": "https://api.github.com/users/bgrant0607/orgs",
        "repos_url": "https://api.github.com/users/bgrant0607/repos",
        "events_url": "https://api.github.com/users/bgrant0607/events{/privacy}",
        "received_events_url": "https://api.github.com/users/bgrant0607/received_events",
        "type": "User",
        "site_admin": false
      },
      "open_issues": 54,
      "closed_issues": 1048,
      "state": "open",
      "created_at": "2016-05-19T18:50:19Z",
      "updated_at": "2016-10-01T20:16:58Z",
      "due_on": "2016-09-20T00:00:00Z",
      "closed_at": null
    },
    "comments": 1,
    "created_at": "2016-10-01T20:16:57Z",
    "updated_at": "2016-10-01T20:33:19Z",
    "closed_at": null,
    "body": "As described here because the provider isn't gce/gke. However, the deployment mechanism itself should work on all the providers. We shouldn't exit, but just branch out to kubeconfig check instead.\r\n\r\ncc @owenhaynes @kubernetes/sig-cluster-federation "
  },
  {
    "url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885",
    "repository_url": "https://api.github.com/repos/kubernetes/kubernetes",
    "labels_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885/labels{/name}",
    "comments_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885/comments",
    "events_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885/events",
    "html_url": "https://github.com/kubernetes/kubernetes/pull/33885",
    "id": 180476023,
    "number": 33885,
    "title": "kubeadm: Pre-pull images to limit time waiting for control plane.",
    "user": {
      "login": "dgoodwin",
      "id": 51265,
      "avatar_url": "https://avatars.githubusercontent.com/u/51265?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/dgoodwin",
      "html_url": "https://github.com/dgoodwin",
      "followers_url": "https://api.github.com/users/dgoodwin/followers",
      "following_url": "https://api.github.com/users/dgoodwin/following{/other_user}",
      "gists_url": "https://api.github.com/users/dgoodwin/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/dgoodwin/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/dgoodwin/subscriptions",
      "organizations_url": "https://api.github.com/users/dgoodwin/orgs",
      "repos_url": "https://api.github.com/users/dgoodwin/repos",
      "events_url": "https://api.github.com/users/dgoodwin/events{/privacy}",
      "received_events_url": "https://api.github.com/users/dgoodwin/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
      {
        "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/cla:%20yes",
        "name": "cla: yes",
        "color": "bfe5bf"
      },
      {
        "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/release-note-none",
        "name": "release-note-none",
        "color": "c2e0c6"
      },
      {
        "url": "https://api.github.com/repos/kubernetes/kubernetes/labels/size/M",
        "name": "size/M",
        "color": "eebb00"
      }
    ],
    "state": "open",
    "locked": false,
    "assignee": {
      "login": "errordeveloper",
      "id": 251467,
      "avatar_url": "https://avatars.githubusercontent.com/u/251467?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/errordeveloper",
      "html_url": "https://github.com/errordeveloper",
      "followers_url": "https://api.github.com/users/errordeveloper/followers",
      "following_url": "https://api.github.com/users/errordeveloper/following{/other_user}",
      "gists_url": "https://api.github.com/users/errordeveloper/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/errordeveloper/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/errordeveloper/subscriptions",
      "organizations_url": "https://api.github.com/users/errordeveloper/orgs",
      "repos_url": "https://api.github.com/users/errordeveloper/repos",
      "events_url": "https://api.github.com/users/errordeveloper/events{/privacy}",
      "received_events_url": "https://api.github.com/users/errordeveloper/received_events",
      "type": "User",
      "site_admin": false
    },
    "assignees": [
      {
        "login": "errordeveloper",
        "id": 251467,
        "avatar_url": "https://avatars.githubusercontent.com/u/251467?v=3",
        "gravatar_id": "",
        "url": "https://api.github.com/users/errordeveloper",
        "html_url": "https://github.com/errordeveloper",
        "followers_url": "https://api.github.com/users/errordeveloper/followers",
        "following_url": "https://api.github.com/users/errordeveloper/following{/other_user}",
        "gists_url": "https://api.github.com/users/errordeveloper/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/errordeveloper/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/errordeveloper/subscriptions",
        "organizations_url": "https://api.github.com/users/errordeveloper/orgs",
        "repos_url": "https://api.github.com/users/errordeveloper/repos",
        "events_url": "https://api.github.com/users/errordeveloper/events{/privacy}",
        "received_events_url": "https://api.github.com/users/errordeveloper/received_events",
        "type": "User",
        "site_admin": false
      }
    ],
    "milestone": null,
    "comments": 3,
    "created_at": "2016-10-01T19:48:29Z",
    "updated_at": "2016-10-01T20:11:44Z",
    "closed_at": null,
    "pull_request": {
      "url": "https://api.github.com/repos/kubernetes/kubernetes/pulls/33885",
      "html_url": "https://github.com/kubernetes/kubernetes/pull/33885",
      "diff_url": "https://github.com/kubernetes/kubernetes/pull/33885.diff",
      "patch_url": "https://github.com/kubernetes/kubernetes/pull/33885.patch"
    },
    "body": "<!--  Thanks for sending a pull request!  Here are some tips for you:\r\n1. If this is your first time, read our contributor guidelines https://github.com/kubernetes/kubernetes/blob/master/CONTRIBUTING.md and developer guide https://github.com/kubernetes/kubernetes/blob/master/docs/devel/development.md\r\n2. If you want *faster* PR reviews, read how: https://github.com/kubernetes/kubernetes/blob/master/docs/devel/faster_reviews.md\r\n3. Follow the instructions for writing a release note: https://github.com/kubernetes/kubernetes/blob/master/docs/devel/pull-requests.md#release-notes\r\n-->\r\n\r\n**What this PR does / why we need it**: kubeadm can sit at the \"waiting for control plane\" for quite awhile leading some users to think something is wrong. Pre-pulling all the images we'll be using in our static pods cuts this time down drastically.\r\n\r\n\r\n**Special notes for your reviewer**: There is a small issue with glog output appearing due to re-use of the dockertools methods to create a client. I have not yet found a way to disable these or alter the log level to avoid them but it may still be possible.\r\n\r\n\r\nUse pre-existing kubernetes docker tools to pull all images we'll be\r\nconfiguring in the static manifests.\r\n\r\nThis should eliminate a large chunk of the time spent at the \"waiting\r\nfor control plane\" state where the user does not know what's happening\r\nunless they're monitoring containers running or kubelet logs.\r\n\r\nThe user will now receive an error quickly if the docker daemon is not\r\nrunning when we go to do a pull.\n\n<!-- Reviewable:start -->\n---\nThis change is [<img src=\"https://reviewable.kubernetes.io/review_button.svg\" height=\"34\" align=\"absmiddle\" alt=\"Reviewable\"/>](https://reviewable.kubernetes.io/reviews/kubernetes/kubernetes/33885)\n<!-- Reviewable:end -->\n"
  }
]
`
