package model

type CovidData struct {
	Regional_id string `bson:"regional_id" json:"regional_id"`
	CovidCases float64 `bson:"covidCases" json:"covidCases"`
	Last_updated string `bson:"last_updated" json:"last_updated"`
}