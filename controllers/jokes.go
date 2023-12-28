package controllers

import (
	"fmt"

	"github.com/ZainJavedDev/gqlgen-jokes-2/dbModels"
	"github.com/ZainJavedDev/gqlgen-jokes-2/graph/model"
	"github.com/ZainJavedDev/gqlgen-jokes-2/utils"
)

func GetAllJokes() []*model.Joke {

	db := utils.ConnectDB()

	dbJokes := []dbModels.Joke{}
	result := db.Find(&dbJokes)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	jokes := []*model.Joke{}

	for _, dbJoke := range dbJokes {
		joke := &model.Joke{
			ID:   fmt.Sprint(dbJoke.ID), // Convert the ID to string
			Text: dbJoke.Text,
		}
		jokes = append(jokes, joke)
	}

	return jokes
}
