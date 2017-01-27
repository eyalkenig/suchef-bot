package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"

	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"github.com/eyalkenig/suchef-bot/server/models"
	concreteProviders "github.com/eyalkenig/suchef-bot/server/providers"
	"strconv"
)

type SuchefServer struct {
	controller      ISuchefController
	adminController ISuchefAdminController
	authorizer      providers.IAuthorizer
	accountID       int64
}

func NewSuchefServer(accountID int64, messengerClient *messenger.Messenger, dataProvider providers.IBotDataProvider, authorizationProvider providers.AuthorizationProvider, adminDataProvider providers.AdminDataProvider) *SuchefServer {
	authorizer := concreteProviders.NewAuthorizer(authorizationProvider)
	controller := NewSuchefController(accountID, messengerClient, dataProvider)
	adminController := NewAdminController(adminDataProvider)
	return &SuchefServer{controller: controller, adminController: adminController, authorizer: authorizer, accountID: accountID}
}

func (suchefServer *SuchefServer) BindMessageReceived() messenger.MessageReceivedHandler {
	return func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		fmt.Println("got message: " + msg.Text)
		err := suchefServer.controller.Handle(suchefServer.accountID, event, opts, msg)
		if err != nil {
			fmt.Println("error handling message: " + err.Error())
		}
		fmt.Println("handled message: " + msg.Text)
	}
}

func (suchefServer *SuchefServer) BindPostbackReceived() messenger.PostbackHandler {
	return func(event messenger.Event, opts messenger.MessageOpts, postback messenger.Postback) {
		fakeMid := fmt.Sprintf("postback_%d", event.ID)
		fmt.Println("got postback: " + fakeMid)
		postbackMessage := messenger.ReceivedMessage{ID: fakeMid, Text: postback.Payload, Seq: -1}
		err := suchefServer.controller.Handle(suchefServer.accountID, event, opts, postbackMessage)
		if err != nil {
			fmt.Println("error handling postback: " + err.Error())
		}
		fmt.Println("handled posback: " + postback.Payload)
	}
}

type addCourseRequest struct {
	Token         string
	Name          string
	Description   string
	MainImageUrl  string
	Diets         []*string
	Sensitivities []*string
	Themes        []*string
}

func (suchefServer *SuchefServer) AddCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIDStr := vars["account_id"]
	accountIDInt, error := strconv.Atoi(accountIDStr)
	accountID := int64(accountIDInt)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var params addCourseRequest
	decoder := json.NewDecoder(r.Body)
	error = decoder.Decode(&params)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	error = suchefServer.authorizer.Authorize(accountID, params.Token)
	if error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newCourse := models.NewCourse(int64(-1), params.Name, params.Description, params.MainImageUrl)
	newCourse.Tags[models.DietMetadataTypeName] = params.Diets
	newCourse.Tags[models.SensitivityMetadataTypeName] = params.Sensitivities
	newCourse.Tags[models.ThemeMetadataTypeName] = params.Themes

	error = suchefServer.adminController.AddCourse(accountID, newCourse)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
