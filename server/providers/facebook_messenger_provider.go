package providers

import (
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
)

type FacebookMessengerProvider struct {
	messengerClient *messenger.Messenger
}

func NewFacebookMessengerProvider(messengerClient *messenger.Messenger) *FacebookMessengerProvider {
	return &FacebookMessengerProvider{messengerClient: messengerClient}
}

func (messengerProvider *FacebookMessengerProvider) SendSimpleMessage(externalUserID, text string) (err error) {
	_, err = messengerProvider.messengerClient.SendSimpleMessage(externalUserID, text)
	return err
}

func (messengerProvider *FacebookMessengerProvider) SendQuickReplyMessage(externalUserID, text string, quickReplies map[string]string) (err error){
	recipient := messenger.Recipient{ID: externalUserID}
	message := messenger.SendMessage{Text: text}
	messageQuery := messenger.MessageQuery{Recipient: recipient, Message: message}
	for title, payload := range quickReplies {
		err = messageQuery.QuickReply(title, payload)
		if err != nil {
			return err
		}
	}
	_, err = messengerProvider.messengerClient.SendMessage(messageQuery)
	return err
}
