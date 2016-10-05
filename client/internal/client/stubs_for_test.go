package client_test

var commitsPagesStub = [][]byte{[]byte(commitsStub)}

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

var commentsPagesStub = [][]byte{[]byte(commentsStub)}

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

var issuesPagesStub = [][]byte{[]byte(issuesStub)}

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
  }
]
`

var pullRequestsPagesStub = [][]byte{[]byte(pullRequestsStub)}

const pullRequestsStub = `
[
{
	 "url": "https://api.github.com/repos/kubernetes/kubernetes/pulls/33885",
	 "id": 87612855,
	 "html_url": "https://github.com/kubernetes/kubernetes/pull/33885",
	 "diff_url": "https://github.com/kubernetes/kubernetes/pull/33885.diff",
	 "patch_url": "https://github.com/kubernetes/kubernetes/pull/33885.patch",
	 "issue_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885",
	 "number": 33885,
	 "state": "open",
	 "locked": false,
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
	 "body": "<!--  Thanks for sending a pull request!  Here are some tips for you:\r\n1. If this is your first time, read our contributor guidelines https://github.com/kubernetes/kubernetes/blob/master/CONTRIBUTING.md and developer guide https://github.com/kubernetes/kubernetes/blob/master/docs/devel/development.md\r\n2. If you want *faster* PR reviews, read how: https://github.com/kubernetes/kubernetes/blob/master/docs/devel/faster_reviews.md\r\n3. Follow the instructions for writing a release note: https://github.com/kubernetes/kubernetes/blob/master/docs/devel/pull-requests.md#release-notes\r\n-->\r\n\r\n**What this PR does / why we need it**: kubeadm can sit at the \"waiting for control plane\" for quite awhile leading some users to think something is wrong. Pre-pulling all the images we'll be using in our static pods cuts this time down drastically.\r\n\r\n\r\n**Special notes for your reviewer**: There is a small issue with glog output appearing due to re-use of the dockertools methods to create a client. I have not yet found a way to disable these or alter the log level to avoid them but it may still be possible.\r\n\r\n\r\nUse pre-existing kubernetes docker tools to pull all images we'll be\r\nconfiguring in the static manifests.\r\n\r\nThis should eliminate a large chunk of the time spent at the \"waiting\r\nfor control plane\" state where the user does not know what's happening\r\nunless they're monitoring containers running or kubelet logs.\r\n\r\nThe user will now receive an error quickly if the docker daemon is not\r\nrunning when we go to do a pull.\n\n<!-- Reviewable:start -->\n---\nThis change is [<img src=\"https://reviewable.kubernetes.io/review_button.svg\" height=\"34\" align=\"absmiddle\" alt=\"Reviewable\"/>](https://reviewable.kubernetes.io/reviews/kubernetes/kubernetes/33885)\n<!-- Reviewable:end -->\n",
	 "created_at": "2016-10-01T19:48:28Z",
	 "updated_at": "2016-10-01T20:11:44Z",
	 "closed_at": null,
	 "merged_at": null,
	 "merge_commit_sha": "95b79b2595d3fdce69074a76200972deffa2e25b",
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
	 "commits_url": "https://api.github.com/repos/kubernetes/kubernetes/pulls/33885/commits",
	 "review_comments_url": "https://api.github.com/repos/kubernetes/kubernetes/pulls/33885/comments",
	 "review_comment_url": "https://api.github.com/repos/kubernetes/kubernetes/pulls/comments{/number}",
	 "comments_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885/comments",
	 "statuses_url": "https://api.github.com/repos/kubernetes/kubernetes/statuses/9283386ab8cc0ad5c700fdc48a882a39dad07870",
	 "head": {
		 "label": "dgoodwin:kubeadm-pre-pull",
		 "ref": "kubeadm-pre-pull",
		 "sha": "9283386ab8cc0ad5c700fdc48a882a39dad07870",
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
		 "repo": {
			 "id": 65905656,
			 "name": "kubernetes",
			 "full_name": "dgoodwin/kubernetes",
			 "owner": {
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
			 "private": false,
			 "html_url": "https://github.com/dgoodwin/kubernetes",
			 "description": "Production-Grade Container Scheduling and Management",
			 "fork": true,
			 "url": "https://api.github.com/repos/dgoodwin/kubernetes",
			 "forks_url": "https://api.github.com/repos/dgoodwin/kubernetes/forks",
			 "keys_url": "https://api.github.com/repos/dgoodwin/kubernetes/keys{/key_id}",
			 "collaborators_url": "https://api.github.com/repos/dgoodwin/kubernetes/collaborators{/collaborator}",
			 "teams_url": "https://api.github.com/repos/dgoodwin/kubernetes/teams",
			 "hooks_url": "https://api.github.com/repos/dgoodwin/kubernetes/hooks",
			 "issue_events_url": "https://api.github.com/repos/dgoodwin/kubernetes/issues/events{/number}",
			 "events_url": "https://api.github.com/repos/dgoodwin/kubernetes/events",
			 "assignees_url": "https://api.github.com/repos/dgoodwin/kubernetes/assignees{/user}",
			 "branches_url": "https://api.github.com/repos/dgoodwin/kubernetes/branches{/branch}",
			 "tags_url": "https://api.github.com/repos/dgoodwin/kubernetes/tags",
			 "blobs_url": "https://api.github.com/repos/dgoodwin/kubernetes/git/blobs{/sha}",
			 "git_tags_url": "https://api.github.com/repos/dgoodwin/kubernetes/git/tags{/sha}",
			 "git_refs_url": "https://api.github.com/repos/dgoodwin/kubernetes/git/refs{/sha}",
			 "trees_url": "https://api.github.com/repos/dgoodwin/kubernetes/git/trees{/sha}",
			 "statuses_url": "https://api.github.com/repos/dgoodwin/kubernetes/statuses/{sha}",
			 "languages_url": "https://api.github.com/repos/dgoodwin/kubernetes/languages",
			 "stargazers_url": "https://api.github.com/repos/dgoodwin/kubernetes/stargazers",
			 "contributors_url": "https://api.github.com/repos/dgoodwin/kubernetes/contributors",
			 "subscribers_url": "https://api.github.com/repos/dgoodwin/kubernetes/subscribers",
			 "subscription_url": "https://api.github.com/repos/dgoodwin/kubernetes/subscription",
			 "commits_url": "https://api.github.com/repos/dgoodwin/kubernetes/commits{/sha}",
			 "git_commits_url": "https://api.github.com/repos/dgoodwin/kubernetes/git/commits{/sha}",
			 "comments_url": "https://api.github.com/repos/dgoodwin/kubernetes/comments{/number}",
			 "issue_comment_url": "https://api.github.com/repos/dgoodwin/kubernetes/issues/comments{/number}",
			 "contents_url": "https://api.github.com/repos/dgoodwin/kubernetes/contents/{+path}",
			 "compare_url": "https://api.github.com/repos/dgoodwin/kubernetes/compare/{base}...{head}",
			 "merges_url": "https://api.github.com/repos/dgoodwin/kubernetes/merges",
			 "archive_url": "https://api.github.com/repos/dgoodwin/kubernetes/{archive_format}{/ref}",
			 "downloads_url": "https://api.github.com/repos/dgoodwin/kubernetes/downloads",
			 "issues_url": "https://api.github.com/repos/dgoodwin/kubernetes/issues{/number}",
			 "pulls_url": "https://api.github.com/repos/dgoodwin/kubernetes/pulls{/number}",
			 "milestones_url": "https://api.github.com/repos/dgoodwin/kubernetes/milestones{/number}",
			 "notifications_url": "https://api.github.com/repos/dgoodwin/kubernetes/notifications{?since,all,participating}",
			 "labels_url": "https://api.github.com/repos/dgoodwin/kubernetes/labels{/name}",
			 "releases_url": "https://api.github.com/repos/dgoodwin/kubernetes/releases{/id}",
			 "deployments_url": "https://api.github.com/repos/dgoodwin/kubernetes/deployments",
			 "created_at": "2016-08-17T12:26:59Z",
			 "updated_at": "2016-08-17T12:30:41Z",
			 "pushed_at": "2016-10-01T19:40:31Z",
			 "git_url": "git://github.com/dgoodwin/kubernetes.git",
			 "ssh_url": "git@github.com:dgoodwin/kubernetes.git",
			 "clone_url": "https://github.com/dgoodwin/kubernetes.git",
			 "svn_url": "https://github.com/dgoodwin/kubernetes",
			 "homepage": "http://kubernetes.io",
			 "size": 329507,
			 "stargazers_count": 0,
			 "watchers_count": 0,
			 "language": "Go",
			 "has_issues": false,
			 "has_downloads": true,
			 "has_wiki": true,
			 "has_pages": false,
			 "forks_count": 0,
			 "mirror_url": null,
			 "open_issues_count": 0,
			 "forks": 0,
			 "open_issues": 0,
			 "watchers": 0,
			 "default_branch": "master"
		 }
	 },
	 "base": {
		 "label": "kubernetes:master",
		 "ref": "master",
		 "sha": "8cdd526913fe1deeb2921092f67f5ebb3ed98fec",
		 "user": {
			 "login": "kubernetes",
			 "id": 13629408,
			 "avatar_url": "https://avatars.githubusercontent.com/u/13629408?v=3",
			 "gravatar_id": "",
			 "url": "https://api.github.com/users/kubernetes",
			 "html_url": "https://github.com/kubernetes",
			 "followers_url": "https://api.github.com/users/kubernetes/followers",
			 "following_url": "https://api.github.com/users/kubernetes/following{/other_user}",
			 "gists_url": "https://api.github.com/users/kubernetes/gists{/gist_id}",
			 "starred_url": "https://api.github.com/users/kubernetes/starred{/owner}{/repo}",
			 "subscriptions_url": "https://api.github.com/users/kubernetes/subscriptions",
			 "organizations_url": "https://api.github.com/users/kubernetes/orgs",
			 "repos_url": "https://api.github.com/users/kubernetes/repos",
			 "events_url": "https://api.github.com/users/kubernetes/events{/privacy}",
			 "received_events_url": "https://api.github.com/users/kubernetes/received_events",
			 "type": "Organization",
			 "site_admin": false
		 },
		 "repo": {
			 "id": 20580498,
			 "name": "kubernetes",
			 "full_name": "kubernetes/kubernetes",
			 "owner": {
				 "login": "kubernetes",
				 "id": 13629408,
				 "avatar_url": "https://avatars.githubusercontent.com/u/13629408?v=3",
				 "gravatar_id": "",
				 "url": "https://api.github.com/users/kubernetes",
				 "html_url": "https://github.com/kubernetes",
				 "followers_url": "https://api.github.com/users/kubernetes/followers",
				 "following_url": "https://api.github.com/users/kubernetes/following{/other_user}",
				 "gists_url": "https://api.github.com/users/kubernetes/gists{/gist_id}",
				 "starred_url": "https://api.github.com/users/kubernetes/starred{/owner}{/repo}",
				 "subscriptions_url": "https://api.github.com/users/kubernetes/subscriptions",
				 "organizations_url": "https://api.github.com/users/kubernetes/orgs",
				 "repos_url": "https://api.github.com/users/kubernetes/repos",
				 "events_url": "https://api.github.com/users/kubernetes/events{/privacy}",
				 "received_events_url": "https://api.github.com/users/kubernetes/received_events",
				 "type": "Organization",
				 "site_admin": false
			 },
			 "private": false,
			 "html_url": "https://github.com/kubernetes/kubernetes",
			 "description": "Production-Grade Container Scheduling and Management",
			 "fork": false,
			 "url": "https://api.github.com/repos/kubernetes/kubernetes",
			 "forks_url": "https://api.github.com/repos/kubernetes/kubernetes/forks",
			 "keys_url": "https://api.github.com/repos/kubernetes/kubernetes/keys{/key_id}",
			 "collaborators_url": "https://api.github.com/repos/kubernetes/kubernetes/collaborators{/collaborator}",
			 "teams_url": "https://api.github.com/repos/kubernetes/kubernetes/teams",
			 "hooks_url": "https://api.github.com/repos/kubernetes/kubernetes/hooks",
			 "issue_events_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/events{/number}",
			 "events_url": "https://api.github.com/repos/kubernetes/kubernetes/events",
			 "assignees_url": "https://api.github.com/repos/kubernetes/kubernetes/assignees{/user}",
			 "branches_url": "https://api.github.com/repos/kubernetes/kubernetes/branches{/branch}",
			 "tags_url": "https://api.github.com/repos/kubernetes/kubernetes/tags",
			 "blobs_url": "https://api.github.com/repos/kubernetes/kubernetes/git/blobs{/sha}",
			 "git_tags_url": "https://api.github.com/repos/kubernetes/kubernetes/git/tags{/sha}",
			 "git_refs_url": "https://api.github.com/repos/kubernetes/kubernetes/git/refs{/sha}",
			 "trees_url": "https://api.github.com/repos/kubernetes/kubernetes/git/trees{/sha}",
			 "statuses_url": "https://api.github.com/repos/kubernetes/kubernetes/statuses/{sha}",
			 "languages_url": "https://api.github.com/repos/kubernetes/kubernetes/languages",
			 "stargazers_url": "https://api.github.com/repos/kubernetes/kubernetes/stargazers",
			 "contributors_url": "https://api.github.com/repos/kubernetes/kubernetes/contributors",
			 "subscribers_url": "https://api.github.com/repos/kubernetes/kubernetes/subscribers",
			 "subscription_url": "https://api.github.com/repos/kubernetes/kubernetes/subscription",
			 "commits_url": "https://api.github.com/repos/kubernetes/kubernetes/commits{/sha}",
			 "git_commits_url": "https://api.github.com/repos/kubernetes/kubernetes/git/commits{/sha}",
			 "comments_url": "https://api.github.com/repos/kubernetes/kubernetes/comments{/number}",
			 "issue_comment_url": "https://api.github.com/repos/kubernetes/kubernetes/issues/comments{/number}",
			 "contents_url": "https://api.github.com/repos/kubernetes/kubernetes/contents/{+path}",
			 "compare_url": "https://api.github.com/repos/kubernetes/kubernetes/compare/{base}...{head}",
			 "merges_url": "https://api.github.com/repos/kubernetes/kubernetes/merges",
			 "archive_url": "https://api.github.com/repos/kubernetes/kubernetes/{archive_format}{/ref}",
			 "downloads_url": "https://api.github.com/repos/kubernetes/kubernetes/downloads",
			 "issues_url": "https://api.github.com/repos/kubernetes/kubernetes/issues{/number}",
			 "pulls_url": "https://api.github.com/repos/kubernetes/kubernetes/pulls{/number}",
			 "milestones_url": "https://api.github.com/repos/kubernetes/kubernetes/milestones{/number}",
			 "notifications_url": "https://api.github.com/repos/kubernetes/kubernetes/notifications{?since,all,participating}",
			 "labels_url": "https://api.github.com/repos/kubernetes/kubernetes/labels{/name}",
			 "releases_url": "https://api.github.com/repos/kubernetes/kubernetes/releases{/id}",
			 "deployments_url": "https://api.github.com/repos/kubernetes/kubernetes/deployments",
			 "created_at": "2014-06-06T22:56:04Z",
			 "updated_at": "2016-10-02T02:24:19Z",
			 "pushed_at": "2016-10-02T02:23:00Z",
			 "git_url": "git://github.com/kubernetes/kubernetes.git",
			 "ssh_url": "git@github.com:kubernetes/kubernetes.git",
			 "clone_url": "https://github.com/kubernetes/kubernetes.git",
			 "svn_url": "https://github.com/kubernetes/kubernetes",
			 "homepage": "http://kubernetes.io",
			 "size": 575716,
			 "stargazers_count": 17251,
			 "watchers_count": 17251,
			 "language": "Go",
			 "has_issues": true,
			 "has_downloads": true,
			 "has_wiki": true,
			 "has_pages": true,
			 "forks_count": 5597,
			 "mirror_url": null,
			 "open_issues_count": 4710,
			 "forks": 5597,
			 "open_issues": 4710,
			 "watchers": 17251,
			 "default_branch": "master"
		 }
	 },
	 "_links": {
		 "self": {
			 "href": "https://api.github.com/repos/kubernetes/kubernetes/pulls/33885"
		 },
		 "html": {
			 "href": "https://github.com/kubernetes/kubernetes/pull/33885"
		 },
		 "issue": {
			 "href": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885"
		 },
		 "comments": {
			 "href": "https://api.github.com/repos/kubernetes/kubernetes/issues/33885/comments"
		 },
		 "review_comments": {
			 "href": "https://api.github.com/repos/kubernetes/kubernetes/pulls/33885/comments"
		 },
		 "review_comment": {
			 "href": "https://api.github.com/repos/kubernetes/kubernetes/pulls/comments{/number}"
		 },
		 "commits": {
			 "href": "https://api.github.com/repos/kubernetes/kubernetes/pulls/33885/commits"
		 },
		 "statuses": {
			 "href": "https://api.github.com/repos/kubernetes/kubernetes/statuses/9283386ab8cc0ad5c700fdc48a882a39dad07870"
		 }
	 }
 }
]
`
