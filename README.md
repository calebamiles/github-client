A very basic client for the GitHub API. Currently basic information relating to

- Pull Requests
- Issues
- Commits

can be obtained. In addition raw JSON can be returned from an GitHub API endpoint
if the caller knows how to properly handle the result. Consumers should only
ever create a client via

```
client.New(repositoryOwner, repositoryName, accessToken)
```

no other way of creating a client.Client is supported. In order to help consumers
test against this libraries fakes generated with [Counterfeiter](https://github.com/maxbrunsfeld/counterfeiter)
have been provided for all major interfaces in `<packagename>/<packageName>fakes`.
This library should be considered to be very experimental and probably should not
be used. Consumers looking for a GitHub client should probably use [the one from GitHub](https://github.com/octokit/go-octokit).
You have been warned!
