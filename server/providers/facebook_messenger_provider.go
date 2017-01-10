package providers

import (
	"github.com/maciekmm/messenger-platform-go-sdk/template"
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

func (messengerProvider *FacebookMessengerProvider) SendQuickReplyMessage(externalUserID, text string, quickReplies map[string]string) (err error) {
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

func (messengerProvider *FacebookMessengerProvider) SendImage(externalUserID, imageURL string) (err error) {
	recipient := messenger.Recipient{ID: externalUserID}
	message := messenger.SendMessage{}
	messageQuery := messenger.MessageQuery{Recipient: recipient, Message: message}

	messageQuery.Image(imageURL)
	_, err = messengerProvider.messengerClient.SendMessage(messageQuery)
	return err
}

func (messengerProvider *FacebookMessengerProvider) SendGenericTemplate(externalUserID string, titleToPhotoURL map[string]string) (err error) {
	messageQuery := messenger.MessageQuery{}
	messageQuery.RecipientID(externalUserID)

	for title, photoURL := range titleToPhotoURL {
		messageQuery.Template(template.GenericTemplate{Title: title,
			ImageURL: photoURL,
			Buttons: []template.Button{
				{
					Type:    template.ButtonTypePostback,
					Payload: "test_" + title,
					Title:   "מרכיבים",
				},
			},
		})
	}

	_, err = messengerProvider.messengerClient.SendMessage(messageQuery)
	return err
}
