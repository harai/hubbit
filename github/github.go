package github

import "strconv"

// IssueAsHashtag converts an issue number to corresponding hashtag.
func IssueAsHashtag(issueNo int) string {
	return "#" + strconv.Itoa(issueNo)
}
