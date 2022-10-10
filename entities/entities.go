package entities

// PlayRequest struct defines the structure of a game result
type PlayRequest struct {
	Player   int8 `json:"player"`
	Computer int8 `json:"computer"`
}

// PlayResponse struct defines the structure of a game result
type PlayResponse struct {
	Results  string `json:"results"`
	Player   int8   `json:"player"`
	Computer int8   `json:"computer"`
}

// RandomChoice struct defines the response gotten from choice service
type RandomChoice struct {
	Id   int8   `json:"id"`
	Name string `json:"name"`
}

// ChoiceRule struct defines the structure of the rules of the Rock Paper Scissors Spock Lizard game
type ChoiceRule struct {
	Id    int8   `json:"id"`
	Name  string `json:"name"`
	Beats [2]int8
}
