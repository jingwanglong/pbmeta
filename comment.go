package pbmeta

import (
	"strings"

	pbprotos "github.com/davyxu/pbmeta/proto"
)

type CommentMeta struct {
	comment *pbprotos.SourceCodeInfo_Location

	trailing string
	leading  string
}

func (self *CommentMeta) TrailingComment() string {
	return self.trailing

}

func (self *CommentMeta) LeadingComment() string {
	return self.leading

}

func (self *CommentMeta) parse() {
	self.trailing = strings.TrimSpace(self.comment.GetTrailingComments())
	self.leading = strings.TrimSpace(self.comment.GetLeadingComments())

	//log.Printf("<%s> <%s>", self.trailing, self.leading)
}

func newCommentMeta(comment *pbprotos.SourceCodeInfo_Location) CommentMeta {
	cm := CommentMeta{
		comment: comment,
	}

	cm.parse()

	return cm
}
