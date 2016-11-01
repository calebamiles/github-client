package serialization

// import (
// 	"time"
//
// 	"github.com/calebamiles/github-client/comments"
// 	"github.com/calebamiles/github-client/prs/pr"
//
// 	"github.com/golang/protobuf/proto"
// 	"github.com/golang/protobuf/ptypes"
// )
//
// func MarshalPullReqest(pull pr.PullRequest) ([]byte, error) {
//
// 	var serializableCommits []*Commit
// 	modelCommits := pull.Commits()
// 	for i := range modelCommits {
//
// 		var serializableModelCommitComments []*Comment
// 		modelCommitComments := modelCommits[i].Comments()
// 		for j := range modelCommitComments {
// 			serializableModelCommitComments = append(serializableModelCommitComments, &Comment{
// 				Body:   modelCommitComments[j].Body(),
// 				Author: modelCommitComments[j].Author(),
// 			})
// 		}
//
// 		commitDate, err := ptypes.TimestampProto(modelCommits[i].Date())
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		serializableCommits = append(serializableCommits, &Commit{
// 			Sha:      modelCommits[i].SHA(),
// 			Date:     commitDate,
// 			Comments: serializableModelCommitComments,
// 		})
// 	}
//
// 	createdAtDate, err := ptypes.TimestampProto(pull.CreatedAt())
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	updatedAtDate, err := ptypes.TimestampProto(pull.UpdatedAt())
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	closedAtDate, err := ptypes.TimestampProto(pull.ClosedAt())
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	mergedAtDate, err := ptypes.TimestampProto(pull.MergedAt())
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var serializableComments []*Comment
// 	modelComments := pull.Comments()
// 	for i := range modelComments {
// 		commentCreatedAtDate, err := ptypes.TimestampProto(modelComments[i].CreatedAt())
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		commentUpdatedAtDate, err := ptypes.TimestampProto(modelComments[i].UpdatedAt())
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		serializableComments = append(serializableComments, &Comment{
// 			Body:      modelComments[i].Body(),
// 			Author:    modelComments[i].Author(),
// 			CreatedAt: commentCreatedAtDate,
// 			UpdatedAt: commentUpdatedAtDate,
// 		})
// 	}
//
// 	var serializableLabels []*Label
// 	modelLabels := pull.Labels()
// 	for i := range modelLabels {
// 		serializableLabels = append(serializableLabels, &Label{
// 			Name: modelLabels[i].Name(),
// 		})
// 	}
//
// 	var serializableFileChanges []*FileChange
// 	modelFileChanges := pull.FilesChanged()
// 	for i := range modelFileChanges {
// 		serializableFileChanges = append(serializableFileChanges, &FileChange{
// 			FileSha:      modelFileChanges[i].FileSHA(),
// 			Filename:     modelFileChanges[i].Filename(),
// 			NumAdditions: int32(modelFileChanges[i].NumAdditions()),
// 			NumChanges:   int32(modelFileChanges[i].NumChanges()),
// 			NumDeletions: int32(modelFileChanges[i].NumDeletions()),
// 		})
// 	}
//
// 	serializablePR := &PullRequest{
// 		Id:           pull.ID(),
// 		Author:       pull.Author(),
// 		Body:         pull.Body(),
// 		Title:        pull.Title(),
// 		Open:         pull.Open(),
// 		Merged:       pull.Merged(),
// 		CreatedAt:    createdAtDate,
// 		UpdatedAt:    updatedAtDate,
// 		ClosedAt:     closedAtDate,
// 		MergedAt:     mergedAtDate,
// 		Reviewers:    pull.Reviewers(),
// 		Commmits:     serializableCommits,
// 		Labels:       serializableLabels,
// 		FilesChanged: serializableFileChanges,
// 	}
//
// 	return proto.Marshal(serializablePR)
// }
//
// func UnmarshalPullRequest(pullRequestBytes []byte) (pr.PullRequest, error) {
// 	serializablePR := &PullRequest{}
// 	prModel = pr.PullRequest
//
// 	err := proto.Unmarshal(pullRequestBytes, serializablePR)
// 	if err != nil {
// 		return nil, err
// 	}
//
// }
//
// func (p *PullRequest) ID() string          { return p.Id }
// func (p *PullRequest) Author() string      { return p.Author }
// func (p *PullRequest) Body() string        { return p.Body }
// func (p *PullRequest) Title() string       { return p.Title }
// func (p *PullRequest) Open() bool          { return p.Open }
// func (p *PullRequest) Merged() bool        { return p.Merged }
// func (p *PullRequest) Reviewers() []string { return p.Reviewers }
//
// func (p *PullRequest) CreatedAt() time.Time {
// 	ts := p.GetCreatedAt()
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the PullRequest
// 		return time.Time{}
// 	}
//
// 	return t
// }
//
// func (p *PullRequest) UpdatedAt() time.Time {
// 	ts := p.GetUpdatedAt()
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the PullRequest
// 		return time.Time{}
// 	}
//
// 	return t
// }
//
// func (p *PullRequest) ClosedAt() time.Time {
// 	ts := p.GetClosedAt()
// 	if ts == nil {
// 		return time.Time{}
// 	}
//
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the PullRequest
// 		return time.Time{}
// 	}
//
// 	return t
// }
//
// func (p *PullRequest) MergedAt() time.Time {
// 	ts := p.GetMergedAt()
// 	if ts == nil {
// 		return time.Time{}
// 	}
//
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the PullRequest
// 		return time.Time{}
// 	}
//
// 	return t
// }
//
// func (c *Comment) Author() string { return c.Author }
// func (c *Comment) Body() string   { return c.Body }
//
// func (c *Comment) CreatedAt() time.Time {
// 	ts := c.GetCreatedAt()
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the Comment
// 		return time.Time{}
// 	}
//
// 	return t
// }
//
// func (c *Comment) UpdatedAt() time.Time {
// 	ts := c.GetUpdatedAt()
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the Comment
// 		return time.Time{}
// 	}
//
// 	return t
// }
//
// func (m *Milestone) Open() bool           { return m.Open }
// func (m *Milestone) Description() string  { return m.Description }
// func (m *Milestone) Title() string        { return m.Title }
// func (m *Milestone) NumIssuesOpen() int   { return int(m.NumIssuesOpen) }
// func (m *Milestone) NumIssuesClosed() int { return int(m.NumIssuesOpen) }
//
// func (m *Milestone) Deadline() time.Time {
// 	ts := m.GetDeadline()
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the Milestone
// 		return time.Time{}
// 	}
//
// 	return t
// }
//
// func (c *Commit) SHA() string                  { return c.Sha }
// func (c *Commit) Author() string               { return c.Author }
// func (c *Commit) ParentSHAs() []string         { return c.ParentShas }
// func (c *Commit) Comments() []comments.Comment { return c.GetComments() }
//
// func (c *Commit) Date() time.Time {
// 	ts := c.GetDate()
// 	t, err := ptypes.Timestamp(ts)
// 	if err != nil {
// 		// that this time is valid should be checked before returning the Milestone
// 		return time.Time{}
// 	}
//
// 	return t
// }
