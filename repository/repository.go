package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ong-gtp/play-service/entities"
	"github.com/ong-gtp/play-service/helper"
)

// Repository interfaces defines the service repository's implementation structure
type Repository interface {
	EvaluateGame(player, opponent int8) (entities.PlayResponse, error)
}

// repository struct implements the Repository interface
type repository struct{}

// NewRepository creats a new instance of repository
func NewRepository() *repository {
	return &repository{}
}

// opponentChoice call the external RANDOM_CHOICE_URL for an opponent
func opponentChoice() (int8, error) {
	u := os.Getenv("RANDOM_CHOICE_URL")
	if u == "" {
		return 0, helper.ErrRandomUrlServiceNotPassed
	}

	response, err := http.Get(u)
	if err != nil {
		helper.Log("error", "error :", err)
		return 0, helper.ErrRandomUrlServiceInvalid
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		helper.Log("error", "error :", err)
		return 0, helper.ErrRandomUrlServiceInvalid
	}

	var choiceRes entities.RandomChoice
	err = json.Unmarshal(responseData, &choiceRes)
	if err != nil {
		helper.Log("error", "error :", err)
		return 0, helper.ErrRandomUrlServiceInvalid
	}
	helper.Log("debug", "random choice responseData", choiceRes.Name)

	return choiceRes.Id, nil
}

// EvaluteGame resovles a game session result
func (r *repository) EvaluateGame(player, opponent int8) (entities.PlayResponse, error) {

	if opponent == 0 {
		o, err := opponentChoice()
		if err != nil {
			return entities.PlayResponse{}, err
		}
		opponent = o
	}

	result := playResult(player, opponent)
	return entities.PlayResponse{Results: result, Player: player, Computer: opponent}, nil
}

// choiceRules defines the rules of the RPSSL game
func choiceRules() map[int8]entities.ChoiceRule {
	rules := make(map[int8]entities.ChoiceRule)
	rules[1] = entities.ChoiceRule{Id: 1, Name: "rock", Beats: [2]int8{3, 5}}     //["scissors", "lizard"]
	rules[2] = entities.ChoiceRule{Id: 2, Name: "paper", Beats: [2]int8{1, 4}}    //["rock", "spock"]
	rules[3] = entities.ChoiceRule{Id: 3, Name: "scissors", Beats: [2]int8{2, 5}} //["paper", "lizard"],
	rules[4] = entities.ChoiceRule{Id: 4, Name: "spock", Beats: [2]int8{1, 3}}    //["rock", "scissors"],
	rules[5] = entities.ChoiceRule{Id: 5, Name: "lizard", Beats: [2]int8{2, 4}}   //["paper", "spock"],
	return rules
}

// playResult resolves winner of a game
func playResult(player, opponent int8) string {
	rule := choiceRules()[player]
	if player == opponent {
		return "tie"
	}
	for _, v := range rule.Beats {
		if v == opponent {
			return "win"
		}
	}
	return "lose"
}
