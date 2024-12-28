package bridge

import (
	"authentication/cmd/entity"
	"authentication/cmd/helper"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Bridge struct {
	helper helper.Helper
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Bridge) logItem(w http.ResponseWriter, entry LogPayload) {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.helper.ErrorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.helper.ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.helper.ErrorJSON(w, err)
		return
	}

	var payload entity.JsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.helper.WriteJSON(w, http.StatusAccepted, payload)

}

func (app *Bridge) sendMail(w http.ResponseWriter, msg MailPayload) {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	// call the mail service
	mailServiceURL := "http://mailer-service/send"

	// post to mail service
	request, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.helper.ErrorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.helper.ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the right status code
	if response.StatusCode != http.StatusAccepted {
		app.helper.ErrorJSON(w, errors.New("error calling mail service"))
		return
	}

	// send back json
	var payload entity.JsonResponse
	payload.Error = false
	payload.Message = "Message sent to " + msg.To

	app.helper.WriteJSON(w, http.StatusAccepted, payload)

}
