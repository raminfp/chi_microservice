package main

import (
	"authentication/cmd/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

type ForgetPayload struct {
	Email string `json:"email"`
}

func (app *Config) ForgetPass(w http.ResponseWriter, forget ForgetPayload) {
	jsonData, _ := json.MarshalIndent(forget, "", "\t")
	fmt.Println(jsonData)

	resData := entity.JsonResponse{
		Error:   false,
		Message: "Forget Pass sended",
	}
	app.Helper.WriteJSON(w, http.StatusOK, resData)
}
