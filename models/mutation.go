package models

type MutationInput struct {
	JobId        string                 `bson:"job_id" json:"job_id"`
	InputProtein string                 `bson:"input_protein" json:"input_protein"`
	Options      map[string]interface{} `bson:"options" json:"options"`
}

type MutationResultInput struct {
	IsBookmark bool `bson:"is_bookmark" json:"is_bookmark"`
}
