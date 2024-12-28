package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Config) ForgetPass(w http.ResponseWriter, forget ForgetPayload) {
	jsonData, _ := json.MarshalIndent(forget, "", "\t")
	fmt.Println(jsonData)

	resData := jsonResponse{
		Error:   false,
		Message: "Forget Pass sended",
	}
	app.writeJSON(w, http.StatusOK, resData)
}
