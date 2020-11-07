package models


type Meme struct {
	OLSentences []string `json:"oLSentences"`
	Difficulty  int      `json:"difficulty"`
	SLSentences []string `json:"sLSentences"`
	Tag         []string `json:"tag"`
	Phrase      []string `json:"phrase"`
}