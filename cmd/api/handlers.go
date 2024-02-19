package main

import (
	"log"
	"net/http"

	"github.com/clydotron/go-micro-log-service/models"

	helpers "github.com/clydotron/toolbox/helpers"
)

// TODO figure out better place to put these
type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *App) WriteLog(w http.ResponseWriter, r *http.Request) {
	var payload jsonPayload
	err := helpers.ReadJSON(w, r, &payload)
	if err != nil {
		log.Println("error reading the json:", err)
		return
	}

	event := models.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	}

	if err := app.DataStore.LogRepo.Insert(event); err != nil {
		helpers.ErrorJSON(w, err)
		return
	}

	log.Println("successfully wrote logs to DB")
	resp := helpers.JsonResponse{
		Error:   false,
		Message: "logged",
	}
	helpers.WriteJSON(w, http.StatusAccepted, resp)
}
