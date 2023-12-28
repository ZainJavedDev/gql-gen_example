package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ZainJavedDev/gqlgen-jokes-2/dbModels"
	"github.com/ZainJavedDev/gqlgen-jokes-2/utils"
)

func MigrationController(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectDB()

	_ = db.Create(&dbModels.Joke{Text: "Funny Joke 3"})

	// db.AutoMigrate(&dbModels.Joke{})

	responseData := map[string]interface{}{
		"message": "Query completed successfully!",
	}

	json.NewEncoder(w).Encode(responseData)
}
