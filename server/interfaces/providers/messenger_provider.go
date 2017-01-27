package providers

type IMessengerProvider interface {
	SendImage(externalUserID, imageURL string) error
	SendSimpleMessage(externalUserID, text string) error
	SendQuickReplyMessage(externalUserID, text string, quickReplies map[string]string) error
	SendGenericTemplate(externalUserID string, titleToPhotoURL map[string]string) error
}
