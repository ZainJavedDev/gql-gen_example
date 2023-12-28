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
			ID:   fmt.Sprint(dbJoke.ID),
			Text: dbJoke.Text,
		}
		jokes = append(jokes, joke)
	}

	return jokes
}

func GetAJoke(id string) *model.Joke {

	db := utils.ConnectDB()

	dbJoke := &dbModels.Joke{}
	result := db.Where("id = ?", id).First(dbJoke)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	joke := &model.Joke{
		ID:   fmt.Sprint(dbJoke.ID),
		Text: dbJoke.Text,
	}

	return joke
}

func CreateNewJoke(input model.JokeInput) *model.Joke {
	db := utils.ConnectDB()
	dbJoke := &dbModels.Joke{
		Text: input.Content,
	}
	db.Create(dbJoke)
	return &model.Joke{
		ID:   fmt.Sprint(dbJoke.ID),
		Text: dbJoke.Text,
	}

}
