package inappropriate_post_report_decision

import "github.com/Nistagram-Organization/nistagram-shared/src/utils/rest_error"

type InappropriatePostReportDecision struct {
	PostID      uint   `json:"post_id"`
	AuthorEmail string `json:"author_email"`
	Delete      bool   `json:"delete"`
	Terminate   bool   `json:"terminate"`
}

func (d *InappropriatePostReportDecision) Validate() rest_error.RestErr {
	if d.Terminate && !d.Delete {
		return rest_error.NewBadRequestError("Cannot terminate author's profile and not delete reported post")
	}
	return nil
}
