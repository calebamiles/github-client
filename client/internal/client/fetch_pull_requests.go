package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/calebamiles/github-client/prs"
	"github.com/calebamiles/github-client/prs/pr"
)

// FetchPullRequestsSince returns all pull requests against the repository, optionally since a non zero time
func (c *DefaultClient) FetchPullRequests() ([]pr.PullRequest, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/pullRequests", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	pullRequestPages, _, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	// make processing these PRs a little bit easier to do
	var jsonPRs []json.RawMessage
	err = json.Unmarshal(pullRequestPages, &jsonPRs)
	if err != nil {
		return nil, err
	}

	basicPullRequests, err := prs.NewBasicPullRequests(jsonPRs)
	/*
		TODO: think carefully about what we do here next
					1. we need to fetch several URLs to build the richer data model,
					   do we do the fetching here?


		one option:
		pullRequests, err := prs.NewPullRequets(prStore, fetcher, basicPullRequests)

		the unknown is where the NewPullRequests function will live because it would
		be somewhat unforuntate to have the prs package learn about our model caching,
		maybe a reasonable pattern would be to have internal/prs know about the datastore




	*/

	//

	return nil, nil
}
