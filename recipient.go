package instabot

// Recipient defines instagram user with instagram_user_id.
// Recipient of a message or action.
type Recipient struct {
	ID string `json:"id"`
}

// PrivateReplyRecipient recipients defines recipient for private replies.
type PrivateReplyRecipient struct {
	CommentID string `json:"comment_id"`
}
