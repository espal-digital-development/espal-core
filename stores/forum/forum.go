package forum

import (
	"strconv"
	"time"
)

// nolint:deadcode
type forumMethods interface {
	TopicsCountAsString() string
	PostsCountAsString() string
}

// Forum database object.
// @synthesize
type Forum struct {
	// TODO :: Track who is online, but on a more global level?
	// TODO :: Forum moderation rights? (post/reply/edit/delete)
	// TODO :: Image/thumbnail linking
	// TODO :: Mechanic for approval before posting/editing/deleting?
	// TODO :: Report Post?
	// TODO :: Permanent/Temporary banning of Users for Forums specific?
	// TODO :: Mark as read table per Forum per Post User
	id                 string
	createdByID        string
	updatedByID        *string
	createdAt          time.Time
	updatedAt          *time.Time
	createdByFirstName *string
	createdBySurname   *string
	updatedByFirstName *string
	updatedBySurname   *string
	active             bool
	sorting            uint
	parentID           *string
	name               string
	topicsCount        uint
	postsCount         uint

	// children []Forum
}

// TableAlias returns the unique resolved table alias for use in queries.
func (f *Forum) TableAlias() string {
	return "fo"
}

// TopicsCountAsString returns the model's TopicsCount as a string.
func (f *Forum) TopicsCountAsString() string {
	return strconv.FormatUint(uint64(f.topicsCount), 10)
}

// PostsCountAsString returns the model's PostsCount as a string.
func (f *Forum) PostsCountAsString() string {
	return strconv.FormatUint(uint64(f.postsCount), 10)
}
