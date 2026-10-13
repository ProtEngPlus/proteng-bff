package models

import "time"

type MutationInput struct {
	JobId        string                 `bson:"job_id" json:"job_id"`
	InputProtein string                 `bson:"input_protein" json:"input_protein"`
	Options      map[string]interface{} `bson:"options" json:"options"`
}

type MutationResultInput struct {
	IsBookmark bool `bson:"is_bookmark" json:"is_bookmark"`
}

type ExperimentalResultInput struct {
	ActualAssayScore float32   `bson:"actual_assay_score" json:"actual_assay_score"`
	Note             string    `bson:"note" json:"note"`
	MeasuredAt       time.Time `bson:"measured_at" json:"measured_at"`
}
