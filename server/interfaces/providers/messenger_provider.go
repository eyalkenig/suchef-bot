package providers

import "github.com/eyalkenig/suchef-bot/server/interfaces/messaging"

type IMessengerProvider interface {
	SendImage(externalUserID, imageURL string) error
	SendSimpleMessage(externalUserID, text string) error
	SendQuickReplyMessage(externalUserID, text string, quickReplies map[string]string) error
	SendGenericTemplate(externalUserID string, quickReplies map[string]string, cards []messaging.ICard) error
}
