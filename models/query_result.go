package models

type QueryResultInput struct {
	JobId        string `bson:"job_id" json:"job_id"`
	InputProtein string `bson:"input_protein" json:"input_protein"`
	IsSelected   bool   `bson:"is_selected" json:"is_selected"`
}
